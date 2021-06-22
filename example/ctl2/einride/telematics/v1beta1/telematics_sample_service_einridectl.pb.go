package telematicsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_telematics_v1beta1_TelematicsSampleService = &cobra.Command{
	Use: "einride.telematics.v1beta1.TelematicsSampleService",
}

var einride_telematics_v1beta1_GetTelematicsSampleRequest v1beta1.GetTelematicsSampleRequest
var einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample = &cobra.Command{
	Use: "GetTelematicsSample",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsSampleService.GetTelematicsSample")
		return nil
	},
}

var einride_telematics_v1beta1_ListTelematicsSamplesRequest v1beta1.ListTelematicsSamplesRequest
var einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples = &cobra.Command{
	Use: "ListTelematicsSamples",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsSampleService.ListTelematicsSamples")
		return nil
	},
}

var einride_telematics_v1beta1_SearchTelematicsSamplesRequest v1beta1.SearchTelematicsSamplesRequest
var einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples = &cobra.Command{
	Use: "SearchTelematicsSamples",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsSampleService.SearchTelematicsSamples")
		return nil
	},
}

var einride_telematics_v1beta1_CreateTelematicsSampleRequest v1beta1.CreateTelematicsSampleRequest
var einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample = &cobra.Command{
	Use: "CreateTelematicsSample",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.telematics.v1beta1.TelematicsSampleService.CreateTelematicsSample")
		return nil
	},
}

func AddTelematicsSampleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsSampleService)
}

func init() {
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples)
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample)
}
