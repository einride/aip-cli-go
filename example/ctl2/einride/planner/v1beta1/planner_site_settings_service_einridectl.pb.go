package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.PlannerSiteSettingsService.
var (
	einride_planner_v1beta1_PlannerSiteSettingsServiceClient v1beta1.PlannerSiteSettingsServiceClient
	einride_planner_v1beta1_PlannerSiteSettingsService       = &cobra.Command{
		Use: "einride.planner.v1beta1.PlannerSiteSettingsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_PlannerSiteSettingsServiceClient = v1beta1.NewPlannerSiteSettingsServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.PlannerSiteSettingsService.GetPlannerSiteSettings.
var (
	einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings_Request v1beta1.GetPlannerSiteSettingsRequest
	einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings         = &cobra.Command{
		Use: "GetPlannerSiteSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.GetPlannerSiteSettings")
			return nil
		},
	}
)

// einride.planner.v1beta1.PlannerSiteSettingsService.BatchGetPlannerSiteSettings.
var (
	einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings_Request v1beta1.BatchGetPlannerSiteSettingsRequest
	einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings         = &cobra.Command{
		Use: "BatchGetPlannerSiteSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.BatchGetPlannerSiteSettings")
			return nil
		},
	}
)

// einride.planner.v1beta1.PlannerSiteSettingsService.UpdatePlannerSiteSettings.
var (
	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request v1beta1.UpdatePlannerSiteSettingsRequest
	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings         = &cobra.Command{
		Use: "UpdatePlannerSiteSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.UpdatePlannerSiteSettings")
			return nil
		},
	}
)

func AddPlannerSiteSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService)
}

func init() {
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings)
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings)
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings)
}
