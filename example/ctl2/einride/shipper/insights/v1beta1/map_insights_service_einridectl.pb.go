package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_MapInsightsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.MapInsightsService",
}

var einride_shipper_insights_v1beta1_BatchGetShipperMapInsightsRequest v1beta1.BatchGetShipperMapInsightsRequest
var einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights = &cobra.Command{
	Use: "BatchGetShipperMapInsights",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetShipperMapInsights")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchGetSiteMapInsightsRequest v1beta1.BatchGetSiteMapInsightsRequest
var einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights = &cobra.Command{
	Use: "BatchGetSiteMapInsights",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetSiteMapInsights")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchGetLaneMapInsightsRequest v1beta1.BatchGetLaneMapInsightsRequest
var einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights = &cobra.Command{
	Use: "BatchGetLaneMapInsights",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetLaneMapInsights")
		return nil
	},
}

func AddMapInsightsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService)
}

func init() {
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights)
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights)
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights)
}
