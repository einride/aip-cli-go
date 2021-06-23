package messengerv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/messenger/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.messenger.v1beta1.SMSMessageService.
var (
	einride_messenger_v1beta1_SMSMessageServiceClient v1beta1.SMSMessageServiceClient
	einride_messenger_v1beta1_SMSMessageService       = &cobra.Command{
		Use:   "einride.messenger.v1beta1.SMSMessageService",
		Short: "SMSMessage service.",
		Long:  "SMSMessage service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_messenger_v1beta1_SMSMessageServiceClient = v1beta1.NewSMSMessageServiceClient(conn)
			return nil
		},
	}
)

// einride.messenger.v1beta1.SMSMessageService.GetSMSMessage.
var (
	einride_messenger_v1beta1_SMSMessageService_GetSMSMessage_Request v1beta1.GetSMSMessageRequest
	einride_messenger_v1beta1_SMSMessageService_GetSMSMessage         = &cobra.Command{
		Use: "GetSMSMessage",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_messenger_v1beta1_SMSMessageServiceClient.GetSMSMessage(cmd.Context(), &einride_messenger_v1beta1_SMSMessageService_GetSMSMessage_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.messenger.v1beta1.SMSMessageService.CreateSMSMessage.
var (
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request v1beta1.CreateSMSMessageRequest
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage         = &cobra.Command{
		Use: "CreateSMSMessage",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_messenger_v1beta1_SMSMessageServiceClient.CreateSMSMessage(cmd.Context(), &einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.messenger.v1beta1.SMSMessageService.ListSMSMessages.
var (
	einride_messenger_v1beta1_SMSMessageService_ListSMSMessages_Request v1beta1.ListSMSMessagesRequest
	einride_messenger_v1beta1_SMSMessageService_ListSMSMessages         = &cobra.Command{
		Use: "ListSMSMessages",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_messenger_v1beta1_SMSMessageServiceClient.ListSMSMessages(cmd.Context(), &einride_messenger_v1beta1_SMSMessageService_ListSMSMessages_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.messenger.v1beta1.SMSMessageService.SendSMSMessage.
var (
	einride_messenger_v1beta1_SMSMessageService_SendSMSMessage_Request v1beta1.SendSMSMessageRequest
	einride_messenger_v1beta1_SMSMessageService_SendSMSMessage         = &cobra.Command{
		Use: "SendSMSMessage",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_messenger_v1beta1_SMSMessageServiceClient.SendSMSMessage(cmd.Context(), &einride_messenger_v1beta1_SMSMessageService_SendSMSMessage_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.messenger.v1beta1.SMSMessageService.SendNewSMSMessage.
var (
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request v1beta1.SendNewSMSMessageRequest
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage         = &cobra.Command{
		Use: "SendNewSMSMessage",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_messenger_v1beta1_SMSMessageServiceClient.SendNewSMSMessage(cmd.Context(), &einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddSMSMessageServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_messenger_v1beta1_SMSMessageService)
}

func init() {
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_GetSMSMessage)

	einride_messenger_v1beta1_SMSMessageService_GetSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_GetSMSMessage_Request.Name, "name", "", "Resource name of the SMS message to retrieve.\nFormat: `tenants/{tenant}/smsMessages/{sms_message}`.")
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage)

	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.Parent, "parent", "", "Resource name of the parent tenant where the new SMS message will be created.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessage = new(v1beta1.SMSMessage)
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessage.Name, "smsMessage.name", "", "The resource name of the SMS.")
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessage.Etag, "smsMessage.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessage.PhoneNumber, "smsMessage.phoneNumber", "", "The phone number the SMS is sent to.")
	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessage.Content = new(v1beta1.SMSMessage_Content)
	// TODO: one-of: PlainText
	// TODO: one-of: ShipmentEtaNotification

	einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage_Request.SmsMessageId, "smsMessageId", "", "The ID to use for the SMSMessage.\nWill become the final component of the site charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.")
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_ListSMSMessages)

	einride_messenger_v1beta1_SMSMessageService_ListSMSMessages.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_ListSMSMessages_Request.Parent, "parent", "", "Resource name of the parent tenant.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_SMSMessageService_ListSMSMessages.Flags().Int32Var(&einride_messenger_v1beta1_SMSMessageService_ListSMSMessages_Request.PageSize, "pageSize", 10, "The maximum number of SMS messages to return.\nThe service may return fewer site chargers than this value.\nIf unspecified, at most 50 site chargers will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_messenger_v1beta1_SMSMessageService_ListSMSMessages.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_ListSMSMessages_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\nWhen paginating, all other parameters must match the call that provided the page token.")
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendSMSMessage)

	einride_messenger_v1beta1_SMSMessageService_SendSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendSMSMessage_Request.Name, "name", "", "Resource name of the SMS message to send.\nFormat: `tenants/{tenant}/smsMessages/{sms_message}`.")
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage)

	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.Parent, "parent", "", "Resource name of the parent tenant where the new SMS message will be created.\nPattern: `tenants/{tenant}`.")

	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessage = new(v1beta1.SMSMessage)
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessage.Name, "smsMessage.name", "", "The resource name of the SMS.")
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessage.Etag, "smsMessage.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessage.PhoneNumber, "smsMessage.phoneNumber", "", "The phone number the SMS is sent to.")
	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessage.Content = new(v1beta1.SMSMessage_Content)
	// TODO: one-of: PlainText
	// TODO: one-of: ShipmentEtaNotification

	einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage.Flags().StringVar(&einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage_Request.SmsMessageId, "smsMessageId", "", "The ID to use for the SMS message.\nWill become the final component of the site charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.")
}
