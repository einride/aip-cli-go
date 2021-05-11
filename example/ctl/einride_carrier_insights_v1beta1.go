package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_carrier_insights_v1beta1_GeocodingService = &cobra.Command{
	Use: "einride.carrier.insights.v1beta1.Geocoding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.insights.v1beta1.Geocoding called")
	},
}

var einride_carrier_insights_v1beta1_GeocodingService_ReverseGeocode = &cobra.Command{
	Use: "ReverseGeocode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReverseGeocode called")
	},
}

var einride_carrier_insights_v1beta1_VehicleInsightsService = &cobra.Command{
	Use: "einride.carrier.insights.v1beta1.VehicleInsights",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.insights.v1beta1.VehicleInsights called")
	},
}

var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleEvents = &cobra.Command{
	Use: "ComputeVehicleEvents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputeVehicleEvents called")
	},
}

var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleStats = &cobra.Command{
	Use: "ComputeVehicleStats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputeVehicleStats called")
	},
}

var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleHeatMap = &cobra.Command{
	Use: "ComputeVehicleHeatMap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputeVehicleHeatMap called")
	},
}

func init() {
	rootCmd.AddCommand(einride_carrier_insights_v1beta1_GeocodingService)
	einride_carrier_insights_v1beta1_GeocodingService.AddCommand(einride_carrier_insights_v1beta1_GeocodingService_ReverseGeocode)
	rootCmd.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService)
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleEvents)
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleStats)
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleHeatMap)
}
