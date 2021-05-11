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

func init() {
	rootCmd.AddCommand(einride_cem_v1beta1_CEMService)
}
