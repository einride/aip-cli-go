package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_ShippingCostService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingCostService",
}

var einride_shipper_insights_v1beta1_QueryHistoricalShippingCostsRequest v1beta1.QueryHistoricalShippingCostsRequest
var einride_shipper_insights_v1beta1_ShippingCostService_QueryHistoricalShippingCosts = &cobra.Command{
	Use: "QueryHistoricalShippingCosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingCostService.QueryHistoricalShippingCosts")
		return nil
	},
}

var einride_shipper_insights_v1beta1_QueryPlannedShippingCostsRequest v1beta1.QueryPlannedShippingCostsRequest
var einride_shipper_insights_v1beta1_ShippingCostService_QueryPlannedShippingCosts = &cobra.Command{
	Use: "QueryPlannedShippingCosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingCostService.QueryPlannedShippingCosts")
		return nil
	},
}

var einride_shipper_insights_v1beta1_GetShipmentShippingCostRequest v1beta1.GetShipmentShippingCostRequest
var einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost = &cobra.Command{
	Use: "GetShipmentShippingCost",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingCostService.GetShipmentShippingCost")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchGetShipmentShippingCostsRequest v1beta1.BatchGetShipmentShippingCostsRequest
var einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts = &cobra.Command{
	Use: "BatchGetShipmentShippingCosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingCostService.BatchGetShipmentShippingCosts")
		return nil
	},
}

func AddShippingCostServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_QueryHistoricalShippingCosts)
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_QueryPlannedShippingCosts)
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost)
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts)
}
