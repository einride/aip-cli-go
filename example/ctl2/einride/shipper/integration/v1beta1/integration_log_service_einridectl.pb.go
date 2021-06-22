package integrationv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.integration.v1beta1.IntegrationLogService.
var (
	einride_shipper_integration_v1beta1_IntegrationLogServiceClient v1beta1.IntegrationLogServiceClient
	einride_shipper_integration_v1beta1_IntegrationLogService       = &cobra.Command{
		Use: "einride.shipper.integration.v1beta1.IntegrationLogService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_integration_v1beta1_IntegrationLogServiceClient = v1beta1.NewIntegrationLogServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationLogService.GetIntegrationLogEntry.
var (
	einride_shipper_integration_v1beta1_IntegrationLogService_GetIntegrationLogEntry_Request v1beta1.GetIntegrationLogEntryRequest
	einride_shipper_integration_v1beta1_IntegrationLogService_GetIntegrationLogEntry         = &cobra.Command{
		Use: "GetIntegrationLogEntry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.GetIntegrationLogEntry")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationLogService.ListIntegrationLogEntries.
var (
	einride_shipper_integration_v1beta1_IntegrationLogService_ListIntegrationLogEntries_Request v1beta1.ListIntegrationLogEntriesRequest
	einride_shipper_integration_v1beta1_IntegrationLogService_ListIntegrationLogEntries         = &cobra.Command{
		Use: "ListIntegrationLogEntries",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.ListIntegrationLogEntries")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationLogService.CreateIntegrationLogEntry.
var (
	einride_shipper_integration_v1beta1_IntegrationLogService_CreateIntegrationLogEntry_Request v1beta1.CreateIntegrationLogEntryRequest
	einride_shipper_integration_v1beta1_IntegrationLogService_CreateIntegrationLogEntry         = &cobra.Command{
		Use: "CreateIntegrationLogEntry",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationLogService.CreateIntegrationLogEntry")
			return nil
		},
	}
)

func AddIntegrationLogServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_GetIntegrationLogEntry)
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_ListIntegrationLogEntries)
	einride_shipper_integration_v1beta1_IntegrationLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService_CreateIntegrationLogEntry)
}
