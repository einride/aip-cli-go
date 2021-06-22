package messengerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/messenger/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_messenger_v1beta1_EmailService = &cobra.Command{
	Use: "einride.messenger.v1beta1.EmailService",
}

var einride_messenger_v1beta1_GetEmailRequest v1beta1.GetEmailRequest
var einride_messenger_v1beta1_EmailService_GetEmail = &cobra.Command{
	Use: "GetEmail",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.EmailService.GetEmail")
		return nil
	},
}

var einride_messenger_v1beta1_CreateEmailRequest v1beta1.CreateEmailRequest
var einride_messenger_v1beta1_EmailService_CreateEmail = &cobra.Command{
	Use: "CreateEmail",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.EmailService.CreateEmail")
		return nil
	},
}

var einride_messenger_v1beta1_ListEmailsRequest v1beta1.ListEmailsRequest
var einride_messenger_v1beta1_EmailService_ListEmails = &cobra.Command{
	Use: "ListEmails",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.EmailService.ListEmails")
		return nil
	},
}

var einride_messenger_v1beta1_SendEmailRequest v1beta1.SendEmailRequest
var einride_messenger_v1beta1_EmailService_SendEmail = &cobra.Command{
	Use: "SendEmail",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.EmailService.SendEmail")
		return nil
	},
}

var einride_messenger_v1beta1_SendNewEmailRequest v1beta1.SendNewEmailRequest
var einride_messenger_v1beta1_EmailService_SendNewEmail = &cobra.Command{
	Use: "SendNewEmail",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.EmailService.SendNewEmail")
		return nil
	},
}

func AddEmailServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_messenger_v1beta1_EmailService)
}

func init() {
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_GetEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_CreateEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_ListEmails)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendEmail)
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendNewEmail)
}
