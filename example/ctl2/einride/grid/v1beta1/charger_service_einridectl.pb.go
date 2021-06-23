package gridv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/grid/v1beta1"
	cobra "github.com/spf13/cobra"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.grid.v1beta1.ChargerService.
var (
	einride_grid_v1beta1_ChargerServiceClient v1beta1.ChargerServiceClient
	einride_grid_v1beta1_ChargerService       = &cobra.Command{
		Use:   "grid.v1beta1.ChargerService",
		Short: "Charger service.",
		Long:  "Charger service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_grid_v1beta1_ChargerServiceClient = v1beta1.NewChargerServiceClient(conn)
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.CreateChargingPool.
var (
	einride_grid_v1beta1_ChargerService_CreateChargingPool_Request v1beta1.CreateChargingPoolRequest
	einride_grid_v1beta1_ChargerService_CreateChargingPool         = &cobra.Command{
		Use: "CreateChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.CreateChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_CreateChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ListChargingPools.
var (
	einride_grid_v1beta1_ChargerService_ListChargingPools_Request v1beta1.ListChargingPoolsRequest
	einride_grid_v1beta1_ChargerService_ListChargingPools         = &cobra.Command{
		Use: "ListChargingPools",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.ListChargingPools(cmd.Context(), &einride_grid_v1beta1_ChargerService_ListChargingPools_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.GetChargingPool.
var (
	einride_grid_v1beta1_ChargerService_GetChargingPool_Request v1beta1.GetChargingPoolRequest
	einride_grid_v1beta1_ChargerService_GetChargingPool         = &cobra.Command{
		Use: "GetChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.GetChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_GetChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ReferenceChargingPool.
var (
	einride_grid_v1beta1_ChargerService_ReferenceChargingPool_Request v1beta1.ReferenceChargingPoolRequest
	einride_grid_v1beta1_ChargerService_ReferenceChargingPool         = &cobra.Command{
		Use: "ReferenceChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.ReferenceChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_ReferenceChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.BatchGetChargingPools.
var (
	einride_grid_v1beta1_ChargerService_BatchGetChargingPools_Request v1beta1.BatchGetChargingPoolsRequest
	einride_grid_v1beta1_ChargerService_BatchGetChargingPools         = &cobra.Command{
		Use: "BatchGetChargingPools",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.BatchGetChargingPools(cmd.Context(), &einride_grid_v1beta1_ChargerService_BatchGetChargingPools_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UpdateChargingPool.
var (
	einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request v1beta1.UpdateChargingPoolRequest
	einride_grid_v1beta1_ChargerService_UpdateChargingPool         = &cobra.Command{
		Use: "UpdateChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.UpdateChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.DeleteChargingPool.
var (
	einride_grid_v1beta1_ChargerService_DeleteChargingPool_Request v1beta1.DeleteChargingPoolRequest
	einride_grid_v1beta1_ChargerService_DeleteChargingPool         = &cobra.Command{
		Use: "DeleteChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.DeleteChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_DeleteChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UndeleteChargingPool.
var (
	einride_grid_v1beta1_ChargerService_UndeleteChargingPool_Request v1beta1.UndeleteChargingPoolRequest
	einride_grid_v1beta1_ChargerService_UndeleteChargingPool         = &cobra.Command{
		Use: "UndeleteChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.UndeleteChargingPool(cmd.Context(), &einride_grid_v1beta1_ChargerService_UndeleteChargingPool_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.GetCharger.
var (
	einride_grid_v1beta1_ChargerService_GetCharger_Request v1beta1.GetChargerRequest
	einride_grid_v1beta1_ChargerService_GetCharger         = &cobra.Command{
		Use: "GetCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.GetCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_GetCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ReferenceCharger.
var (
	einride_grid_v1beta1_ChargerService_ReferenceCharger_Request v1beta1.ReferenceChargerRequest
	einride_grid_v1beta1_ChargerService_ReferenceCharger         = &cobra.Command{
		Use: "ReferenceCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.ReferenceCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_ReferenceCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.CreateCharger.
var (
	einride_grid_v1beta1_ChargerService_CreateCharger_Request v1beta1.CreateChargerRequest
	einride_grid_v1beta1_ChargerService_CreateCharger         = &cobra.Command{
		Use: "CreateCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.CreateCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_CreateCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UpdateCharger.
var (
	einride_grid_v1beta1_ChargerService_UpdateCharger_Request v1beta1.UpdateChargerRequest
	einride_grid_v1beta1_ChargerService_UpdateCharger         = &cobra.Command{
		Use: "UpdateCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.UpdateCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_UpdateCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.DeleteCharger.
var (
	einride_grid_v1beta1_ChargerService_DeleteCharger_Request v1beta1.DeleteChargerRequest
	einride_grid_v1beta1_ChargerService_DeleteCharger         = &cobra.Command{
		Use: "DeleteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.DeleteCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_DeleteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UndeleteCharger.
var (
	einride_grid_v1beta1_ChargerService_UndeleteCharger_Request v1beta1.UndeleteChargerRequest
	einride_grid_v1beta1_ChargerService_UndeleteCharger         = &cobra.Command{
		Use: "UndeleteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.UndeleteCharger(cmd.Context(), &einride_grid_v1beta1_ChargerService_UndeleteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ListChargers.
var (
	einride_grid_v1beta1_ChargerService_ListChargers_Request v1beta1.ListChargersRequest
	einride_grid_v1beta1_ChargerService_ListChargers         = &cobra.Command{
		Use: "ListChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.ListChargers(cmd.Context(), &einride_grid_v1beta1_ChargerService_ListChargers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.BatchGetChargers.
var (
	einride_grid_v1beta1_ChargerService_BatchGetChargers_Request v1beta1.BatchGetChargersRequest
	einride_grid_v1beta1_ChargerService_BatchGetChargers         = &cobra.Command{
		Use: "BatchGetChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_grid_v1beta1_ChargerServiceClient.BatchGetChargers(cmd.Context(), &einride_grid_v1beta1_ChargerService_BatchGetChargers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddChargerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_grid_v1beta1_ChargerService)
}

func init() {
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_CreateChargingPool)

	einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool = new(v1beta1.ChargingPool)
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.Name, "chargingPool.name", "", "The resource name of the charging pool.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.Etag, "chargingPool.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.EvsePoolId, "chargingPool.evsePoolId", "", "The charging pool's globally unique EVSE Pool ID.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.LatLng = new(latlng.LatLng)
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().Float64Var(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.LatLng.Latitude, "chargingPool.latLng.latitude", 10, "The latitude in degrees. It must be in the range [-90.0, +90.0].")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().Float64Var(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.LatLng.Longitude, "chargingPool.latLng.longitude", 10, "The longitude in degrees. It must be in the range [-180.0, +180.0].")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.DisplayName, "chargingPool.displayName", "", "The charging pool's display name.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.Address, "chargingPool.address", "", "The charging pool's address.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.City, "chargingPool.city", "", "The city the charging pool is located in.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.PostalCode, "chargingPool.postalCode", "", "The postal code of the charging pool.")
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.RegionCode, "chargingPool.regionCode", "", "The Unicode CLDR region code of the charging pool.")
	// TODO: enum ExternalApi
	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPool.ExternalApiReferenceId, "chargingPool.externalApiReferenceId", "", "Reference ID for this charging pool used by the external API.")

	einride_grid_v1beta1_ChargerService_CreateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateChargingPool_Request.ChargingPoolId, "chargingPoolId", "", "The ID to use for the charging pool, which will become the final component of\nthe charging pool's resource name.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ListChargingPools)

	einride_grid_v1beta1_ChargerService_ListChargingPools.Flags().Int32Var(&einride_grid_v1beta1_ChargerService_ListChargingPools_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.")

	einride_grid_v1beta1_ChargerService_ListChargingPools.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ListChargingPools_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_GetChargingPool)

	einride_grid_v1beta1_ChargerService_GetChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_GetChargingPool_Request.Name, "name", "", "Resource name of the charging pool to retrieve.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ReferenceChargingPool)

	// TODO: enum ExternalApi

	einride_grid_v1beta1_ChargerService_ReferenceChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ReferenceChargingPool_Request.ReferenceId, "referenceId", "", "External reference ID of the charging pool.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_BatchGetChargingPools)

	einride_grid_v1beta1_ChargerService_BatchGetChargingPools.Flags().StringSliceVar(&einride_grid_v1beta1_ChargerService_BatchGetChargingPools_Request.Names, "names", []string{}, "Resource names of the charging pools to retrieve.\nA maximum of 1000 charging pools can be retrieved in a batch.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UpdateChargingPool)

	einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool = new(v1beta1.ChargingPool)
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.Name, "chargingPool.name", "", "The resource name of the charging pool.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.Etag, "chargingPool.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.EvsePoolId, "chargingPool.evsePoolId", "", "The charging pool's globally unique EVSE Pool ID.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.LatLng = new(latlng.LatLng)
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().Float64Var(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.LatLng.Latitude, "chargingPool.latLng.latitude", 10, "The latitude in degrees. It must be in the range [-90.0, +90.0].")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().Float64Var(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.LatLng.Longitude, "chargingPool.latLng.longitude", 10, "The longitude in degrees. It must be in the range [-180.0, +180.0].")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.DisplayName, "chargingPool.displayName", "", "The charging pool's display name.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.Address, "chargingPool.address", "", "The charging pool's address.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.City, "chargingPool.city", "", "The city the charging pool is located in.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.PostalCode, "chargingPool.postalCode", "", "The postal code of the charging pool.")
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.RegionCode, "chargingPool.regionCode", "", "The Unicode CLDR region code of the charging pool.")
	// TODO: enum ExternalApi
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.ChargingPool.ExternalApiReferenceId, "chargingPool.externalApiReferenceId", "", "Reference ID for this charging pool used by the external API.")

	einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_grid_v1beta1_ChargerService_UpdateChargingPool.Flags().StringSliceVar(&einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_DeleteChargingPool)

	einride_grid_v1beta1_ChargerService_DeleteChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_DeleteChargingPool_Request.Name, "name", "", "Resource name of the charging pool to delete.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UndeleteChargingPool)

	einride_grid_v1beta1_ChargerService_UndeleteChargingPool.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UndeleteChargingPool_Request.Name, "name", "", "Resource name of the charging pool to undelete.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_GetCharger)

	einride_grid_v1beta1_ChargerService_GetCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_GetCharger_Request.Name, "name", "", "Resource name of the charger to retrieve.\nFormat: `chargingPools/{charging_pool}/chargers/{charger}`.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ReferenceCharger)

	einride_grid_v1beta1_ChargerService_ReferenceCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ReferenceCharger_Request.Parent, "parent", "", "Resource name of the parent charging pool to look up the charger for.\nPattern: `chargingPools/{charging_pool}`.")

	// TODO: enum ExternalApi

	einride_grid_v1beta1_ChargerService_ReferenceCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ReferenceCharger_Request.ChargingPoolReferenceId, "chargingPoolReferenceId", "", "External reference ID of the charging pool.")

	einride_grid_v1beta1_ChargerService_ReferenceCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ReferenceCharger_Request.ReferenceId, "referenceId", "", "External reference ID of the charger.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_CreateCharger)

	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Parent, "parent", "", "Resource name of the parent charging pool where this charger will be created.\nPattern: `chargingPools/{charging_pool}`.")

	einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger = new(v1beta1.Charger)
	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger.Name, "charger.name", "", "The resource name of the charger.")
	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger.Etag, "charger.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger.EvseId, "charger.evseId", "", "The charger's EVSE ID (Eletrical Vehicle Supply Equipment ID).\nThe EVSEID must be globally unique.")
	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger.DisplayName, "charger.displayName", "", "The charger's display name.")
	// TODO: list: Connectors message
	// TODO: enum ExternalApi
	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.Charger.ExternalApiReferenceId, "charger.externalApiReferenceId", "", "Reference ID for this charger used by the external API.")

	einride_grid_v1beta1_ChargerService_CreateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_CreateCharger_Request.ChargerId, "chargerId", "", "The ID to use for the charger.\nWill become the final component of the charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.\nThe ID should be 3-63 characters and valid characters are /[a-zA-Z0-9]/.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UpdateCharger)

	einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger = new(v1beta1.Charger)
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger.Name, "charger.name", "", "The resource name of the charger.")
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger.Etag, "charger.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger.EvseId, "charger.evseId", "", "The charger's EVSE ID (Eletrical Vehicle Supply Equipment ID).\nThe EVSEID must be globally unique.")
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger.DisplayName, "charger.displayName", "", "The charger's display name.")
	// TODO: list: Connectors message
	// TODO: enum ExternalApi
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.Charger.ExternalApiReferenceId, "charger.externalApiReferenceId", "", "Reference ID for this charger used by the external API.")

	einride_grid_v1beta1_ChargerService_UpdateCharger_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_grid_v1beta1_ChargerService_UpdateCharger.Flags().StringSliceVar(&einride_grid_v1beta1_ChargerService_UpdateCharger_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_DeleteCharger)

	einride_grid_v1beta1_ChargerService_DeleteCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_DeleteCharger_Request.Name, "name", "", "Resource name of the charger to delete.\nFormat: `chargingPools/{charging_pool}/chargers/{charger}`.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UndeleteCharger)

	einride_grid_v1beta1_ChargerService_UndeleteCharger.Flags().StringVar(&einride_grid_v1beta1_ChargerService_UndeleteCharger_Request.Name, "name", "", "Resource name of the charger to undelete.\nFormat: `chargingPools/{charging_pool}/chargers/{charger}`.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ListChargers)

	einride_grid_v1beta1_ChargerService_ListChargers.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ListChargers_Request.Parent, "parent", "", "Resource name of the parent charging pool.\nPattern: `chargingPools/{charging_pool}`.")

	einride_grid_v1beta1_ChargerService_ListChargers.Flags().Int32Var(&einride_grid_v1beta1_ChargerService_ListChargers_Request.PageSize, "pageSize", 10, "The maximum number of chargers to return.\nThe service may return fewer chargers than this value.\nIf unspecified, at most 50 chargers will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_grid_v1beta1_ChargerService_ListChargers.Flags().StringVar(&einride_grid_v1beta1_ChargerService_ListChargers_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\nWhen paginating, all other parameters must match the call that provided the page token.")
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_BatchGetChargers)

	einride_grid_v1beta1_ChargerService_BatchGetChargers.Flags().StringVar(&einride_grid_v1beta1_ChargerService_BatchGetChargers_Request.Parent, "parent", "", "Resource name of the parent charging pool shared by all chargers being retrieved.\nIf set, the parent of all chargers specified in `names` must match this field.\nPattern: `chargingPools/{charging_pool}`.")

	einride_grid_v1beta1_ChargerService_BatchGetChargers.Flags().StringSliceVar(&einride_grid_v1beta1_ChargerService_BatchGetChargers_Request.Names, "names", []string{}, "Resource names of the chargers to retrieve.\nA maximum of 1000 chargers can be retrieved in a batch.")
}
