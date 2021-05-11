package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_messenger_v1beta1_EmailService = &cobra.Command{
	Use: "einride.messenger.v1beta1.Email",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.messenger.v1beta1.Email called")
	},
}

var einride_messenger_v1beta1_EmailService_GetEmail = &cobra.Command{
	Use: "GetEmail",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetEmail called")
	},
}

var einride_messenger_v1beta1_EmailService_CreateEmail = &cobra.Command{
	Use: "CreateEmail",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateEmail called")
	},
}

var einride_messenger_v1beta1_EmailService_ListEmails = &cobra.Command{
	Use: "ListEmails",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListEmails called")
	},
}

var einride_messenger_v1beta1_EmailService_SendEmail = &cobra.Command{
	Use: "SendEmail",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SendEmail called")
	},
}

var einride_messenger_v1beta1_EmailService_SendNewEmail = &cobra.Command{
	Use: "SendNewEmail",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SendNewEmail called")
	},
}

var einride_messenger_v1beta1_SMSMessageService = &cobra.Command{
	Use: "einride.messenger.v1beta1.SMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.messenger.v1beta1.SMSMessage called")
	},
}

var einride_messenger_v1beta1_SMSMessageService_GetSMSMessage = &cobra.Command{
	Use: "GetSMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetSMSMessage called")
	},
}

var einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage = &cobra.Command{
	Use: "CreateSMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateSMSMessage called")
	},
}

var einride_messenger_v1beta1_SMSMessageService_ListSMSMessages = &cobra.Command{
	Use: "ListSMSMessages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListSMSMessages called")
	},
}

var einride_messenger_v1beta1_SMSMessageService_SendSMSMessage = &cobra.Command{
	Use: "SendSMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SendSMSMessage called")
	},
}

var einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage = &cobra.Command{
	Use: "SendNewSMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SendNewSMSMessage called")
	},
}

func init() {
	rootCmd.AddCommand(einride_messenger_v1beta1_EmailService)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_GetEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_CreateEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_ListEmails)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendNewEmail)
	rootCmd.AddCommand(einride_messenger_v1beta1_SMSMessageService)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_GetSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_ListSMSMessages)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage)
}
