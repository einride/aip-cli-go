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

func init() {
	rootCmd.AddCommand(einride_upload_v1beta1_FileService)
}
