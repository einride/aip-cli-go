package messengerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/messenger/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_messenger_v1beta1_SMSMessageService = &cobra.Command{
	Use: "einride.messenger.v1beta1.SMSMessageService",
}

var einride_messenger_v1beta1_GetSMSMessageRequest v1beta1.GetSMSMessageRequest
var einride_messenger_v1beta1_SMSMessageService_GetSMSMessage = &cobra.Command{
	Use: "GetSMSMessage",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.SMSMessageService.GetSMSMessage")
		return nil
	},
}

var einride_messenger_v1beta1_CreateSMSMessageRequest v1beta1.CreateSMSMessageRequest
var einride_messenger_v1beta1_SMSMessageService_CreateSMSMessage = &cobra.Command{
	Use: "CreateSMSMessage",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.SMSMessageService.CreateSMSMessage")
		return nil
	},
}

var einride_messenger_v1beta1_ListSMSMessagesRequest v1beta1.ListSMSMessagesRequest
var einride_messenger_v1beta1_SMSMessageService_ListSMSMessages = &cobra.Command{
	Use: "ListSMSMessages",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.SMSMessageService.ListSMSMessages")
		return nil
	},
}

var einride_messenger_v1beta1_SendSMSMessageRequest v1beta1.SendSMSMessageRequest
var einride_messenger_v1beta1_SMSMessageService_SendSMSMessage = &cobra.Command{
	Use: "SendSMSMessage",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.SMSMessageService.SendSMSMessage")
		return nil
	},
}

var einride_messenger_v1beta1_SendNewSMSMessageRequest v1beta1.SendNewSMSMessageRequest
var einride_messenger_v1beta1_SMSMessageService_SendNewSMSMessage = &cobra.Command{
	Use: "SendNewSMSMessage",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.messenger.v1beta1.SMSMessageService.SendNewSMSMessage")
		return nil
	},
}

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
