package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_LaneService = &cobra.Command{
	Use: "einride.shipper.v1beta1.LaneService",
}

var einride_shipper_v1beta1_CreateLaneRequest v1beta1.CreateLaneRequest
var einride_shipper_v1beta1_LaneService_CreateLane = &cobra.Command{
	Use: "CreateLane",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.CreateLane")
		return nil
	},
}

var einride_shipper_v1beta1_GetLaneRequest v1beta1.GetLaneRequest
var einride_shipper_v1beta1_LaneService_GetLane = &cobra.Command{
	Use: "GetLane",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.GetLane")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetLanesRequest v1beta1.BatchGetLanesRequest
var einride_shipper_v1beta1_LaneService_BatchGetLanes = &cobra.Command{
	Use: "BatchGetLanes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.BatchGetLanes")
		return nil
	},
}

var einride_shipper_v1beta1_SearchLanesRequest v1beta1.SearchLanesRequest
var einride_shipper_v1beta1_LaneService_SearchLanes = &cobra.Command{
	Use: "SearchLanes",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.SearchLanes")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateLaneRequest v1beta1.UpdateLaneRequest
var einride_shipper_v1beta1_LaneService_UpdateLane = &cobra.Command{
	Use: "UpdateLane",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.UpdateLane")
		return nil
	},
}

var einride_shipper_v1beta1_ReferenceLaneRequest v1beta1.ReferenceLaneRequest
var einride_shipper_v1beta1_LaneService_ReferenceLane = &cobra.Command{
	Use: "ReferenceLane",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.LaneService.ReferenceLane")
		return nil
	},
}

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
