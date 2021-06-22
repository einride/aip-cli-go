package carrierv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_carrier_v1beta1_VehicleService = &cobra.Command{
	Use: "einride.carrier.v1beta1.VehicleService",
}

var einride_carrier_v1beta1_CreateVehicleRequest v1beta1.CreateVehicleRequest
var einride_carrier_v1beta1_VehicleService_CreateVehicle = &cobra.Command{
	Use: "CreateVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.CreateVehicle")
		return nil
	},
}

var einride_carrier_v1beta1_GetVehicleRequest v1beta1.GetVehicleRequest
var einride_carrier_v1beta1_VehicleService_GetVehicle = &cobra.Command{
	Use: "GetVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.GetVehicle")
		return nil
	},
}

var einride_carrier_v1beta1_BatchGetVehiclesRequest v1beta1.BatchGetVehiclesRequest
var einride_carrier_v1beta1_VehicleService_BatchGetVehicles = &cobra.Command{
	Use: "BatchGetVehicles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.BatchGetVehicles")
		return nil
	},
}

var einride_carrier_v1beta1_UpdateVehicleRequest v1beta1.UpdateVehicleRequest
var einride_carrier_v1beta1_VehicleService_UpdateVehicle = &cobra.Command{
	Use: "UpdateVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.UpdateVehicle")
		return nil
	},
}

var einride_carrier_v1beta1_ListVehiclesRequest v1beta1.ListVehiclesRequest
var einride_carrier_v1beta1_VehicleService_ListVehicles = &cobra.Command{
	Use: "ListVehicles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.ListVehicles")
		return nil
	},
}

var einride_carrier_v1beta1_DeleteVehicleRequest v1beta1.DeleteVehicleRequest
var einride_carrier_v1beta1_VehicleService_DeleteVehicle = &cobra.Command{
	Use: "DeleteVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.DeleteVehicle")
		return nil
	},
}

var einride_carrier_v1beta1_UndeleteVehicleRequest v1beta1.UndeleteVehicleRequest
var einride_carrier_v1beta1_VehicleService_UndeleteVehicle = &cobra.Command{
	Use: "UndeleteVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.VehicleService.UndeleteVehicle")
		return nil
	},
}

func AddVehicleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_VehicleService)
}

func init() {
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_CreateVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_GetVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_BatchGetVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UpdateVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_ListVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_DeleteVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UndeleteVehicle)
}
