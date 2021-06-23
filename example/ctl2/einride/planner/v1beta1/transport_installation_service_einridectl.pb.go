package plannerv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.planner.v1beta1.TransportInstallationService.
var (
	einride_planner_v1beta1_TransportInstallationServiceClient v1beta1.TransportInstallationServiceClient
	einride_planner_v1beta1_TransportInstallationService       = &cobra.Command{
		Use:   "planner.v1beta1.TransportInstallationService",
		Short: "Transport installation service.",
		Long:  "Transport installation service.",
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
			response, err := einride_planner_v1beta1_TransportInstallationServiceClient.CreateTransportInstallation(cmd.Context(), &einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planner_v1beta1_TransportInstallationServiceClient.GetTransportInstallation(cmd.Context(), &einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planner_v1beta1_TransportInstallationServiceClient.UpdateTransportInstallation(cmd.Context(), &einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planner_v1beta1_TransportInstallationServiceClient.SearchTransportInstallations(cmd.Context(), &einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddTransportInstallationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportInstallationService)
}

func init() {
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation)

	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request.TransportInstallation = new(v1beta1.TransportInstallation)
	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request.TransportInstallation.Name, "transportInstallation.name", "", "The resource name of the transport installation.")
	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request.TransportInstallation.Etag, "transportInstallation.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request.TransportInstallation.DisplayName, "transportInstallation.displayName", "", "Display name of the transport installation.\nE.g. \"Oatly Landskrona, Trensums & Seafrigo\".")

	einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_CreateTransportInstallation_Request.TransportInstallationId, "transportInstallationId", "", "The ID to use for the transport installation, which will become the final component\nof the transport installation's resource name.\n\nIf a transport installation ID is not provided, a transport installation ID will be\nautomatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation)

	einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_GetTransportInstallation_Request.Name, "name", "", "Resource name of the transport installation to retrieve.\nPattern: `transportInstallations/{transport_installation}`.")
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation)

	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.TransportInstallation = new(v1beta1.TransportInstallation)
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.TransportInstallation.Name, "transportInstallation.name", "", "The resource name of the transport installation.")
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.TransportInstallation.Etag, "transportInstallation.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.TransportInstallation.DisplayName, "transportInstallation.displayName", "", "Display name of the transport installation.\nE.g. \"Oatly Landskrona, Trensums & Seafrigo\".")

	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation.Flags().StringSliceVar(&einride_planner_v1beta1_TransportInstallationService_UpdateTransportInstallation_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_TransportInstallationService.AddCommand(einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations)

	einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations.Flags().Int32Var(&einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations.Flags().StringVar(&einride_planner_v1beta1_TransportInstallationService_SearchTransportInstallations_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
