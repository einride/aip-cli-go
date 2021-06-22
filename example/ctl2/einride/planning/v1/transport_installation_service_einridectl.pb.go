package planningv1

import (
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planning_v1_TransportInstallationService = &cobra.Command{
	Use: "einride.planning.v1.TransportInstallationService",
}

var einride_planning_v1_CreateTransportInstallationRequest v1.CreateTransportInstallationRequest
var einride_planning_v1_TransportInstallationService_CreateTransportInstallation = &cobra.Command{
	Use: "CreateTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TransportInstallationService.CreateTransportInstallation")
		return nil
	},
}

var einride_planning_v1_GetTransportInstallationRequest v1.GetTransportInstallationRequest
var einride_planning_v1_TransportInstallationService_GetTransportInstallation = &cobra.Command{
	Use: "GetTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TransportInstallationService.GetTransportInstallation")
		return nil
	},
}

var einride_planning_v1_UpdateTransportInstallationRequest v1.UpdateTransportInstallationRequest
var einride_planning_v1_TransportInstallationService_UpdateTransportInstallation = &cobra.Command{
	Use: "UpdateTransportInstallation",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TransportInstallationService.UpdateTransportInstallation")
		return nil
	},
}

var einride_planning_v1_SearchTransportInstallationsRequest v1.SearchTransportInstallationsRequest
var einride_planning_v1_TransportInstallationService_SearchTransportInstallations = &cobra.Command{
	Use: "SearchTransportInstallations",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TransportInstallationService.SearchTransportInstallations")
		return nil
	},
}

func AddTransportInstallationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_TransportInstallationService)
}

func init() {
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_CreateTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_GetTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_UpdateTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_SearchTransportInstallations)
}
