package prospectiveshipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/prospect/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_prospect_shipper_v1beta1_ProspectiveShipperService = &cobra.Command{
	Use: "einride.prospect.shipper.v1beta1.ProspectiveShipperService",
}

var einride_prospect_shipper_v1beta1_CreateProspectiveShipperRequest v1beta1.CreateProspectiveShipperRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper = &cobra.Command{
	Use: "CreateProspectiveShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipper")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_GetProspectiveShipperRequest v1beta1.GetProspectiveShipperRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper = &cobra.Command{
	Use: "GetProspectiveShipper",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipper")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_ListProspectiveShippersRequest v1beta1.ListProspectiveShippersRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers = &cobra.Command{
	Use: "ListProspectiveShippers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShippers")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_CreateProspectiveShipperFileRequest v1beta1.CreateProspectiveShipperFileRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile = &cobra.Command{
	Use: "CreateProspectiveShipperFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperFile")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_GetProspectiveShipperFileRequest v1beta1.GetProspectiveShipperFileRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile = &cobra.Command{
	Use: "GetProspectiveShipperFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperFile")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_GenerateFileUploadURIRequest v1beta1.GenerateFileUploadURIRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI = &cobra.Command{
	Use: "GenerateFileUploadURI",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GenerateFileUploadURI")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_CreateProspectiveShipperDatasetRequest v1beta1.CreateProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset = &cobra.Command{
	Use: "CreateProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.CreateProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_UpdateProspectiveShipperDatasetRequest v1beta1.UpdateProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset = &cobra.Command{
	Use: "UpdateProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.UpdateProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_ValidateProspectiveShipperDatasetRequest v1beta1.ValidateProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset = &cobra.Command{
	Use: "ValidateProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ValidateProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_IngestProspectiveShipperDatasetRequest v1beta1.IngestProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset = &cobra.Command{
	Use: "IngestProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.IngestProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_GetProspectiveShipperDatasetRequest v1beta1.GetProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset = &cobra.Command{
	Use: "GetProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.GetProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_ListProspectiveShipperDatasetsRequest v1beta1.ListProspectiveShipperDatasetsRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets = &cobra.Command{
	Use: "ListProspectiveShipperDatasets",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperDatasets")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_ArchiveProspectiveShipperDatasetRequest v1beta1.ArchiveProspectiveShipperDatasetRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset = &cobra.Command{
	Use: "ArchiveProspectiveShipperDataset",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ArchiveProspectiveShipperDataset")
		return nil
	},
}

var einride_prospect_shipper_v1beta1_ListProspectiveShipperValidationErrorsRequest v1beta1.ListProspectiveShipperValidationErrorsRequest
var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors = &cobra.Command{
	Use: "ListProspectiveShipperValidationErrors",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.prospect.shipper.v1beta1.ProspectiveShipperService.ListProspectiveShipperValidationErrors")
		return nil
	},
}

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
}
