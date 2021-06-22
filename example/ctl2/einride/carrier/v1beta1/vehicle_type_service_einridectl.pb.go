package carrierv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_carrier_v1beta1_VehicleTypeService = &cobra.Command{
	Use: "einride.carrier.v1beta1.VehicleTypeService",
}

var einride_carrier_v1beta1_CreateVehicleTypeRequest v1beta1.CreateVehicleTypeRequest
var einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType = &cobra.Command{
	Use: "CreateVehicleType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.CreateVehicleType")
		return nil
	},
}

var einride_carrier_v1beta1_GetVehicleTypeRequest v1beta1.GetVehicleTypeRequest
var einride_carrier_v1beta1_VehicleTypeService_GetVehicleType = &cobra.Command{
	Use: "GetVehicleType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.GetVehicleType")
		return nil
	},
}

var einride_carrier_v1beta1_BatchGetVehicleTypesRequest v1beta1.BatchGetVehicleTypesRequest
var einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes = &cobra.Command{
	Use: "BatchGetVehicleTypes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.BatchGetVehicleTypes")
		return nil
	},
}

var einride_carrier_v1beta1_UpdateVehicleTypeRequest v1beta1.UpdateVehicleTypeRequest
var einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType = &cobra.Command{
	Use: "UpdateVehicleType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.UpdateVehicleType")
		return nil
	},
}

var einride_carrier_v1beta1_ListVehicleTypesRequest v1beta1.ListVehicleTypesRequest
var einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes = &cobra.Command{
	Use: "ListVehicleTypes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.ListVehicleTypes")
		return nil
	},
}

var einride_carrier_v1beta1_DeleteVehicleTypeRequest v1beta1.DeleteVehicleTypeRequest
var einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType = &cobra.Command{
	Use: "DeleteVehicleType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.DeleteVehicleType")
		return nil
	},
}

var einride_carrier_v1beta1_UndeleteVehicleTypeRequest v1beta1.UndeleteVehicleTypeRequest
var einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType = &cobra.Command{
	Use: "UndeleteVehicleType",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleTypeService.UndeleteVehicleType")
		return nil
	},
}

func AddVehicleTypeServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_VehicleTypeService)
}

func init() {
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_GetVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType)
}
