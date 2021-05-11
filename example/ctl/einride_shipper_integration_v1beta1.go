package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_shipper_integration_v1beta1_IntegrationBucketService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationBucket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.integration.v1beta1.IntegrationBucket called")
	},
}
var einride_shipper_integration_v1beta1_IntegrationFileLogService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationFileLog",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.integration.v1beta1.IntegrationFileLog called")
	},
}
var einride_shipper_integration_v1beta1_IntegrationFileService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationFile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.integration.v1beta1.IntegrationFile called")
	},
}
var einride_shipper_integration_v1beta1_IntegrationLogService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationLog",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.integration.v1beta1.IntegrationLog called")
	},
}

func init() {
	rootCmd.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService)
	rootCmd.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileLogService)
	rootCmd.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService)
	rootCmd.AddCommand(einride_shipper_integration_v1beta1_IntegrationLogService)
}
