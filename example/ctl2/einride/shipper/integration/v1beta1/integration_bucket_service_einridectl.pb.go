package integrationv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.Parent, "parent", "", "Resource name of the parent shipper where this integration bucket will be created.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket = new(v1beta1.IntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket.Name, "integrationBucket.name", "", "The resource name of the integration bucket.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket.Etag, "integrationBucket.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket.StorageBucket, "integrationBucket.storageBucket", "", "The Cloud Storage bucket name.")
	// TODO: enum FileFormat
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket.DisplayName, "integrationBucket.displayName", "", "The display name of the bucket.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().BoolVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucket.AutomaticIngestion, "integrationBucket.automaticIngestion", false, "Flag indicating if files should be automatically ingested upon being added\nto the bucket.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_CreateIntegrationBucket_Request.IntegrationBucketId, "integrationBucketId", "", "The ID to use for the integration bucket, which will become the final\ncomponent of the integration bucket type's resource name.\n\nIf a integration bucket ID is not provided, a integration bucket ID will\nbe automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket)

	einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_GetIntegrationBucket_Request.Name, "name", "", "Resource name of of the integration bucket to retrieve.")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets)

	einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all integration buckets being retrieved.\nIf this is set, the parent of all of the integration buckets specified in `names`\nmust match this field.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets.Flags().StringSliceVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_BatchGetIntegrationBuckets_Request.Names, "names", []string{}, "Resource names of the integration buckets to retrieve.\nA maximum of 1000 integration buckets can be retrieved in a batch.")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket)

	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket = new(v1beta1.IntegrationBucket)
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket.Name, "integrationBucket.name", "", "The resource name of the integration bucket.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket.Etag, "integrationBucket.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket.StorageBucket, "integrationBucket.storageBucket", "", "The Cloud Storage bucket name.")
	// TODO: enum FileFormat
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket.DisplayName, "integrationBucket.displayName", "", "The display name of the bucket.")
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().BoolVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.IntegrationBucket.AutomaticIngestion, "integrationBucket.automaticIngestion", false, "Flag indicating if files should be automatically ingested upon being added\nto the bucket.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket.Flags().StringSliceVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UpdateIntegrationBucket_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets)

	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets_Request.Parent, "parent", "", "Resource name of the parent shipper.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_ListIntegrationBuckets_Request.Filter, "filter", "", "A filter.\n\nCurrently the only supported format is:\n\n`storage_bucket = \"<string>\"`")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket)

	einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_DeleteIntegrationBucket_Request.Name, "name", "", "Resource name of the integration bucket to delete.")
	einride_shipper_integration_v1beta1_IntegrationBucketService.AddCommand(einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket)

	einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationBucketService_UndeleteIntegrationBucket_Request.Name, "name", "", "Resource name of the integration bucket to undelete.")
}
