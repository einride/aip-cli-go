package accountv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_account_v1beta1_UserService = &cobra.Command{
	Use: "einride.account.v1beta1.UserService",
}

var einride_account_v1beta1_CreateUserRequest v1beta1.CreateUserRequest
var einride_account_v1beta1_UserService_CreateUser = &cobra.Command{
	Use: "CreateUser",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.UserService.CreateUser")
		return nil
	},
}

var einride_account_v1beta1_GetUserRequest v1beta1.GetUserRequest
var einride_account_v1beta1_UserService_GetUser = &cobra.Command{
	Use: "GetUser",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.UserService.GetUser")
		return nil
	},
}

var einride_account_v1beta1_BatchGetUsersRequest v1beta1.BatchGetUsersRequest
var einride_account_v1beta1_UserService_BatchGetUsers = &cobra.Command{
	Use: "BatchGetUsers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.UserService.BatchGetUsers")
		return nil
	},
}

var einride_account_v1beta1_UpdateUserRequest v1beta1.UpdateUserRequest
var einride_account_v1beta1_UserService_UpdateUser = &cobra.Command{
	Use: "UpdateUser",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.UserService.UpdateUser")
		return nil
	},
}

var einride_account_v1beta1_ListUsersRequest v1beta1.ListUsersRequest
var einride_account_v1beta1_UserService_ListUsers = &cobra.Command{
	Use: "ListUsers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.UserService.ListUsers")
		return nil
	},
}

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
