package prospectiveshipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/prospect/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient v1beta1.ProspectiveShipperServiceClient
	einride_prospect_shipper_v1beta1_ProspectiveShipperService       = &cobra.Command{
		Use: "einride.prospect.shipper.v1beta1.ProspectiveShipperService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient = v1beta1.NewProspectiveShipperServiceClient(conn)
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipper.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request v1beta1.CreateProspectiveShipperRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper         = &cobra.Command{
		Use: "CreateProspectiveShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.CreateProspectiveShipper(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipper.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper_Request v1beta1.GetProspectiveShipperRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper         = &cobra.Command{
		Use: "GetProspectiveShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.GetProspectiveShipper(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShippers.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers_Request v1beta1.ListProspectiveShippersRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers         = &cobra.Command{
		Use: "ListProspectiveShippers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.ListProspectiveShippers(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperFile.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request v1beta1.CreateProspectiveShipperFileRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile         = &cobra.Command{
		Use: "CreateProspectiveShipperFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.CreateProspectiveShipperFile(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperFile.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile_Request v1beta1.GetProspectiveShipperFileRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile         = &cobra.Command{
		Use: "GetProspectiveShipperFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.GetProspectiveShipperFile(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.GenerateFileUploadURI.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request v1beta1.GenerateFileUploadURIRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI         = &cobra.Command{
		Use: "GenerateFileUploadURI",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.GenerateFileUploadURI(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request v1beta1.CreateProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset         = &cobra.Command{
		Use: "CreateProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.CreateProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.UpdateProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request v1beta1.UpdateProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset         = &cobra.Command{
		Use: "UpdateProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.UpdateProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.ValidateProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset_Request v1beta1.ValidateProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset         = &cobra.Command{
		Use: "ValidateProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.ValidateProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.IngestProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset_Request v1beta1.IngestProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset         = &cobra.Command{
		Use: "IngestProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.IngestProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset_Request v1beta1.GetProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset         = &cobra.Command{
		Use: "GetProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.GetProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperDatasets.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets_Request v1beta1.ListProspectiveShipperDatasetsRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets         = &cobra.Command{
		Use: "ListProspectiveShipperDatasets",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.ListProspectiveShipperDatasets(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.ArchiveProspectiveShipperDataset.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset_Request v1beta1.ArchiveProspectiveShipperDatasetRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset         = &cobra.Command{
		Use: "ArchiveProspectiveShipperDataset",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.ArchiveProspectiveShipperDataset(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperValidationErrors.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors_Request v1beta1.ListProspectiveShipperValidationErrorsRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors         = &cobra.Command{
		Use: "ListProspectiveShipperValidationErrors",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.ListProspectiveShipperValidationErrors(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.SetIamPolicy.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy_Request v1.SetIamPolicyRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy         = &cobra.Command{
		Use: "SetIamPolicy",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.SetIamPolicy(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetIamPolicy.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy_Request v1.GetIamPolicyRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy         = &cobra.Command{
		Use: "GetIamPolicy",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.GetIamPolicy(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ProspectiveShipperService.TestIamPermissions.
var (
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions_Request v1.TestIamPermissionsRequest
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions         = &cobra.Command{
		Use: "TestIamPermissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_prospect_shipper_v1beta1_ProspectiveShipperServiceClient.TestIamPermissions(cmd.Context(), &einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddProspectiveShipperServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService)
}

func init() {
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request.ProspectiveShipper = new(v1beta1.ProspectiveShipper)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request.ProspectiveShipper.Name, "prospectiveShipper.name", "", "The resource name of the shipper.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request.ProspectiveShipper.DisplayName, "prospectiveShipper.displayName", "", "Display name of the shipper.\nFor example: \"Lemon Beverages Ltd.\"")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper_Request.ProspectiveShipperId, "prospectiveShipperId", "", "The ID to use for the shipper, which will become the final component of\nthe shipper's resource name.\n\nIf a shipper ID is not provided, a shipper ID will be automatically\ngenerated.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-zA-Z0-9]/.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper_Request.Name, "name", "", "Resource name of the shipper to retrieve.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request.ProspectiveShipperFile = new(v1beta1.ProspectiveShipperFile)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request.ProspectiveShipperFile.Name, "prospectiveShipperFile.name", "", "The resource name of the file.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request.ProspectiveShipperFile.DisplayName, "prospectiveShipperFile.displayName", "", "The display name of the file.\nFor example: \"einride-logo.png\"")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request.ProspectiveShipperFile.MimeType, "prospectiveShipperFile.mimeType", "", "MIME type of the file.\n\nMust be a valid IANA media type.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile_Request.Parent, "parent", "", "Resource name of the parent shipper.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile_Request.Name, "name", "", "Resource name of the shipper to retrieve.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request.Name, "name", "", "Resource name of the file to sign upload URI for.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request.Ttl = new(durationpb.Duration)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI.Flags().Int64Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request.Ttl.Seconds, "ttl.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI_Request.Ttl.Nanos, "ttl.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request.ProspectiveShipperDataset = new(v1beta1.ProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request.ProspectiveShipperDataset.Name, "prospectiveShipperDataset.name", "", "The resource name of the dataset.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request.ProspectiveShipperDataset.DisplayName, "prospectiveShipperDataset.displayName", "", "Display name of the dataset.\nFor example: \"First dump\"")
	// TODO: one-of: PowerBiImport
	// TODO: one-of: FmpIntegrationImport
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request.ProspectiveShipperDataset.Description, "prospectiveShipperDataset.description", "", "Description of the dataset.\nFor example: \"2019 data from Lemon Beverages Ltd.\"")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset_Request.Parent, "parent", "", "Resource name of the parent shipper.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.ProspectiveShipperDataset = new(v1beta1.ProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.ProspectiveShipperDataset.Name, "prospectiveShipperDataset.name", "", "The resource name of the dataset.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.ProspectiveShipperDataset.DisplayName, "prospectiveShipperDataset.displayName", "", "Display name of the dataset.\nFor example: \"First dump\"")
	// TODO: one-of: PowerBiImport
	// TODO: one-of: FmpIntegrationImport
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.ProspectiveShipperDataset.Description, "prospectiveShipperDataset.description", "", "Description of the dataset.\nFor example: \"2019 data from Lemon Beverages Ltd.\"")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset.Flags().StringSliceVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset_Request.Name, "name", "", "Resource name of the dataset to validate.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset_Request.Name, "name", "", "Resource name of the dataset to start ingestion for.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset_Request.Name, "name", "", "Resource name of the dataset to get.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets_Request.Parent, "parent", "", "Resource name of the parent shipper.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset_Request.Name, "name", "", "Resource name of the dataset to delete.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors_Request.Parent, "parent", "", "Resource name of the parent dataset.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy_Request.Policy = new(v1.Policy)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy_Request.Policy.Version, "policy.version", 10, "Specifies the format of the policy.\n\nValid values are 0, 1, and 3. Requests specifying an invalid value will be\nrejected.\n\nOperations affecting conditional bindings must specify version 3. This can\nbe either setting a conditional policy, modifying a conditional binding,\nor removing a binding (conditional or unconditional) from the stored\nconditional policy.\nOperations on non-conditional policies may specify any valid value or\nleave the field unset.\n\nIf no etag is provided in the call to `setIamPolicy`, version compliance\nchecks against the stored policy is skipped.")
	// TODO: list: Bindings message
	// TODO: bytes Etag
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy_Request.Options = new(v1.GetPolicyOptions)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy.Flags().Int32Var(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy_Request.Options.RequestedPolicyVersion, "options.requestedPolicyVersion", 10, "Optional. The policy format version to be returned.\n\nValid values are 0, 1, and 3. Requests specifying an invalid value will be\nrejected.\n\nRequests for policies with any conditional bindings must specify version 3.\nPolicies without any conditional bindings may specify any valid value or\nleave the field unset.")
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions)

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions.Flags().StringVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.")

	einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions.Flags().StringSliceVar(&einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions_Request.Permissions, "permissions", []string{}, "The set of permissions to check for the `resource`. Permissions with\nwildcards (such as '*' or 'storage.*') are not allowed. For more\ninformation see\n[IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).")
}
