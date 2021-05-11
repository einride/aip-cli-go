package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_prospect_shipper_v1beta1_ProspectiveShipperService = &cobra.Command{
	Use: "einride.prospect.shipper.v1beta1.ProspectiveShipper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.prospect.shipper.v1beta1.ProspectiveShipper called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipper = &cobra.Command{
	Use: "CreateProspectiveShipper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateProspectiveShipper called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipper = &cobra.Command{
	Use: "GetProspectiveShipper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetProspectiveShipper called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShippers = &cobra.Command{
	Use: "ListProspectiveShippers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListProspectiveShippers called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperFile = &cobra.Command{
	Use: "CreateProspectiveShipperFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateProspectiveShipperFile called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperFile = &cobra.Command{
	Use: "GetProspectiveShipperFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetProspectiveShipperFile called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GenerateFileUploadURI = &cobra.Command{
	Use: "GenerateFileUploadURI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GenerateFileUploadURI called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_CreateProspectiveShipperDataset = &cobra.Command{
	Use: "CreateProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_UpdateProspectiveShipperDataset = &cobra.Command{
	Use: "UpdateProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ValidateProspectiveShipperDataset = &cobra.Command{
	Use: "ValidateProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ValidateProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_IngestProspectiveShipperDataset = &cobra.Command{
	Use: "IngestProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("IngestProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_GetProspectiveShipperDataset = &cobra.Command{
	Use: "GetProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperDatasets = &cobra.Command{
	Use: "ListProspectiveShipperDatasets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListProspectiveShipperDatasets called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ArchiveProspectiveShipperDataset = &cobra.Command{
	Use: "ArchiveProspectiveShipperDataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ArchiveProspectiveShipperDataset called")
	},
}

var einride_prospect_shipper_v1beta1_ProspectiveShipperService_ListProspectiveShipperValidationErrors = &cobra.Command{
	Use: "ListProspectiveShipperValidationErrors",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListProspectiveShipperValidationErrors called")
	},
}

func init() {
	rootCmd.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService)
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
