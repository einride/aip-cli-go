package integrationv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_integration_v1beta1_IntegrationFileService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationFileService",
}

var einride_shipper_integration_v1beta1_CreateIntegrationFileRequest v1beta1.CreateIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile = &cobra.Command{
	Use: "CreateIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.CreateIntegrationFile")
		return nil
	},
}

var einride_shipper_integration_v1beta1_GetIntegrationFileRequest v1beta1.GetIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile = &cobra.Command{
	Use: "GetIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.GetIntegrationFile")
		return nil
	},
}

var einride_shipper_integration_v1beta1_BatchGetIntegrationFilesRequest v1beta1.BatchGetIntegrationFilesRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles = &cobra.Command{
	Use: "BatchGetIntegrationFiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.BatchGetIntegrationFiles")
		return nil
	},
}

var einride_shipper_integration_v1beta1_ListIntegrationFilesRequest v1beta1.ListIntegrationFilesRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles = &cobra.Command{
	Use: "ListIntegrationFiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.ListIntegrationFiles")
		return nil
	},
}

var einride_shipper_integration_v1beta1_IngestIntegrationFileRequest v1beta1.IngestIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile = &cobra.Command{
	Use: "IngestIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.IngestIntegrationFile")
		return nil
	},
}

var einride_shipper_integration_v1beta1_IngestAsyncIntegrationFileRequest v1beta1.IngestAsyncIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile = &cobra.Command{
	Use: "IngestAsyncIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.IngestAsyncIntegrationFile")
		return nil
	},
}

var einride_shipper_integration_v1beta1_DownloadIntegrationFileRequest v1beta1.DownloadIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile = &cobra.Command{
	Use: "DownloadIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.DownloadIntegrationFile")
		return nil
	},
}

var einride_shipper_integration_v1beta1_UploadIntegrationFileRequest v1beta1.UploadIntegrationFileRequest
var einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile = &cobra.Command{
	Use: "UploadIntegrationFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.UploadIntegrationFile")
		return nil
	},
}

func AddIntegrationFileServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile)
}
