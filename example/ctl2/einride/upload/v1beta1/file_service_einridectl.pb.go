package uploadv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/upload/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_upload_v1beta1_FileService = &cobra.Command{
	Use: "einride.upload.v1beta1.FileService",
}

var einride_upload_v1beta1_CreateFileRequest v1beta1.CreateFileRequest
var einride_upload_v1beta1_FileService_CreateFile = &cobra.Command{
	Use: "CreateFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.CreateFile")
		return nil
	},
}

var einride_upload_v1beta1_GetFileRequest v1beta1.GetFileRequest
var einride_upload_v1beta1_FileService_GetFile = &cobra.Command{
	Use: "GetFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.GetFile")
		return nil
	},
}

var einride_upload_v1beta1_BatchGetFilesRequest v1beta1.BatchGetFilesRequest
var einride_upload_v1beta1_FileService_BatchGetFiles = &cobra.Command{
	Use: "BatchGetFiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.BatchGetFiles")
		return nil
	},
}

var einride_upload_v1beta1_ListFilesRequest v1beta1.ListFilesRequest
var einride_upload_v1beta1_FileService_ListFiles = &cobra.Command{
	Use: "ListFiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.ListFiles")
		return nil
	},
}

var einride_upload_v1beta1_UploadFileRequest v1beta1.UploadFileRequest
var einride_upload_v1beta1_FileService_UploadFile = &cobra.Command{
	Use: "UploadFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.UploadFile")
		return nil
	},
}

var einride_upload_v1beta1_DownloadFileRequest v1beta1.DownloadFileRequest
var einride_upload_v1beta1_FileService_DownloadFile = &cobra.Command{
	Use: "DownloadFile",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.DownloadFile")
		return nil
	},
}

var einride_upload_v1beta1_BatchDownloadFilesRequest v1beta1.BatchDownloadFilesRequest
var einride_upload_v1beta1_FileService_BatchDownloadFiles = &cobra.Command{
	Use: "BatchDownloadFiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.upload.v1beta1.FileService.BatchDownloadFiles")
		return nil
	},
}

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
