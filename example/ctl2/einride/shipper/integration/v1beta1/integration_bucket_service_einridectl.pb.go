package integrationv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_integration_v1beta1_IntegrationBucketService = &cobra.Command{
	Use: "einride.shipper.integration.v1beta1.IntegrationBucketService",
}

var einride_shipper_integration_v1beta1_CreateIntegrationBucketRequest v1beta1.CreateIntegrationBucketRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket = &cobra.Command{
	Use: "CreateIntegrationBucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.CreateIntegrationBucket")
		return nil
	},
}

var einride_shipper_integration_v1beta1_GetIntegrationBucketRequest v1beta1.GetIntegrationBucketRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket = &cobra.Command{
	Use: "GetIntegrationBucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.GetIntegrationBucket")
		return nil
	},
}

var einride_shipper_integration_v1beta1_BatchGetIntegrationBucketsRequest v1beta1.BatchGetIntegrationBucketsRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets = &cobra.Command{
	Use: "BatchGetIntegrationBuckets",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.BatchGetIntegrationBuckets")
		return nil
	},
}

var einride_shipper_integration_v1beta1_UpdateIntegrationBucketRequest v1beta1.UpdateIntegrationBucketRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket = &cobra.Command{
	Use: "UpdateIntegrationBucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.UpdateIntegrationBucket")
		return nil
	},
}

var einride_shipper_integration_v1beta1_ListIntegrationBucketsRequest v1beta1.ListIntegrationBucketsRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets = &cobra.Command{
	Use: "ListIntegrationBuckets",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.ListIntegrationBuckets")
		return nil
	},
}

var einride_shipper_integration_v1beta1_DeleteIntegrationBucketRequest v1beta1.DeleteIntegrationBucketRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket = &cobra.Command{
	Use: "DeleteIntegrationBucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.DeleteIntegrationBucket")
		return nil
	},
}

var einride_shipper_integration_v1beta1_UndeleteIntegrationBucketRequest v1beta1.UndeleteIntegrationBucketRequest
var einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket = &cobra.Command{
	Use: "UndeleteIntegrationBucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.UndeleteIntegrationBucket")
		return nil
	},
}

func AddIntegrationBucketServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket)
}
