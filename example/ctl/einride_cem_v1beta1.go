package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_cem_v1beta1_CEMService = &cobra.Command{
	Use: "einride.cem.v1beta1.CEM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.cem.v1beta1.CEM called")
	},
}

var einride_cem_v1beta1_CEMService_CalculateResult = &cobra.Command{
	Use: "CalculateResult",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CalculateResult called")
	},
}

var einride_cem_v1beta1_CEMService_BatchCalculateResults = &cobra.Command{
	Use: "BatchCalculateResults",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchCalculateResults called")
	},
}

func init() {
	rootCmd.AddCommand(einride_cem_v1beta1_CEMService)
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_CalculateResult)
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_BatchCalculateResults)
}
