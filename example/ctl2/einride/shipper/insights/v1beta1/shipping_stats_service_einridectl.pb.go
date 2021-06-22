package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_ShippingStatsService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.ShippingStatsService",
}

var einride_shipper_insights_v1beta1_BatchSnapshotShipperShippingStatsRequest v1beta1.BatchSnapshotShipperShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotShipperShippingStats = &cobra.Command{
	Use: "BatchSnapshotShipperShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.BatchSnapshotShipperShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchSnapshotSiteShippingStatsRequest v1beta1.BatchSnapshotSiteShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotSiteShippingStats = &cobra.Command{
	Use: "BatchSnapshotSiteShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.BatchSnapshotSiteShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchSnapshotLaneShippingStatsRequest v1beta1.BatchSnapshotLaneShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotLaneShippingStats = &cobra.Command{
	Use: "BatchSnapshotLaneShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.BatchSnapshotLaneShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_AggregateShipperShippingStatsRequest v1beta1.AggregateShipperShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_AggregateShipperShippingStats = &cobra.Command{
	Use: "AggregateShipperShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.AggregateShipperShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_AggregateSiteShippingStatsRequest v1beta1.AggregateSiteShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_AggregateSiteShippingStats = &cobra.Command{
	Use: "AggregateSiteShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.AggregateSiteShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_AggregateLaneShippingStatsRequest v1beta1.AggregateLaneShippingStatsRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_AggregateLaneShippingStats = &cobra.Command{
	Use: "AggregateLaneShippingStats",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.AggregateLaneShippingStats")
		return nil
	},
}

var einride_shipper_insights_v1beta1_QuerySiteHourlyPerformanceRequest v1beta1.QuerySiteHourlyPerformanceRequest
var einride_shipper_insights_v1beta1_ShippingStatsService_QuerySiteHourlyPerformance = &cobra.Command{
	Use: "QuerySiteHourlyPerformance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.ShippingStatsService.QuerySiteHourlyPerformance")
		return nil
	},
}

func AddShippingStatsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService)
}

func init() {
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotShipperShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotSiteShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_BatchSnapshotLaneShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_AggregateShipperShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_AggregateSiteShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_AggregateLaneShippingStats)
	einride_shipper_insights_v1beta1_ShippingStatsService.AddCommand(einride_shipper_insights_v1beta1_ShippingStatsService_QuerySiteHourlyPerformance)
}
