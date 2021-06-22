package carrierv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_carrier_v1beta1_CarrierService = &cobra.Command{
	Use: "einride.carrier.v1beta1.CarrierService",
}

var einride_carrier_v1beta1_CreateCarrierRequest v1beta1.CreateCarrierRequest
var einride_carrier_v1beta1_CarrierService_CreateCarrier = &cobra.Command{
	Use: "CreateCarrier",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.CreateCarrier")
		return nil
	},
}

var einride_carrier_v1beta1_ListCarriersRequest v1beta1.ListCarriersRequest
var einride_carrier_v1beta1_CarrierService_ListCarriers = &cobra.Command{
	Use: "ListCarriers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.ListCarriers")
		return nil
	},
}

var einride_carrier_v1beta1_GetCarrierRequest v1beta1.GetCarrierRequest
var einride_carrier_v1beta1_CarrierService_GetCarrier = &cobra.Command{
	Use: "GetCarrier",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.GetCarrier")
		return nil
	},
}

var einride_carrier_v1beta1_BatchGetCarriersRequest v1beta1.BatchGetCarriersRequest
var einride_carrier_v1beta1_CarrierService_BatchGetCarriers = &cobra.Command{
	Use: "BatchGetCarriers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.BatchGetCarriers")
		return nil
	},
}

var einride_carrier_v1beta1_UpdateCarrierRequest v1beta1.UpdateCarrierRequest
var einride_carrier_v1beta1_CarrierService_UpdateCarrier = &cobra.Command{
	Use: "UpdateCarrier",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.UpdateCarrier")
		return nil
	},
}

var einride_carrier_v1beta1_DeleteCarrierRequest v1beta1.DeleteCarrierRequest
var einride_carrier_v1beta1_CarrierService_DeleteCarrier = &cobra.Command{
	Use: "DeleteCarrier",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.DeleteCarrier")
		return nil
	},
}

var einride_carrier_v1beta1_UndeleteCarrierRequest v1beta1.UndeleteCarrierRequest
var einride_carrier_v1beta1_CarrierService_UndeleteCarrier = &cobra.Command{
	Use: "UndeleteCarrier",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.v1beta1.CarrierService.UndeleteCarrier")
		return nil
	},
}

func AddCarrierServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_CarrierService)
}

func init() {
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_CreateCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_ListCarriers)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_GetCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_BatchGetCarriers)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UpdateCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_DeleteCarrier)
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UndeleteCarrier)
}
