package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.SiteChargerService.
var (
	einride_shipper_v1beta1_SiteChargerServiceClient v1beta1.SiteChargerServiceClient
	einride_shipper_v1beta1_SiteChargerService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.SiteChargerService",
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.GetSiteCharger")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.CreateSiteCharger")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.UpdateSiteCharger")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.DeleteSiteCharger")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.UndeleteSiteCharger")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.ListSiteChargers")
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
			log.Println("einride.shipper.v1beta1.SiteChargerService.BatchGetSiteChargers")
			return nil
		},
	}
)

func AddSiteChargerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteChargerService)
}

func init() {
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_GetSiteCharger)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_ListSiteChargers)
	einride_shipper_v1beta1_SiteChargerService.AddCommand(einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers)
}
