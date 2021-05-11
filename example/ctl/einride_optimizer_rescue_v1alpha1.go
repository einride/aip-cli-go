package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_optimizer_rescue_v1alpha1_RescueService = &cobra.Command{
	Use: "einride.optimizer.rescue.v1alpha1.Rescue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.optimizer.rescue.v1alpha1.Rescue called")
	},
}

func init() {
	rootCmd.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService)
}
