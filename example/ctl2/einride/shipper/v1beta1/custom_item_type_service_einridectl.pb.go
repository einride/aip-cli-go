package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.CustomItemTypeService.
var (
	einride_shipper_v1beta1_CustomItemTypeServiceClient v1beta1.CustomItemTypeServiceClient
	einride_shipper_v1beta1_CustomItemTypeService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.CustomItemTypeService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_CustomItemTypeServiceClient = v1beta1.NewCustomItemTypeServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.GetCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType_Request v1beta1.GetCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType         = &cobra.Command{
		Use: "GetCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.GetCustomItemType")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.CreateCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request v1beta1.CreateCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType         = &cobra.Command{
		Use: "CreateCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.CreateCustomItemType")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.UpdateCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request v1beta1.UpdateCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType         = &cobra.Command{
		Use: "UpdateCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.UpdateCustomItemType")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.DeleteCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType_Request v1beta1.DeleteCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType         = &cobra.Command{
		Use: "DeleteCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.DeleteCustomItemType")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.UndeleteCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType_Request v1beta1.UndeleteCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType         = &cobra.Command{
		Use: "UndeleteCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.UndeleteCustomItemType")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.ListCustomItemTypes.
var (
	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request v1beta1.ListCustomItemTypesRequest
	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes         = &cobra.Command{
		Use: "ListCustomItemTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.ListCustomItemTypes")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.BatchGetCustomItemTypes.
var (
	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes_Request v1beta1.BatchGetCustomItemTypesRequest
	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes         = &cobra.Command{
		Use: "BatchGetCustomItemTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.BatchGetCustomItemTypes")
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.ReferenceCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType_Request v1beta1.ReferenceCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType         = &cobra.Command{
		Use: "ReferenceCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.CustomItemTypeService.ReferenceCustomItemType")
			return nil
		},
	}
)

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
