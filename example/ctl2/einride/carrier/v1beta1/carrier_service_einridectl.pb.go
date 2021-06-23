package carrierv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.carrier.v1beta1.CarrierService.
var (
	einride_carrier_v1beta1_CarrierServiceClient v1beta1.CarrierServiceClient
	einride_carrier_v1beta1_CarrierService       = &cobra.Command{
		Use:   "einride.carrier.v1beta1.CarrierService",
		Short: "Carrier service.",
		Long:  "Carrier service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_carrier_v1beta1_CarrierServiceClient = v1beta1.NewCarrierServiceClient(conn)
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.CreateCarrier.
var (
	einride_carrier_v1beta1_CarrierService_CreateCarrier_Request v1beta1.CreateCarrierRequest
	einride_carrier_v1beta1_CarrierService_CreateCarrier         = &cobra.Command{
		Use: "CreateCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.CreateCarrier(cmd.Context(), &einride_carrier_v1beta1_CarrierService_CreateCarrier_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.ListCarriers.
var (
	einride_carrier_v1beta1_CarrierService_ListCarriers_Request v1beta1.ListCarriersRequest
	einride_carrier_v1beta1_CarrierService_ListCarriers         = &cobra.Command{
		Use: "ListCarriers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.ListCarriers(cmd.Context(), &einride_carrier_v1beta1_CarrierService_ListCarriers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.GetCarrier.
var (
	einride_carrier_v1beta1_CarrierService_GetCarrier_Request v1beta1.GetCarrierRequest
	einride_carrier_v1beta1_CarrierService_GetCarrier         = &cobra.Command{
		Use: "GetCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.GetCarrier(cmd.Context(), &einride_carrier_v1beta1_CarrierService_GetCarrier_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.BatchGetCarriers.
var (
	einride_carrier_v1beta1_CarrierService_BatchGetCarriers_Request v1beta1.BatchGetCarriersRequest
	einride_carrier_v1beta1_CarrierService_BatchGetCarriers         = &cobra.Command{
		Use: "BatchGetCarriers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.BatchGetCarriers(cmd.Context(), &einride_carrier_v1beta1_CarrierService_BatchGetCarriers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.UpdateCarrier.
var (
	einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request v1beta1.UpdateCarrierRequest
	einride_carrier_v1beta1_CarrierService_UpdateCarrier         = &cobra.Command{
		Use: "UpdateCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.UpdateCarrier(cmd.Context(), &einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.DeleteCarrier.
var (
	einride_carrier_v1beta1_CarrierService_DeleteCarrier_Request v1beta1.DeleteCarrierRequest
	einride_carrier_v1beta1_CarrierService_DeleteCarrier         = &cobra.Command{
		Use: "DeleteCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.DeleteCarrier(cmd.Context(), &einride_carrier_v1beta1_CarrierService_DeleteCarrier_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.CarrierService.UndeleteCarrier.
var (
	einride_carrier_v1beta1_CarrierService_UndeleteCarrier_Request v1beta1.UndeleteCarrierRequest
	einride_carrier_v1beta1_CarrierService_UndeleteCarrier         = &cobra.Command{
		Use: "UndeleteCarrier",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_CarrierServiceClient.UndeleteCarrier(cmd.Context(), &einride_carrier_v1beta1_CarrierService_UndeleteCarrier_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddCarrierServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_CarrierService)
}

func init() {
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_CreateCarrier)

	einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier = new(v1beta1.Carrier)
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Name, "carrier.name", "", "The resource name of the carrier.")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Etag, "carrier.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.DisplayName, "carrier.displayName", "", "Display name of the carrier.\nFor example: \"Johnson Trucking Co.\"")
	einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter = new(v1beta1.Address)
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter.DisplayName, "carrier.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter.Street, "carrier.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter.PostalCode, "carrier.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter.City, "carrier.headquarter.city", "", "For example: Göteborg.")
	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.Carrier.Headquarter.RegionCode, "carrier.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_carrier_v1beta1_CarrierService_CreateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_CreateCarrier_Request.CarrierId, "carrierId", "", "The ID to use for the carrier, which will become the final component of\nthe carrier's resource name.\n\nIf a carrier ID is not provided, a carrier ID will be automatically\ngenerated.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-zA-Z0-9]/.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_ListCarriers)

	einride_carrier_v1beta1_CarrierService_ListCarriers.Flags().Int32Var(&einride_carrier_v1beta1_CarrierService_ListCarriers_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_carrier_v1beta1_CarrierService_ListCarriers.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_ListCarriers_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_GetCarrier)

	einride_carrier_v1beta1_CarrierService_GetCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_GetCarrier_Request.Name, "name", "", "Resource name of the carrier to retrieve.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_BatchGetCarriers)

	einride_carrier_v1beta1_CarrierService_BatchGetCarriers.Flags().StringSliceVar(&einride_carrier_v1beta1_CarrierService_BatchGetCarriers_Request.Names, "names", []string{}, "Resource names of the carriers to retrieve.\nA maximum of 1000 carriers can be retrieved in a batch.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UpdateCarrier)

	einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier = new(v1beta1.Carrier)
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Name, "carrier.name", "", "The resource name of the carrier.")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Etag, "carrier.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.DisplayName, "carrier.displayName", "", "Display name of the carrier.\nFor example: \"Johnson Trucking Co.\"")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter = new(v1beta1.Address)
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter.DisplayName, "carrier.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter.Street, "carrier.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter.PostalCode, "carrier.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter.City, "carrier.headquarter.city", "", "For example: Göteborg.")
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.Carrier.Headquarter.RegionCode, "carrier.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_carrier_v1beta1_CarrierService_UpdateCarrier.Flags().StringSliceVar(&einride_carrier_v1beta1_CarrierService_UpdateCarrier_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_DeleteCarrier)

	einride_carrier_v1beta1_CarrierService_DeleteCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_DeleteCarrier_Request.Name, "name", "", "Resource name of the carrier to delete.")
	einride_carrier_v1beta1_CarrierService.AddCommand(einride_carrier_v1beta1_CarrierService_UndeleteCarrier)

	einride_carrier_v1beta1_CarrierService_UndeleteCarrier.Flags().StringVar(&einride_carrier_v1beta1_CarrierService_UndeleteCarrier_Request.Name, "name", "", "Resource name of the carrier to undelete.")
}
