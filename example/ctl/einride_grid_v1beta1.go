package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_grid_v1beta1_ChargerService = &cobra.Command{
	Use: "einride.grid.v1beta1.Charger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.grid.v1beta1.Charger called")
	},
}

func init() {
	rootCmd.AddCommand(einride_grid_v1beta1_ChargerService)
}
