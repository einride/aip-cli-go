package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipmentClaimService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentClaimService",
}

var einride_shipper_v1beta1_CreateShipmentClaimRequest v1beta1.CreateShipmentClaimRequest
var einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim = &cobra.Command{
	Use: "CreateShipmentClaim",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentClaimService.CreateShipmentClaim")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipmentClaimRequest v1beta1.GetShipmentClaimRequest
var einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim = &cobra.Command{
	Use: "GetShipmentClaim",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentClaimService.GetShipmentClaim")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipmentClaimRequest v1beta1.UpdateShipmentClaimRequest
var einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim = &cobra.Command{
	Use: "UpdateShipmentClaim",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentClaimService.UpdateShipmentClaim")
		return nil
	},
}

var einride_shipper_v1beta1_ListShipmentClaimsRequest v1beta1.ListShipmentClaimsRequest
var einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims = &cobra.Command{
	Use: "ListShipmentClaims",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentClaimService.ListShipmentClaims")
		return nil
	},
}

func AddShipmentClaimServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentClaimService)
}

func init() {
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims)
}
