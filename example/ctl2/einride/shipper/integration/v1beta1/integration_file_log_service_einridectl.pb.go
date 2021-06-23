package integrationv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient.GetIntegrationFileLogEntry(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient.ListIntegrationFileLogEntries(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient.CreateIntegrationFileLogEntry(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileLogServiceClient.BatchCreateIntegrationFileLogEntries(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddIntegrationFileLogServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry)

	einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_GetIntegrationFileLogEntry_Request.Name, "name", "", "Resource name of the integration file log entry to retrieve.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries)

	einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries_Request.Parent, "parent", "", "Resource name of the parent integration file.")

	einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries_Request.PageSize, "pageSize", 10, "The maximum number of results to return.")

	einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_ListIntegrationFileLogEntries_Request.PageToken, "pageToken", "", "A page token, received from a previous List call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry)

	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.Parent, "parent", "", "Resource name of the parent integration file where this integration file log entry will be created.")

	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry = new(v1beta1.IntegrationFileLogEntry)
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.Name, "integrationFileLogEntry.name", "", "The resource name of the ingestion log.")
	// TODO: enum Severity
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.Text, "integrationFileLogEntry.text", "", "Free-text message in log entry.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.ShipmentReferenceId, "integrationFileLogEntry.shipmentReferenceId", "", "The shipment reference ID, if the log relates to a specific shipment.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.LineItemReferenceId, "integrationFileLogEntry.lineItemReferenceId", "", "The line item reference ID, if the log relates to a specific line item.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.SiteReferenceId, "integrationFileLogEntry.siteReferenceId", "", "The site reference ID, if the log relates to a specific site.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().Int64Var(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.OffsetStart, "integrationFileLogEntry.offsetStart", 10, "The related byte offset start in the file (inclusive).")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().Int64Var(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.OffsetEnd, "integrationFileLogEntry.offsetEnd", 10, "The related byte offset end in the file (exclusive).")
	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntry.FieldPath, "integrationFileLogEntry.fieldPath", "", "The related field path.\nFor example: \"shipment.pickup_earliest_time\"")

	einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_CreateIntegrationFileLogEntry_Request.IntegrationFileLogEntryId, "integrationFileLogEntryId", "", "The ID to use for the integration file log entry, which will become the final component of\nthe integration file log entry's resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be a valid unix nanosecond timestamp.")
	einride_shipper_integration_v1beta1_IntegrationFileLogService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries)

	einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileLogService_BatchCreateIntegrationFileLogEntries_Request.Parent, "parent", "", "Resource name of the parent integration file shared by all resources being created.\n\nIf this is set, the parent field in the CreateIntegrationFileLogEntry\nmessages must either be empty or match this field.")

	// TODO: list: Requests message
}
