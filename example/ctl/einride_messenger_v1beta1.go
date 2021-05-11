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
var einride_messenger_v1beta1_SMSMessageService = &cobra.Command{
	Use: "einride.messenger.v1beta1.SMSMessage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.messenger.v1beta1.SMSMessage called")
	},
}

func init() {
	rootCmd.AddCommand(einride_messenger_v1beta1_EmailService)
	rootCmd.AddCommand(einride_messenger_v1beta1_SMSMessageService)
}
