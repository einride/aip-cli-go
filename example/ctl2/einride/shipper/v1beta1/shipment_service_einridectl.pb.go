package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentService",
}

var einride_shipper_v1beta1_CreateShipmentRequest v1beta1.CreateShipmentRequest
var einride_shipper_v1beta1_ShipmentService_CreateShipment = &cobra.Command{
	Use: "CreateShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.CreateShipment")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipmentRequest v1beta1.GetShipmentRequest
var einride_shipper_v1beta1_ShipmentService_GetShipment = &cobra.Command{
	Use: "GetShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.GetShipment")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetShipmentsRequest v1beta1.BatchGetShipmentsRequest
var einride_shipper_v1beta1_ShipmentService_BatchGetShipments = &cobra.Command{
	Use: "BatchGetShipments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.BatchGetShipments")
		return nil
	},
}

var einride_shipper_v1beta1_SearchShipmentsRequest v1beta1.SearchShipmentsRequest
var einride_shipper_v1beta1_ShipmentService_SearchShipments = &cobra.Command{
	Use: "SearchShipments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.SearchShipments")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipmentRequest v1beta1.UpdateShipmentRequest
var einride_shipper_v1beta1_ShipmentService_UpdateShipment = &cobra.Command{
	Use: "UpdateShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.UpdateShipment")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteShipmentRequest v1beta1.DeleteShipmentRequest
var einride_shipper_v1beta1_ShipmentService_DeleteShipment = &cobra.Command{
	Use: "DeleteShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.DeleteShipment")
		return nil
	},
}

var einride_shipper_v1beta1_UndeleteShipmentRequest v1beta1.UndeleteShipmentRequest
var einride_shipper_v1beta1_ShipmentService_UndeleteShipment = &cobra.Command{
	Use: "UndeleteShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.UndeleteShipment")
		return nil
	},
}

var einride_shipper_v1beta1_ReferenceShipmentRequest v1beta1.ReferenceShipmentRequest
var einride_shipper_v1beta1_ShipmentService_ReferenceShipment = &cobra.Command{
	Use: "ReferenceShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.ReferenceShipment")
		return nil
	},
}

var einride_shipper_v1beta1_PlanShipmentRequest v1beta1.PlanShipmentRequest
var einride_shipper_v1beta1_ShipmentService_PlanShipment = &cobra.Command{
	Use: "PlanShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.PlanShipment")
		return nil
	},
}

var einride_shipper_v1beta1_UnplanShipmentRequest v1beta1.UnplanShipmentRequest
var einride_shipper_v1beta1_ShipmentService_UnplanShipment = &cobra.Command{
	Use: "UnplanShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.UnplanShipment")
		return nil
	},
}

var einride_shipper_v1beta1_LoadShipmentRequest v1beta1.LoadShipmentRequest
var einride_shipper_v1beta1_ShipmentService_LoadShipment = &cobra.Command{
	Use: "LoadShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.LoadShipment")
		return nil
	},
}

var einride_shipper_v1beta1_LoadShipmentUnknownVehicleRequest v1beta1.LoadShipmentUnknownVehicleRequest
var einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle = &cobra.Command{
	Use: "LoadShipmentUnknownVehicle",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.LoadShipmentUnknownVehicle")
		return nil
	},
}

var einride_shipper_v1beta1_TransitShipmentRequest v1beta1.TransitShipmentRequest
var einride_shipper_v1beta1_ShipmentService_TransitShipment = &cobra.Command{
	Use: "TransitShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.TransitShipment")
		return nil
	},
}

var einride_shipper_v1beta1_ArriveShipmentRequest v1beta1.ArriveShipmentRequest
var einride_shipper_v1beta1_ShipmentService_ArriveShipment = &cobra.Command{
	Use: "ArriveShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.ArriveShipment")
		return nil
	},
}

var einride_shipper_v1beta1_UnloadShipmentRequest v1beta1.UnloadShipmentRequest
var einride_shipper_v1beta1_ShipmentService_UnloadShipment = &cobra.Command{
	Use: "UnloadShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.UnloadShipment")
		return nil
	},
}

var einride_shipper_v1beta1_DeliverShipmentRequest v1beta1.DeliverShipmentRequest
var einride_shipper_v1beta1_ShipmentService_DeliverShipment = &cobra.Command{
	Use: "DeliverShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.DeliverShipment")
		return nil
	},
}

var einride_shipper_v1beta1_HandleShipmentExternallyRequest v1beta1.HandleShipmentExternallyRequest
var einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally = &cobra.Command{
	Use: "HandleShipmentExternally",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.HandleShipmentExternally")
		return nil
	},
}

var einride_shipper_v1beta1_RejectShipmentRequest v1beta1.RejectShipmentRequest
var einride_shipper_v1beta1_ShipmentService_RejectShipment = &cobra.Command{
	Use: "RejectShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.RejectShipment")
		return nil
	},
}

var einride_shipper_v1beta1_UnrejectShipmentRequest v1beta1.UnrejectShipmentRequest
var einride_shipper_v1beta1_ShipmentService_UnrejectShipment = &cobra.Command{
	Use: "UnrejectShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentService.UnrejectShipment")
		return nil
	},
}

func AddShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_CreateShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_GetShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_BatchGetShipments)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_SearchShipments)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UpdateShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeleteShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UndeleteShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ReferenceShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_PlanShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnplanShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_TransitShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ArriveShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnloadShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeliverShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_RejectShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnrejectShipment)
}
