package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_ETruckRecommendationService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ETruckRecommendationService",
}

var einride_shipper_insights_v1beta1_CreateETruckRecommendationRequest v1beta1.CreateETruckRecommendationRequest
var einride_shipper_insights_v1beta1_ETruckRecommendationService_CreateETruckRecommendation = &cobra.Command{
	Use: "CreateETruckRecommendation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ETruckRecommendationService.CreateETruckRecommendation")
		return nil
	},
}

var einride_shipper_insights_v1beta1_GetETruckRecommendationRequest v1beta1.GetETruckRecommendationRequest
var einride_shipper_insights_v1beta1_ETruckRecommendationService_GetETruckRecommendation = &cobra.Command{
	Use: "GetETruckRecommendation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ETruckRecommendationService.GetETruckRecommendation")
		return nil
	},
}

var einride_shipper_insights_v1beta1_ListETruckRecommendationsRequest v1beta1.ListETruckRecommendationsRequest
var einride_shipper_insights_v1beta1_ETruckRecommendationService_ListETruckRecommendations = &cobra.Command{
	Use: "ListETruckRecommendations",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ETruckRecommendationService.ListETruckRecommendations")
		return nil
	},
}

func AddETruckRecommendationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ETruckRecommendationService)
}

func init() {
	einride_shipper_insights_v1beta1_ETruckRecommendationService.AddCommand(einride_shipper_insights_v1beta1_ETruckRecommendationService_CreateETruckRecommendation)
	einride_shipper_insights_v1beta1_ETruckRecommendationService.AddCommand(einride_shipper_insights_v1beta1_ETruckRecommendationService_GetETruckRecommendation)
	einride_shipper_insights_v1beta1_ETruckRecommendationService.AddCommand(einride_shipper_insights_v1beta1_ETruckRecommendationService_ListETruckRecommendations)
}
