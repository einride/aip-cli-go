package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_telematics_v1beta1_TelematicsDeviceService = &cobra.Command{
	Use: "einride.telematics.v1beta1.TelematicsDevice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.telematics.v1beta1.TelematicsDevice called")
	},
}
var einride_telematics_v1beta1_TelematicsSampleService = &cobra.Command{
	Use: "einride.telematics.v1beta1.TelematicsSample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.telematics.v1beta1.TelematicsSample called")
	},
}
var einride_telematics_v1beta1_VehicleConnectionService = &cobra.Command{
	Use: "einride.telematics.v1beta1.VehicleConnection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.telematics.v1beta1.VehicleConnection called")
	},
}

func init() {
	rootCmd.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService)
	rootCmd.AddCommand(einride_telematics_v1beta1_TelematicsSampleService)
	rootCmd.AddCommand(einride_telematics_v1beta1_VehicleConnectionService)
}
