package integrationv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.integration.v1beta1.IntegrationFileService.
var (
	einride_shipper_integration_v1beta1_IntegrationFileServiceClient v1beta1.IntegrationFileServiceClient
	einride_shipper_integration_v1beta1_IntegrationFileService       = &cobra.Command{
		Use: "einride.shipper.integration.v1beta1.IntegrationFileService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_integration_v1beta1_IntegrationFileServiceClient = v1beta1.NewIntegrationFileServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.CreateIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request v1beta1.CreateIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile         = &cobra.Command{
		Use: "CreateIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.CreateIntegrationFile")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.GetIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile_Request v1beta1.GetIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile         = &cobra.Command{
		Use: "GetIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.GetIntegrationFile")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.BatchGetIntegrationFiles.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles_Request v1beta1.BatchGetIntegrationFilesRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles         = &cobra.Command{
		Use: "BatchGetIntegrationFiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.BatchGetIntegrationFiles")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.ListIntegrationFiles.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request v1beta1.ListIntegrationFilesRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles         = &cobra.Command{
		Use: "ListIntegrationFiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.ListIntegrationFiles")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.IngestIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile_Request v1beta1.IngestIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile         = &cobra.Command{
		Use: "IngestIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.IngestIntegrationFile")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.IngestAsyncIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile_Request v1beta1.IngestAsyncIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile         = &cobra.Command{
		Use: "IngestAsyncIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.IngestAsyncIntegrationFile")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.DownloadIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile_Request v1beta1.DownloadIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile         = &cobra.Command{
		Use: "DownloadIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.DownloadIntegrationFile")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationFileService.UploadIntegrationFile.
var (
	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request v1beta1.UploadIntegrationFileRequest
	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile         = &cobra.Command{
		Use: "UploadIntegrationFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationFileService.UploadIntegrationFile")
			return nil
		},
	}
)

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
