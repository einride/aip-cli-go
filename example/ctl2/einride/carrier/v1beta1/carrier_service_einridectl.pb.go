package carrierv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.carrier.v1beta1.CarrierService.
var (
	einride_carrier_v1beta1_CarrierServiceClient v1beta1.CarrierServiceClient
	einride_carrier_v1beta1_CarrierService       = &cobra.Command{
		Use: "einride.carrier.v1beta1.CarrierService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_carrier_v1beta1_CarrierServiceClient = v1beta1.NewCarrierServiceClient(conn)
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.CreateCarrier.
var (
	einride_carrier_v1beta1_CarrierService_CreateCarrier_Request v1beta1.CreateCarrierRequest
	einride_carrier_v1beta1_CarrierService_CreateCarrier         = &cobra.Command{
		Use: "CreateCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.CreateCarrier")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.ListCarriers.
var (
	einride_carrier_v1beta1_CarrierService_ListCarriers_Request v1beta1.ListCarriersRequest
	einride_carrier_v1beta1_CarrierService_ListCarriers         = &cobra.Command{
		Use: "ListCarriers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.ListCarriers")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.GetCarrier.
var (
	einride_carrier_v1beta1_CarrierService_GetCarrier_Request v1beta1.GetCarrierRequest
	einride_carrier_v1beta1_CarrierService_GetCarrier         = &cobra.Command{
		Use: "GetCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.GetCarrier")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.BatchGetCarriers.
var (
	einride_carrier_v1beta1_CarrierService_BatchGetCarriers_Request v1beta1.BatchGetCarriersRequest
	einride_carrier_v1beta1_CarrierService_BatchGetCarriers         = &cobra.Command{
		Use: "BatchGetCarriers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.BatchGetCarriers")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.UpdateCarrier.
var (
	einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request v1beta1.UpdateCarrierRequest
	einride_carrier_v1beta1_CarrierService_UpdateCarrier         = &cobra.Command{
		Use: "UpdateCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.UpdateCarrier")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.DeleteCarrier.
var (
	einride_carrier_v1beta1_CarrierService_DeleteCarrier_Request v1beta1.DeleteCarrierRequest
	einride_carrier_v1beta1_CarrierService_DeleteCarrier         = &cobra.Command{
		Use: "DeleteCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.DeleteCarrier")
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.UndeleteCarrier.
var (
	einride_carrier_v1beta1_CarrierService_UndeleteCarrier_Request v1beta1.UndeleteCarrierRequest
	einride_carrier_v1beta1_CarrierService_UndeleteCarrier         = &cobra.Command{
		Use: "UndeleteCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.CarrierService.UndeleteCarrier")
			return nil
		},
	}
)

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
