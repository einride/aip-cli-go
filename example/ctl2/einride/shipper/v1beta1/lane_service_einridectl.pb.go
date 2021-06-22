package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.LaneService.
var (
	einride_shipper_v1beta1_LaneServiceClient v1beta1.LaneServiceClient
	einride_shipper_v1beta1_LaneService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.LaneService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_LaneServiceClient = v1beta1.NewLaneServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.CreateLane.
var (
	einride_shipper_v1beta1_LaneService_CreateLane_Request v1beta1.CreateLaneRequest
	einride_shipper_v1beta1_LaneService_CreateLane         = &cobra.Command{
		Use: "CreateLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.CreateLane")
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.GetLane.
var (
	einride_shipper_v1beta1_LaneService_GetLane_Request v1beta1.GetLaneRequest
	einride_shipper_v1beta1_LaneService_GetLane         = &cobra.Command{
		Use: "GetLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.GetLane")
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.BatchGetLanes.
var (
	einride_shipper_v1beta1_LaneService_BatchGetLanes_Request v1beta1.BatchGetLanesRequest
	einride_shipper_v1beta1_LaneService_BatchGetLanes         = &cobra.Command{
		Use: "BatchGetLanes",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.BatchGetLanes")
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.SearchLanes.
var (
	einride_shipper_v1beta1_LaneService_SearchLanes_Request v1beta1.SearchLanesRequest
	einride_shipper_v1beta1_LaneService_SearchLanes         = &cobra.Command{
		Use: "SearchLanes",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.SearchLanes")
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.UpdateLane.
var (
	einride_shipper_v1beta1_LaneService_UpdateLane_Request v1beta1.UpdateLaneRequest
	einride_shipper_v1beta1_LaneService_UpdateLane         = &cobra.Command{
		Use: "UpdateLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.UpdateLane")
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.ReferenceLane.
var (
	einride_shipper_v1beta1_LaneService_ReferenceLane_Request v1beta1.ReferenceLaneRequest
	einride_shipper_v1beta1_LaneService_ReferenceLane         = &cobra.Command{
		Use: "ReferenceLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.LaneService.ReferenceLane")
			return nil
		},
	}
)

func AddLaneServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_LaneService)
}

func init() {
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_CreateLane)
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_GetLane)
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_BatchGetLanes)
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_SearchLanes)
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_UpdateLane)
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_ReferenceLane)
}
