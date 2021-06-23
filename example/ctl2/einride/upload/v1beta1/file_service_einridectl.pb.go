package uploadv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/upload/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

// einride.upload.v1beta1.FileService.
var (
	einride_upload_v1beta1_FileServiceClient v1beta1.FileServiceClient
	einride_upload_v1beta1_FileService       = &cobra.Command{
		Use:   "einride.upload.v1beta1.FileService",
		Short: "An uploaded file.",
		Long:  "An uploaded file.",
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
			response, err := einride_upload_v1beta1_FileServiceClient.CreateFile(cmd.Context(), &einride_upload_v1beta1_FileService_CreateFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.GetFile(cmd.Context(), &einride_upload_v1beta1_FileService_GetFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.BatchGetFiles(cmd.Context(), &einride_upload_v1beta1_FileService_BatchGetFiles_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.ListFiles(cmd.Context(), &einride_upload_v1beta1_FileService_ListFiles_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.UploadFile(cmd.Context(), &einride_upload_v1beta1_FileService_UploadFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.DownloadFile(cmd.Context(), &einride_upload_v1beta1_FileService_DownloadFile_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_upload_v1beta1_FileServiceClient.BatchDownloadFiles(cmd.Context(), &einride_upload_v1beta1_FileService_BatchDownloadFiles_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddFileServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_upload_v1beta1_FileService)
}

func init() {
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_CreateFile)

	einride_upload_v1beta1_FileService_CreateFile.Flags().StringVar(&einride_upload_v1beta1_FileService_CreateFile_Request.Parent, "parent", "", "Resource name of the parent tenant where this file will be created.")

	einride_upload_v1beta1_FileService_CreateFile_Request.File = new(v1beta1.File)
	einride_upload_v1beta1_FileService_CreateFile.Flags().StringVar(&einride_upload_v1beta1_FileService_CreateFile_Request.File.Name, "file.name", "", "The resource name of the file.")
	einride_upload_v1beta1_FileService_CreateFile.Flags().StringVar(&einride_upload_v1beta1_FileService_CreateFile_Request.File.DisplayName, "file.displayName", "", "The display name of the file.\nFor example: \"einride-logo.png\"")
	einride_upload_v1beta1_FileService_CreateFile.Flags().StringVar(&einride_upload_v1beta1_FileService_CreateFile_Request.File.MimeType, "file.mimeType", "", "MIME type of the file.\n\nMust be a valid IANA media type.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_GetFile)

	einride_upload_v1beta1_FileService_GetFile.Flags().StringVar(&einride_upload_v1beta1_FileService_GetFile_Request.Name, "name", "", "Resource name of the file to retrieve.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchGetFiles)

	einride_upload_v1beta1_FileService_BatchGetFiles.Flags().StringVar(&einride_upload_v1beta1_FileService_BatchGetFiles_Request.Parent, "parent", "", "Resource name of the parent tenant shared by all files being retrieved.\nUse `\"tenants/-\"` to batch get files for multiple tenants.\nIf not a wildcard, the parent of all of the files specified in `names`\nmust match this field.")

	einride_upload_v1beta1_FileService_BatchGetFiles.Flags().StringSliceVar(&einride_upload_v1beta1_FileService_BatchGetFiles_Request.Names, "names", []string{}, "Resource names of the files to get.\nA maximum of 1000 files can be retrieved in a batch.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_ListFiles)

	einride_upload_v1beta1_FileService_ListFiles.Flags().StringVar(&einride_upload_v1beta1_FileService_ListFiles_Request.Parent, "parent", "", "Resource name of the parent tenant.")

	einride_upload_v1beta1_FileService_ListFiles.Flags().Int32Var(&einride_upload_v1beta1_FileService_ListFiles_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_upload_v1beta1_FileService_ListFiles.Flags().StringVar(&einride_upload_v1beta1_FileService_ListFiles_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_UploadFile)

	einride_upload_v1beta1_FileService_UploadFile.Flags().StringVar(&einride_upload_v1beta1_FileService_UploadFile_Request.Name, "name", "", "Resource name of the file to sign upload URI for.")

	einride_upload_v1beta1_FileService_UploadFile_Request.Ttl = new(durationpb.Duration)
	einride_upload_v1beta1_FileService_UploadFile.Flags().Int64Var(&einride_upload_v1beta1_FileService_UploadFile_Request.Ttl.Seconds, "ttl.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_upload_v1beta1_FileService_UploadFile.Flags().Int32Var(&einride_upload_v1beta1_FileService_UploadFile_Request.Ttl.Nanos, "ttl.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_DownloadFile)

	einride_upload_v1beta1_FileService_DownloadFile.Flags().StringVar(&einride_upload_v1beta1_FileService_DownloadFile_Request.Name, "name", "", "Resource name of the file to sign a download URI for.")

	einride_upload_v1beta1_FileService_DownloadFile_Request.Ttl = new(durationpb.Duration)
	einride_upload_v1beta1_FileService_DownloadFile.Flags().Int64Var(&einride_upload_v1beta1_FileService_DownloadFile_Request.Ttl.Seconds, "ttl.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_upload_v1beta1_FileService_DownloadFile.Flags().Int32Var(&einride_upload_v1beta1_FileService_DownloadFile_Request.Ttl.Nanos, "ttl.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_upload_v1beta1_FileService.AddCommand(einride_upload_v1beta1_FileService_BatchDownloadFiles)

	einride_upload_v1beta1_FileService_BatchDownloadFiles.Flags().StringVar(&einride_upload_v1beta1_FileService_BatchDownloadFiles_Request.Parent, "parent", "", "Resource name of the parent tenant shared by all files being retrieved.\nUse `\"tenants/-\"` to batch get files for multiple tenants.\nIf not a wildcard, the parent of all of the files specified in `names`\nmust match this field.")

	einride_upload_v1beta1_FileService_BatchDownloadFiles.Flags().StringSliceVar(&einride_upload_v1beta1_FileService_BatchDownloadFiles_Request.Names, "names", []string{}, "Resource names of the files to sign download URIs for.\nMaximum 1000 files can be signed in the same request.")

	einride_upload_v1beta1_FileService_BatchDownloadFiles_Request.Ttl = new(durationpb.Duration)
	einride_upload_v1beta1_FileService_BatchDownloadFiles.Flags().Int64Var(&einride_upload_v1beta1_FileService_BatchDownloadFiles_Request.Ttl.Seconds, "ttl.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_upload_v1beta1_FileService_BatchDownloadFiles.Flags().Int32Var(&einride_upload_v1beta1_FileService_BatchDownloadFiles_Request.Ttl.Nanos, "ttl.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
}
