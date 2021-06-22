package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_ShippingEmissionsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingEmissionsService",
}

var einride_shipper_insights_v1beta1_QueryHistoricalShippingEmissionsRequest v1beta1.QueryHistoricalShippingEmissionsRequest
var einride_shipper_insights_v1beta1_ShippingEmissionsService_QueryHistoricalShippingEmissions = &cobra.Command{
	Use: "QueryHistoricalShippingEmissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.QueryHistoricalShippingEmissions")
		return nil
	},
}

var einride_shipper_insights_v1beta1_QueryPlannedShippingEmissionsRequest v1beta1.QueryPlannedShippingEmissionsRequest
var einride_shipper_insights_v1beta1_ShippingEmissionsService_QueryPlannedShippingEmissions = &cobra.Command{
	Use: "QueryPlannedShippingEmissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.QueryPlannedShippingEmissions")
		return nil
	},
}

var einride_shipper_insights_v1beta1_GetShipmentShippingEmissionRequest v1beta1.GetShipmentShippingEmissionRequest
var einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission = &cobra.Command{
	Use: "GetShipmentShippingEmission",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.GetShipmentShippingEmission")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchGetShipmentShippingEmissionsRequest v1beta1.BatchGetShipmentShippingEmissionsRequest
var einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions = &cobra.Command{
	Use: "BatchGetShipmentShippingEmissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.BatchGetShipmentShippingEmissions")
		return nil
	},
}

func AddShippingEmissionsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_QueryHistoricalShippingEmissions)
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_QueryPlannedShippingEmissions)
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission)
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions)
}
