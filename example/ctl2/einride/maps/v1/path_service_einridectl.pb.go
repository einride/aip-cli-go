package mapsv1

import (
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_maps_v1_PathService = &cobra.Command{
	Use: "einride.maps.v1.PathService",
}

var einride_maps_v1_ComputePathSummaryMatrixRequest v1.ComputePathSummaryMatrixRequest
var einride_maps_v1_PathService_ComputePathSummaryMatrix = &cobra.Command{
	Use: "ComputePathSummaryMatrix",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.maps.v1.PathService.ComputePathSummaryMatrix")
		return nil
	},
}

var einride_maps_v1_BatchComputePathSummaryMatricesRequest v1.BatchComputePathSummaryMatricesRequest
var einride_maps_v1_PathService_BatchComputePathSummaryMatrices = &cobra.Command{
	Use: "BatchComputePathSummaryMatrices",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.maps.v1.PathService.BatchComputePathSummaryMatrices")
		return nil
	},
}

func AddPathServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_maps_v1_PathService)
}

func init() {
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_ComputePathSummaryMatrix)
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_BatchComputePathSummaryMatrices)
}
