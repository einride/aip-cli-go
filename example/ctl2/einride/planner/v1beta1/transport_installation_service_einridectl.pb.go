package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_TransportInstallationService = &cobra.Command{
	Use: "einride.planner.v1beta1.TransportInstallationService",
}

var einride_planner_v1beta1_CreateTransportInstallationRequest v1beta1.CreateTransportInstallationRequest
var einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation = &cobra.Command{
	Use: "CreateTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportInstallationService.CreateTransportInstallation")
		return nil
	},
}

var einride_planner_v1beta1_GetTransportInstallationRequest v1beta1.GetTransportInstallationRequest
var einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation = &cobra.Command{
	Use: "GetTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportInstallationService.GetTransportInstallation")
		return nil
	},
}

var einride_planner_v1beta1_UpdateTransportInstallationRequest v1beta1.UpdateTransportInstallationRequest
var einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation = &cobra.Command{
	Use: "UpdateTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportInstallationService.UpdateTransportInstallation")
		return nil
	},
}

var einride_planner_v1beta1_SearchTransportInstallationsRequest v1beta1.SearchTransportInstallationsRequest
var einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations = &cobra.Command{
	Use: "SearchTransportInstallations",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportInstallationService.SearchTransportInstallations")
		return nil
	},
}

func AddTransportInstallationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportInstallationService)
}

func init() {
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations)
}
