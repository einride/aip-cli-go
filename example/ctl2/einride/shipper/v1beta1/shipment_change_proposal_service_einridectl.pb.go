package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipmentChangeProposalService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentChangeProposalService",
}

var einride_shipper_v1beta1_CreateShipmentChangeProposalRequest v1beta1.CreateShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal = &cobra.Command{
	Use: "CreateShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.CreateShipmentChangeProposal")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipmentChangeProposalRequest v1beta1.GetShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal = &cobra.Command{
	Use: "GetShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.GetShipmentChangeProposal")
		return nil
	},
}

var einride_shipper_v1beta1_ListShipmentChangeProposalsRequest v1beta1.ListShipmentChangeProposalsRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals = &cobra.Command{
	Use: "ListShipmentChangeProposals",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.ListShipmentChangeProposals")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipmentChangeProposalRequest v1beta1.UpdateShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal = &cobra.Command{
	Use: "UpdateShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.UpdateShipmentChangeProposal")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteShipmentChangeProposalRequest v1beta1.DeleteShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal = &cobra.Command{
	Use: "DeleteShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.DeleteShipmentChangeProposal")
		return nil
	},
}

var einride_shipper_v1beta1_ApproveShipmentChangeProposalRequest v1beta1.ApproveShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal = &cobra.Command{
	Use: "ApproveShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.ApproveShipmentChangeProposal")
		return nil
	},
}

var einride_shipper_v1beta1_RejectShipmentChangeProposalRequest v1beta1.RejectShipmentChangeProposalRequest
var einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal = &cobra.Command{
	Use: "RejectShipmentChangeProposal",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.RejectShipmentChangeProposal")
		return nil
	},
}

func AddShipmentChangeProposalServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService)
}

func init() {
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal)
	einride_shipper_v1beta1_ShipmentChangeProposalService.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal)
}
