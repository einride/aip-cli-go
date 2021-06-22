package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipmentAttachmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentAttachmentService",
}

var einride_shipper_v1beta1_CreateShipmentAttachmentRequest v1beta1.CreateShipmentAttachmentRequest
var einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment = &cobra.Command{
	Use: "CreateShipmentAttachment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.CreateShipmentAttachment")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipmentAttachmentRequest v1beta1.GetShipmentAttachmentRequest
var einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment = &cobra.Command{
	Use: "GetShipmentAttachment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.GetShipmentAttachment")
		return nil
	},
}

var einride_shipper_v1beta1_ListShipmentAttachmentsRequest v1beta1.ListShipmentAttachmentsRequest
var einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments = &cobra.Command{
	Use: "ListShipmentAttachments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentAttachmentService.ListShipmentAttachments")
		return nil
	},
}

func AddShipmentAttachmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_CreateShipmentAttachment)
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_GetShipmentAttachment)
	einride_shipper_v1beta1_ShipmentAttachmentService.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService_ListShipmentAttachments)
}
