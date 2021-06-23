package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalServiceClient v1beta1.ShipmentChangeProposalServiceClient
	einride_shipper_v1beta1_ShipmentChangeProposalService       = &cobra.Command{
		Use:   "einride.shipper.v1beta1.ShipmentChangeProposalService",
		Short: "Shipment proposal service.",
		Long:  "Shipment proposal service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentChangeProposalServiceClient = v1beta1.NewShipmentChangeProposalServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.CreateShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request v1beta1.CreateShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal         = &cobra.Command{
		Use: "CreateShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.CreateShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.GetShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal_Request v1beta1.GetShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal         = &cobra.Command{
		Use: "GetShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.GetShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.ListShipmentChangeProposals.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request v1beta1.ListShipmentChangeProposalsRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals         = &cobra.Command{
		Use: "ListShipmentChangeProposals",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.ListShipmentChangeProposals(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.UpdateShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request v1beta1.UpdateShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal         = &cobra.Command{
		Use: "UpdateShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.UpdateShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.DeleteShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal_Request v1beta1.DeleteShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal         = &cobra.Command{
		Use: "DeleteShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.DeleteShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.ApproveShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal_Request v1beta1.ApproveShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal         = &cobra.Command{
		Use: "ApproveShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.ApproveShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.RejectShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal_Request v1beta1.RejectShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal         = &cobra.Command{
		Use: "RejectShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentChangeProposalServiceClient.RejectShipmentChangeProposal(cmd.Context(), &einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipmentChangeProposalServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService)
}

func init() {
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request.Parent, "parent", "", "Resource name of the parent shipment where this shipment change proposal will be created.\nPattern: `shippers/{shipper}/shipments/{shipment}`.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request.ShipmentChangeProposal = new(v1beta1.ShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request.ShipmentChangeProposal.Name, "shipmentChangeProposal.name", "", "The resource name of the shipment proposal.")
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request.ShipmentChangeProposal.Etag, "shipmentChangeProposal.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: list: ShipmentChanges message

	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request.ShipmentChangeProposalId, "shipmentChangeProposalId", "", "The ID to use for the shipment change proposal, which will become the final component of\nthe shipment change proposal's resource name.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-z][0-9]-/.")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal_Request.Name, "name", "", "Resource name of the shipment change proposal to retrieve.\nFormat: `shippers/{shipper}/shipments/{shipment}/changeProposals/{change_proposal}`.")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals)

	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request.Parent, "parent", "", "Resource name of the parent shipment.\nPattern: `shippers/{shipper}/shipments/{shipment}`.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals.Flags().BoolVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request.ShowDeleted, "showDeleted", false, "Flag indicating whether to include deleted results.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request.Filter, "filter", "", "Filter the shipment change proposals.\nIf not provided, no filtering is applied.\n\nCurrently, the only supported format is:\n`user = \"<user resource name>\"`")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request.ShipmentChangeProposal = new(v1beta1.ShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request.ShipmentChangeProposal.Name, "shipmentChangeProposal.name", "", "The resource name of the shipment proposal.")
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request.ShipmentChangeProposal.Etag, "shipmentChangeProposal.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: list: ShipmentChanges message

	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal_Request.Name, "name", "", "Resource name of the shipment change proposal to delete.\nFormat: `shippers/{shipper}/shipments/{shipment}/changeProposals/{change_proposal}`.")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal_Request.Name, "name", "", "Resource name of the shipment change proposal to approve.\nFormat: `shippers/{shipper}/shipments/{shipment}/changeProposals/{change_proposal}`.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal_Request.Note, "note", "", "The approval note.")
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal)

	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal_Request.Name, "name", "", "Resource name of the shipment change proposal to reject.\nFormat: `shippers/{shipper}/shipments/{shipment}/changeProposals/{change_proposal}`.")

	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal.Flags().StringVar(&einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal_Request.Note, "note", "", "The rejection note.")
}
