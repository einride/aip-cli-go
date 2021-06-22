package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_SiteChargerService = &cobra.Command{
	Use: "einride.shipper.v1beta1.SiteChargerService",
}

var einride_shipper_v1beta1_GetSiteChargerRequest v1beta1.GetSiteChargerRequest
var einride_shipper_v1beta1_SiteChargerService_GetSiteCharger = &cobra.Command{
	Use: "GetSiteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.GetSiteCharger")
		return nil
	},
}

var einride_shipper_v1beta1_CreateSiteChargerRequest v1beta1.CreateSiteChargerRequest
var einride_shipper_v1beta1_SiteChargerService_CreateSiteCharger = &cobra.Command{
	Use: "CreateSiteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.CreateSiteCharger")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateSiteChargerRequest v1beta1.UpdateSiteChargerRequest
var einride_shipper_v1beta1_SiteChargerService_UpdateSiteCharger = &cobra.Command{
	Use: "UpdateSiteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.UpdateSiteCharger")
		return nil
	},
}

var einride_shipper_v1beta1_DeleteSiteChargerRequest v1beta1.DeleteSiteChargerRequest
var einride_shipper_v1beta1_SiteChargerService_DeleteSiteCharger = &cobra.Command{
	Use: "DeleteSiteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.DeleteSiteCharger")
		return nil
	},
}

var einride_shipper_v1beta1_UndeleteSiteChargerRequest v1beta1.UndeleteSiteChargerRequest
var einride_shipper_v1beta1_SiteChargerService_UndeleteSiteCharger = &cobra.Command{
	Use: "UndeleteSiteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.UndeleteSiteCharger")
		return nil
	},
}

var einride_shipper_v1beta1_ListSiteChargersRequest v1beta1.ListSiteChargersRequest
var einride_shipper_v1beta1_SiteChargerService_ListSiteChargers = &cobra.Command{
	Use: "ListSiteChargers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.ListSiteChargers")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetSiteChargersRequest v1beta1.BatchGetSiteChargersRequest
var einride_shipper_v1beta1_SiteChargerService_BatchGetSiteChargers = &cobra.Command{
	Use: "BatchGetSiteChargers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteChargerService.BatchGetSiteChargers")
		return nil
	},
}

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
