package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	log "log"
)

// einride.planner.v1beta1.DefaultScheduleService.
var (
	einride_planner_v1beta1_DefaultScheduleServiceClient v1beta1.DefaultScheduleServiceClient
	einride_planner_v1beta1_DefaultScheduleService       = &cobra.Command{
		Use: "einride.planner.v1beta1.DefaultScheduleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_DefaultScheduleServiceClient = v1beta1.NewDefaultScheduleServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.GetDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule_Request v1beta1.GetDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule         = &cobra.Command{
		Use: "GetDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.GetDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.UpdateDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request v1beta1.UpdateDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule         = &cobra.Command{
		Use: "UpdateDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.UpdateDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule_Request v1beta1.BookDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule         = &cobra.Command{
		Use: "BookDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedules.
var (
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules_Request v1beta1.BookDefaultSchedulesRequest
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules         = &cobra.Command{
		Use: "BookDefaultSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedules")
			return nil
		},
	}
)

func AddDefaultScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_DefaultScheduleService)
}

func init() {
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule)

	einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule_Request.Name, "name", "", "Resource name of the default schedule to retrieve.\nPattern: `transportInstallations/{transport_installation}/defaultSchedule`.")
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule)

	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule = new(v1beta1.DefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.Name, "defaultSchedule.name", "", "The resource name of the default schedules.\nFor example: `\"transportInstallations/transportInstallation/defaultSchedule`.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.Etag, "defaultSchedule.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule = new(v1beta1.WeeklyDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Monday, "defaultSchedule.weeklyDefaultSchedule.monday", "", "Resource name of the schedule to execute on Mondays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Tuesday, "defaultSchedule.weeklyDefaultSchedule.tuesday", "", "Resource name of the schedule to execute on Tuesdays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Wednesday, "defaultSchedule.weeklyDefaultSchedule.wednesday", "", "Resource name of the schedule to execute on Wednesdays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Thursday, "defaultSchedule.weeklyDefaultSchedule.thursday", "", "Resource name of the schedule to execute on Thursdays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Friday, "defaultSchedule.weeklyDefaultSchedule.friday", "", "Resource name of the schedule to execute on Fridays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Saturday, "defaultSchedule.weeklyDefaultSchedule.saturday", "", "Resource name of the schedule to execute on Saturdays.")
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.DefaultSchedule.WeeklyDefaultSchedule.Sunday, "defaultSchedule.weeklyDefaultSchedule.sunday", "", "Resource name of the schedule to execute on Sundays.")

	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule.Flags().StringSliceVar(&einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule)

	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule.Flags().StringVar(&einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule_Request.Name, "name", "", "Resource name of the default schedules to apply.\nPattern: `transportInstallations/{transport_installation}/defaultSchedule`.")

	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule.Flags().Int32Var(&einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule_Request.Days, "days", 10, "How many days in the future to book the default schedules for.\nIf not specified, schedules will be booked 14 days ahead.\nThe maximum value for this field is 30 days.")
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules)

	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules.Flags().Int32Var(&einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules_Request.Days, "days", 10, "How many days in the future to book the default schedules for.\nIf not specified, schedules will be booked 14 days ahead.\nThe maximum value for this field is 30 days.")
}
