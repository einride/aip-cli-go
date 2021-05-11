package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_maps_v1_GeocodingService = &cobra.Command{
	Use: "einride.maps.v1.Geocoding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.maps.v1.Geocoding called")
	},
}
var einride_maps_v1_PathService = &cobra.Command{
	Use: "einride.maps.v1.Path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.maps.v1.Path called")
	},
}

func init() {
	rootCmd.AddCommand(einride_maps_v1_GeocodingService)
	rootCmd.AddCommand(einride_maps_v1_PathService)
}
