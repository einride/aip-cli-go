package telematicsv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.telematics.v1beta1.VehicleConnectionService.
var (
	einride_telematics_v1beta1_VehicleConnectionServiceClient v1beta1.VehicleConnectionServiceClient
	einride_telematics_v1beta1_VehicleConnectionService       = &cobra.Command{
		Use:   "telematics.v1beta1.VehicleConnectionService",
		Short: "Vehicle connection service.",
		Long:  "Vehicle connection service.\n\nConnections between telematics devices, vehicles and shipments ensure that\ntelematics samples can be stored and queried by device, vehicle and\nshipment.",
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
			response, err := einride_telematics_v1beta1_VehicleConnectionServiceClient.ListVehicleConnections(cmd.Context(), &einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_telematics_v1beta1_VehicleConnectionServiceClient.ConnectTelematicsDevice(cmd.Context(), &einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_telematics_v1beta1_VehicleConnectionServiceClient.DisconnectTelematicsDevice(cmd.Context(), &einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddVehicleConnectionServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_VehicleConnectionService)
}

func init() {
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections)

	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request.Parent, "parent", "", "Resource name of the parent vehicle.")

	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections.Flags().Int32Var(&einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request.PageSize, "pageSize", 10, "The maximum number of results to return.")

	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request.PageToken, "pageToken", "", "A page token, received from a previous List call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_ListVehicleConnections_Request.Filter, "filter", "", "A filter.\n\nCurrently the only supported format is:\n\n`telematics_device = \"<string>\"`")
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice)

	einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice_Request.Vehicle, "vehicle", "", "Resource name of the vehicle.")

	einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_ConnectTelematicsDevice_Request.TelematicsDevice, "telematicsDevice", "", "Resource name of the telematics device to connect.")
	einride_telematics_v1beta1_VehicleConnectionService.AddCommand(einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice)

	einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice_Request.Vehicle, "vehicle", "", "Resource name of the vehicle.")

	einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_VehicleConnectionService_DisconnectTelematicsDevice_Request.TelematicsDevice, "telematicsDevice", "", "Resource name of the telematics device to disconnect.")
}
