package plannerv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.planner.v1beta1.PlannerSiteSettingsService.
var (
	einride_planner_v1beta1_PlannerSiteSettingsServiceClient v1beta1.PlannerSiteSettingsServiceClient
	einride_planner_v1beta1_PlannerSiteSettingsService       = &cobra.Command{
		Use:   "einride.planner.v1beta1.PlannerSiteSettingsService",
		Short: "planner site settings service.",
		Long:  "planner site settings service.",
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
			response, err := einride_planner_v1beta1_PlannerSiteSettingsServiceClient.GetPlannerSiteSettings(cmd.Context(), &einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planner_v1beta1_PlannerSiteSettingsServiceClient.BatchGetPlannerSiteSettings(cmd.Context(), &einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planner_v1beta1_PlannerSiteSettingsServiceClient.UpdatePlannerSiteSettings(cmd.Context(), &einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddPlannerSiteSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService)
}

func init() {
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings)

	einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings.Flags().StringVar(&einride_planner_v1beta1_PlannerSiteSettingsService_GetPlannerSiteSettings_Request.Name, "name", "", "Resource name of the planner site settings to retrieve.")
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings)

	einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings.Flags().StringVar(&einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings_Request.Parent, "parent", "", "Resource name of the parent site shared by all planner site settings being\nretrieved.\n\nThe parent of all of the sites specified in `names` must match this field.")

	einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings.Flags().StringSliceVar(&einride_planner_v1beta1_PlannerSiteSettingsService_BatchGetPlannerSiteSettings_Request.Names, "names", []string{}, "Resource names of the planner site settings to retrieve.\nA maximum of 1000 resources can be retrieved in a batch.")
	einride_planner_v1beta1_PlannerSiteSettingsService.AddCommand(einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings)

	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request.PlannerSiteSettings = new(v1beta1.PlannerSiteSettings)
	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings.Flags().StringVar(&einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request.PlannerSiteSettings.Name, "plannerSiteSettings.name", "", "The resource name.")
	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings.Flags().StringVar(&einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request.PlannerSiteSettings.Etag, "plannerSiteSettings.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	// TODO: enum Mode

	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings.Flags().StringSliceVar(&einride_planner_v1beta1_PlannerSiteSettingsService_UpdatePlannerSiteSettings_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
}
