package accountv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.account.v1beta1.UserService.
var (
	einride_account_v1beta1_UserServiceClient v1beta1.UserServiceClient
	einride_account_v1beta1_UserService       = &cobra.Command{
		Use: "einride.account.v1beta1.UserService",
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
			log.Println("einride.account.v1beta1.UserService.CreateUser")
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
			log.Println("einride.account.v1beta1.UserService.GetUser")
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
			log.Println("einride.account.v1beta1.UserService.BatchGetUsers")
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
			log.Println("einride.account.v1beta1.UserService.UpdateUser")
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
			log.Println("einride.account.v1beta1.UserService.ListUsers")
			return nil
		},
	}
)

func AddUserServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_UserService)
}

func init() {
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_CreateUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_GetUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_BatchGetUsers)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_UpdateUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_ListUsers)
}
