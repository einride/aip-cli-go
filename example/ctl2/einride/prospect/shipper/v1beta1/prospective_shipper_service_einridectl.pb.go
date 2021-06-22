package prospectiveshipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/prospect/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	log "log"
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipper")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipper")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShippers")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperFile")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperFile")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GenerateFileUploadURI")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.UpdateProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ValidateProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.IngestProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperDatasets")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ArchiveProspectiveShipperDataset")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperValidationErrors")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.SetIamPolicy")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetIamPolicy")
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
			log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.TestIamPermissions")
			return nil
		},
	}
)

func AddProspectiveShipperServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService)
}

func init() {
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_SetIamPolicy)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetIamPolicy)
	einride_prospect_shipper_v1beta1_ProspectiveShipperService.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService_TestIamPermissions)
}
