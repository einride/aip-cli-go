package shipperinsightsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.insights.v1beta1.MapInsightsService.
var (
	einride_shipper_insights_v1beta1_MapInsightsServiceClient v1beta1.MapInsightsServiceClient
	einride_shipper_insights_v1beta1_MapInsightsService       = &cobra.Command{
		Use: "einride.shipper.insights.v1beta1.MapInsightsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_insights_v1beta1_MapInsightsServiceClient = v1beta1.NewMapInsightsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetShipperMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights_Request v1beta1.BatchGetShipperMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights         = &cobra.Command{
		Use: "BatchGetShipperMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetShipperMapInsights")
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetSiteMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights_Request v1beta1.BatchGetSiteMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights         = &cobra.Command{
		Use: "BatchGetSiteMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetSiteMapInsights")
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetLaneMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights_Request v1beta1.BatchGetLaneMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights         = &cobra.Command{
		Use: "BatchGetLaneMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.insights.v1beta1.MapInsightsService.BatchGetLaneMapInsights")
			return nil
		},
	}
)

func AddMapInsightsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService)
}

func init() {
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights)
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights)
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights)
}
