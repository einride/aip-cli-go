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
var einride_carrier_insights_v1beta1_VehicleInsightsService = &cobra.Command{
	Use: "einride.carrier.insights.v1beta1.VehicleInsights",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.carrier.insights.v1beta1.VehicleInsights called")
	},
}

func init() {
	rootCmd.AddCommand(einride_carrier_insights_v1beta1_GeocodingService)
	rootCmd.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService)
}
