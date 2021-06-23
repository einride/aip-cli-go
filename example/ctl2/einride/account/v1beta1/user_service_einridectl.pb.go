package accountv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.account.v1beta1.UserService.
var (
	einride_account_v1beta1_UserServiceClient v1beta1.UserServiceClient
	einride_account_v1beta1_UserService       = &cobra.Command{
		Use:   "account.v1beta1.UserService",
		Short: "User service.",
		Long:  "User service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_account_v1beta1_UserServiceClient = v1beta1.NewUserServiceClient(conn)
			return nil
		},
	}
)

// einride.account.v1beta1.UserService.CreateUser.
var (
	einride_account_v1beta1_UserService_CreateUser_Request v1beta1.CreateUserRequest
	einride_account_v1beta1_UserService_CreateUser         = &cobra.Command{
		Use: "CreateUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_UserServiceClient.CreateUser(cmd.Context(), &einride_account_v1beta1_UserService_CreateUser_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.UserService.GetUser.
var (
	einride_account_v1beta1_UserService_GetUser_Request v1beta1.GetUserRequest
	einride_account_v1beta1_UserService_GetUser         = &cobra.Command{
		Use: "GetUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_UserServiceClient.GetUser(cmd.Context(), &einride_account_v1beta1_UserService_GetUser_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.UserService.BatchGetUsers.
var (
	einride_account_v1beta1_UserService_BatchGetUsers_Request v1beta1.BatchGetUsersRequest
	einride_account_v1beta1_UserService_BatchGetUsers         = &cobra.Command{
		Use: "BatchGetUsers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_UserServiceClient.BatchGetUsers(cmd.Context(), &einride_account_v1beta1_UserService_BatchGetUsers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.UserService.UpdateUser.
var (
	einride_account_v1beta1_UserService_UpdateUser_Request v1beta1.UpdateUserRequest
	einride_account_v1beta1_UserService_UpdateUser         = &cobra.Command{
		Use: "UpdateUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_UserServiceClient.UpdateUser(cmd.Context(), &einride_account_v1beta1_UserService_UpdateUser_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.UserService.ListUsers.
var (
	einride_account_v1beta1_UserService_ListUsers_Request v1beta1.ListUsersRequest
	einride_account_v1beta1_UserService_ListUsers         = &cobra.Command{
		Use: "ListUsers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_UserServiceClient.ListUsers(cmd.Context(), &einride_account_v1beta1_UserService_ListUsers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddUserServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_UserService)
}

func init() {
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_CreateUser)

	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.Parent, "parent", "", "Resource name of the parent tenant where this user will be created.")

	einride_account_v1beta1_UserService_CreateUser_Request.User = new(v1beta1.User)
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.Name, "user.name", "", "The resource name of the user.")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.Etag, "user.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.DisplayName, "user.displayName", "", "The user's display name.")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.Email, "user.email", "", "The user's email.")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.PhoneNumber, "user.phoneNumber", "", "The user's phone number.")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.PhotoUrl, "user.photoUrl", "", "The user's photo URL.\n(-- api-linter: core::0140::uri=disabled\n    aip.dev/not-precedent: Backwards compatibility. --)")
	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.User.Password, "user.password", "", "The user's password. Only used to set a password when creating a user.")

	einride_account_v1beta1_UserService_CreateUser.Flags().StringVar(&einride_account_v1beta1_UserService_CreateUser_Request.UserId, "userId", "", "The user ID of the user to create.\nThis field is deprecated. The user ID is always a system-generated random UUID.\nThis field will be removed in a future schema update.")
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_GetUser)

	einride_account_v1beta1_UserService_GetUser.Flags().StringVar(&einride_account_v1beta1_UserService_GetUser_Request.Name, "name", "", "Resource name of the user to retrieve.")
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_BatchGetUsers)

	einride_account_v1beta1_UserService_BatchGetUsers.Flags().StringVar(&einride_account_v1beta1_UserService_BatchGetUsers_Request.Parent, "parent", "", "Resource name of the parent tenant shared by all users being retrieved.\nIf this is set, the parent of all of the users specified in `names`\nmust match this field.")

	einride_account_v1beta1_UserService_BatchGetUsers.Flags().StringSliceVar(&einride_account_v1beta1_UserService_BatchGetUsers_Request.Names, "names", []string{}, "Resource names of the users to retrieve.\nA maximum of 1000 users can be retrieved in a batch.")
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_UpdateUser)

	einride_account_v1beta1_UserService_UpdateUser_Request.User = new(v1beta1.User)
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.Name, "user.name", "", "The resource name of the user.")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.Etag, "user.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.DisplayName, "user.displayName", "", "The user's display name.")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.Email, "user.email", "", "The user's email.")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.PhoneNumber, "user.phoneNumber", "", "The user's phone number.")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.PhotoUrl, "user.photoUrl", "", "The user's photo URL.\n(-- api-linter: core::0140::uri=disabled\n    aip.dev/not-precedent: Backwards compatibility. --)")
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringVar(&einride_account_v1beta1_UserService_UpdateUser_Request.User.Password, "user.password", "", "The user's password. Only used to set a password when creating a user.")

	einride_account_v1beta1_UserService_UpdateUser_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_account_v1beta1_UserService_UpdateUser.Flags().StringSliceVar(&einride_account_v1beta1_UserService_UpdateUser_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_ListUsers)

	einride_account_v1beta1_UserService_ListUsers.Flags().StringVar(&einride_account_v1beta1_UserService_ListUsers_Request.Parent, "parent", "", "Resource name of the parent tenant.")

	einride_account_v1beta1_UserService_ListUsers.Flags().Int32Var(&einride_account_v1beta1_UserService_ListUsers_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_account_v1beta1_UserService_ListUsers.Flags().StringVar(&einride_account_v1beta1_UserService_ListUsers_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
