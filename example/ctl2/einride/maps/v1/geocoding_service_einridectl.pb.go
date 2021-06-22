package mapsv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	log "log"
)

// einride.maps.v1.GeocodingService.
var (
	einride_maps_v1_GeocodingServiceClient v1.GeocodingServiceClient
	einride_maps_v1_GeocodingService       = &cobra.Command{
		Use: "einride.maps.v1.GeocodingService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_maps_v1_GeocodingServiceClient = v1.NewGeocodingServiceClient(conn)
			return nil
		},
	}
)

// einride.maps.v1.GeocodingService.ReverseGeocode.
var (
	einride_maps_v1_GeocodingService_ReverseGeocode_Request v1.ReverseGeocodeRequest
	einride_maps_v1_GeocodingService_ReverseGeocode         = &cobra.Command{
		Use: "ReverseGeocode",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.maps.v1.GeocodingService.ReverseGeocode")
			return nil
		},
	}
)

func AddGeocodingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_maps_v1_GeocodingService)
}

func init() {
	einride_maps_v1_GeocodingService.AddCommand(einride_maps_v1_GeocodingService_ReverseGeocode)

	einride_maps_v1_GeocodingService_ReverseGeocode_Request.LatLng = new(latlng.LatLng)
	einride_maps_v1_GeocodingService_ReverseGeocode.Flags().Float64Var(&einride_maps_v1_GeocodingService_ReverseGeocode_Request.LatLng.Latitude, "latLng.latitude", 10, "The latitude in degrees. It must be in the range [-90.0, +90.0].")
	einride_maps_v1_GeocodingService_ReverseGeocode.Flags().Float64Var(&einride_maps_v1_GeocodingService_ReverseGeocode_Request.LatLng.Longitude, "latLng.longitude", 10, "The longitude in degrees. It must be in the range [-180.0, +180.0].")
}
