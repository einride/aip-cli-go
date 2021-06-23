package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.CustomItemTypeService.
var (
	einride_shipper_v1beta1_CustomItemTypeServiceClient v1beta1.CustomItemTypeServiceClient
	einride_shipper_v1beta1_CustomItemTypeService       = &cobra.Command{
		Use:   "einride.shipper.v1beta1.CustomItemTypeService",
		Short: "Custom item type service.",
		Long:  "Custom item type service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_CustomItemTypeServiceClient = v1beta1.NewCustomItemTypeServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.GetCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType_Request v1beta1.GetCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType         = &cobra.Command{
		Use: "GetCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.GetCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.CreateCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request v1beta1.CreateCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType         = &cobra.Command{
		Use: "CreateCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.CreateCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.UpdateCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request v1beta1.UpdateCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType         = &cobra.Command{
		Use: "UpdateCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.UpdateCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.DeleteCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType_Request v1beta1.DeleteCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType         = &cobra.Command{
		Use: "DeleteCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.DeleteCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.UndeleteCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType_Request v1beta1.UndeleteCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType         = &cobra.Command{
		Use: "UndeleteCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.UndeleteCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.ListCustomItemTypes.
var (
	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request v1beta1.ListCustomItemTypesRequest
	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes         = &cobra.Command{
		Use: "ListCustomItemTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.ListCustomItemTypes(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.BatchGetCustomItemTypes.
var (
	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes_Request v1beta1.BatchGetCustomItemTypesRequest
	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes         = &cobra.Command{
		Use: "BatchGetCustomItemTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.BatchGetCustomItemTypes(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.CustomItemTypeService.ReferenceCustomItemType.
var (
	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType_Request v1beta1.ReferenceCustomItemTypeRequest
	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType         = &cobra.Command{
		Use: "ReferenceCustomItemType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_CustomItemTypeServiceClient.ReferenceCustomItemType(cmd.Context(), &einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddCustomItemTypeServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_CustomItemTypeService)
}

func init() {
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_GetCustomItemType_Request.Name, "name", "", "Resource name of the custom item type to retrieve.\nFormat: `shippers/{shipper}/customItemTypes/{custom_item_type}`.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.Parent, "parent", "", "Resource name of the parent shipper where this custom item type will be created.\nPattern: `shippers/{shipper}`.")

	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType = new(v1beta1.CustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.Name, "customItemType.name", "", "The resource name of the custom item type.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.Etag, "customItemType.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.ReferenceId, "customItemType.referenceId", "", "A unique reference ID that customer uses to refer to this\ncustom item type. Must be URL-safe.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().BoolVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.TemperatureDependent, "customItemType.temperatureDependent", false, "Flag indicating if the item type is temperature dependent.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().Float32Var(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.MinTemperatureCelsius, "customItemType.minTemperatureCelsius", 10, "Min allowed temperature of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().Float32Var(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.MaxTemperatureCelsius, "customItemType.maxTemperatureCelsius", 10, "Max allowed temperature of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().BoolVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.PhDependent, "customItemType.phDependent", false, "Flag indicating if the item type is pH-dependent.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().Int32Var(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.MinPh, "customItemType.minPh", 10, "Min allowed pH level of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().Int32Var(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemType.MaxPh, "customItemType.maxPh", 10, "Max allowed pH level of the item type.")

	einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_CreateCustomItemType_Request.CustomItemTypeId, "customItemTypeId", "", "The ID to use for the custom item type.\nWill become the final component of the custom item type's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.\nThe ID should be 3-63 characters and valid characters are /[1-9A-HJ-NP-Za-km-z]/.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType = new(v1beta1.CustomItemType)
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.Name, "customItemType.name", "", "The resource name of the custom item type.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.Etag, "customItemType.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.ReferenceId, "customItemType.referenceId", "", "A unique reference ID that customer uses to refer to this\ncustom item type. Must be URL-safe.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().BoolVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.TemperatureDependent, "customItemType.temperatureDependent", false, "Flag indicating if the item type is temperature dependent.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().Float32Var(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.MinTemperatureCelsius, "customItemType.minTemperatureCelsius", 10, "Min allowed temperature of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().Float32Var(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.MaxTemperatureCelsius, "customItemType.maxTemperatureCelsius", 10, "Max allowed temperature of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().BoolVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.PhDependent, "customItemType.phDependent", false, "Flag indicating if the item type is pH-dependent.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().Int32Var(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.MinPh, "customItemType.minPh", 10, "Min allowed pH level of the item type.")
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().Int32Var(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.CustomItemType.MaxPh, "customItemType.maxPh", 10, "Max allowed pH level of the item type.")

	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType.Flags().StringSliceVar(&einride_shipper_v1beta1_CustomItemTypeService_UpdateCustomItemType_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_DeleteCustomItemType_Request.Name, "name", "", "Resource name of the custom item type to delete.\nFormat: `shippers/{shipper}/customItemTypes/{custom_item_type}`.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_UndeleteCustomItemType_Request.Name, "name", "", "Resource name of the custom item type to undelete.\nFormat: `shippers/{shipper}/customItemTypes/{custom_item_type}`.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes)

	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request.Parent, "parent", "", "Resource name of the parent shipper.\nPattern: `shippers/{shipper}`.")

	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes.Flags().Int32Var(&einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request.PageSize, "pageSize", 10, "The maximum number of custom item types to return.\nThe service may return fewer custom item types than this value.\nIf unspecified, at most 50 custom item types will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_ListCustomItemTypes_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\nWhen paginating, all other parameters must match the call that provided the page token.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes)

	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all custom item types being retrieved.\nIf set, the parent of all custom item types specified in `names` must match this field.\nPattern: `shippers/{shipper}`.")

	einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes.Flags().StringSliceVar(&einride_shipper_v1beta1_CustomItemTypeService_BatchGetCustomItemTypes_Request.Names, "names", []string{}, "Resource names of the custom item types to retrieve.\nA maximum of 1000 custom item types can be retrieved in a batch.")
	einride_shipper_v1beta1_CustomItemTypeService.AddCommand(einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType)

	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType_Request.Parent, "parent", "", "Resource name of the shipper to look up the custom item type for.")

	einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType.Flags().StringVar(&einride_shipper_v1beta1_CustomItemTypeService_ReferenceCustomItemType_Request.ReferenceId, "referenceId", "", "Reference ID of the custom item type to look up.")
}
