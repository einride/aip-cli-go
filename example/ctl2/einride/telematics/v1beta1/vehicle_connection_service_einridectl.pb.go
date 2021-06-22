package telematicsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.telematics.v1beta1.VehicleConnectionService.
var (
	einride_telematics_v1beta1_VehicleConnectionServiceClient v1beta1.VehicleConnectionServiceClient
	einride_telematics_v1beta1_VehicleConnectionService       = &cobra.Command{
		Use: "einride.telematics.v1beta1.VehicleConnectionService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_telematics_v1beta1_VehicleConnectionServiceClient = v1beta1.NewVehicleConnectionServiceClient(conn)
			return nil
		},
	}
)

// einride.telematics.v1beta1.VehicleConnectionService.ListVehicleConnections.
var (
	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request v1beta1.ListVehicleConnectionsRequest
	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections         = &cobra.Command{
		Use: "ListVehicleConnections",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.VehicleConnectionService.ListVehicleConnections")
			return nil
		},
	}
)

// einride.telematics.v1beta1.VehicleConnectionService.ConnectTelematicsDevice.
var (
	einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice_Request v1beta1.ConnectTelematicsDeviceRequest
	einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice         = &cobra.Command{
		Use: "ConnectTelematicsDevice",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.VehicleConnectionService.ConnectTelematicsDevice")
			return nil
		},
	}
)

// einride.telematics.v1beta1.VehicleConnectionService.DisconnectTelematicsDevice.
var (
	einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice_Request v1beta1.DisconnectTelematicsDeviceRequest
	einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice         = &cobra.Command{
		Use: "DisconnectTelematicsDevice",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.VehicleConnectionService.DisconnectTelematicsDevice")
			return nil
		},
	}
)

func AddVehicleConnectionServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_VehicleConnectionService)
}

func init() {
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice)
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice)
}
