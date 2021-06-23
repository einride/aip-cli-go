package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.ShipmentClaimService.
var (
	einride_shipper_v1beta1_ShipmentClaimServiceClient v1beta1.ShipmentClaimServiceClient
	einride_shipper_v1beta1_ShipmentClaimService       = &cobra.Command{
		Use:   "shipper.v1beta1.ShipmentClaimService",
		Short: "Shipment resource service.",
		Long:  "Shipment resource service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentClaimServiceClient = v1beta1.NewShipmentClaimServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.CreateShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request v1beta1.CreateShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim         = &cobra.Command{
		Use: "CreateShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentClaimServiceClient.CreateShipmentClaim(cmd.Context(), &einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.GetShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim_Request v1beta1.GetShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim         = &cobra.Command{
		Use: "GetShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentClaimServiceClient.GetShipmentClaim(cmd.Context(), &einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.UpdateShipmentClaim.
var (
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request v1beta1.UpdateShipmentClaimRequest
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim         = &cobra.Command{
		Use: "UpdateShipmentClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentClaimServiceClient.UpdateShipmentClaim(cmd.Context(), &einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentClaimService.ListShipmentClaims.
var (
	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request v1beta1.ListShipmentClaimsRequest
	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims         = &cobra.Command{
		Use: "ListShipmentClaims",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentClaimServiceClient.ListShipmentClaims(cmd.Context(), &einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipmentClaimServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentClaimService)
}

func init() {
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim)

	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.Parent, "parent", "", "Resource name of the parent shipment where this shipment claim will be created.")

	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.ShipmentClaim = new(v1beta1.ShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.ShipmentClaim.Name, "shipmentClaim.name", "", "The resource name of the shipment claim.")
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.ShipmentClaim.DisplayName, "shipmentClaim.displayName", "", "The display name of the claim.")
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.ShipmentClaim.Description, "shipmentClaim.description", "", "Free text description of the claim.")
	einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentClaimService_CreateShipmentClaim_Request.ShipmentClaim.Images, "shipmentClaim.images", []string{}, "Resource names of image files attached to the claim.")
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim)

	einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_GetShipmentClaim_Request.Name, "name", "", "Resource name of the shipment claim to retrieve.")
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim)

	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.ShipmentClaim = new(v1beta1.ShipmentClaim)
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.ShipmentClaim.Name, "shipmentClaim.name", "", "The resource name of the shipment claim.")
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.ShipmentClaim.DisplayName, "shipmentClaim.displayName", "", "The display name of the claim.")
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.ShipmentClaim.Description, "shipmentClaim.description", "", "Free text description of the claim.")
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.ShipmentClaim.Images, "shipmentClaim.images", []string{}, "Resource names of image files attached to the claim.")

	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentClaimService_UpdateShipmentClaim_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_ShipmentClaimService.AddCommand(einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims)

	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request.Parent, "parent", "", "Resource name of the parent shipment.")

	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims.Flags().StringVar(&einride_shipper_v1beta1_ShipmentClaimService_ListShipmentClaims_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
