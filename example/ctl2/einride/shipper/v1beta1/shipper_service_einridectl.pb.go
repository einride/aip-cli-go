package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipperService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipperService",
}

var einride_shipper_v1beta1_CreateShipperRequest v1beta1.CreateShipperRequest
var einride_shipper_v1beta1_ShipperService_CreateShipper = &cobra.Command{
	Use: "CreateShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.CreateShipper")
		return nil
	},
}

var einride_shipper_v1beta1_ListShippersRequest v1beta1.ListShippersRequest
var einride_shipper_v1beta1_ShipperService_ListShippers = &cobra.Command{
	Use: "ListShippers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.ListShippers")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipperRequest v1beta1.GetShipperRequest
var einride_shipper_v1beta1_ShipperService_GetShipper = &cobra.Command{
	Use: "GetShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.GetShipper")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetShippersRequest v1beta1.BatchGetShippersRequest
var einride_shipper_v1beta1_ShipperService_BatchGetShippers = &cobra.Command{
	Use: "BatchGetShippers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.BatchGetShippers")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipperRequest v1beta1.UpdateShipperRequest
var einride_shipper_v1beta1_ShipperService_UpdateShipper = &cobra.Command{
	Use: "UpdateShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.UpdateShipper")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteShipperRequest v1beta1.DeleteShipperRequest
var einride_shipper_v1beta1_ShipperService_DeleteShipper = &cobra.Command{
	Use: "DeleteShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.DeleteShipper")
		return nil
	},
}

var einride_shipper_v1beta1_UndeleteShipperRequest v1beta1.UndeleteShipperRequest
var einride_shipper_v1beta1_ShipperService_UndeleteShipper = &cobra.Command{
	Use: "UndeleteShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperService.UndeleteShipper")
		return nil
	},
}

func AddShipperServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperService)
}

func init() {
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_CreateShipper)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_ListShippers)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_GetShipper)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_BatchGetShippers)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_UpdateShipper)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_DeleteShipper)
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_UndeleteShipper)
}
