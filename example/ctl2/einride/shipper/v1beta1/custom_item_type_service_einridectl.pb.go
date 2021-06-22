package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_CustomItemTypeService = &cobra.Command{
	Use: "einride.shipper.v1beta1.CustomItemTypeService",
}

var einride_shipper_v1beta1_GetCustomItemTypeRequest v1beta1.GetCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType = &cobra.Command{
	Use: "GetCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.GetCustomItemType")
		return nil
	},
}

var einride_shipper_v1beta1_CreateCustomItemTypeRequest v1beta1.CreateCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType = &cobra.Command{
	Use: "CreateCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.CreateCustomItemType")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateCustomItemTypeRequest v1beta1.UpdateCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType = &cobra.Command{
	Use: "UpdateCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.UpdateCustomItemType")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteCustomItemTypeRequest v1beta1.DeleteCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType = &cobra.Command{
	Use: "DeleteCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.DeleteCustomItemType")
		return nil
	},
}

var einride_shipper_v1beta1_UndeleteCustomItemTypeRequest v1beta1.UndeleteCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType = &cobra.Command{
	Use: "UndeleteCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.UndeleteCustomItemType")
		return nil
	},
}

var einride_shipper_v1beta1_ListCustomItemTypesRequest v1beta1.ListCustomItemTypesRequest
var einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes = &cobra.Command{
	Use: "ListCustomItemTypes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.ListCustomItemTypes")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetCustomItemTypesRequest v1beta1.BatchGetCustomItemTypesRequest
var einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes = &cobra.Command{
	Use: "BatchGetCustomItemTypes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.BatchGetCustomItemTypes")
		return nil
	},
}

var einride_shipper_v1beta1_ReferenceCustomItemTypeRequest v1beta1.ReferenceCustomItemTypeRequest
var einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType = &cobra.Command{
	Use: "ReferenceCustomItemType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.CustomItemTypeService.ReferenceCustomItemType")
		return nil
	},
}

func AddCustomItemTypeServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_CustomItemTypeService)
}

func init() {
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes)
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType)
}
