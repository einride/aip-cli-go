package integrationv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_integration_v1beta1_IntegrationFileLogService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationFileLogService",
}

var einride_shipper_integration_v1beta1_GetIntegrationFileLogEntryRequest v1beta1.GetIntegrationFileLogEntryRequest
var einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry = &cobra.Command{
	Use: "GetIntegrationFileLogEntry",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.GetIntegrationFileLogEntry")
		return nil
	},
}

var einride_shipper_integration_v1beta1_ListIntegrationFileLogEntriesRequest v1beta1.ListIntegrationFileLogEntriesRequest
var einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries = &cobra.Command{
	Use: "ListIntegrationFileLogEntries",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.ListIntegrationFileLogEntries")
		return nil
	},
}

var einride_shipper_integration_v1beta1_CreateIntegrationFileLogEntryRequest v1beta1.CreateIntegrationFileLogEntryRequest
var einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry = &cobra.Command{
	Use: "CreateIntegrationFileLogEntry",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.CreateIntegrationFileLogEntry")
		return nil
	},
}

var einride_shipper_integration_v1beta1_BatchCreateIntegrationFileLogEntriesRequest v1beta1.BatchCreateIntegrationFileLogEntriesRequest
var einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries = &cobra.Command{
	Use: "BatchCreateIntegrationFileLogEntries",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.BatchCreateIntegrationFileLogEntries")
		return nil
	},
}

func AddIntegrationFileLogServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries)
}
