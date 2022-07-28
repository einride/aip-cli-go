package aipcli

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewIAMModuleCommand returns a *cobra.Command for standard IAM operations.
func NewIAMModuleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iam",
		Short: "IAM policies and access control",
		Long: strings.TrimSpace(`
Manages Identity and Access Management (IAM) policies.

Any implementation of an API that offers access control features
implements the google.iam.v1.IAMPolicy interface.

Access control is applied when a principal (user or service account), takes
some action on a resource exposed by a service. Resources, identified by
URI-like names, are the unit of access control specification. Service
implementations can choose the granularity of access control and the
supported permissions for their resources.

For example one database service may allow access control to be
specified only at the Table level, whereas another might allow access control
to also be specified at the Column level.

This is intentionally not a CRUD style API because access control policies
are created and deleted implicitly with the resources to which they are
attached.
		`),
		Annotations: map[string]string{
			moduleNameAnnotation: "iam",
		},
	}
	cmd.AddCommand(newSetIAMPolicyCommand())
	cmd.AddCommand(newGetIAMPolicyCommand())
	cmd.AddCommand(newAddIAMPolicyBindingCommand())
	cmd.AddCommand(newRemoveIAMPolicyBindingCommand())
	return cmd
}

func newGetIAMPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-iam-policy",
		Short: "get the IAM policy for a resource",
		Long: strings.TrimSpace(`
Get the IAM policy associated with a resource. 

The output includes an "etag" identifier that is used to check for 
concurrent policy updates. An edited policy should include the same 
etag so that set-iam-policy applies the changes to the correct policy 
version.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
	}
	resource := cmd.Flags().String("resource", "", "the resource for which the policy is being requested")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iam.GetIamPolicyRequest{Resource: *resource},
			&iam.Policy{},
		)
	}
	return cmd
}

func newSetIAMPolicyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-iam-policy",
		Short: "set the IAM policy for a resource",
		Long: strings.TrimSpace(`
Set the IAM policy associated with a resource.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
	}
	resource := cmd.Flags().String("resource", "", "the resource for which the policy is being specified")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	policyFile := cmd.Flags().String("policy-file", "", "path to a local JSON file containing a valid policy")
	_ = cmd.MarkFlagRequired("policy-file")
	_ = cmd.MarkFlagFilename("policy-file", "json")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(*policyFile)
		if err != nil {
			return err
		}
		var policy iam.Policy
		if err := protojson.Unmarshal(data, &policy); err != nil {
			return err
		}
		return invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iam.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iam.Policy{},
		)
	}
	return cmd
}

func newAddIAMPolicyBindingCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-iam-policy-binding",
		Short: "add an IAM policy binding to the IAM policy of a resource",
		Long: strings.TrimSpace(`
Adds an IAM policy binding to the IAM policy of a resource. 

One binding consists of a member, a role, and an optional condition.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
	}
	resource := cmd.Flags().String("resource", "", "the resource for which the policy binding is being added")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	member := cmd.Flags().String("member", "", "principal to add the binding for")
	_ = cmd.MarkFlagRequired("member")
	_ = cmd.RegisterFlagCompletionFunc("member", completeMember)
	role := cmd.Flags().String("role", "", "role name to assign to the principal")
	_ = cmd.MarkFlagRequired("role")
	_ = cmd.RegisterFlagCompletionFunc("role", completeRole)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var policy iam.Policy
		if err := invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iam.GetIamPolicyRequest{Resource: *resource},
			&policy,
		); err != nil {
			return err
		}
		addBinding(&policy, *member, *role)
		return invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iam.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iam.Policy{},
		)
	}
	return cmd
}

func newRemoveIAMPolicyBindingCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-iam-policy-binding",
		Short: "remove an IAM policy binding from the IAM policy of a resource",
		Long: strings.TrimSpace(`
Removes an IAM policy binding from the IAM policy of a resource. 
One binding consists of a member, a role, and an optional condition.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
	}
	resource := cmd.Flags().String("resource", "", "the resource for which the policy binding is being removed")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	member := cmd.Flags().String("member", "", "principal to remove the binding for")
	_ = cmd.MarkFlagRequired("member")
	_ = cmd.RegisterFlagCompletionFunc("member", completeMember)
	role := cmd.Flags().String("role", "", "role name to remove the principal from")
	_ = cmd.MarkFlagRequired("role")
	_ = cmd.RegisterFlagCompletionFunc("role", completeRole)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var policy iam.Policy
		if err := invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iam.GetIamPolicyRequest{Resource: *resource},
			&policy,
		); err != nil {
			return err
		}
		removeBinding(&policy, *member, *role)
		return invoke(
			cmd.Context(),
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iam.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iam.Policy{},
		)
	}
	return cmd
}

func completeResource(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return cobra.AppendActiveHelp(
		nil,
		"Please enter a valid resource name. Example: resources/1234",
	), cobra.ShellCompDirectiveNoFileComp
}

func completeMember(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return cobra.AppendActiveHelp(
		nil,
		"Please enter a valid member. Example: email:foo@example.com",
	), cobra.ShellCompDirectiveNoFileComp
}

func completeRole(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	// TODO: Use the QueryGrantableRoles method to auto-complete grantable roles.
	return cobra.AppendActiveHelp(
		nil,
		"Please enter a valid role name. Example: roles/example.admin",
	), cobra.ShellCompDirectiveNoFileComp
}

func addBinding(policy *iam.Policy, member, role string) {
	// look for existing binding with this role and member
	for _, binding := range policy.Bindings {
		if binding.Role == role {
			for _, bindingMember := range binding.Members {
				if bindingMember == member {
					// already have a binding with this role and member
					return
				}
			}
			// already have a binding with this role, but not the member
			binding.Members = append(binding.Members, member)
			return
		}
	}
	// add a new binding with this role and member
	policy.Bindings = append(policy.Bindings, &iam.Binding{
		Role:    role,
		Members: []string{member},
	})
}

func removeBinding(policy *iam.Policy, member, role string) {
	for _, binding := range policy.Bindings {
		if binding.Role == role {
			binding.Members = removeMember(binding.Members, member)
			if len(binding.Members) == 0 {
				policy.Bindings = removeRole(policy.Bindings, role)
			}
			return
		}
	}
}

func removeMember(members []string, member string) []string {
	for i, candidate := range members {
		if candidate == member {
			return append(members[:i], members[i+1:]...)
		}
	}
	return members
}

func removeRole(bindings []*iam.Binding, role string) []*iam.Binding {
	for i, binding := range bindings {
		if binding.Role == role {
			return append(bindings[:i], bindings[i+1:]...)
		}
	}
	return bindings
}
