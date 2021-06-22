package mapsv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.maps.v1.PathService.
var (
	einride_maps_v1_PathServiceClient v1.PathServiceClient
	einride_maps_v1_PathService       = &cobra.Command{
		Use: "einride.maps.v1.PathService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_maps_v1_PathServiceClient = v1.NewPathServiceClient(conn)
			return nil
		},
	}
)

// einride.maps.v1.PathService.ComputePathSummaryMatrix.
var (
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request v1.ComputePathSummaryMatrixRequest
	einride_maps_v1_PathService_ComputePathSummaryMatrix         = &cobra.Command{
		Use: "ComputePathSummaryMatrix",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.maps.v1.PathService.ComputePathSummaryMatrix")
			return nil
		},
	}
)

// einride.maps.v1.PathService.BatchComputePathSummaryMatrices.
var (
	einride_maps_v1_PathService_BatchComputePathSummaryMatrices_Request v1.BatchComputePathSummaryMatricesRequest
	einride_maps_v1_PathService_BatchComputePathSummaryMatrices         = &cobra.Command{
		Use: "BatchComputePathSummaryMatrices",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.maps.v1.PathService.BatchComputePathSummaryMatrices")
			return nil
		},
	}
)

func AddPathServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_maps_v1_PathService)
}

func init() {
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_ComputePathSummaryMatrix)
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_BatchComputePathSummaryMatrices)
}
