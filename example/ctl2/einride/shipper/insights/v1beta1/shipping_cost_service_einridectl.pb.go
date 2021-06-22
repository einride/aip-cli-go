package shipperinsightsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.insights.v1beta1.ShippingCostService.
var (
	einride_shipper_insights_v1beta1_ShippingCostServiceClient v1beta1.ShippingCostServiceClient
	einride_shipper_insights_v1beta1_ShippingCostService       = &cobra.Command{
		Use: "einride.shipper.insights.v1beta1.ShippingCostService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_insights_v1beta1_ShippingCostServiceClient = v1beta1.NewShippingCostServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.ShippingCostService.GetShipmentShippingCost.
var (
	einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost_Request v1beta1.GetShipmentShippingCostRequest
	einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost         = &cobra.Command{
		Use: "GetShipmentShippingCost",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.ShippingCostService.GetShipmentShippingCost")
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.ShippingCostService.BatchGetShipmentShippingCosts.
var (
	einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts_Request v1beta1.BatchGetShipmentShippingCostsRequest
	einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts         = &cobra.Command{
		Use: "BatchGetShipmentShippingCosts",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.ShippingCostService.BatchGetShipmentShippingCosts")
			return nil
		},
	}
)

func AddShippingCostServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost)

	einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost.Flags().StringVar(&einride_shipper_insights_v1beta1_ShippingCostService_GetShipmentShippingCost_Request.Name, "name", "", "Resource name of the shipment shipping cost to retrieve.")
	einride_shipper_insights_v1beta1_ShippingCostService.AddCommand(einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts)

	einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts.Flags().StringVar(&einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts_Request.Parent, "parent", "", "Resource name of the parent shipment shared by all shipment shipping costs.\nThe last segment of the parent resource name may be a wildcard to get\nshipment shipping costs for multiple shipments.")

	einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts.Flags().StringSliceVar(&einride_shipper_insights_v1beta1_ShippingCostService_BatchGetShipmentShippingCosts_Request.Names, "names", []string{}, "Resource names of the shipment shipping costs.")
}
