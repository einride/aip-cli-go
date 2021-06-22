package rescuev1alpha1

import (
	v1alpha1 "github.com/einride/proto/gen/go/einride/optimizer/rescue/v1alpha1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_optimizer_rescue_v1alpha1_RescueService = &cobra.Command{
	Use: "einride.optimizer.rescue.v1alpha1.RescueService",
}

var einride_optimizer_rescue_v1alpha1_ComputeVehiclePlansRequest v1alpha1.ComputeVehiclePlansRequest
var einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans = &cobra.Command{
	Use: "ComputeVehiclePlans",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.optimizer.rescue.v1alpha1.RescueService.ComputeVehiclePlans")
		return nil
	},
}

func AddRescueServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService)
}

func init() {
	einride_optimizer_rescue_v1alpha1_RescueService.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans)
}
