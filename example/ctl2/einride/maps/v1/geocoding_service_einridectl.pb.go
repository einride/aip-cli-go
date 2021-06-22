package mapsv1

import (
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_maps_v1_GeocodingService = &cobra.Command{
	Use: "einride.maps.v1.GeocodingService",
}

var einride_maps_v1_ReverseGeocodeRequest v1.ReverseGeocodeRequest
var einride_maps_v1_GeocodingService_ReverseGeocode = &cobra.Command{
	Use: "ReverseGeocode",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.maps.v1.GeocodingService.ReverseGeocode")
		return nil
	},
}

func AddGeocodingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_maps_v1_GeocodingService)
}

func init() {
	einride_maps_v1_GeocodingService.AddCommand(einride_maps_v1_GeocodingService_ReverseGeocode)
}
