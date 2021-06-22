package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_ShippingForecastService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingForecastService",
}

var einride_shipper_insights_v1beta1_ForecastShipperShippingRequest v1beta1.ForecastShipperShippingRequest
var einride_shipper_insights_v1beta1_ShippingForecastService_ForecastShipperShipping = &cobra.Command{
	Use: "ForecastShipperShipping",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingForecastService.ForecastShipperShipping")
		return nil
	},
}

var einride_shipper_insights_v1beta1_ForecastSiteShippingRequest v1beta1.ForecastSiteShippingRequest
var einride_shipper_insights_v1beta1_ShippingForecastService_ForecastSiteShipping = &cobra.Command{
	Use: "ForecastSiteShipping",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingForecastService.ForecastSiteShipping")
		return nil
	},
}

func AddShippingForecastServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingForecastService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingForecastService.AddCommand(einride_shipper_insights_v1beta1_ShippingForecastService_ForecastShipperShipping)
	einride_shipper_insights_v1beta1_ShippingForecastService.AddCommand(einride_shipper_insights_v1beta1_ShippingForecastService_ForecastSiteShipping)
}
