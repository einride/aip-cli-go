package telematicsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.telematics.v1beta1.TelematicsSampleService.
var (
	einride_telematics_v1beta1_TelematicsSampleServiceClient v1beta1.TelematicsSampleServiceClient
	einride_telematics_v1beta1_TelematicsSampleService       = &cobra.Command{
		Use: "einride.telematics.v1beta1.TelematicsSampleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_telematics_v1beta1_TelematicsSampleServiceClient = v1beta1.NewTelematicsSampleServiceClient(conn)
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.GetTelematicsSample.
var (
	einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample_Request v1beta1.GetTelematicsSampleRequest
	einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample         = &cobra.Command{
		Use: "GetTelematicsSample",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.GetTelematicsSample")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.ListTelematicsSamples.
var (
	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples_Request v1beta1.ListTelematicsSamplesRequest
	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples         = &cobra.Command{
		Use: "ListTelematicsSamples",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.ListTelematicsSamples")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.SearchTelematicsSamples.
var (
	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples_Request v1beta1.SearchTelematicsSamplesRequest
	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples         = &cobra.Command{
		Use: "SearchTelematicsSamples",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.SearchTelematicsSamples")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.CreateTelematicsSample.
var (
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request v1beta1.CreateTelematicsSampleRequest
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample         = &cobra.Command{
		Use: "CreateTelematicsSample",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.CreateTelematicsSample")
			return nil
		},
	}
)

func AddTelematicsSampleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsSampleService)
}

func init() {
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample)
}
