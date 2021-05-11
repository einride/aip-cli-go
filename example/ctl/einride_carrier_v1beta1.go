package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_carrier_v1beta1_CarrierService = &cobra.Command{
	Use: "einride.carrier.v1beta1.Carrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.Carrier called")
	},
}
var einride_carrier_v1beta1_VehicleService = &cobra.Command{
	Use: "einride.carrier.v1beta1.Vehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.Vehicle called")
	},
}
var einride_carrier_v1beta1_VehicleTypeService = &cobra.Command{
	Use: "einride.carrier.v1beta1.VehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.VehicleType called")
	},
}

func init() {
	rootCmd.AddCommand(einride_carrier_v1beta1_CarrierService)
	rootCmd.AddCommand(einride_carrier_v1beta1_VehicleService)
	rootCmd.AddCommand(einride_carrier_v1beta1_VehicleTypeService)
}
