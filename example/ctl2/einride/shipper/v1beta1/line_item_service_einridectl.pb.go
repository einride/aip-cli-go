package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_LineItemService = &cobra.Command{
	Use: "einride.shipper.v1beta1.LineItemService",
}

var einride_shipper_v1beta1_CreateLineItemRequest v1beta1.CreateLineItemRequest
var einride_shipper_v1beta1_LineItemService_CreateLineItem = &cobra.Command{
	Use: "CreateLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.CreateLineItem")
		return nil
	},
}

var einride_shipper_v1beta1_GetLineItemRequest v1beta1.GetLineItemRequest
var einride_shipper_v1beta1_LineItemService_GetLineItem = &cobra.Command{
	Use: "GetLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.GetLineItem")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetLineItemsRequest v1beta1.BatchGetLineItemsRequest
var einride_shipper_v1beta1_LineItemService_BatchGetLineItems = &cobra.Command{
	Use: "BatchGetLineItems",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.BatchGetLineItems")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateLineItemRequest v1beta1.UpdateLineItemRequest
var einride_shipper_v1beta1_LineItemService_UpdateLineItem = &cobra.Command{
	Use: "UpdateLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.UpdateLineItem")
		return nil
	},
}

var einride_shipper_v1beta1_ReferenceLineItemRequest v1beta1.ReferenceLineItemRequest
var einride_shipper_v1beta1_LineItemService_ReferenceLineItem = &cobra.Command{
	Use: "ReferenceLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.ReferenceLineItem")
		return nil
	},
}

var einride_shipper_v1beta1_ListLineItemsRequest v1beta1.ListLineItemsRequest
var einride_shipper_v1beta1_LineItemService_ListLineItems = &cobra.Command{
	Use: "ListLineItems",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.ListLineItems")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteLineItemRequest v1beta1.DeleteLineItemRequest
var einride_shipper_v1beta1_LineItemService_DeleteLineItem = &cobra.Command{
	Use: "DeleteLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.DeleteLineItem")
		return nil
	},
}

var einride_shipper_v1beta1_UndeleteLineItemRequest v1beta1.UndeleteLineItemRequest
var einride_shipper_v1beta1_LineItemService_UndeleteLineItem = &cobra.Command{
	Use: "UndeleteLineItem",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LineItemService.UndeleteLineItem")
		return nil
	},
}

func AddLineItemServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_LineItemService)
}

func init() {
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_CreateLineItem)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_GetLineItem)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_BatchGetLineItems)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_UpdateLineItem)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_ReferenceLineItem)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_ListLineItems)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_DeleteLineItem)
	einride_shipper_v1beta1_LineItemService.AddCommand(einride_shipper_v1beta1_LineItemService_UndeleteLineItem)
}
