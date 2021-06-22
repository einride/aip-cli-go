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

	einride_messenger_v1beta1_EmailService_GetEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_GetEmail_Request.Name, "name", "", "Resource name of the email to retrieve.\nFormat: `tenants/{tenant}/emails/{email}`.")
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_CreateEmail)

	einride_messenger_v1beta1_EmailService_CreateEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_CreateEmail_Request.Parent, "parent", "", "Resource name of the parent tenant where the new email will be created.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_EmailService_CreateEmail_Request.Email = new(v1beta1.Email)
	einride_messenger_v1beta1_EmailService_CreateEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_CreateEmail_Request.Email.Name, "email.name", "", "The resource name of the email.")
	einride_messenger_v1beta1_EmailService_CreateEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_CreateEmail_Request.Email.Etag, "email.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: list: EmailAddresses message
	// TODO: list: CcEmailAddresses message
	// TODO: list: BccEmailAddresses message
	einride_messenger_v1beta1_EmailService_CreateEmail_Request.Email.Content = new(v1beta1.Email_Content)
	// TODO: one-of: PlainText

	einride_messenger_v1beta1_EmailService_CreateEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_CreateEmail_Request.EmailId, "emailId", "", "The ID to use for the email.\nWill become the final component of the site charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.")
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_ListEmails)

	einride_messenger_v1beta1_EmailService_ListEmails.Flags().StringVar(&einride_messenger_v1beta1_EmailService_ListEmails_Request.Parent, "parent", "", "Resource name of the parent tenant.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_EmailService_ListEmails.Flags().Int32Var(&einride_messenger_v1beta1_EmailService_ListEmails_Request.PageSize, "pageSize", 10, "The maximum number of emails to return.\nThe service may return fewer site chargers than this value.\nIf unspecified, at most 50 site chargers will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_messenger_v1beta1_EmailService_ListEmails.Flags().StringVar(&einride_messenger_v1beta1_EmailService_ListEmails_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\nWhen paginating, all other parameters must match the call that provided the page token.")
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendEmail)

	einride_messenger_v1beta1_EmailService_SendEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_SendEmail_Request.Name, "name", "", "Resource name of the email to send.\nFormat: `tenants/{tenant}/emails/{email}`.")
	einride_messenger_v1beta1_EmailService.AddCommand(einride_messenger_v1beta1_EmailService_SendNewEmail)

	einride_messenger_v1beta1_EmailService_SendNewEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_SendNewEmail_Request.Parent, "parent", "", "Resource name of the parent tenant where the new email will be created.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_EmailService_SendNewEmail_Request.Email = new(v1beta1.Email)
	einride_messenger_v1beta1_EmailService_SendNewEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_SendNewEmail_Request.Email.Name, "email.name", "", "The resource name of the email.")
	einride_messenger_v1beta1_EmailService_SendNewEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_SendNewEmail_Request.Email.Etag, "email.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: list: EmailAddresses message
	// TODO: list: CcEmailAddresses message
	// TODO: list: BccEmailAddresses message
	einride_messenger_v1beta1_EmailService_SendNewEmail_Request.Email.Content = new(v1beta1.Email_Content)
	// TODO: one-of: PlainText

	einride_messenger_v1beta1_EmailService_SendNewEmail.Flags().StringVar(&einride_messenger_v1beta1_EmailService_SendNewEmail_Request.EmailId, "emailId", "", "The ID to use for the email.\nWill become the final component of the site charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.")
}
