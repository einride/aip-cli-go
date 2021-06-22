package messengerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/messenger/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.messenger.v1beta1.SMSMessageService.
var (
	einride_messenger_v1beta1_SMSMessageServiceClient v1beta1.SMSMessageServiceClient
	einride_messenger_v1beta1_SMSMessageService       = &cobra.Command{
		Use: "einride.messenger.v1beta1.SMSMessageService",
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
			log.Println("einride.messenger.v1beta1.SMSMessageService.GetSMSMessage")
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
			log.Println("einride.messenger.v1beta1.SMSMessageService.CreateSMSMessage")
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
			log.Println("einride.messenger.v1beta1.SMSMessageService.ListSMSMessages")
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
			log.Println("einride.messenger.v1beta1.SMSMessageService.SendSMSMessage")
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
			log.Println("einride.messenger.v1beta1.SMSMessageService.SendNewSMSMessage")
			return nil
		},
	}
)

func AddSMSMessageServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_messenger_v1beta1_SMSMessageService)
}

func init() {
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_GetSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_ListSMSMessages)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendSMSMessage)
	einride_messenger_v1beta1_SMSMessageService.AddCommand(einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage)
}
