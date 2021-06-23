package rescuev1alpha1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1alpha1 "github.com/einride/proto/gen/go/einride/optimizer/rescue/v1alpha1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// einride.optimizer.rescue.v1alpha1.RescueService.
var (
	einride_optimizer_rescue_v1alpha1_RescueServiceClient v1alpha1.RescueServiceClient
	einride_optimizer_rescue_v1alpha1_RescueService       = &cobra.Command{
		Use:   "einride.optimizer.rescue.v1alpha1.RescueService",
		Short: "Rescue service.",
		Long:  "Rescue service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_optimizer_rescue_v1alpha1_RescueServiceClient = v1alpha1.NewRescueServiceClient(conn)
			return nil
		},
	}
)

// einride.optimizer.rescue.v1alpha1.RescueService.ComputeVehiclePlans.
var (
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request v1alpha1.ComputeVehiclePlansRequest
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans         = &cobra.Command{
		Use: "ComputeVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_optimizer_rescue_v1alpha1_RescueServiceClient.ComputeVehiclePlans(cmd.Context(), &einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddRescueServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService)
}

func init() {
	einride_optimizer_rescue_v1alpha1_RescueService.AddCommand(einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans)

	// TODO: list: Shipments message

	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle = new(v1alpha1.Vehicle)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().StringVar(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.StartPlaceId, "vehicle.startPlaceId", "", "Identifier of the place where the vehicle is at at the start of this\nplanning cycle.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().StringVar(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.EndPlaceId, "vehicle.endPlaceId", "", "Identifier of the place where this vehicle must be at at the end of\nthis planning cycle. If not specified the does not return to the depot (open-ended route).")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod = new(v1alpha1.TimePeriod)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.StartTime = new(timestamppb.Timestamp)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.StartTime.Seconds, "vehicle.availabilityTimePeriod.startTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int32Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.StartTime.Nanos, "vehicle.availabilityTimePeriod.startTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.EndTime = new(timestamppb.Timestamp)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.EndTime.Seconds, "vehicle.availabilityTimePeriod.endTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int32Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.AvailabilityTimePeriod.EndTime.Nanos, "vehicle.availabilityTimePeriod.endTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery = new(v1alpha1.Battery)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Float64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery.CapacityJ, "vehicle.battery.capacityJ", 10, "The total capacity of the battery.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Float64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery.StartSocFraction, "vehicle.battery.startSocFraction", 10, "State of charge at the start of the planning cycle, expressed as a\nratio of capacity_joules.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Float64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery.MinSocFraction, "vehicle.battery.minSocFraction", 10, "Minimum state of charge over the planning cycle, expressed as a ratio of\ncapacity_joules.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Float64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery.MaxSocFraction, "vehicle.battery.maxSocFraction", 10, "Maximum state of charge over the planning cycle, expressed as a ratio of\ncapacity_joules.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Float64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.Battery.MaxChargingPowerW, "vehicle.battery.maxChargingPowerW", 10, "The maximum charging power the battery accepts.")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.MaxDriveDuration = new(durationpb.Duration)
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int64Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.MaxDriveDuration.Seconds, "vehicle.maxDriveDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans.Flags().Int32Var(&einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Vehicle.MaxDriveDuration.Nanos, "vehicle.maxDriveDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")

	// TODO: list: Places message

	// TODO: list: Chargers message

	// TODO: list: PathSummaries message

	einride_optimizer_rescue_v1alpha1_RescueService_ComputeVehiclePlans_Request.Sequence = new(v1alpha1.Sequence)
	// TODO: list: Activities message
}
