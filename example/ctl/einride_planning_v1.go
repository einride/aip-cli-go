package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_planning_v1_AllocateShipmentsService = &cobra.Command{
	Use: "einride.planning.v1.AllocateShipments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planning.v1.AllocateShipments called")
	},
}
var einride_planning_v1_ScheduleBookingService = &cobra.Command{
	Use: "einride.planning.v1.ScheduleBooking",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planning.v1.ScheduleBooking called")
	},
}
var einride_planning_v1_ScheduleService = &cobra.Command{
	Use: "einride.planning.v1.Schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planning.v1.Schedule called")
	},
}
var einride_planning_v1_TaskService = &cobra.Command{
	Use: "einride.planning.v1.Task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planning.v1.Task called")
	},
}
var einride_planning_v1_TransportInstallationService = &cobra.Command{
	Use: "einride.planning.v1.TransportInstallation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.planning.v1.TransportInstallation called")
	},
}

func init() {
	rootCmd.AddCommand(einride_planning_v1_AllocateShipmentsService)
	rootCmd.AddCommand(einride_planning_v1_ScheduleBookingService)
	rootCmd.AddCommand(einride_planning_v1_ScheduleService)
	rootCmd.AddCommand(einride_planning_v1_TaskService)
	rootCmd.AddCommand(einride_planning_v1_TransportInstallationService)
}
