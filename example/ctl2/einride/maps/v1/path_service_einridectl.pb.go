package mapsv1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	v1 "github.com/einride/proto/gen/go/einride/maps/v1"
	cobra "github.com/spf13/cobra"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	log "log"
)

// einride.maps.v1.PathService.
var (
	einride_maps_v1_PathServiceClient v1.PathServiceClient
	einride_maps_v1_PathService       = &cobra.Command{
		Use: "einride.maps.v1.PathService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_maps_v1_PathServiceClient = v1.NewPathServiceClient(conn)
			return nil
		},
	}
)

// einride.maps.v1.PathService.ComputePathSummaryMatrix.
var (
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request v1.ComputePathSummaryMatrixRequest
	einride_maps_v1_PathService_ComputePathSummaryMatrix         = &cobra.Command{
		Use: "ComputePathSummaryMatrix",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.maps.v1.PathService.ComputePathSummaryMatrix")
			return nil
		},
	}
)

// einride.maps.v1.PathService.BatchComputePathSummaryMatrices.
var (
	einride_maps_v1_PathService_BatchComputePathSummaryMatrices_Request v1.BatchComputePathSummaryMatricesRequest
	einride_maps_v1_PathService_BatchComputePathSummaryMatrices         = &cobra.Command{
		Use: "BatchComputePathSummaryMatrices",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.maps.v1.PathService.BatchComputePathSummaryMatrices")
			return nil
		},
	}
)

func AddPathServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_maps_v1_PathService)
}

func init() {
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_ComputePathSummaryMatrix)

	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType = new(v1beta1.VehicleType)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().StringVar(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.Name, "vehicleType.name", "", "The resource name of the vehicle type.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().StringVar(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.Etag, "vehicleType.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: enum Category
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().StringVar(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.DisplayName, "vehicleType.displayName", "", "Display name for the vehicle type.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.HeightMetres = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.HeightMetres.Value, "vehicleType.heightMetres.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.LengthMetres = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.LengthMetres.Value, "vehicleType.lengthMetres.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.WidthMetres = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.WidthMetres.Value, "vehicleType.widthMetres.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.AxleCount = new(wrapperspb.Int32Value)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Int32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.AxleCount.Value, "vehicleType.axleCount.value", 10, "The int32 value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.CurbWeightKilograms = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.CurbWeightKilograms.Value, "vehicleType.curbWeightKilograms.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.GrossWeightKilograms = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.GrossWeightKilograms.Value, "vehicleType.grossWeightKilograms.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.TrailerCount = new(wrapperspb.Int32Value)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Int32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.TrailerCount.Value, "vehicleType.trailerCount.value", 10, "The int32 value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.TotalBatteryCapacityJoules = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.TotalBatteryCapacityJoules.Value, "vehicleType.totalBatteryCapacityJoules.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre.Value, "vehicleType.averageFuelConsumptionLitresPerKilometre.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.LoadCapacityEurPallets = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.LoadCapacityEurPallets.Value, "vehicleType.loadCapacityEurPallets.value", 10, "The float value.")
	// TODO: enum VehicleClass
	// TODO: enum Drivetrain
	// TODO: enum DriveType
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre.Value, "vehicleType.energyConsumptionEmptyJoulesPerKilometre.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre.Value, "vehicleType.energyConsumptionPayloadJoulesPerKilogramKilometre.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond.Value, "vehicleType.energyConsumptionAuxJoulesPerSecond.value", 10, "The float value.")
	einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_maps_v1_PathService_ComputePathSummaryMatrix.Flags().Float32Var(&einride_maps_v1_PathService_ComputePathSummaryMatrix_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond.Value, "vehicleType.batteryMaxDesiredChargingPowerJoulesPerSecond.value", 10, "The float value.")
	// TODO: list: LoadCapacities message

	// TODO: list: StartLocations message

	// TODO: list: EndLocations message
	einride_maps_v1_PathService.AddCommand(einride_maps_v1_PathService_BatchComputePathSummaryMatrices)

	// TODO: list: VehicleTypes message

	// TODO: list: StartLocations message

	// TODO: list: EndLocations message
}
