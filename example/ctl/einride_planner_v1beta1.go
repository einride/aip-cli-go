package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_planner_v1beta1_DefaultScheduleService = &cobra.Command{
	Use: "einride.planner.v1beta1.DefaultSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.DefaultSchedule called")
	},
}
var einride_planner_v1beta1_PlannerSiteSettingsService = &cobra.Command{
	Use: "einride.planner.v1beta1.PlannerSiteSettings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.PlannerSiteSettings called")
	},
}
var einride_planner_v1beta1_ScheduleBookingService = &cobra.Command{
	Use: "einride.planner.v1beta1.ScheduleBooking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.ScheduleBooking called")
	},
}
var einride_planner_v1beta1_ShipmentAllocationService = &cobra.Command{
	Use: "einride.planner.v1beta1.ShipmentAllocation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.ShipmentAllocation called")
	},
}
var einride_planner_v1beta1_TransportInstallationService = &cobra.Command{
	Use: "einride.planner.v1beta1.TransportInstallation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.TransportInstallation called")
	},
}
var einride_planner_v1beta1_TransportScheduleService = &cobra.Command{
	Use: "einride.planner.v1beta1.TransportSchedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.TransportSchedule called")
	},
}
var einride_planner_v1beta1_VehiclePlanService = &cobra.Command{
	Use: "einride.planner.v1beta1.VehiclePlan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planner.v1beta1.VehiclePlan called")
	},
}

func init() {
	rootCmd.AddCommand(einride_planner_v1beta1_DefaultScheduleService)
	rootCmd.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService)
	rootCmd.AddCommand(einride_planner_v1beta1_ScheduleBookingService)
	rootCmd.AddCommand(einride_planner_v1beta1_ShipmentAllocationService)
	rootCmd.AddCommand(einride_planner_v1beta1_TransportInstallationService)
	rootCmd.AddCommand(einride_planner_v1beta1_TransportScheduleService)
	rootCmd.AddCommand(einride_planner_v1beta1_VehiclePlanService)
}
