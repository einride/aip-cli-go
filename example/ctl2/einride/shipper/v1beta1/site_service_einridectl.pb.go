package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.SiteService.
var (
	einride_shipper_v1beta1_SiteServiceClient v1beta1.SiteServiceClient
	einride_shipper_v1beta1_SiteService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.SiteService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_SiteServiceClient = v1beta1.NewSiteServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.CreateSite.
var (
	einride_shipper_v1beta1_SiteService_CreateSite_Request v1beta1.CreateSiteRequest
	einride_shipper_v1beta1_SiteService_CreateSite         = &cobra.Command{
		Use: "CreateSite",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.CreateSite")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.GetSite.
var (
	einride_shipper_v1beta1_SiteService_GetSite_Request v1beta1.GetSiteRequest
	einride_shipper_v1beta1_SiteService_GetSite         = &cobra.Command{
		Use: "GetSite",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.GetSite")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.BatchGetSites.
var (
	einride_shipper_v1beta1_SiteService_BatchGetSites_Request v1beta1.BatchGetSitesRequest
	einride_shipper_v1beta1_SiteService_BatchGetSites         = &cobra.Command{
		Use: "BatchGetSites",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.BatchGetSites")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.UpdateSite.
var (
	einride_shipper_v1beta1_SiteService_UpdateSite_Request v1beta1.UpdateSiteRequest
	einride_shipper_v1beta1_SiteService_UpdateSite         = &cobra.Command{
		Use: "UpdateSite",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.UpdateSite")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.ReferenceSite.
var (
	einride_shipper_v1beta1_SiteService_ReferenceSite_Request v1beta1.ReferenceSiteRequest
	einride_shipper_v1beta1_SiteService_ReferenceSite         = &cobra.Command{
		Use: "ReferenceSite",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.ReferenceSite")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteService.SearchSites.
var (
	einride_shipper_v1beta1_SiteService_SearchSites_Request v1beta1.SearchSitesRequest
	einride_shipper_v1beta1_SiteService_SearchSites         = &cobra.Command{
		Use: "SearchSites",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteService.SearchSites")
			return nil
		},
	}
)

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
