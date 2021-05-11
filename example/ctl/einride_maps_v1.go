package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_maps_v1_GeocodingService = &cobra.Command{
	Use: "einride.maps.v1.Geocoding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.maps.v1.Geocoding called")
	},
}

var einride_maps_v1_GeocodingService_ReverseGeocode = &cobra.Command{
	Use: "ReverseGeocode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReverseGeocode called")
	},
}

var einride_maps_v1_PathService = &cobra.Command{
	Use: "einride.maps.v1.Path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.maps.v1.Path called")
	},
}

var einride_maps_v1_PathService_ComputePathSummaryMatrix = &cobra.Command{
	Use: "ComputePathSummaryMatrix",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputePathSummaryMatrix called")
	},
}

var einride_maps_v1_PathService_BatchComputePathSummaryMatrices = &cobra.Command{
	Use: "BatchComputePathSummaryMatrices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchComputePathSummaryMatrices called")
	},
}

func init() {
	rootCmd.AddCommand(einride_maps_v1_GeocodingService)
	einride_maps_v1_GeocodingService.AddCommand(einride_maps_v1_GeocodingService_ReverseGeocode)
	rootCmd.AddCommand(einride_maps_v1_PathService)
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_ComputePathSummaryMatrix)
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_BatchComputePathSummaryMatrices)
}
