package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.TransportScheduleService.
var (
	einride_planner_v1beta1_TransportScheduleServiceClient v1beta1.TransportScheduleServiceClient
	einride_planner_v1beta1_TransportScheduleService       = &cobra.Command{
		Use: "einride.planner.v1beta1.TransportScheduleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_TransportScheduleServiceClient = v1beta1.NewTransportScheduleServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.CreateTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request v1beta1.CreateTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule         = &cobra.Command{
		Use: "CreateTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.CreateTransportSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.GetTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule_Request v1beta1.GetTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule         = &cobra.Command{
		Use: "GetTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.GetTransportSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.UpdateTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request v1beta1.UpdateTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule         = &cobra.Command{
		Use: "UpdateTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.UpdateTransportSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.ListTransportSchedules.
var (
	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request v1beta1.ListTransportSchedulesRequest
	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules         = &cobra.Command{
		Use: "ListTransportSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.ListTransportSchedules")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.BatchGetTransportSchedules.
var (
	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules_Request v1beta1.BatchGetTransportSchedulesRequest
	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules         = &cobra.Command{
		Use: "BatchGetTransportSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.BatchGetTransportSchedules")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.DeleteTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule_Request v1beta1.DeleteTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule         = &cobra.Command{
		Use: "DeleteTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.DeleteTransportSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.UndeleteTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule_Request v1beta1.UndeleteTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule         = &cobra.Command{
		Use: "UndeleteTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.TransportScheduleService.UndeleteTransportSchedule")
			return nil
		},
	}
)

func AddTransportScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportScheduleService)
}

func init() {
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule)
}
