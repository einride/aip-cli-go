package rescuev1alpha1

import (
	ctl "github.com/einride/ctl"
	v1alpha1 "github.com/einride/proto/gen/go/einride/optimizer/rescue/v1alpha1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.optimizer.rescue.v1alpha1.RescueService.
var (
	einride_optimizer_rescue_v1alpha1_RescueServiceClient v1alpha1.RescueServiceClient
	einride_optimizer_rescue_v1alpha1_RescueService       = &cobra.Command{
		Use: "einride.optimizer.rescue.v1alpha1.RescueService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_optimizer_rescue_v1alpha1_RescueServiceClient = v1alpha1.NewRescueServiceClient(conn)
			return nil
		},
	}
)

// einride.optimizer.rescue.v1alpha1.RescueService.ComputeVehiclePlans.
var (
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request v1alpha1.ComputeVehiclePlansRequest
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans         = &cobra.Command{
		Use: "ComputeVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.optimizer.rescue.v1alpha1.RescueService.ComputeVehiclePlans")
			return nil
		},
	}
)

func AddRescueServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService)
}

func init() {
	einride_optimizer_rescue_v1alpha1_RescueService.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans)
}
