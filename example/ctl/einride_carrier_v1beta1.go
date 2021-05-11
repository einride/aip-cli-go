package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_carrier_v1beta1_CarrierService = &cobra.Command{
	Use: "einride.carrier.v1beta1.Carrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.Carrier called")
	},
}

var einride_carrier_v1beta1_CarrierService_CreateCarrier = &cobra.Command{
	Use: "CreateCarrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateCarrier called")
	},
}

var einride_carrier_v1beta1_CarrierService_ListCarriers = &cobra.Command{
	Use: "ListCarriers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListCarriers called")
	},
}

var einride_carrier_v1beta1_CarrierService_GetCarrier = &cobra.Command{
	Use: "GetCarrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetCarrier called")
	},
}

var einride_carrier_v1beta1_CarrierService_BatchGetCarriers = &cobra.Command{
	Use: "BatchGetCarriers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetCarriers called")
	},
}

var einride_carrier_v1beta1_CarrierService_UpdateCarrier = &cobra.Command{
	Use: "UpdateCarrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateCarrier called")
	},
}

var einride_carrier_v1beta1_CarrierService_DeleteCarrier = &cobra.Command{
	Use: "DeleteCarrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteCarrier called")
	},
}

var einride_carrier_v1beta1_CarrierService_UndeleteCarrier = &cobra.Command{
	Use: "UndeleteCarrier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UndeleteCarrier called")
	},
}

var einride_carrier_v1beta1_VehicleService = &cobra.Command{
	Use: "einride.carrier.v1beta1.Vehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.Vehicle called")
	},
}

var einride_carrier_v1beta1_VehicleService_CreateVehicle = &cobra.Command{
	Use: "CreateVehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateVehicle called")
	},
}

var einride_carrier_v1beta1_VehicleService_GetVehicle = &cobra.Command{
	Use: "GetVehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetVehicle called")
	},
}

var einride_carrier_v1beta1_VehicleService_BatchGetVehicles = &cobra.Command{
	Use: "BatchGetVehicles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetVehicles called")
	},
}

var einride_carrier_v1beta1_VehicleService_UpdateVehicle = &cobra.Command{
	Use: "UpdateVehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateVehicle called")
	},
}

var einride_carrier_v1beta1_VehicleService_ListVehicles = &cobra.Command{
	Use: "ListVehicles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListVehicles called")
	},
}

var einride_carrier_v1beta1_VehicleService_DeleteVehicle = &cobra.Command{
	Use: "DeleteVehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteVehicle called")
	},
}

var einride_carrier_v1beta1_VehicleService_UndeleteVehicle = &cobra.Command{
	Use: "UndeleteVehicle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UndeleteVehicle called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService = &cobra.Command{
	Use: "einride.carrier.v1beta1.VehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.v1beta1.VehicleType called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType = &cobra.Command{
	Use: "CreateVehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateVehicleType called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_GetVehicleType = &cobra.Command{
	Use: "GetVehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetVehicleType called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes = &cobra.Command{
	Use: "BatchGetVehicleTypes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetVehicleTypes called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType = &cobra.Command{
	Use: "UpdateVehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateVehicleType called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes = &cobra.Command{
	Use: "ListVehicleTypes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListVehicleTypes called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType = &cobra.Command{
	Use: "DeleteVehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteVehicleType called")
	},
}

var einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType = &cobra.Command{
	Use: "UndeleteVehicleType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UndeleteVehicleType called")
	},
}

func init() {
	rootCmd.AddCommand(einride_carrier_v1beta1_CarrierService)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_CreateCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_ListCarriers)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_GetCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_BatchGetCarriers)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UpdateCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_DeleteCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UndeleteCarrier)
	rootCmd.AddCommand(einride_carrier_v1beta1_VehicleService)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_CreateVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_GetVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_BatchGetVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UpdateVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_ListVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_DeleteVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UndeleteVehicle)
	rootCmd.AddCommand(einride_carrier_v1beta1_VehicleTypeService)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_GetVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType)
}
