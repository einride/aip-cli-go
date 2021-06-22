package integrationv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.integration.v1beta1.IntegrationFileLogService.
var (
	einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient v1beta1.IntegrationFileLogServiceClient
	einride_shipper_integration_v1beta1_IntegrationFileLogService       = &cobra.Command{
		Use: "einride.shipper.integration.v1beta1.IntegrationFileLogService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient = v1beta1.NewIntegrationFileLogServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileLogService.GetIntegrationFileLogEntry.
var (
	einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry_Request v1beta1.GetIntegrationFileLogEntryRequest
	einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry         = &cobra.Command{
		Use: "GetIntegrationFileLogEntry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.GetIntegrationFileLogEntry")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileLogService.ListIntegrationFileLogEntries.
var (
	einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries_Request v1beta1.ListIntegrationFileLogEntriesRequest
	einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries         = &cobra.Command{
		Use: "ListIntegrationFileLogEntries",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.ListIntegrationFileLogEntries")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileLogService.CreateIntegrationFileLogEntry.
var (
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request v1beta1.CreateIntegrationFileLogEntryRequest
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry         = &cobra.Command{
		Use: "CreateIntegrationFileLogEntry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.CreateIntegrationFileLogEntry")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileLogService.BatchCreateIntegrationFileLogEntries.
var (
	einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries_Request v1beta1.BatchCreateIntegrationFileLogEntriesRequest
	einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries         = &cobra.Command{
		Use: "BatchCreateIntegrationFileLogEntries",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileLogService.BatchCreateIntegrationFileLogEntries")
			return nil
		},
	}
)

func AddIntegrationFileLogServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry)
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries)
}
