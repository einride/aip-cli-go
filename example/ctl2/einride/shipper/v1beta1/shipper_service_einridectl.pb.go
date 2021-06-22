package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipperService.
var (
	einride_shipper_v1beta1_ShipperServiceClient v1beta1.ShipperServiceClient
	einride_shipper_v1beta1_ShipperService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipperService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipperServiceClient = v1beta1.NewShipperServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.CreateShipper.
var (
	einride_shipper_v1beta1_ShipperService_CreateShipper_Request v1beta1.CreateShipperRequest
	einride_shipper_v1beta1_ShipperService_CreateShipper         = &cobra.Command{
		Use: "CreateShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.CreateShipper")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.ListShippers.
var (
	einride_shipper_v1beta1_ShipperService_ListShippers_Request v1beta1.ListShippersRequest
	einride_shipper_v1beta1_ShipperService_ListShippers         = &cobra.Command{
		Use: "ListShippers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.ListShippers")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.GetShipper.
var (
	einride_shipper_v1beta1_ShipperService_GetShipper_Request v1beta1.GetShipperRequest
	einride_shipper_v1beta1_ShipperService_GetShipper         = &cobra.Command{
		Use: "GetShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.GetShipper")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.BatchGetShippers.
var (
	einride_shipper_v1beta1_ShipperService_BatchGetShippers_Request v1beta1.BatchGetShippersRequest
	einride_shipper_v1beta1_ShipperService_BatchGetShippers         = &cobra.Command{
		Use: "BatchGetShippers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.BatchGetShippers")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.UpdateShipper.
var (
	einride_shipper_v1beta1_ShipperService_UpdateShipper_Request v1beta1.UpdateShipperRequest
	einride_shipper_v1beta1_ShipperService_UpdateShipper         = &cobra.Command{
		Use: "UpdateShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.UpdateShipper")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.DeleteShipper.
var (
	einride_shipper_v1beta1_ShipperService_DeleteShipper_Request v1beta1.DeleteShipperRequest
	einride_shipper_v1beta1_ShipperService_DeleteShipper         = &cobra.Command{
		Use: "DeleteShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.DeleteShipper")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.UndeleteShipper.
var (
	einride_shipper_v1beta1_ShipperService_UndeleteShipper_Request v1beta1.UndeleteShipperRequest
	einride_shipper_v1beta1_ShipperService_UndeleteShipper         = &cobra.Command{
		Use: "UndeleteShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperService.UndeleteShipper")
			return nil
		},
	}
)

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
