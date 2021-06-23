package telematicsv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.telematics.v1beta1.TelematicsDeviceService.
var (
	einride_telematics_v1beta1_TelematicsDeviceServiceClient v1beta1.TelematicsDeviceServiceClient
	einride_telematics_v1beta1_TelematicsDeviceService       = &cobra.Command{
		Use: "einride.telematics.v1beta1.TelematicsDeviceService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_telematics_v1beta1_TelematicsDeviceServiceClient = v1beta1.NewTelematicsDeviceServiceClient(conn)
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsDeviceService.GetTelematicsDevice.
var (
	einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice_Request v1beta1.GetTelematicsDeviceRequest
	einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice         = &cobra.Command{
		Use: "GetTelematicsDevice",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_telematics_v1beta1_TelematicsDeviceServiceClient.GetTelematicsDevice(cmd.Context(), &einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsDeviceService.ListTelematicsDevices.
var (
	einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices_Request v1beta1.ListTelematicsDevicesRequest
	einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices         = &cobra.Command{
		Use: "ListTelematicsDevices",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_telematics_v1beta1_TelematicsDeviceServiceClient.ListTelematicsDevices(cmd.Context(), &einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsDeviceService.CreateTelematicsDevice.
var (
	einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice_Request v1beta1.CreateTelematicsDeviceRequest
	einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice         = &cobra.Command{
		Use: "CreateTelematicsDevice",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_telematics_v1beta1_TelematicsDeviceServiceClient.CreateTelematicsDevice(cmd.Context(), &einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddTelematicsDeviceServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService)
}

func init() {
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice)

	einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice_Request.Name, "name", "", "Resource name of the telematics device to fetch")
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices)

	einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices_Request.PageSize, "pageSize", 10, "The maximum number of results to return.")

	einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices.Flags().StringVar(&einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices_Request.PageToken, "pageToken", "", "A page token, received from a previous List call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice)

	einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice_Request.TelematicsDevice = new(v1beta1.TelematicsDevice)
	einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice_Request.TelematicsDevice.Name, "telematicsDevice.name", "", "The resource name of the device.")
	// TODO: enum Category
	einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice.Flags().StringVar(&einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice_Request.TelematicsDevice.Id, "telematicsDevice.id", "", "ID is a generic identifier for various types of devices.\nFor Aplicom devices it will be the IMEI of the device. For\nBalena devices it will be a UUID. The ID must be unique, that is the only requirement.\nIMEI (International Mobile Equipment Identity) code.\nExample: \"352099001761481\".\nBalena device UUID.\nFor example: \"123e4567-e89b-12d3-a456-426655440000\"")
}
