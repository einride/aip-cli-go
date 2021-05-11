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

var einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice = &cobra.Command{
	Use: "GetTelematicsDevice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetTelematicsDevice called")
	},
}

var einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices = &cobra.Command{
	Use: "ListTelematicsDevices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListTelematicsDevices called")
	},
}

var einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice = &cobra.Command{
	Use: "CreateTelematicsDevice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateTelematicsDevice called")
	},
}

var einride_telematics_v1beta1_TelematicsSampleService = &cobra.Command{
	Use: "einride.telematics.v1beta1.TelematicsSample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.telematics.v1beta1.TelematicsSample called")
	},
}

var einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample = &cobra.Command{
	Use: "GetTelematicsSample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetTelematicsSample called")
	},
}

var einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples = &cobra.Command{
	Use: "ListTelematicsSamples",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListTelematicsSamples called")
	},
}

var einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples = &cobra.Command{
	Use: "SearchTelematicsSamples",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SearchTelematicsSamples called")
	},
}

var einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample = &cobra.Command{
	Use: "CreateTelematicsSample",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateTelematicsSample called")
	},
}

var einride_telematics_v1beta1_VehicleConnectionService = &cobra.Command{
	Use: "einride.telematics.v1beta1.VehicleConnection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.telematics.v1beta1.VehicleConnection called")
	},
}

var einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections = &cobra.Command{
	Use: "ListVehicleConnections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListVehicleConnections called")
	},
}

var einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice = &cobra.Command{
	Use: "ConnectTelematicsDevice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ConnectTelematicsDevice called")
	},
}

var einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice = &cobra.Command{
	Use: "DisconnectTelematicsDevice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DisconnectTelematicsDevice called")
	},
}

func init() {
	rootCmd.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice)
	rootCmd.AddCommand(einride_telematics_v1beta1_TelematicsSampleService)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample)
	rootCmd.AddCommand(einride_telematics_v1beta1_VehicleConnectionService)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice)
}
