package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_PlannerSiteSettingsService = &cobra.Command{
	Use: "einride.planner.v1beta1.PlannerSiteSettingsService",
}

var einride_planner_v1beta1_GetPlannerSiteSettingsRequest v1beta1.GetPlannerSiteSettingsRequest
var einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings = &cobra.Command{
	Use: "GetPlannerSiteSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.GetPlannerSiteSettings")
		return nil
	},
}

var einride_planner_v1beta1_BatchGetPlannerSiteSettingsRequest v1beta1.BatchGetPlannerSiteSettingsRequest
var einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings = &cobra.Command{
	Use: "BatchGetPlannerSiteSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.BatchGetPlannerSiteSettings")
		return nil
	},
}

var einride_planner_v1beta1_UpdatePlannerSiteSettingsRequest v1beta1.UpdatePlannerSiteSettingsRequest
var einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings = &cobra.Command{
	Use: "UpdatePlannerSiteSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.PlannerSiteSettingsService.UpdatePlannerSiteSettings")
		return nil
	},
}

func AddPlannerSiteSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService)
}

func init() {
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings)
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings)
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings)
}
