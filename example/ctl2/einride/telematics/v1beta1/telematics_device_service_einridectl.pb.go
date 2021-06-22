package telematicsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_telematics_v1beta1_TelematicsDeviceService = &cobra.Command{
	Use: "einride.telematics.v1beta1.TelematicsDeviceService",
}

var einride_telematics_v1beta1_GetTelematicsDeviceRequest v1beta1.GetTelematicsDeviceRequest
var einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice = &cobra.Command{
	Use: "GetTelematicsDevice",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsDeviceService.GetTelematicsDevice")
		return nil
	},
}

var einride_telematics_v1beta1_ListTelematicsDevicesRequest v1beta1.ListTelematicsDevicesRequest
var einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices = &cobra.Command{
	Use: "ListTelematicsDevices",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsDeviceService.ListTelematicsDevices")
		return nil
	},
}

var einride_telematics_v1beta1_CreateTelematicsDeviceRequest v1beta1.CreateTelematicsDeviceRequest
var einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice = &cobra.Command{
	Use: "CreateTelematicsDevice",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsDeviceService.CreateTelematicsDevice")
		return nil
	},
}

func AddTelematicsDeviceServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService)
}

func init() {
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice)
}
