package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
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
			log.Println("einride.shipper.v1beta1.OatlyShipmentService.NextShipment")
			return nil
		},
	}
)

func AddOatlyShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_OatlyShipmentService)
}

func init() {
	einride_shipper_v1beta1_OatlyShipmentService.AddCommand(einride_shipper_v1beta1_OatlyShipmentService_NextShipment)
}
