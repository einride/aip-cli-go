package carrierinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_carrier_insights_v1beta1_VehicleInsightsService = &cobra.Command{
	Use: "einride.carrier.insights.v1beta1.VehicleInsightsService",
}

var einride_carrier_insights_v1beta1_ComputeVehicleEventsRequest v1beta1.ComputeVehicleEventsRequest
var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleEvents = &cobra.Command{
	Use: "ComputeVehicleEvents",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.insights.v1beta1.VehicleInsightsService.ComputeVehicleEvents")
		return nil
	},
}

var einride_carrier_insights_v1beta1_ComputeVehicleStatsRequest v1beta1.ComputeVehicleStatsRequest
var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleStats = &cobra.Command{
	Use: "ComputeVehicleStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.insights.v1beta1.VehicleInsightsService.ComputeVehicleStats")
		return nil
	},
}

var einride_carrier_insights_v1beta1_ComputeVehicleHeatMapRequest v1beta1.ComputeVehicleHeatMapRequest
var einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleHeatMap = &cobra.Command{
	Use: "ComputeVehicleHeatMap",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.carrier.insights.v1beta1.VehicleInsightsService.ComputeVehicleHeatMap")
		return nil
	},
}

func AddVehicleInsightsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService)
}

func init() {
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleEvents)
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleStats)
	einride_carrier_insights_v1beta1_VehicleInsightsService.AddCommand(einride_carrier_insights_v1beta1_VehicleInsightsService_ComputeVehicleHeatMap)
}
