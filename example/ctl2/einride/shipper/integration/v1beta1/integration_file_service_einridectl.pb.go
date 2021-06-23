package integrationv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/integration/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

// einride.shipper.integration.v1beta1.IntegrationFileService.
var (
	einride_shipper_integration_v1beta1_IntegrationFileServiceClient v1beta1.IntegrationFileServiceClient
	einride_shipper_integration_v1beta1_IntegrationFileService       = &cobra.Command{
		Use:   "einride.shipper.integration.v1beta1.IntegrationFileService",
		Short: "Transport data integration service.",
		Long:  "Transport data integration service.",
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.CreateIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.GetIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.BatchGetIntegrationFiles(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.ListIntegrationFiles(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.IngestIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.IngestAsyncIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.DownloadIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_integration_v1beta1_IntegrationFileServiceClient.UploadIntegrationFile(cmd.Context(), &einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddIntegrationFileServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService)
}

func init() {
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.Parent, "parent", "", "Resource name of the parent integration bucket where this integration file will be created.")

	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile = new(v1beta1.IntegrationFile)
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.Name, "integrationFile.name", "", "The resource name of the integration file.")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.Etag, "integrationFile.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.StorageObject, "integrationFile.storageObject", "", "The storage object path.\nFor example: \"2019011519537.xml\"")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringSliceVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.Shipments, "integrationFile.shipments", []string{}, "Resource names of the shipments successfully ingested from the file.")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringSliceVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.ShipmentReferenceIds, "integrationFile.shipmentReferenceIds", []string{}, "Shipment reference IDs in the file.\nSome shipment reference IDs may not have been successfully ingested.")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.WarningCount, "integrationFile.warningCount", 10, "Number of warnings encountered when the file was last ingested.")
	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFile.ErrorCount, "integrationFile.errorCount", 10, "Number of errors encountered when the file was last ingested.")

	einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_CreateIntegrationFile_Request.IntegrationFileId, "integrationFileId", "", "The ID to use for the integration file, which will become the final\ncomponent of the integration file type's resource name.\n\nIf a integration file ID is not provided, a integration file ID will\nbe automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9_.]/.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_GetIntegrationFile_Request.Name, "name", "", "Resource name of the integration file to retrieve.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles)

	einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles_Request.Parent, "parent", "", "Resource name of the parent integration bucket shared by all integration files being retrieved.\nIf this is set, the parent of all of the integration files specified in `names`\nmust match this field.")

	einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles.Flags().StringSliceVar(&einride_shipper_integration_v1beta1_IntegrationFileService_BatchGetIntegrationFiles_Request.Names, "names", []string{}, "Resource names of the integration files to retrieve.\nA maximum of 1000 integration files can be retrieved in a batch.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles)

	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request.Parent, "parent", "", "Resource name of the parent integration bucket.")

	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request.PageSize, "pageSize", 10, "The maximum number of files to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 files will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request.OrderBy, "orderBy", "", "How the results should be sorted. Presently, the only permitted\nvalues are:\n\n- `\"file asc\"` (default)\n- `\"create_time asc\"`\n- `\"create_time desc\"`\n- `\"update_time asc\"`\n- `\"update_time desc\"`\n\nIf not specified, the default ordering will be used.")

	einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_ListIntegrationFiles_Request.Filter, "filter", "", "A filter.\n\nCurrently the only supported format is:\n\n`storage_object = \"<string>\"`")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_IngestIntegrationFile_Request.Name, "name", "", "Resource name of the integration file to ingest.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_IngestAsyncIntegrationFile_Request.Name, "name", "", "Resource name of the integration file to ingest.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile_Request.Name, "name", "", "Resource name of the integration file to download.")

	einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile.Flags().Int64Var(&einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile_Request.OffsetStart, "offsetStart", 10, "Start of range to download (inclusive).")

	einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile.Flags().Int64Var(&einride_shipper_integration_v1beta1_IntegrationFileService_DownloadIntegrationFile_Request.OffsetEnd, "offsetEnd", 10, "End of range to download (exclusive), or the end of the file if not specified.")
	einride_shipper_integration_v1beta1_IntegrationFileService.AddCommand(einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile)

	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile.Flags().StringVar(&einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request.Name, "name", "", "Resource name of the file to sign upload URL for.")

	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request.Ttl = new(durationpb.Duration)
	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile.Flags().Int64Var(&einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request.Ttl.Seconds, "ttl.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile.Flags().Int32Var(&einride_shipper_integration_v1beta1_IntegrationFileService_UploadIntegrationFile_Request.Ttl.Nanos, "ttl.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
}
