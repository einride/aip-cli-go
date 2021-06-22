package telematicsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_telematics_v1beta1_VehicleConnectionService = &cobra.Command{
	Use: "einride.telematics.v1beta1.VehicleConnectionService",
}

var einride_telematics_v1beta1_ListVehicleConnectionsRequest v1beta1.ListVehicleConnectionsRequest
var einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections = &cobra.Command{
	Use: "ListVehicleConnections",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.VehicleConnectionService.ListVehicleConnections")
		return nil
	},
}

var einride_telematics_v1beta1_ConnectTelematicsDeviceRequest v1beta1.ConnectTelematicsDeviceRequest
var einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice = &cobra.Command{
	Use: "ConnectTelematicsDevice",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.VehicleConnectionService.ConnectTelematicsDevice")
		return nil
	},
}

var einride_telematics_v1beta1_DisconnectTelematicsDeviceRequest v1beta1.DisconnectTelematicsDeviceRequest
var einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice = &cobra.Command{
	Use: "DisconnectTelematicsDevice",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.VehicleConnectionService.DisconnectTelematicsDevice")
		return nil
	},
}

func AddVehicleConnectionServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_VehicleConnectionService)
}

func init() {
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice)
}
