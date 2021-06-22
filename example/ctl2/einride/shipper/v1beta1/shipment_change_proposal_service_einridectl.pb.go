package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalServiceClient v1beta1.ShipmentChangeProposalServiceClient
	einride_shipper_v1beta1_ShipmentChangeProposalService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipmentChangeProposalService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentChangeProposalServiceClient = v1beta1.NewShipmentChangeProposalServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.CreateShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal_Request v1beta1.CreateShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_CreateShipmentChangeProposal         = &cobra.Command{
		Use: "CreateShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.CreateShipmentChangeProposal")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.GetShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal_Request v1beta1.GetShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_GetShipmentChangeProposal         = &cobra.Command{
		Use: "GetShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.GetShipmentChangeProposal")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.ListShipmentChangeProposals.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals_Request v1beta1.ListShipmentChangeProposalsRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_ListShipmentChangeProposals         = &cobra.Command{
		Use: "ListShipmentChangeProposals",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.ListShipmentChangeProposals")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.UpdateShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal_Request v1beta1.UpdateShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_UpdateShipmentChangeProposal         = &cobra.Command{
		Use: "UpdateShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.UpdateShipmentChangeProposal")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.DeleteShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal_Request v1beta1.DeleteShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_DeleteShipmentChangeProposal         = &cobra.Command{
		Use: "DeleteShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.DeleteShipmentChangeProposal")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.ApproveShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal_Request v1beta1.ApproveShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_ApproveShipmentChangeProposal         = &cobra.Command{
		Use: "ApproveShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.ApproveShipmentChangeProposal")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentChangeProposalService.RejectShipmentChangeProposal.
var (
	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal_Request v1beta1.RejectShipmentChangeProposalRequest
	einride_shipper_v1beta1_ShipmentChangeProposalService_RejectShipmentChangeProposal         = &cobra.Command{
		Use: "RejectShipmentChangeProposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentChangeProposalService.RejectShipmentChangeProposal")
			return nil
		},
	}
)

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
