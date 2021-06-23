package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.ShipperService.
var (
	einride_shipper_v1beta1_ShipperServiceClient v1beta1.ShipperServiceClient
	einride_shipper_v1beta1_ShipperService       = &cobra.Command{
		Use:   "shipper.v1beta1.ShipperService",
		Short: "Shipper service.",
		Long:  "Shipper service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipperServiceClient = v1beta1.NewShipperServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.CreateShipper.
var (
	einride_shipper_v1beta1_ShipperService_CreateShipper_Request v1beta1.CreateShipperRequest
	einride_shipper_v1beta1_ShipperService_CreateShipper         = &cobra.Command{
		Use: "CreateShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.CreateShipper(cmd.Context(), &einride_shipper_v1beta1_ShipperService_CreateShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.ListShippers.
var (
	einride_shipper_v1beta1_ShipperService_ListShippers_Request v1beta1.ListShippersRequest
	einride_shipper_v1beta1_ShipperService_ListShippers         = &cobra.Command{
		Use: "ListShippers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.ListShippers(cmd.Context(), &einride_shipper_v1beta1_ShipperService_ListShippers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.GetShipper.
var (
	einride_shipper_v1beta1_ShipperService_GetShipper_Request v1beta1.GetShipperRequest
	einride_shipper_v1beta1_ShipperService_GetShipper         = &cobra.Command{
		Use: "GetShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.GetShipper(cmd.Context(), &einride_shipper_v1beta1_ShipperService_GetShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.BatchGetShippers.
var (
	einride_shipper_v1beta1_ShipperService_BatchGetShippers_Request v1beta1.BatchGetShippersRequest
	einride_shipper_v1beta1_ShipperService_BatchGetShippers         = &cobra.Command{
		Use: "BatchGetShippers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.BatchGetShippers(cmd.Context(), &einride_shipper_v1beta1_ShipperService_BatchGetShippers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.UpdateShipper.
var (
	einride_shipper_v1beta1_ShipperService_UpdateShipper_Request v1beta1.UpdateShipperRequest
	einride_shipper_v1beta1_ShipperService_UpdateShipper         = &cobra.Command{
		Use: "UpdateShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.UpdateShipper(cmd.Context(), &einride_shipper_v1beta1_ShipperService_UpdateShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.DeleteShipper.
var (
	einride_shipper_v1beta1_ShipperService_DeleteShipper_Request v1beta1.DeleteShipperRequest
	einride_shipper_v1beta1_ShipperService_DeleteShipper         = &cobra.Command{
		Use: "DeleteShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.DeleteShipper(cmd.Context(), &einride_shipper_v1beta1_ShipperService_DeleteShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperService.UndeleteShipper.
var (
	einride_shipper_v1beta1_ShipperService_UndeleteShipper_Request v1beta1.UndeleteShipperRequest
	einride_shipper_v1beta1_ShipperService_UndeleteShipper         = &cobra.Command{
		Use: "UndeleteShipper",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipperServiceClient.UndeleteShipper(cmd.Context(), &einride_shipper_v1beta1_ShipperService_UndeleteShipper_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipperServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperService)
}

func init() {
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_CreateShipper)

	einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper = new(v1beta1.Shipper)
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Name, "shipper.name", "", "The resource name of the shipper.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Etag, "shipper.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.DisplayName, "shipper.displayName", "", "Display name of the shipper.\nFor example: \"Lemon Beverages Ltd.\"")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.DefaultTimeZone, "shipper.defaultTimeZone", "", "The shipper's default time zone. Used for e.g. stats aggregations.\nMust be on the IANA Time Zone Database format.\nDefaults to UTC if not specified.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.DefaultSite, "shipper.defaultSite", "", "Resource name of the shipper's default site.\nLikely a well-known site with plenty of shipping activity.")
	einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter = new(v1beta1.Address)
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter.DisplayName, "shipper.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter.Street, "shipper.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter.PostalCode, "shipper.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter.City, "shipper.headquarter.city", "", "For example: Göteborg.")
	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.Shipper.Headquarter.RegionCode, "shipper.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_shipper_v1beta1_ShipperService_CreateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_CreateShipper_Request.ShipperId, "shipperId", "", "The ID to use for the shipper, which will become the final component of\nthe shipper's resource name.\n\nIf a shipper ID is not provided, a shipper ID will be automatically\ngenerated.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-zA-Z0-9]/.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_ListShippers)

	einride_shipper_v1beta1_ShipperService_ListShippers.Flags().Int32Var(&einride_shipper_v1beta1_ShipperService_ListShippers_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_shipper_v1beta1_ShipperService_ListShippers.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_ListShippers_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_GetShipper)

	einride_shipper_v1beta1_ShipperService_GetShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_GetShipper_Request.Name, "name", "", "Resource name of the shipper to retrieve.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_BatchGetShippers)

	einride_shipper_v1beta1_ShipperService_BatchGetShippers.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipperService_BatchGetShippers_Request.Names, "names", []string{}, "Resource names of the shippers to retrieve.\nA maximum of 1000 shippers can be retrieved in a batch.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_UpdateShipper)

	einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper = new(v1beta1.Shipper)
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Name, "shipper.name", "", "The resource name of the shipper.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Etag, "shipper.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.DisplayName, "shipper.displayName", "", "Display name of the shipper.\nFor example: \"Lemon Beverages Ltd.\"")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.DefaultTimeZone, "shipper.defaultTimeZone", "", "The shipper's default time zone. Used for e.g. stats aggregations.\nMust be on the IANA Time Zone Database format.\nDefaults to UTC if not specified.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.DefaultSite, "shipper.defaultSite", "", "Resource name of the shipper's default site.\nLikely a well-known site with plenty of shipping activity.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter = new(v1beta1.Address)
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter.DisplayName, "shipper.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter.Street, "shipper.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter.PostalCode, "shipper.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter.City, "shipper.headquarter.city", "", "For example: Göteborg.")
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.Shipper.Headquarter.RegionCode, "shipper.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipperService_UpdateShipper.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipperService_UpdateShipper_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_DeleteShipper)

	einride_shipper_v1beta1_ShipperService_DeleteShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_DeleteShipper_Request.Name, "name", "", "Resource name of the shipper to delete.")
	einride_shipper_v1beta1_ShipperService.AddCommand(einride_shipper_v1beta1_ShipperService_UndeleteShipper)

	einride_shipper_v1beta1_ShipperService_UndeleteShipper.Flags().StringVar(&einride_shipper_v1beta1_ShipperService_UndeleteShipper_Request.Name, "name", "", "Resource name of the shipper to undelete.")
}
