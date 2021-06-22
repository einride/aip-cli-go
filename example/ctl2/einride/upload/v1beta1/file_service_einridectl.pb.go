package uploadv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/upload/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.upload.v1beta1.FileService.
var (
	einride_upload_v1beta1_FileServiceClient v1beta1.FileServiceClient
	einride_upload_v1beta1_FileService       = &cobra.Command{
		Use: "einride.upload.v1beta1.FileService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_upload_v1beta1_FileServiceClient = v1beta1.NewFileServiceClient(conn)
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.CreateFile.
var (
	einride_upload_v1beta1_FileService_CreateFile_Request v1beta1.CreateFileRequest
	einride_upload_v1beta1_FileService_CreateFile         = &cobra.Command{
		Use: "CreateFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.CreateFile")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.GetFile.
var (
	einride_upload_v1beta1_FileService_GetFile_Request v1beta1.GetFileRequest
	einride_upload_v1beta1_FileService_GetFile         = &cobra.Command{
		Use: "GetFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.GetFile")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.BatchGetFiles.
var (
	einride_upload_v1beta1_FileService_BatchGetFiles_Request v1beta1.BatchGetFilesRequest
	einride_upload_v1beta1_FileService_BatchGetFiles         = &cobra.Command{
		Use: "BatchGetFiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.BatchGetFiles")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.ListFiles.
var (
	einride_upload_v1beta1_FileService_ListFiles_Request v1beta1.ListFilesRequest
	einride_upload_v1beta1_FileService_ListFiles         = &cobra.Command{
		Use: "ListFiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.ListFiles")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.UploadFile.
var (
	einride_upload_v1beta1_FileService_UploadFile_Request v1beta1.UploadFileRequest
	einride_upload_v1beta1_FileService_UploadFile         = &cobra.Command{
		Use: "UploadFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.UploadFile")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.DownloadFile.
var (
	einride_upload_v1beta1_FileService_DownloadFile_Request v1beta1.DownloadFileRequest
	einride_upload_v1beta1_FileService_DownloadFile         = &cobra.Command{
		Use: "DownloadFile",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.DownloadFile")
			return nil
		},
	}
)

// einride.upload.v1beta1.FileService.BatchDownloadFiles.
var (
	einride_upload_v1beta1_FileService_BatchDownloadFiles_Request v1beta1.BatchDownloadFilesRequest
	einride_upload_v1beta1_FileService_BatchDownloadFiles         = &cobra.Command{
		Use: "BatchDownloadFiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.upload.v1beta1.FileService.BatchDownloadFiles")
			return nil
		},
	}
)

func AddFileServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_upload_v1beta1_FileService)
}

func init() {
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_CreateFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_GetFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchGetFiles)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_ListFiles)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_UploadFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_DownloadFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchDownloadFiles)
}
