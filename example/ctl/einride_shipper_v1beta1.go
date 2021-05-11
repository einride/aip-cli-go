package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_shipper_v1beta1_CustomItemTypeService = &cobra.Command{
	Use: "einride.shipper.v1beta1.CustomItemType",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.CustomItemType called")
	},
}
var einride_shipper_v1beta1_LaneService = &cobra.Command{
	Use: "einride.shipper.v1beta1.Lane",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.Lane called")
	},
}
var einride_shipper_v1beta1_LineItemService = &cobra.Command{
	Use: "einride.shipper.v1beta1.LineItem",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.LineItem called")
	},
}
var einride_shipper_v1beta1_OatlyShipmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.OatlyShipment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.OatlyShipment called")
	},
}
var einride_shipper_v1beta1_OpeningHoursService = &cobra.Command{
	Use: "einride.shipper.v1beta1.OpeningHours",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.OpeningHours called")
	},
}
var einride_shipper_v1beta1_ShipmentAttachmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentAttachment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipmentAttachment called")
	},
}
var einride_shipper_v1beta1_ShipmentChangeProposalService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentChangeProposal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipmentChangeProposal called")
	},
}
var einride_shipper_v1beta1_ShipmentClaimService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentClaim",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipmentClaim called")
	},
}
var einride_shipper_v1beta1_ShipmentEventService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentEvent",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipmentEvent called")
	},
}
var einride_shipper_v1beta1_ShipmentService = &cobra.Command{
	Use: "einride.shipper.v1beta1.Shipment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.Shipment called")
	},
}
var einride_shipper_v1beta1_ShipperService = &cobra.Command{
	Use: "einride.shipper.v1beta1.Shipper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.Shipper called")
	},
}
var einride_shipper_v1beta1_ShipperSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipperSettings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipperSettings called")
	},
}
var einride_shipper_v1beta1_ShipperUserSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipperUserSettings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.ShipperUserSettings called")
	},
}
var einride_shipper_v1beta1_SiteChargerService = &cobra.Command{
	Use: "einride.shipper.v1beta1.SiteCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.SiteCharger called")
	},
}
var einride_shipper_v1beta1_SiteNotificationSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.SiteNotificationSettings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.SiteNotificationSettings called")
	},
}
var einride_shipper_v1beta1_SiteService = &cobra.Command{
	Use: "einride.shipper.v1beta1.Site",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.shipper.v1beta1.Site called")
	},
}

func init() {
	rootCmd.AddCommand(einride_shipper_v1beta1_CustomItemTypeService)
	rootCmd.AddCommand(einride_shipper_v1beta1_LaneService)
	rootCmd.AddCommand(einride_shipper_v1beta1_LineItemService)
	rootCmd.AddCommand(einride_shipper_v1beta1_OatlyShipmentService)
	rootCmd.AddCommand(einride_shipper_v1beta1_OpeningHoursService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipmentAttachmentService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipmentChangeProposalService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipmentClaimService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipmentEventService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipmentService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipperService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipperSettingsService)
	rootCmd.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService)
	rootCmd.AddCommand(einride_shipper_v1beta1_SiteChargerService)
	rootCmd.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService)
	rootCmd.AddCommand(einride_shipper_v1beta1_SiteService)
}
