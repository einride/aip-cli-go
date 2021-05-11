package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_shipper_insights_v1beta1_ETruckRecommendationService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ETruckRecommendation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.ETruckRecommendation called")
	},
}
var einride_shipper_insights_v1beta1_MapInsightsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.MapInsights",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.MapInsights called")
	},
}
var einride_shipper_insights_v1beta1_ShippingCostService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingCost",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.ShippingCost called")
	},
}
var einride_shipper_insights_v1beta1_ShippingEmissionsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingEmissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.ShippingEmissions called")
	},
}
var einride_shipper_insights_v1beta1_ShippingForecastService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingForecast",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.ShippingForecast called")
	},
}
var einride_shipper_insights_v1beta1_ShippingStatsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingStats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.ShippingStats called")
	},
}
var einride_shipper_insights_v1beta1_TransformationAssessmentService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.TransformationAssessment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.insights.v1beta1.TransformationAssessment called")
	},
}

func init() {
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_ETruckRecommendationService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_ShippingForecastService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService)
	rootCmd.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService)
}
