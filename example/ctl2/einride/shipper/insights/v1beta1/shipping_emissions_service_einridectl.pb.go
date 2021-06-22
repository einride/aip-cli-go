package shipperinsightsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.insights.v1beta1.ShippingEmissionsService.
var (
	einride_shipper_insights_v1beta1_ShippingEmissionsServiceClient v1beta1.ShippingEmissionsServiceClient
	einride_shipper_insights_v1beta1_ShippingEmissionsService       = &cobra.Command{
		Use: "einride.shipper.insights.v1beta1.ShippingEmissionsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_insights_v1beta1_ShippingEmissionsServiceClient = v1beta1.NewShippingEmissionsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.ShippingEmissionsService.GetShipmentShippingEmission.
var (
	einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission_Request v1beta1.GetShipmentShippingEmissionRequest
	einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission         = &cobra.Command{
		Use: "GetShipmentShippingEmission",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.GetShipmentShippingEmission")
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.ShippingEmissionsService.BatchGetShipmentShippingEmissions.
var (
	einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions_Request v1beta1.BatchGetShipmentShippingEmissionsRequest
	einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions         = &cobra.Command{
		Use: "BatchGetShipmentShippingEmissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.ShippingEmissionsService.BatchGetShipmentShippingEmissions")
			return nil
		},
	}
)

func AddShippingEmissionsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission)
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions)
}
