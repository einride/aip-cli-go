package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.shipper.v1beta1.OatlyShipmentService.
var (
	einride_shipper_v1beta1_OatlyShipmentServiceClient v1beta1.OatlyShipmentServiceClient
	einride_shipper_v1beta1_OatlyShipmentService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.OatlyShipmentService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_OatlyShipmentServiceClient = v1beta1.NewOatlyShipmentServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.OatlyShipmentService.NextShipment.
var (
	einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request v1beta1.NextShipmentRequest
	einride_shipper_v1beta1_OatlyShipmentService_NextShipment         = &cobra.Command{
		Use: "NextShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_OatlyShipmentServiceClient.NextShipment(cmd.Context(), &einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddOatlyShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_OatlyShipmentService)
}

func init() {
	einride_shipper_v1beta1_OatlyShipmentService.AddCommand(einride_shipper_v1beta1_OatlyShipmentService_NextShipment)

	einride_shipper_v1beta1_OatlyShipmentService_NextShipment.Flags().StringVar(&einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request.Parent, "parent", "", "Resource name of the shipper owning the next shipment to be retrieved.")

	einride_shipper_v1beta1_OatlyShipmentService_NextShipment.Flags().StringVar(&einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request.DestinationSite, "destinationSite", "", "Resource name of the destination site of the shipment to be retrieved.")

	einride_shipper_v1beta1_OatlyShipmentService_NextShipment.Flags().StringVar(&einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request.OriginSite, "originSite", "", "Resource name of the origin site of the shipment to be retrieved.")

	einride_shipper_v1beta1_OatlyShipmentService_NextShipment.Flags().Float32Var(&einride_shipper_v1beta1_OatlyShipmentService_NextShipment_Request.MaxLoadQuantity, "maxLoadQuantity", 10, "Don't consider shipments with a total load quantity above this value.\nIf not set, all shipments are considered.")

	// TODO: enum MaxLoadQuantityUnit
}
