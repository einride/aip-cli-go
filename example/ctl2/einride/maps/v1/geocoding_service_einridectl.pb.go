package mapsv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
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
}
