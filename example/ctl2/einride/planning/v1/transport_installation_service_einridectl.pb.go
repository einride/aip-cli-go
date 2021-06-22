package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planning.v1.TransportInstallationService.
var (
	einride_planning_v1_TransportInstallationServiceClient v1.TransportInstallationServiceClient
	einride_planning_v1_TransportInstallationService       = &cobra.Command{
		Use: "einride.planning.v1.TransportInstallationService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_TransportInstallationServiceClient = v1.NewTransportInstallationServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.TransportInstallationService.CreateTransportInstallation.
var (
	einride_planning_v1_TransportInstallationService_CreateTransportInstallation_Request v1.CreateTransportInstallationRequest
	einride_planning_v1_TransportInstallationService_CreateTransportInstallation         = &cobra.Command{
		Use: "CreateTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TransportInstallationService.CreateTransportInstallation")
			return nil
		},
	}
)

// einride.planning.v1.TransportInstallationService.GetTransportInstallation.
var (
	einride_planning_v1_TransportInstallationService_GetTransportInstallation_Request v1.GetTransportInstallationRequest
	einride_planning_v1_TransportInstallationService_GetTransportInstallation         = &cobra.Command{
		Use: "GetTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TransportInstallationService.GetTransportInstallation")
			return nil
		},
	}
)

// einride.planning.v1.TransportInstallationService.UpdateTransportInstallation.
var (
	einride_planning_v1_TransportInstallationService_UpdateTransportInstallation_Request v1.UpdateTransportInstallationRequest
	einride_planning_v1_TransportInstallationService_UpdateTransportInstallation         = &cobra.Command{
		Use: "UpdateTransportInstallation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TransportInstallationService.UpdateTransportInstallation")
			return nil
		},
	}
)

// einride.planning.v1.TransportInstallationService.SearchTransportInstallations.
var (
	einride_planning_v1_TransportInstallationService_SearchTransportInstallations_Request v1.SearchTransportInstallationsRequest
	einride_planning_v1_TransportInstallationService_SearchTransportInstallations         = &cobra.Command{
		Use: "SearchTransportInstallations",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TransportInstallationService.SearchTransportInstallations")
			return nil
		},
	}
)

func AddTransportInstallationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_TransportInstallationService)
}

func init() {
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_CreateTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_GetTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_UpdateTransportInstallation)
	einride_planning_v1_TransportInstallationService.AddCommand(einride_planning_v1_TransportInstallationService_SearchTransportInstallations)
}
