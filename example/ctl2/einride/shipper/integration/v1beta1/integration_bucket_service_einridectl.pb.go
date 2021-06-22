package integrationv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketServiceClient v1beta1.IntegrationBucketServiceClient
	einride_shipper_integration_v1beta1_IntegrationBucketService       = &cobra.Command{
		Use: "einride.shipper.integration.v1beta1.IntegrationBucketService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_integration_v1beta1_IntegrationBucketServiceClient = v1beta1.NewIntegrationBucketServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.CreateIntegrationBucket.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request v1beta1.CreateIntegrationBucketRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket         = &cobra.Command{
		Use: "CreateIntegrationBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.CreateIntegrationBucket")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.GetIntegrationBucket.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket_Request v1beta1.GetIntegrationBucketRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket         = &cobra.Command{
		Use: "GetIntegrationBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.GetIntegrationBucket")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.BatchGetIntegrationBuckets.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets_Request v1beta1.BatchGetIntegrationBucketsRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets         = &cobra.Command{
		Use: "BatchGetIntegrationBuckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.BatchGetIntegrationBuckets")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.UpdateIntegrationBucket.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request v1beta1.UpdateIntegrationBucketRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket         = &cobra.Command{
		Use: "UpdateIntegrationBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.UpdateIntegrationBucket")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.ListIntegrationBuckets.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets_Request v1beta1.ListIntegrationBucketsRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets         = &cobra.Command{
		Use: "ListIntegrationBuckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.ListIntegrationBuckets")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.DeleteIntegrationBucket.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket_Request v1beta1.DeleteIntegrationBucketRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket         = &cobra.Command{
		Use: "DeleteIntegrationBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.DeleteIntegrationBucket")
			return nil
		},
	}
)

// einride.shipper.integration.v1beta1.IntegrationBucketService.UndeleteIntegrationBucket.
var (
	einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket_Request v1beta1.UndeleteIntegrationBucketRequest
	einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket         = &cobra.Command{
		Use: "UndeleteIntegrationBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.integration.v1beta1.IntegrationBucketService.UndeleteIntegrationBucket")
			return nil
		},
	}
)

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
