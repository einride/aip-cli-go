package cemv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/cem/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

// einride.cem.v1beta1.CEMService.
var (
	einride_cem_v1beta1_CEMServiceClient v1beta1.CEMServiceClient
	einride_cem_v1beta1_CEMService       = &cobra.Command{
		Use: "einride.cem.v1beta1.CEMService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_cem_v1beta1_CEMServiceClient = v1beta1.NewCEMServiceClient(conn)
			return nil
		},
	}
)

// einride.cem.v1beta1.CEMService.CalculateResult.
var (
	einride_cem_v1beta1_CEMService_CalculateResult_Request v1beta1.CalculateResultRequest
	einride_cem_v1beta1_CEMService_CalculateResult         = &cobra.Command{
		Use: "CalculateResult",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_cem_v1beta1_CEMServiceClient.CalculateResult(cmd.Context(), &einride_cem_v1beta1_CEMService_CalculateResult_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.cem.v1beta1.CEMService.BatchCalculateResults.
var (
	einride_cem_v1beta1_CEMService_BatchCalculateResults_Request v1beta1.BatchCalculateResultsRequest
	einride_cem_v1beta1_CEMService_BatchCalculateResults         = &cobra.Command{
		Use: "BatchCalculateResults",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_cem_v1beta1_CEMServiceClient.BatchCalculateResults(cmd.Context(), &einride_cem_v1beta1_CEMService_BatchCalculateResults_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddCEMServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_cem_v1beta1_CEMService)
}

func init() {
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_CalculateResult)

	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec = new(v1beta1.Spec)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().StringVar(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.DisplayName, "spec.displayName", "", "Display name.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario = new(v1beta1.Spec_Scenario)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().StringVar(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DisplayName, "spec.scenario.displayName", "", "Display name.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.TruckCount, "spec.scenario.truckCount", 10, "Count.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.TrucksPerDriver, "spec.scenario.trucksPerDriver", 10, "Trucks per driver.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.TrailersPerTruck, "spec.scenario.trailersPerTruck", 10, "Trailers per truck.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverRegularTimePerWeekdayDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverRegularTimePerWeekdayDuration.Seconds, "spec.scenario.driverRegularTimePerWeekdayDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverRegularTimePerWeekdayDuration.Nanos, "spec.scenario.driverRegularTimePerWeekdayDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverOvertimePerWeekdayDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverOvertimePerWeekdayDuration.Seconds, "spec.scenario.driverOvertimePerWeekdayDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverOvertimePerWeekdayDuration.Nanos, "spec.scenario.driverOvertimePerWeekdayDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverTimePerWeekendDayDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverTimePerWeekendDayDuration.Seconds, "spec.scenario.driverTimePerWeekendDayDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DriverTimePerWeekendDayDuration.Nanos, "spec.scenario.driverTimePerWeekendDayDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DrivenDistancePerTruckPerWeekdayDayKm, "spec.scenario.drivenDistancePerTruckPerWeekdayDayKm", 10, "Driven distance per truck per weekday day km.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.DrivenDistancePerTruckPerWeekendDayKm, "spec.scenario.drivenDistancePerTruckPerWeekendDayKm", 10, "Driven distance per truck per weekend day km.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayRegularDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayRegularDuration.Seconds, "spec.scenario.supportStaffTimePerTruckPerWeekdayDayRegularDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayRegularDuration.Nanos, "spec.scenario.supportStaffTimePerTruckPerWeekdayDayRegularDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayOvertimeDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayOvertimeDuration.Seconds, "spec.scenario.supportStaffTimePerTruckPerWeekdayDayOvertimeDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekdayDayOvertimeDuration.Nanos, "spec.scenario.supportStaffTimePerTruckPerWeekdayDayOvertimeDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekendDayDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekendDayDuration.Seconds, "spec.scenario.supportStaffTimePerTruckPerWeekendDayDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.SupportStaffTimePerTruckPerWeekendDayDuration.Nanos, "spec.scenario.supportStaffTimePerTruckPerWeekendDayDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.PalletsTransportedAllTrucksPerWeekdayDay, "spec.scenario.palletsTransportedAllTrucksPerWeekdayDay", 10, "Pallets transported all trucks per weekday day.\nFor post-processing cost.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.PalletsTransportedAllTrucksPerWeekendDay, "spec.scenario.palletsTransportedAllTrucksPerWeekendDay", 10, "Pallets transported all trucks per weekend day.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.FuelEnergySource1Percent, "spec.scenario.fuelEnergySource1Percent", 10, "Share fuel energy type1.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Scenario.OverheadMarginPercent, "spec.scenario.overheadMarginPercent", 10, "Overhead margin.\nExpressed as a proportion of calculated cost. For adding a confidence interval.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle = new(v1beta1.Spec_Vehicle)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.LifeSpanDistanceKm, "spec.vehicle.lifeSpanDistanceKm", 10, "Life span distance.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.LifeSpanMaximumDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.LifeSpanMaximumDuration.Seconds, "spec.vehicle.lifeSpanMaximumDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.LifeSpanMaximumDuration.Nanos, "spec.vehicle.lifeSpanMaximumDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedAcquisitionCostTruckMu, "spec.vehicle.fixedAcquisitionCostTruckMu", 10, "Fixed acquisition cost truck.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedAcquisitionCostTrailerMu, "spec.vehicle.fixedAcquisitionCostTrailerMu", 10, "Fixed acquisition cost trailer.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedTruckSalvageValuePercent, "spec.vehicle.fixedTruckSalvageValuePercent", 10, "Fixed truck prv.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedTrailerSalvageValuePercent, "spec.vehicle.fixedTrailerSalvageValuePercent", 10, "Fixed trailer prv.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedDepreciationPercent, "spec.vehicle.fixedDepreciationPercent", 10, "Fixed depreciation percent - age controlled depreciation.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedInsuranceMuPerYear, "spec.vehicle.fixedInsuranceMuPerYear", 10, "Fixed insurance.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedVehicleTaxMuPerYear, "spec.vehicle.fixedVehicleTaxMuPerYear", 10, "Fixed vehicle tax.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedOtherCostsMuPerYear, "spec.vehicle.fixedOtherCostsMuPerYear", 10, "Fixed other costs.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.FixedTollMuPerYear, "spec.vehicle.fixedTollMuPerYear", 10, "Fixed toll.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TruckMaintenanceDistanceMuPerKm, "spec.vehicle.truckMaintenanceDistanceMuPerKm", 10, "Truck maintenance distance.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TruckSteeringTireCount, "spec.vehicle.truckSteeringTireCount", 10, "Truck steering number tires.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TruckAxle2TireCount, "spec.vehicle.truckAxle2TireCount", 10, "Truck axle2 number tires.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TruckAxle3TireCount, "spec.vehicle.truckAxle3TireCount", 10, "Truck axle3 number tires.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TrailerTireCount, "spec.vehicle.trailerTireCount", 10, "Trailer number tires.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceNewSteeringMuPerTire, "spec.vehicle.tirePriceNewSteeringMuPerTire", 10, "Tire price new steering.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceRetreadedAxle2MuPerTire, "spec.vehicle.tirePriceRetreadedAxle2MuPerTire", 10, "Tire price retreaded axle2.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceRetreadedAxle3MuPerTire, "spec.vehicle.tirePriceRetreadedAxle3MuPerTire", 10, "Tire price retreaded axle3.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceRetreadedTrailerMuPerTire, "spec.vehicle.tirePriceRetreadedTrailerMuPerTire", 10, "Tire price retreaded trailer.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceNewAxle2MuPerTire, "spec.vehicle.tirePriceNewAxle2MuPerTire", 10, "Tire price new axle2.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceNewAxle3MuPerTire, "spec.vehicle.tirePriceNewAxle3MuPerTire", 10, "Tire price new axle3.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceNewTrailerMuPerTire, "spec.vehicle.tirePriceNewTrailerMuPerTire", 10, "Tire price new trailer.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeRetreadedAxle2Km, "spec.vehicle.tireLifetimeRetreadedAxle2Km", 10, "Tire lifetime retreaded axle2.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeRetreadedAxle3Km, "spec.vehicle.tireLifetimeRetreadedAxle3Km", 10, "Tire lifetime retreaded axle3.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeRetreadedTrailerKm, "spec.vehicle.tireLifetimeRetreadedTrailerKm", 10, "Tire lifetime retreaded trailer.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeNewSteeringKm, "spec.vehicle.tireLifetimeNewSteeringKm", 10, "Tire lifetime new steering.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeNewAxle2Km, "spec.vehicle.tireLifetimeNewAxle2Km", 10, "Tire lifetime new axle2.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeNewAxle3Km, "spec.vehicle.tireLifetimeNewAxle3Km", 10, "Tire lifetime new axle3.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireLifetimeNewTrailerKm, "spec.vehicle.tireLifetimeNewTrailerKm", 10, "Tire lifetime new trailer.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TiresPerCycleNewCount, "spec.vehicle.tiresPerCycleNewCount", 10, "Tires per cycle new.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TiresPerCycleRetreadedCount, "spec.vehicle.tiresPerCycleRetreadedCount", 10, "Tires per cycle retreaded.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TirePriceChangeAndBalanceMuPerTire, "spec.vehicle.tirePriceChangeAndBalanceMuPerTire", 10, "Tire price change n balance.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TruckEnergyConsumptionJPerKm, "spec.vehicle.truckEnergyConsumptionJPerKm", 10, "Truck consumption distance.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Vehicle.TireDamageFrequencyNewKm, "spec.vehicle.tireDamageFrequencyNewKm", 10, "Tire damage frequency new.\nHow frequently a new tire will blow out in terms of distance.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary = new(v1beta1.Spec_Salary)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.DriverSalaryRegularWeekdayMuPerHour, "spec.salary.driverSalaryRegularWeekdayMuPerHour", 10, "Driver salary r weekday.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.DriverSalaryOvertimeWeekdayMuPerHour, "spec.salary.driverSalaryOvertimeWeekdayMuPerHour", 10, "Driver salary o weekday.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.DriverSalaryRegularWeekendMuPerHour, "spec.salary.driverSalaryRegularWeekendMuPerHour", 10, "Driver salary r weekend.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.SupportStaffSalaryWeekdayMuPerHour, "spec.salary.supportStaffSalaryWeekdayMuPerHour", 10, "Support staff salary weekday.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.SupportStaffSalaryOvertimeMuPerHour, "spec.salary.supportStaffSalaryOvertimeMuPerHour", 10, "Support staff salary overtime.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Salary.SupportStaffSalaryWeekendMuPerHour, "spec.salary.supportStaffSalaryWeekendMuPerHour", 10, "Support staff salary weekend.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy = new(v1beta1.Spec_FuelAndEnergy)
	// TODO: enum EnergyType
	// TODO: enum EnergySource1
	// TODO: enum EnergySource2
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.OtherEngineLiquidsConsumptionLPerFuelL, "spec.fuelAndEnergy.otherEngineLiquidsConsumptionLPerFuelL", 10, "Fuel additive consumption.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.OtherEngineLiquidsPriceMuPerL, "spec.fuelAndEnergy.otherEngineLiquidsPriceMuPerL", 10, "Fuel additive price.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.FuelEnergyInputEfficiencyPercent, "spec.fuelAndEnergy.fuelEnergyInputEfficiencyPercent", 10, "Fuel energy input efficiency.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.DieselUnitPriceMuPerL, "spec.fuelAndEnergy.dieselUnitPriceMuPerL", 10, "Diesel unit price.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.HvoUnitPriceMultiplier, "spec.fuelAndEnergy.hvoUnitPriceMultiplier", 10, "HVO unit price multiplier.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.HvoConsumptionMultiplier, "spec.fuelAndEnergy.hvoConsumptionMultiplier", 10, "HVO consumption multiplier.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.HvoUnitPriceMuPerL, "spec.fuelAndEnergy.hvoUnitPriceMuPerL", 10, "HVO unit price.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.CommercialUnitPriceMuPerJ, "spec.fuelAndEnergy.commercialUnitPriceMuPerJ", 10, "Commercial unit price.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.FuelAndEnergy.NoncommercialUnitPriceMuPerJ, "spec.fuelAndEnergy.noncommercialUnitPriceMuPerJ", 10, "Noncommercial unit price.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions = new(v1beta1.Spec_Emissions)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.DieselUnitEmissionsCo2KgPerL, "spec.emissions.dieselUnitEmissionsCo2KgPerL", 10, "Diesel unit emissions.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.HvoConsumptionMultiplier, "spec.emissions.hvoConsumptionMultiplier", 10, "Hvo consumption multiplier.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.HvoEmissionsMultiplier, "spec.emissions.hvoEmissionsMultiplier", 10, "Hvo emission multiplier.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.HvoUnitEmissionsCo2KgPerL, "spec.emissions.hvoUnitEmissionsCo2KgPerL", 10, "Hvo unit emissions.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.NoncommercialUnitEmissionsCo2KgPerJ, "spec.emissions.noncommercialUnitEmissionsCo2KgPerJ", 10, "Noncommercial unit emissions.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.CommercialUnitEmissionsCo2KgPerJ, "spec.emissions.commercialUnitEmissionsCo2KgPerJ", 10, "Commercial unit emissions.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Emissions.TotalVehicleEmbodiedEmissionsCo2Kg, "spec.emissions.totalVehicleEmbodiedEmissionsCo2Kg", 10, "Total vehicle embodied emissions.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Utilization = new(v1beta1.Spec_Utilization)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Utilization.UtilizationOperatingDays, "spec.utilization.utilizationOperatingDays", 10, "Utilization operating days.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance = new(v1beta1.Spec_Finance)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.LeaseResidualValuePercent, "spec.finance.leaseResidualValuePercent", 10, "Percentage of acquisition value left at end of lease.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.DownPaymentPercent, "spec.finance.downPaymentPercent", 10, "Down payment, as a percentage of acquisition value.")
	einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.LeaseDuration = new(durationpb.Duration)
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int64Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.LeaseDuration.Seconds, "spec.finance.leaseDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Int32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.LeaseDuration.Nanos, "spec.finance.leaseDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.InterestRateLeasePercentPerYear, "spec.finance.interestRateLeasePercentPerYear", 10, "Yearly interest rate on lease of truck.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.InterestRateLoanDownpaymentPercentPerYear, "spec.finance.interestRateLoanDownpaymentPercentPerYear", 10, "Yearly interest rate on loan for downpayment.")
	einride_cem_v1beta1_CEMService_CalculateResult.Flags().Float32Var(&einride_cem_v1beta1_CEMService_CalculateResult_Request.Spec.Finance.InterestPriorVehicleDeliveryMu, "spec.finance.interestPriorVehicleDeliveryMu", 10, "Total interest paid prior to recieving the vehicle.")
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_BatchCalculateResults)

	// TODO: list: Specs message
}
