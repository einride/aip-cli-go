package telematicsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
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
			log.Println("einride.telematics.v1beta1.TelematicsDeviceService.GetTelematicsDevice")
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
			log.Println("einride.telematics.v1beta1.TelematicsDeviceService.ListTelematicsDevices")
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
			log.Println("einride.telematics.v1beta1.TelematicsDeviceService.CreateTelematicsDevice")
			return nil
		},
	}
)

func AddTelematicsDeviceServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService)
}

func init() {
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_GetTelematicsDevice)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_ListTelematicsDevices)
	einride_telematics_v1beta1_TelematicsDeviceService.AddCommand(einride_telematics_v1beta1_TelematicsDeviceService_CreateTelematicsDevice)
}
