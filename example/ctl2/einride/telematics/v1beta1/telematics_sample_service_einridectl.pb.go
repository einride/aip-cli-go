package telematicsv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/telematics/v1beta1"
	cobra "github.com/spf13/cobra"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	log "log"
)

// einride.telematics.v1beta1.TelematicsSampleService.
var (
	einride_telematics_v1beta1_TelematicsSampleServiceClient v1beta1.TelematicsSampleServiceClient
	einride_telematics_v1beta1_TelematicsSampleService       = &cobra.Command{
		Use: "einride.telematics.v1beta1.TelematicsSampleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_telematics_v1beta1_TelematicsSampleServiceClient = v1beta1.NewTelematicsSampleServiceClient(conn)
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.GetTelematicsSample.
var (
	einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample_Request v1beta1.GetTelematicsSampleRequest
	einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample         = &cobra.Command{
		Use: "GetTelematicsSample",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.GetTelematicsSample")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.ListTelematicsSamples.
var (
	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples_Request v1beta1.ListTelematicsSamplesRequest
	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples         = &cobra.Command{
		Use: "ListTelematicsSamples",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.ListTelematicsSamples")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.SearchTelematicsSamples.
var (
	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples_Request v1beta1.SearchTelematicsSamplesRequest
	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples         = &cobra.Command{
		Use: "SearchTelematicsSamples",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.SearchTelematicsSamples")
			return nil
		},
	}
)

// einride.telematics.v1beta1.TelematicsSampleService.CreateTelematicsSample.
var (
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request v1beta1.CreateTelematicsSampleRequest
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample         = &cobra.Command{
		Use: "CreateTelematicsSample",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.telematics.v1beta1.TelematicsSampleService.CreateTelematicsSample")
			return nil
		},
	}
)

func AddTelematicsSampleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_telematics_v1beta1_TelematicsSampleService)
}

func init() {
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample)

	einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_GetTelematicsSample_Request.Name, "name", "", "Resource name of the telematics sample to retrieve.")
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples)

	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples_Request.Parent, "parent", "", "Resource name of the parent telematics device.")

	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_ListTelematicsSamples_Request.PageToken, "pageToken", "", "A page token, received from a previous List call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples)

	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples_Request.Parent, "parent", "", "Resource name of the parent telematics device.\nCan be `\"telematicsDevices/-\"` to search across all devices.")

	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_SearchTelematicsSamples_Request.PageToken, "pageToken", "", "A page token, received from a previous List call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	// TODO: one-of: Vehicle

	// TODO: one-of: Shipment
	einride_telematics_v1beta1_TelematicsSampleService.AddCommand(einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample)

	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.Parent, "parent", "", "Resource name of the parent telematics device where this telematics sample will be created.")

	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample = new(v1beta1.TelematicsSample)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.Name, "telematicsSample.name", "", "The resource name of the telematics sample.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TelematicsDevice, "telematicsSample.telematicsDevice", "", "Resource name of the device that generated the sample.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.DeviceTime = new(timestamppb.Timestamp)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Int64Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.DeviceTime.Seconds, "telematicsSample.deviceTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.DeviceTime.Nanos, "telematicsSample.deviceTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssTime = new(timestamppb.Timestamp)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Int64Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssTime.Seconds, "telematicsSample.gnssTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssTime.Nanos, "telematicsSample.gnssTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LatitudeDegrees = new(wrapperspb.DoubleValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float64Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LatitudeDegrees.Value, "telematicsSample.latitudeDegrees.value", 10, "The double value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LongitudeDegrees = new(wrapperspb.DoubleValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float64Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LongitudeDegrees.Value, "telematicsSample.longitudeDegrees.value", 10, "The double value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AltitudeMetres = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AltitudeMetres.Value, "telematicsSample.altitudeMetres.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.HeadingRadians = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.HeadingRadians.Value, "telematicsSample.headingRadians.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.PositionHorizontalAccuracyMetres = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.PositionHorizontalAccuracyMetres.Value, "telematicsSample.positionHorizontalAccuracyMetres.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.PositionVerticalAccuracyMetres = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.PositionVerticalAccuracyMetres.Value, "telematicsSample.positionVerticalAccuracyMetres.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.HasLocationFix = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.HasLocationFix.Value, "telematicsSample.hasLocationFix.value", false, "The bool value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssSatelliteCount = new(wrapperspb.Int32Value)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Int32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssSatelliteCount.Value, "telematicsSample.gnssSatelliteCount.value", 10, "The int32 value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssJammingDetected = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssJammingDetected.Value, "telematicsSample.gnssJammingDetected.value", false, "The bool value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.UplinkBandwidthUsageMegabitsPerSecond = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.UplinkBandwidthUsageMegabitsPerSecond.Value, "telematicsSample.uplinkBandwidthUsageMegabitsPerSecond.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.DownlinkBandwidthUsageMegabitsPerSecond = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.DownlinkBandwidthUsageMegabitsPerSecond.Value, "telematicsSample.downlinkBandwidthUsageMegabitsPerSecond.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.Rssi = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.Rssi.Value, "telematicsSample.rssi.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalChargePercent = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalChargePercent.Value, "telematicsSample.totalChargePercent.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalBatteryVoltageVolts = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalBatteryVoltageVolts.Value, "telematicsSample.totalBatteryVoltageVolts.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalBatteryCurrentAmperes = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalBatteryCurrentAmperes.Value, "telematicsSample.totalBatteryCurrentAmperes.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.IsChargeCableConnected = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.IsChargeCableConnected.Value, "telematicsSample.isChargeCableConnected.value", false, "The bool value.")
	// TODO: list: Batteries message
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ChargingCurrentAmperes = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ChargingCurrentAmperes.Value, "telematicsSample.chargingCurrentAmperes.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ChargingVoltageVolts = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ChargingVoltageVolts.Value, "telematicsSample.chargingVoltageVolts.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.SpeedKilometresPerHour = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.SpeedKilometresPerHour.Value, "telematicsSample.speedKilometresPerHour.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.SteerAngleRadians = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.SteerAngleRadians.Value, "telematicsSample.steerAngleRadians.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.MotorOutputSpeedRpm = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.MotorOutputSpeedRpm.Value, "telematicsSample.motorOutputSpeedRpm.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AppliedThrottlePercent = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AppliedThrottlePercent.Value, "telematicsSample.appliedThrottlePercent.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.MotorOutputTorqueNewtonmetres = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.MotorOutputTorqueNewtonmetres.Value, "telematicsSample.motorOutputTorqueNewtonmetres.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AppliedServiceBrakePercent = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AppliedServiceBrakePercent.Value, "telematicsSample.appliedServiceBrakePercent.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ForwardAccelerationMetresPerSecondSquared = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ForwardAccelerationMetresPerSecondSquared.Value, "telematicsSample.forwardAccelerationMetresPerSecondSquared.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ReverseGear = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ReverseGear.Value, "telematicsSample.reverseGear.value", false, "The bool value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ParkBrake = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.ParkBrake.Value, "telematicsSample.parkBrake.value", false, "The bool value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AxleWeightKilograms = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.AxleWeightKilograms.Value, "telematicsSample.axleWeightKilograms.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadingBayDoorOpen = new(wrapperspb.BoolValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().BoolVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadingBayDoorOpen.Value, "telematicsSample.loadingBayDoorOpen.value", false, "The bool value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TirePressurePascals = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TirePressurePascals.Value, "telematicsSample.tirePressurePascals.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadCarrierTemperatureCelsius = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadCarrierTemperatureCelsius.Value, "telematicsSample.loadCarrierTemperatureCelsius.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadCarrierPh = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.LoadCarrierPh.Value, "telematicsSample.loadCarrierPh.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.OutsideNoiseDecibels = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.OutsideNoiseDecibels.Value, "telematicsSample.outsideNoiseDecibels.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.BatteryInsulationResistanceOhm = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.BatteryInsulationResistanceOhm.Value, "telematicsSample.batteryInsulationResistanceOhm.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalVehicleDistanceKm = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.TotalVehicleDistanceKm.Value, "telematicsSample.totalVehicleDistanceKm.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssSpeedKilometresPerHour = new(wrapperspb.FloatValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().Float32Var(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.GnssSpeedKilometresPerHour.Value, "telematicsSample.gnssSpeedKilometresPerHour.value", 10, "The float value.")
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.VehicleId = new(wrapperspb.StringValue)
	einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample.Flags().StringVar(&einride_telematics_v1beta1_TelematicsSampleService_CreateTelematicsSample_Request.TelematicsSample.VehicleId.Value, "telematicsSample.vehicleId.value", "", "The string value.")
	// TODO: list: Errors message
}
