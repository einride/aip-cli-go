package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.shipper.v1beta1.ShipmentAttachmentService.
var (
	einride_shipper_v1beta1_ShipmentAttachmentServiceClient v1beta1.ShipmentAttachmentServiceClient
	einride_shipper_v1beta1_ShipmentAttachmentService       = &cobra.Command{
		Use:   "shipper.v1beta1.ShipmentAttachmentService",
		Short: "Shipment attachment service.",
		Long:  "Shipment attachment service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentAttachmentServiceClient = v1beta1.NewShipmentAttachmentServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentAttachmentService.CreateShipmentAttachment.
var (
	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request v1beta1.CreateShipmentAttachmentRequest
	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment         = &cobra.Command{
		Use: "CreateShipmentAttachment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentAttachmentServiceClient.CreateShipmentAttachment(cmd.Context(), &einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentAttachmentService.GetShipmentAttachment.
var (
	einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment_Request v1beta1.GetShipmentAttachmentRequest
	einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment         = &cobra.Command{
		Use: "GetShipmentAttachment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentAttachmentServiceClient.GetShipmentAttachment(cmd.Context(), &einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentAttachmentService.ListShipmentAttachments.
var (
	einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments_Request v1beta1.ListShipmentAttachmentsRequest
	einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments         = &cobra.Command{
		Use: "ListShipmentAttachments",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentAttachmentServiceClient.ListShipmentAttachments(cmd.Context(), &einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipmentAttachmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment)

	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request.Parent, "parent", "", "Resource name of the parent shipment where this shipment attachment will be created.")

	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request.ShipmentAttachment = new(v1beta1.ShipmentAttachment)
	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request.ShipmentAttachment.Name, "shipmentAttachment.name", "", "The resource name of the shipment attachment.")
	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request.ShipmentAttachment.Description, "shipmentAttachment.description", "", "Free text description of the attachment.")
	einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment_Request.ShipmentAttachment.Files, "shipmentAttachment.files", []string{}, "Resource names of files attached to the attachment.")
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment)

	einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment_Request.Name, "name", "", "Resource name of the shipment attachment to retrieve.")
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments)

	einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments_Request.Parent, "parent", "", "Resource name of the parent shipment.")

	einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
