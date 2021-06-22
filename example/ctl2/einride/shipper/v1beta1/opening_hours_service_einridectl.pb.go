package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_OpeningHoursService = &cobra.Command{
	Use: "einride.shipper.v1beta1.OpeningHoursService",
}

var einride_shipper_v1beta1_GetShipperWeeklyOpeningHoursRequest v1beta1.GetShipperWeeklyOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_GetShipperWeeklyOpeningHours = &cobra.Command{
	Use: "GetShipperWeeklyOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.GetShipperWeeklyOpeningHours")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipperWeeklyOpeningHoursRequest v1beta1.UpdateShipperWeeklyOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_UpdateShipperWeeklyOpeningHours = &cobra.Command{
	Use: "UpdateShipperWeeklyOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.UpdateShipperWeeklyOpeningHours")
		return nil
	},
}

var einride_shipper_v1beta1_GetSiteWeeklyOpeningHoursRequest v1beta1.GetSiteWeeklyOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_GetSiteWeeklyOpeningHours = &cobra.Command{
	Use: "GetSiteWeeklyOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.GetSiteWeeklyOpeningHours")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateSiteWeeklyOpeningHoursRequest v1beta1.UpdateSiteWeeklyOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_UpdateSiteWeeklyOpeningHours = &cobra.Command{
	Use: "UpdateSiteWeeklyOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.UpdateSiteWeeklyOpeningHours")
		return nil
	},
}

var einride_shipper_v1beta1_GetSiteOpeningHoursOverrideRequest v1beta1.GetSiteOpeningHoursOverrideRequest
var einride_shipper_v1beta1_OpeningHoursService_GetSiteOpeningHoursOverride = &cobra.Command{
	Use: "GetSiteOpeningHoursOverride",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.GetSiteOpeningHoursOverride")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateSiteOpeningHoursOverrideRequest v1beta1.UpdateSiteOpeningHoursOverrideRequest
var einride_shipper_v1beta1_OpeningHoursService_UpdateSiteOpeningHoursOverride = &cobra.Command{
	Use: "UpdateSiteOpeningHoursOverride",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.UpdateSiteOpeningHoursOverride")
		return nil
	},
}

var einride_shipper_v1beta1_GetSiteOpeningHoursRequest v1beta1.GetSiteOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_GetSiteOpeningHours = &cobra.Command{
	Use: "GetSiteOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.GetSiteOpeningHours")
		return nil
	},
}

var einride_shipper_v1beta1_BatchGetSiteOpeningHoursRequest v1beta1.BatchGetSiteOpeningHoursRequest
var einride_shipper_v1beta1_OpeningHoursService_BatchGetSiteOpeningHours = &cobra.Command{
	Use: "BatchGetSiteOpeningHours",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.OpeningHoursService.BatchGetSiteOpeningHours")
		return nil
	},
}

func AddOpeningHoursServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_OpeningHoursService)
}

func init() {
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_GetShipperWeeklyOpeningHours)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_UpdateShipperWeeklyOpeningHours)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_GetSiteWeeklyOpeningHours)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_UpdateSiteWeeklyOpeningHours)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_GetSiteOpeningHoursOverride)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_UpdateSiteOpeningHoursOverride)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_GetSiteOpeningHours)
	einride_shipper_v1beta1_OpeningHoursService.AddCommand(einride_shipper_v1beta1_OpeningHoursService_BatchGetSiteOpeningHours)
}
