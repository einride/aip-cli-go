package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsServiceClient v1beta1.SiteNotificationSettingsServiceClient
	einride_shipper_v1beta1_SiteNotificationSettingsService       = &cobra.Command{
		Use:   "shipper.v1beta1.SiteNotificationSettingsService",
		Short: "Site notification settings service.",
		Long:  "Site notification settings service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_SiteNotificationSettingsServiceClient = v1beta1.NewSiteNotificationSettingsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.GetSiteNotificationSettings.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings_Request v1beta1.GetSiteNotificationSettingsRequest
	einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings         = &cobra.Command{
		Use: "GetSiteNotificationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteNotificationSettingsServiceClient.GetSiteNotificationSettings(cmd.Context(), &einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.UpdateSiteNotificationSettings.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request v1beta1.UpdateSiteNotificationSettingsRequest
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings         = &cobra.Command{
		Use: "UpdateSiteNotificationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_SiteNotificationSettingsServiceClient.UpdateSiteNotificationSettings(cmd.Context(), &einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddSiteNotificationSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService)
}

func init() {
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings)

	einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings.Flags().StringVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings_Request.Name, "name", "", "Resource name of the site notification settings to retrieve.\nFormat: `shippers/{shipper}/sites/{site}/notificationSettings`.")
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings)

	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings = new(v1beta1.SiteNotificationSettings)
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().StringVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.Name, "siteNotificationSettings.name", "", "Resource name of the settings object.")
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.DurationBeforeArrival = new(durationpb.Duration)
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().Int64Var(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.DurationBeforeArrival.Seconds, "siteNotificationSettings.durationBeforeArrival.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().Int32Var(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.DurationBeforeArrival.Nanos, "siteNotificationSettings.durationBeforeArrival.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.EtaPhoneNumbers, "siteNotificationSettings.etaPhoneNumbers", []string{}, "Phone numbers to send ETA notifications to.")
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.EtaEmailAddresses, "siteNotificationSettings.etaEmailAddresses", []string{}, "Email addresses to send ETA notifications to.")
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.SiteNotificationSettings.VoiceEtaPhoneNumbers, "siteNotificationSettings.voiceEtaPhoneNumbers", []string{}, "Phone numbers to send voice ETA notifications to.")

	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
}
