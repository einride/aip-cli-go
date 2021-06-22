package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.TransportInstallationService.
var (
	einride_planner_v1beta1_TransportInstallationServiceClient v1beta1.TransportInstallationServiceClient
	einride_planner_v1beta1_TransportInstallationService       = &cobra.Command{
		Use: "einride.planner.v1beta1.TransportInstallationService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_TransportInstallationServiceClient = v1beta1.NewTransportInstallationServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportInstallationService.CreateTransportInstallation.
var (
	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request v1beta1.CreateTransportInstallationRequest
	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation         = &cobra.Command{
		Use: "CreateTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportInstallationService.CreateTransportInstallation")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportInstallationService.GetTransportInstallation.
var (
	einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation_Request v1beta1.GetTransportInstallationRequest
	einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation         = &cobra.Command{
		Use: "GetTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportInstallationService.GetTransportInstallation")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportInstallationService.UpdateTransportInstallation.
var (
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request v1beta1.UpdateTransportInstallationRequest
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation         = &cobra.Command{
		Use: "UpdateTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportInstallationService.UpdateTransportInstallation")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportInstallationService.SearchTransportInstallations.
var (
	einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations_Request v1beta1.SearchTransportInstallationsRequest
	einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations         = &cobra.Command{
		Use: "SearchTransportInstallations",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportInstallationService.SearchTransportInstallations")
			return nil
		},
	}
)

func AddTransportInstallationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportInstallationService)
}

func init() {
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation)
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations)
}
