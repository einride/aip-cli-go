package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_SiteService = &cobra.Command{
	Use: "einride.shipper.v1beta1.SiteService",
}

var einride_shipper_v1beta1_CreateSiteRequest v1beta1.CreateSiteRequest
var einride_shipper_v1beta1_SiteService_CreateSite = &cobra.Command{
	Use: "CreateSite",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.CreateSite")
		return nil
	},
}

var einride_shipper_v1beta1_GetSiteRequest v1beta1.GetSiteRequest
var einride_shipper_v1beta1_SiteService_GetSite = &cobra.Command{
	Use: "GetSite",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.GetSite")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetSitesRequest v1beta1.BatchGetSitesRequest
var einride_shipper_v1beta1_SiteService_BatchGetSites = &cobra.Command{
	Use: "BatchGetSites",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.BatchGetSites")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateSiteRequest v1beta1.UpdateSiteRequest
var einride_shipper_v1beta1_SiteService_UpdateSite = &cobra.Command{
	Use: "UpdateSite",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.UpdateSite")
		return nil
	},
}

var einride_shipper_v1beta1_ReferenceSiteRequest v1beta1.ReferenceSiteRequest
var einride_shipper_v1beta1_SiteService_ReferenceSite = &cobra.Command{
	Use: "ReferenceSite",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.ReferenceSite")
		return nil
	},
}

var einride_shipper_v1beta1_SearchSitesRequest v1beta1.SearchSitesRequest
var einride_shipper_v1beta1_SiteService_SearchSites = &cobra.Command{
	Use: "SearchSites",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteService.SearchSites")
		return nil
	},
}

func AddSiteServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteService)
}

func init() {
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_CreateSite)
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_GetSite)
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_BatchGetSites)
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_UpdateSite)
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_ReferenceSite)
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_SearchSites)
}
