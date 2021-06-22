package messengerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/messenger/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.messenger.v1beta1.EmailService.
var (
	einride_messenger_v1beta1_EmailServiceClient v1beta1.EmailServiceClient
	einride_messenger_v1beta1_EmailService       = &cobra.Command{
		Use: "einride.messenger.v1beta1.EmailService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_messenger_v1beta1_EmailServiceClient = v1beta1.NewEmailServiceClient(conn)
			return nil
		},
	}
)

// einride.messenger.v1beta1.EmailService.GetEmail.
var (
	einride_messenger_v1beta1_EmailService_GetEmail_Request v1beta1.GetEmailRequest
	einride_messenger_v1beta1_EmailService_GetEmail         = &cobra.Command{
		Use: "GetEmail",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.messenger.v1beta1.EmailService.GetEmail")
			return nil
		},
	}
)

// einride.messenger.v1beta1.EmailService.CreateEmail.
var (
	einride_messenger_v1beta1_EmailService_CreateEmail_Request v1beta1.CreateEmailRequest
	einride_messenger_v1beta1_EmailService_CreateEmail         = &cobra.Command{
		Use: "CreateEmail",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.messenger.v1beta1.EmailService.CreateEmail")
			return nil
		},
	}
)

// einride.messenger.v1beta1.EmailService.ListEmails.
var (
	einride_messenger_v1beta1_EmailService_ListEmails_Request v1beta1.ListEmailsRequest
	einride_messenger_v1beta1_EmailService_ListEmails         = &cobra.Command{
		Use: "ListEmails",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.messenger.v1beta1.EmailService.ListEmails")
			return nil
		},
	}
)

// einride.messenger.v1beta1.EmailService.SendEmail.
var (
	einride_messenger_v1beta1_EmailService_SendEmail_Request v1beta1.SendEmailRequest
	einride_messenger_v1beta1_EmailService_SendEmail         = &cobra.Command{
		Use: "SendEmail",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.messenger.v1beta1.EmailService.SendEmail")
			return nil
		},
	}
)

// einride.messenger.v1beta1.EmailService.SendNewEmail.
var (
	einride_messenger_v1beta1_EmailService_SendNewEmail_Request v1beta1.SendNewEmailRequest
	einride_messenger_v1beta1_EmailService_SendNewEmail         = &cobra.Command{
		Use: "SendNewEmail",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.messenger.v1beta1.EmailService.SendNewEmail")
			return nil
		},
	}
)

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
