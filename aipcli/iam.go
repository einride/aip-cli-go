package aipcli

import (
	"os"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewIAMModuleCommand returns a *cobra.Command for standard IAM operations.
func NewIAMModuleCommand(name string, config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Identity and Access Management (IAM)",
		Long: strings.TrimSpace(`
Manage Identity and Access Management (IAM) policies.

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
			moduleAnnotation: name,
		},
	}
	initPersistentFlags(cmd)
	setConfig(cmd, config)
	cmd.AddCommand(newSetIAMPolicyCommand(config))
	cmd.AddCommand(newGetIAMPolicyCommand(config))
	cmd.AddCommand(newAddIAMPolicyBindingCommand(config))
	cmd.AddCommand(newRemoveIAMPolicyBindingCommand(config))
	return cmd
}

func newGetIAMPolicyCommand(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-iam-policy",
		Short: "Get the IAM policy for a resource",
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
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			initContext(cmd, config)
		},
		Annotations: map[string]string{
			methodAnnotation: "google.iam.v1.IAMPolicy.GetIamPolicy",
		},
	}
	initPersistentFlags(cmd)
	setConfig(cmd, config)
	resource := cmd.Flags().String("resource", "", "Resource for which the policy is being requested")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		return invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iampb.GetIamPolicyRequest{Resource: *resource},
			&iampb.Policy{},
		)
	}
	return cmd
}

func newSetIAMPolicyCommand(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-iam-policy",
		Short: "Set the IAM policy for a resource",
		Long: strings.TrimSpace(`
Set the IAM policy associated with a resource.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.

  ▪ The specified policy contains invalid or non-existent 
    roles or members.

  ▪ The active account does not have permission to 
    access the given resource's IAM policies.
`),
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			initContext(cmd, config)
		},
		Annotations: map[string]string{
			methodAnnotation: "google.iam.v1.IAMPolicy.SetIAMPolicy",
		},
	}
	initPersistentFlags(cmd)
	setConfig(cmd, config)
	resource := cmd.Flags().String("resource", "", "Resource name for which the policy is being specified")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	_ = cmd.Flags().SetAnnotation("resource", flagArgumentAnnotation, []string{})
	policyFile := cmd.Flags().String("policy-file", "", "Path to a local JSON file containing a valid policy")
	_ = cmd.MarkFlagRequired("policy-file")
	_ = cmd.MarkFlagFilename("policy-file", "json")
	_ = cmd.Flags().SetAnnotation("policy-file", flagArgumentAnnotation, []string{})
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		data, err := os.ReadFile(*policyFile)
		if err != nil {
			return err
		}
		var policy iampb.Policy
		if err := protojson.Unmarshal(data, &policy); err != nil {
			return err
		}
		return invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iampb.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iampb.Policy{},
		)
	}
	return cmd
}

func newAddIAMPolicyBindingCommand(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-iam-policy-binding",
		Short: "Add an IAM policy binding to the IAM policy of a resource",
		Long: strings.TrimSpace(`
Adds an IAM policy binding to the IAM policy of a resource. 

One binding consists of a member, a role, and an optional condition.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			initContext(cmd, config)
		},
		Annotations: map[string]string{
			methodAnnotation: "add-iam-policy-binding", // custom method
		},
	}
	initPersistentFlags(cmd)
	setConfig(cmd, config)
	resource := cmd.Flags().String("resource", "", "Resource name for which the policy binding is being added")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	_ = cmd.Flags().SetAnnotation("resource", flagArgumentAnnotation, nil)
	member := cmd.Flags().String("member", "", "Principal to add the binding for")
	_ = cmd.MarkFlagRequired("member")
	_ = cmd.RegisterFlagCompletionFunc("member", completeMember)
	_ = cmd.Flags().SetAnnotation("member", flagArgumentAnnotation, nil)
	role := cmd.Flags().String("role", "", "Role name to assign to the principal")
	_ = cmd.MarkFlagRequired("role")
	_ = cmd.RegisterFlagCompletionFunc("role", completeRole)
	_ = cmd.Flags().SetAnnotation("role", flagArgumentAnnotation, nil)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		var policy iampb.Policy
		if err := invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iampb.GetIamPolicyRequest{Resource: *resource},
			&policy,
		); err != nil {
			return err
		}
		addBinding(&policy, *member, *role)
		return invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iampb.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iampb.Policy{},
		)
	}
	return cmd
}

func newRemoveIAMPolicyBindingCommand(config Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-iam-policy-binding",
		Short: "Remove an IAM policy binding from the IAM policy of a resource",
		Long: strings.TrimSpace(`
Removes an IAM policy binding from the IAM policy of a resource. 
One binding consists of a member, a role, and an optional condition.

This command can fail for the following reasons:

  ▪ The resource specified does not exist.
  ▪ The active account does not have permission to access the given
	resource's IAM policies.
`),
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			initContext(cmd, config)
		},
		Annotations: map[string]string{
			methodAnnotation: "remove-iam-policy-binding", // custom method
		},
	}
	initPersistentFlags(cmd)
	setConfig(cmd, config)
	resource := cmd.Flags().String("resource", "", "Resource name for which the policy binding is being removed")
	_ = cmd.MarkFlagRequired("resource")
	_ = cmd.RegisterFlagCompletionFunc("resource", completeResource)
	_ = cmd.Flags().SetAnnotation("resource", flagArgumentAnnotation, nil)
	member := cmd.Flags().String("member", "", "Principal to remove the binding for")
	_ = cmd.MarkFlagRequired("member")
	_ = cmd.RegisterFlagCompletionFunc("member", completeMember)
	_ = cmd.Flags().SetAnnotation("member", flagArgumentAnnotation, nil)
	role := cmd.Flags().String("role", "", "Role name to remove the principal from")
	_ = cmd.MarkFlagRequired("role")
	_ = cmd.RegisterFlagCompletionFunc("role", completeRole)
	_ = cmd.Flags().SetAnnotation("role", flagArgumentAnnotation, nil)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		var policy iampb.Policy
		if err := invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/GetIamPolicy",
			&iampb.GetIamPolicyRequest{Resource: *resource},
			&policy,
		); err != nil {
			return err
		}
		removeBinding(&policy, *member, *role)
		return invoke(
			cmd,
			"/google.iam.v1.IAMPolicy/SetIamPolicy",
			&iampb.SetIamPolicyRequest{Resource: *resource, Policy: &policy},
			&iampb.Policy{},
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

func addBinding(policy *iampb.Policy, member, role string) {
	// look for existing binding with this role and member
	for _, binding := range policy.GetBindings() {
		if binding.GetRole() == role {
			for _, bindingMember := range binding.GetMembers() {
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
	policy.Bindings = append(policy.Bindings, &iampb.Binding{
		Role:    role,
		Members: []string{member},
	})
}

func removeBinding(policy *iampb.Policy, member, role string) {
	for _, binding := range policy.GetBindings() {
		if binding.GetRole() == role {
			binding.Members = removeMember(binding.GetMembers(), member)
			if len(binding.GetMembers()) == 0 {
				policy.Bindings = removeRole(policy.GetBindings(), role)
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

func removeRole(bindings []*iampb.Binding, role string) []*iampb.Binding {
	for i, binding := range bindings {
		if binding.GetRole() == role {
			return append(bindings[:i], bindings[i+1:]...)
		}
	}
	return bindings
}
