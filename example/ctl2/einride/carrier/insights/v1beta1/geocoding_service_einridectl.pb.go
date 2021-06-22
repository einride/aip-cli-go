package carrierinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_carrier_insights_v1beta1_GeocodingService = &cobra.Command{
	Use: "einride.carrier.insights.v1beta1.GeocodingService",
}

var einride_carrier_insights_v1beta1_ReverseGeocodeRequest v1beta1.ReverseGeocodeRequest
var einride_carrier_insights_v1beta1_GeocodingService_ReverseGeocode = &cobra.Command{
	Use: "ReverseGeocode",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.insights.v1beta1.GeocodingService.ReverseGeocode")
		return nil
	},
}

func AddGeocodingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_insights_v1beta1_GeocodingService)
}

func init() {
	einride_carrier_insights_v1beta1_GeocodingService.AddCommand(einride_carrier_insights_v1beta1_GeocodingService_ReverseGeocode)
}
