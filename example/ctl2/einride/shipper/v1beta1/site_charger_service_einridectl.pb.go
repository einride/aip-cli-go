package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.SiteChargerService.
var (
	einride_shipper_v1beta1_SiteChargerServiceClient v1beta1.SiteChargerServiceClient
	einride_shipper_v1beta1_SiteChargerService       = &cobra.Command{
		Use:   "einride.shipper.v1beta1.SiteChargerService",
		Short: "Site charger service.",
		Long:  "Site charger service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_SiteChargerServiceClient = v1beta1.NewSiteChargerServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.GetSiteCharger.
var (
	einride_shipper_v1beta1_SiteChargerService_GetSiteCharger_Request v1beta1.GetSiteChargerRequest
	einride_shipper_v1beta1_SiteChargerService_GetSiteCharger         = &cobra.Command{
		Use: "GetSiteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.GetSiteCharger(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_GetSiteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.CreateSiteCharger.
var (
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request v1beta1.CreateSiteChargerRequest
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger         = &cobra.Command{
		Use: "CreateSiteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.CreateSiteCharger(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.UpdateSiteCharger.
var (
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request v1beta1.UpdateSiteChargerRequest
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger         = &cobra.Command{
		Use: "UpdateSiteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.UpdateSiteCharger(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.DeleteSiteCharger.
var (
	einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger_Request v1beta1.DeleteSiteChargerRequest
	einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger         = &cobra.Command{
		Use: "DeleteSiteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.DeleteSiteCharger(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.UndeleteSiteCharger.
var (
	einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger_Request v1beta1.UndeleteSiteChargerRequest
	einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger         = &cobra.Command{
		Use: "UndeleteSiteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.UndeleteSiteCharger(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.ListSiteChargers.
var (
	einride_shipper_v1beta1_SiteChargerService_ListSiteChargers_Request v1beta1.ListSiteChargersRequest
	einride_shipper_v1beta1_SiteChargerService_ListSiteChargers         = &cobra.Command{
		Use: "ListSiteChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.ListSiteChargers(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_ListSiteChargers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteChargerService.BatchGetSiteChargers.
var (
	einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers_Request v1beta1.BatchGetSiteChargersRequest
	einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers         = &cobra.Command{
		Use: "BatchGetSiteChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteChargerServiceClient.BatchGetSiteChargers(cmd.Context(), &einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddSiteChargerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteChargerService)
}

func init() {
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_GetSiteCharger)

	einride_shipper_v1beta1_SiteChargerService_GetSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_GetSiteCharger_Request.Name, "name", "", "Resource name of the site charger to retrieve.\nFormat: `shippers/{shipper}/sites/{site}/chargers/{charger}`.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger)

	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.Parent, "parent", "", "Resource name of the parent site where this site charger will be created.\nPattern: `shippers/{shipper}/sites/{site}`.")

	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteCharger = new(v1beta1.SiteCharger)
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteCharger.Name, "siteCharger.name", "", "The resource name.")
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteCharger.Etag, "siteCharger.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteCharger.DisplayName, "siteCharger.displayName", "", "Free-text display name for the charger.")
	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().Float32Var(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteCharger.PowerWatts, "siteCharger.powerWatts", 10, "The power of the charger (W).")

	einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger_Request.SiteChargerId, "siteChargerId", "", "The ID to use for the site charger.\nWill become the final component of the site charger's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.\nThe ID should be 3-63 characters and valid characters are /[a-zA-Z0-9]/.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger)

	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.SiteCharger = new(v1beta1.SiteCharger)
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.SiteCharger.Name, "siteCharger.name", "", "The resource name.")
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.SiteCharger.Etag, "siteCharger.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.SiteCharger.DisplayName, "siteCharger.displayName", "", "Free-text display name for the charger.")
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger.Flags().Float32Var(&einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.SiteCharger.PowerWatts, "siteCharger.powerWatts", 10, "The power of the charger (W).")

	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger)

	einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger_Request.Name, "name", "", "Resource name of the site charger to delete.\nFormat: `shippers/{shipper}/sites/{site}/chargers/{charger}`.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger)

	einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger_Request.Name, "name", "", "Resource name of the site charger to undelete.\nFormat: `shippers/{shipper}/sites/{site}/chargers/{charger}`.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_ListSiteChargers)

	einride_shipper_v1beta1_SiteChargerService_ListSiteChargers.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_ListSiteChargers_Request.Parent, "parent", "", "Resource name of the parent site.\nPattern: `shippers/{shipper}/sites/{site}`.")

	einride_shipper_v1beta1_SiteChargerService_ListSiteChargers.Flags().Int32Var(&einride_shipper_v1beta1_SiteChargerService_ListSiteChargers_Request.PageSize, "pageSize", 10, "The maximum number of site chargers to return.\nThe service may return fewer site chargers than this value.\nIf unspecified, at most 50 site chargers will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_SiteChargerService_ListSiteChargers.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_ListSiteChargers_Request.PageToken, "pageToken", "", "A page token, received from a previous call.\nProvide this to retrieve the subsequent page.\nWhen paginating, all other parameters must match the call that provided the page token.")
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers)

	einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers.Flags().StringVar(&einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers_Request.Parent, "parent", "", "Resource name of the parent site shared by all site chargers being retrieved.\nIf set, the parent of all site chargers specified in `names` must match this field.\nPattern: `shippers/{shipper}/sites/{site}`.")

	einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers_Request.Names, "names", []string{}, "Resource names of the site chargers to retrieve.\nA maximum of 1000 site chargers can be retrieved in a batch.")
}
