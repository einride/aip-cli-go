package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_upload_v1beta1_FileService = &cobra.Command{
	Use: "einride.upload.v1beta1.File",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.upload.v1beta1.File called")
	},
}

var einride_upload_v1beta1_FileService_CreateFile = &cobra.Command{
	Use: "CreateFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateFile called")
	},
}

var einride_upload_v1beta1_FileService_GetFile = &cobra.Command{
	Use: "GetFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetFile called")
	},
}

var einride_upload_v1beta1_FileService_BatchGetFiles = &cobra.Command{
	Use: "BatchGetFiles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetFiles called")
	},
}

var einride_upload_v1beta1_FileService_ListFiles = &cobra.Command{
	Use: "ListFiles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListFiles called")
	},
}

var einride_upload_v1beta1_FileService_UploadFile = &cobra.Command{
	Use: "UploadFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UploadFile called")
	},
}

var einride_upload_v1beta1_FileService_DownloadFile = &cobra.Command{
	Use: "DownloadFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DownloadFile called")
	},
}

var einride_upload_v1beta1_FileService_BatchDownloadFiles = &cobra.Command{
	Use: "BatchDownloadFiles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchDownloadFiles called")
	},
}

func init() {
	rootCmd.AddCommand(einride_upload_v1beta1_FileService)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_CreateFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_GetFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchGetFiles)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_ListFiles)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_UploadFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_DownloadFile)
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchDownloadFiles)
}
