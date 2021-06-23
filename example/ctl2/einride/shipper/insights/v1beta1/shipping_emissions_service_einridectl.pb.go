package shipperinsightsv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.shipper.insights.v1beta1.ShippingEmissionsService.
var (
	einride_shipper_insights_v1beta1_ShippingEmissionsServiceClient v1beta1.ShippingEmissionsServiceClient
	einride_shipper_insights_v1beta1_ShippingEmissionsService       = &cobra.Command{
		Use:   "einride.shipper.insights.v1beta1.ShippingEmissionsService",
		Short: "Shipping emission service.",
		Long:  "Shipping emission service.",
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
			response, err := einride_shipper_insights_v1beta1_ShippingEmissionsServiceClient.GetShipmentShippingEmission(cmd.Context(), &einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_insights_v1beta1_ShippingEmissionsServiceClient.BatchGetShipmentShippingEmissions(cmd.Context(), &einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShippingEmissionsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission)

	einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission.Flags().StringVar(&einride_shipper_insights_v1beta1_ShippingEmissionsService_GetShipmentShippingEmission_Request.Name, "name", "", "Resource name of the shipment shipping cost to retrieve.")
	einride_shipper_insights_v1beta1_ShippingEmissionsService.AddCommand(einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions)

	einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions.Flags().StringVar(&einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions_Request.Parent, "parent", "", "Resource name of the parent shipment shared by all shipment shipping costs.\nThe last segment of the parent resource name may be a wildcard to get\nshipment shipping costs for multiple shipments.")

	einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions.Flags().StringSliceVar(&einride_shipper_insights_v1beta1_ShippingEmissionsService_BatchGetShipmentShippingEmissions_Request.Names, "names", []string{}, "Resource names of the shipment shipping costs.")
}
