package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipmentClaimService.
var (
	einride_shipper_v1beta1_ShipmentClaimServiceClient v1beta1.ShipmentClaimServiceClient
	einride_shipper_v1beta1_ShipmentClaimService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipmentClaimService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentClaimServiceClient = v1beta1.NewShipmentClaimServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.CreateShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request v1beta1.CreateShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim         = &cobra.Command{
		Use: "CreateShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentClaimService.CreateShipmentClaim")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.GetShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim_Request v1beta1.GetShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim         = &cobra.Command{
		Use: "GetShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentClaimService.GetShipmentClaim")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.UpdateShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request v1beta1.UpdateShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim         = &cobra.Command{
		Use: "UpdateShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentClaimService.UpdateShipmentClaim")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.ListShipmentClaims.
var (
	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request v1beta1.ListShipmentClaimsRequest
	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims         = &cobra.Command{
		Use: "ListShipmentClaims",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentClaimService.ListShipmentClaims")
			return nil
		},
	}
)

func AddShipmentClaimServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentClaimService)
}

func init() {
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims)
}
