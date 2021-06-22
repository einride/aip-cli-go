package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipmentAttachmentService.
var (
	einride_shipper_v1beta1_ShipmentAttachmentServiceClient v1beta1.ShipmentAttachmentServiceClient
	einride_shipper_v1beta1_ShipmentAttachmentService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipmentAttachmentService",
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
			log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.CreateShipmentAttachment")
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
			log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.GetShipmentAttachment")
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
			log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.ListShipmentAttachments")
			return nil
		},
	}
)

func AddShipmentAttachmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment)
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment)
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments)
}
