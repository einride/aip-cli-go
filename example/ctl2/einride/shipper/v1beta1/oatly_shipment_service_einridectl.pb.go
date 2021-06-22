package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_OatlyShipmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.OatlyShipmentService",
}

var einride_shipper_v1beta1_NextShipmentRequest v1beta1.NextShipmentRequest
var einride_shipper_v1beta1_OatlyShipmentService_NextShipment = &cobra.Command{
	Use: "NextShipment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OatlyShipmentService.NextShipment")
		return nil
	},
}

func AddOatlyShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_OatlyShipmentService)
}

func init() {
	einride_shipper_v1beta1_OatlyShipmentService.AddCommand(einride_shipper_v1beta1_OatlyShipmentService_NextShipment)
}
