package integrationv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_integration_v1beta1_IntegrationLogService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationLogService",
}

var einride_shipper_integration_v1beta1_GetIntegrationLogEntryRequest v1beta1.GetIntegrationLogEntryRequest
var einride_shipper_integration_v1beta1_IntegrationLogService_GetIntegrationLogEntry = &cobra.Command{
	Use: "GetIntegrationLogEntry",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.GetIntegrationLogEntry")
		return nil
	},
}

var einride_shipper_integration_v1beta1_ListIntegrationLogEntriesRequest v1beta1.ListIntegrationLogEntriesRequest
var einride_shipper_integration_v1beta1_IntegrationLogService_ListIntegrationLogEntries = &cobra.Command{
	Use: "ListIntegrationLogEntries",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.ListIntegrationLogEntries")
		return nil
	},
}

var einride_shipper_integration_v1beta1_CreateIntegrationLogEntryRequest v1beta1.CreateIntegrationLogEntryRequest
var einride_shipper_integration_v1beta1_IntegrationLogService_CreateIntegrationLogEntry = &cobra.Command{
	Use: "CreateIntegrationLogEntry",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.CreateIntegrationLogEntry")
		return nil
	},
}

func AddIntegrationLogServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_GetIntegrationLogEntry)
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_ListIntegrationLogEntries)
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_CreateIntegrationLogEntry)
}
