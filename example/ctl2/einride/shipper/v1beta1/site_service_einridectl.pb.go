package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Parent, "parent", "", "Resource name of the parent resource where this site will be created.")

	einride_shipper_v1beta1_SiteService_CreateSite_Request.Site = new(v1beta1.Site)
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Name, "site.name", "", "The resource name of the site.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Etag, "site.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.IntegrationFile, "site.integrationFile", "", "Resource name of the integration file the site was ingested from, if the site was ingested.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.ReferenceId, "site.referenceId", "", "The reference ID for this site from e.g. TMS or ERP systems.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.DisplayName, "site.displayName", "", "The display name.\nFor example: \"Västers lager\"")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Address, "site.address", "", "The site's address.\nFor example: Ringvägen 30")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.PostalCode, "site.postalCode", "", "postal_code is on the swedish postal code format.\nFor example: \"820 40\"")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.City, "site.city", "", "For example: \"Göteborg\"")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.CountryCode, "site.countryCode", "", "Country code on ISO 3166-1 alpha-2 format.\nFor example: \"SE\"\n(-- api-linter: core::0143::standardized-codes=disabled\n    aip.dev/not-precedent: Backwards compatibility. --)")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().Float64Var(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.LatitudeDegrees, "site.latitudeDegrees", 10, "The latitude of the site's location.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().Float64Var(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.LongitudeDegrees, "site.longitudeDegrees", 10, "The longitude of the site's location.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().BoolVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Hub, "site.hub", false, "Flag indicating if this site is considered a shipping hub.")
	einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter = new(v1beta1.Address)
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter.DisplayName, "site.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter.Street, "site.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter.PostalCode, "site.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter.City, "site.headquarter.city", "", "For example: Göteborg.")
	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.Site.Headquarter.RegionCode, "site.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_shipper_v1beta1_SiteService_CreateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_CreateSite_Request.SiteId, "siteId", "", "The ID to use for the site, which will become the final component of\nthe site's resource name.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-z][0-9]-/.")
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_GetSite)

	einride_shipper_v1beta1_SiteService_GetSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_GetSite_Request.Name, "name", "", "Resource name of the site to retrieve.")
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_BatchGetSites)

	einride_shipper_v1beta1_SiteService_BatchGetSites.Flags().StringVar(&einride_shipper_v1beta1_SiteService_BatchGetSites_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all sites being retrieved.\nThe parent of all of the sites specified in `names`\nmust match this field.")

	einride_shipper_v1beta1_SiteService_BatchGetSites.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteService_BatchGetSites_Request.Names, "names", []string{}, "Resource names of the sites to retrieve.\nA maximum of 1000 sites can be retrieved in a batch.")
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_UpdateSite)

	einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site = new(v1beta1.Site)
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Name, "site.name", "", "The resource name of the site.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Etag, "site.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.IntegrationFile, "site.integrationFile", "", "Resource name of the integration file the site was ingested from, if the site was ingested.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.ReferenceId, "site.referenceId", "", "The reference ID for this site from e.g. TMS or ERP systems.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.DisplayName, "site.displayName", "", "The display name.\nFor example: \"Västers lager\"")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Address, "site.address", "", "The site's address.\nFor example: Ringvägen 30")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.PostalCode, "site.postalCode", "", "postal_code is on the swedish postal code format.\nFor example: \"820 40\"")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.City, "site.city", "", "For example: \"Göteborg\"")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.CountryCode, "site.countryCode", "", "Country code on ISO 3166-1 alpha-2 format.\nFor example: \"SE\"\n(-- api-linter: core::0143::standardized-codes=disabled\n    aip.dev/not-precedent: Backwards compatibility. --)")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().Float64Var(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.LatitudeDegrees, "site.latitudeDegrees", 10, "The latitude of the site's location.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().Float64Var(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.LongitudeDegrees, "site.longitudeDegrees", 10, "The longitude of the site's location.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().BoolVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Hub, "site.hub", false, "Flag indicating if this site is considered a shipping hub.")
	einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter = new(v1beta1.Address)
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter.DisplayName, "site.headquarter.displayName", "", "Who its intended for, can be a multiline string with C/O.\nFor example: Olle Svensson / Example Company AB / etc.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter.Street, "site.headquarter.street", "", "The street.\nFor example: Ringvägen 30.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter.PostalCode, "site.headquarter.postalCode", "", "The postal code.\nFor example: 820 40")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter.City, "site.headquarter.city", "", "For example: Göteborg.")
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.Site.Headquarter.RegionCode, "site.headquarter.regionCode", "", "Fields representing individual countries or nations must use the Unicode CLDR region codes,\nsuch as US or CH.\nFor region codes see: https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry\nFor example: \"SE\"")

	einride_shipper_v1beta1_SiteService_UpdateSite_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_SiteService_UpdateSite.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteService_UpdateSite_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_ReferenceSite)

	einride_shipper_v1beta1_SiteService_ReferenceSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_ReferenceSite_Request.Parent, "parent", "", "Resource name of the shipper to look up the site for.")

	einride_shipper_v1beta1_SiteService_ReferenceSite.Flags().StringVar(&einride_shipper_v1beta1_SiteService_ReferenceSite_Request.ReferenceId, "referenceId", "", "Reference ID of the site to retreive.")
	einride_shipper_v1beta1_SiteService.AddCommand(einride_shipper_v1beta1_SiteService_SearchSites)

	einride_shipper_v1beta1_SiteService_SearchSites.Flags().StringVar(&einride_shipper_v1beta1_SiteService_SearchSites_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all sites being searched for.")

	einride_shipper_v1beta1_SiteService_SearchSites.Flags().StringVar(&einride_shipper_v1beta1_SiteService_SearchSites_Request.FreeTextQuery, "freeTextQuery", "", "Text search.")

	// TODO: enum SiteType

	einride_shipper_v1beta1_SiteService_SearchSites.Flags().Int32Var(&einride_shipper_v1beta1_SiteService_SearchSites_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_SiteService_SearchSites.Flags().StringVar(&einride_shipper_v1beta1_SiteService_SearchSites_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_shipper_v1beta1_SiteService_SearchSites.Flags().BoolVar(&einride_shipper_v1beta1_SiteService_SearchSites_Request.Hubs, "hubs", false, "Flag indicating if search results should only contain hubs.\nWhen unset, hubs and non-hubs are returned.")
}
