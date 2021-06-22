package ctl2

import (
	v1beta1 "github.com/einride/ctl/example/ctl2/einride/account/v1beta1"
	v1beta11 "github.com/einride/ctl/example/ctl2/einride/authz/v1beta1"
	v1beta12 "github.com/einride/ctl/example/ctl2/einride/carrier/v1beta1"
	v1beta13 "github.com/einride/ctl/example/ctl2/einride/cem/v1beta1"
	v1beta14 "github.com/einride/ctl/example/ctl2/einride/grid/v1beta1"
	v1 "github.com/einride/ctl/example/ctl2/einride/maps/v1"
	v1beta15 "github.com/einride/ctl/example/ctl2/einride/messenger/v1beta1"
	v1alpha1 "github.com/einride/ctl/example/ctl2/einride/optimizer/rescue/v1alpha1"
	v11 "github.com/einride/ctl/example/ctl2/einride/optimizer/v1"
	v1beta16 "github.com/einride/ctl/example/ctl2/einride/planner/v1beta1"
	v12 "github.com/einride/ctl/example/ctl2/einride/planning/v1"
	v1beta17 "github.com/einride/ctl/example/ctl2/einride/prospect/shipper/v1beta1"
	v1beta18 "github.com/einride/ctl/example/ctl2/einride/shipper/insights/v1beta1"
	v1beta19 "github.com/einride/ctl/example/ctl2/einride/shipper/integration/v1beta1"
	v1beta110 "github.com/einride/ctl/example/ctl2/einride/shipper/v1beta1"
	v1beta111 "github.com/einride/ctl/example/ctl2/einride/telematics/v1beta1"
	v1beta112 "github.com/einride/ctl/example/ctl2/einride/upload/v1beta1"
	cobra "github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "einridectl2",
}

func init() {
	flags := Command.PersistentFlags()
	flags.Bool("insecure", false, "make insecure client connection, must be used with 'address' option")
	flags.String("address", "", "address to connect to")
	flags.String("token", "", "bearer token used by client")
	flags.String("project", "einride-dev", "GCP project")
	v1beta1.AddAuthenticationServiceCommand(Command)
	v1beta1.AddFeatureFlagsServiceCommand(Command)
	v1beta1.AddTenantServiceCommand(Command)
	v1beta1.AddUserServiceCommand(Command)
	v1beta11.AddAuthorizationServiceCommand(Command)
	v1beta12.AddCarrierServiceCommand(Command)
	v1beta12.AddVehicleServiceCommand(Command)
	v1beta12.AddVehicleTypeServiceCommand(Command)
	v1beta13.AddCEMServiceCommand(Command)
	v1beta14.AddChargerServiceCommand(Command)
	v1.AddGeocodingServiceCommand(Command)
	v1.AddPathServiceCommand(Command)
	v1beta15.AddEmailServiceCommand(Command)
	v1beta15.AddSMSMessageServiceCommand(Command)
	v1alpha1.AddRescueServiceCommand(Command)
	v11.AddOptimizerServiceCommand(Command)
	v1beta16.AddDefaultScheduleServiceCommand(Command)
	v1beta16.AddPlannerSiteSettingsServiceCommand(Command)
	v1beta16.AddScheduleBookingServiceCommand(Command)
	v1beta16.AddShipmentAllocationServiceCommand(Command)
	v1beta16.AddTransportInstallationServiceCommand(Command)
	v1beta16.AddTransportScheduleServiceCommand(Command)
	v1beta16.AddVehiclePlanServiceCommand(Command)
	v12.AddAllocateShipmentsServiceCommand(Command)
	v12.AddScheduleBookingServiceCommand(Command)
	v12.AddScheduleServiceCommand(Command)
	v12.AddTaskServiceCommand(Command)
	v12.AddTransportInstallationServiceCommand(Command)
	v1beta17.AddElectrificationPlannerServiceCommand(Command)
	v1beta17.AddProspectiveShipperServiceCommand(Command)
	v1beta18.AddMapInsightsServiceCommand(Command)
	v1beta18.AddShippingCostServiceCommand(Command)
	v1beta18.AddShippingEmissionsServiceCommand(Command)
	v1beta19.AddIntegrationBucketServiceCommand(Command)
	v1beta19.AddIntegrationFileLogServiceCommand(Command)
	v1beta19.AddIntegrationFileServiceCommand(Command)
	v1beta19.AddIntegrationLogServiceCommand(Command)
	v1beta110.AddCustomItemTypeServiceCommand(Command)
	v1beta110.AddLaneServiceCommand(Command)
	v1beta110.AddOatlyShipmentServiceCommand(Command)
	v1beta110.AddShipmentAttachmentServiceCommand(Command)
	v1beta110.AddShipmentChangeProposalServiceCommand(Command)
	v1beta110.AddShipmentClaimServiceCommand(Command)
	v1beta110.AddShipmentEventServiceCommand(Command)
	v1beta110.AddShipmentServiceCommand(Command)
	v1beta110.AddShipperServiceCommand(Command)
	v1beta110.AddShipperSettingsServiceCommand(Command)
	v1beta110.AddShipperUserSettingsServiceCommand(Command)
	v1beta110.AddSiteChargerServiceCommand(Command)
	v1beta110.AddSiteNotificationSettingsServiceCommand(Command)
	v1beta110.AddSiteServiceCommand(Command)
	v1beta111.AddTelematicsDeviceServiceCommand(Command)
	v1beta111.AddTelematicsSampleServiceCommand(Command)
	v1beta111.AddVehicleConnectionServiceCommand(Command)
	v1beta112.AddFileServiceCommand(Command)
	AddCompletion(Command)
}
