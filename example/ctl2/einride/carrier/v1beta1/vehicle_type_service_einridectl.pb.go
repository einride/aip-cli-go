package carrierv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// einride.carrier.v1beta1.VehicleTypeService.
var (
	einride_carrier_v1beta1_VehicleTypeServiceClient v1beta1.VehicleTypeServiceClient
	einride_carrier_v1beta1_VehicleTypeService       = &cobra.Command{
		Use:   "carrier.v1beta1.VehicleTypeService",
		Short: "Vehicle type service.",
		Long:  "Vehicle type service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_carrier_v1beta1_VehicleTypeServiceClient = v1beta1.NewVehicleTypeServiceClient(conn)
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.CreateVehicleType.
var (
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request v1beta1.CreateVehicleTypeRequest
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType         = &cobra.Command{
		Use: "CreateVehicleType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.CreateVehicleType(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.GetVehicleType.
var (
	einride_carrier_v1beta1_VehicleTypeService_GetVehicleType_Request v1beta1.GetVehicleTypeRequest
	einride_carrier_v1beta1_VehicleTypeService_GetVehicleType         = &cobra.Command{
		Use: "GetVehicleType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.GetVehicleType(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_GetVehicleType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.BatchGetVehicleTypes.
var (
	einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes_Request v1beta1.BatchGetVehicleTypesRequest
	einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes         = &cobra.Command{
		Use: "BatchGetVehicleTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.BatchGetVehicleTypes(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.UpdateVehicleType.
var (
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request v1beta1.UpdateVehicleTypeRequest
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType         = &cobra.Command{
		Use: "UpdateVehicleType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.UpdateVehicleType(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.ListVehicleTypes.
var (
	einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes_Request v1beta1.ListVehicleTypesRequest
	einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes         = &cobra.Command{
		Use: "ListVehicleTypes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.ListVehicleTypes(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.DeleteVehicleType.
var (
	einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType_Request v1beta1.DeleteVehicleTypeRequest
	einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType         = &cobra.Command{
		Use: "DeleteVehicleType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.DeleteVehicleType(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleTypeService.UndeleteVehicleType.
var (
	einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType_Request v1beta1.UndeleteVehicleTypeRequest
	einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType         = &cobra.Command{
		Use: "UndeleteVehicleType",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_carrier_v1beta1_VehicleTypeServiceClient.UndeleteVehicleType(cmd.Context(), &einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddVehicleTypeServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_VehicleTypeService)
}

func init() {
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType)

	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.Parent, "parent", "", "Resource name of the parent carrier where this vehicle type will be created.")

	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType = new(v1beta1.VehicleType)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.Name, "vehicleType.name", "", "The resource name of the vehicle type.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.Etag, "vehicleType.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: enum Category
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.DisplayName, "vehicleType.displayName", "", "Display name for the vehicle type.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.HeightMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.HeightMetres.Value, "vehicleType.heightMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.LengthMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.LengthMetres.Value, "vehicleType.lengthMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.WidthMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.WidthMetres.Value, "vehicleType.widthMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.AxleCount = new(wrapperspb.Int32Value)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Int32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.AxleCount.Value, "vehicleType.axleCount.value", 10, "The int32 value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.CurbWeightKilograms = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.CurbWeightKilograms.Value, "vehicleType.curbWeightKilograms.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.GrossWeightKilograms = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.GrossWeightKilograms.Value, "vehicleType.grossWeightKilograms.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.TrailerCount = new(wrapperspb.Int32Value)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Int32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.TrailerCount.Value, "vehicleType.trailerCount.value", 10, "The int32 value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.TotalBatteryCapacityJoules = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.TotalBatteryCapacityJoules.Value, "vehicleType.totalBatteryCapacityJoules.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre.Value, "vehicleType.averageFuelConsumptionLitresPerKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.LoadCapacityEurPallets = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.LoadCapacityEurPallets.Value, "vehicleType.loadCapacityEurPallets.value", 10, "The float value.")
	// TODO: enum VehicleClass
	// TODO: enum Drivetrain
	// TODO: enum DriveType
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre.Value, "vehicleType.energyConsumptionEmptyJoulesPerKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre.Value, "vehicleType.energyConsumptionPayloadJoulesPerKilogramKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond.Value, "vehicleType.energyConsumptionAuxJoulesPerSecond.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond.Value, "vehicleType.batteryMaxDesiredChargingPowerJoulesPerSecond.value", 10, "The float value.")
	// TODO: list: LoadCapacities message

	einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType_Request.VehicleTypeId, "vehicleTypeId", "", "The ID to use for the vehicle type, which will become the final component\nof the vehicle type's resource name.\n\nIf a vehicle type ID is not provided, a vehicle type ID will be\nautomatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_GetVehicleType)

	einride_carrier_v1beta1_VehicleTypeService_GetVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_GetVehicleType_Request.Name, "name", "", "Resource name of the vehicle type to retrieve.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes)

	einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes_Request.Parent, "parent", "", "Resource name of the parent carrier shared by all vehicle types being retrieved.\nIf this is set, the parent of all of the vehicle types specified in `names`\nmust match this field.")

	einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes.Flags().StringSliceVar(&einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes_Request.Names, "names", []string{}, "Resource names of the vehicle types to retrieve.\nA maximum of 1000 vehicle types can be retrieved in a batch.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType)

	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType = new(v1beta1.VehicleType)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.Name, "vehicleType.name", "", "The resource name of the vehicle type.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.Etag, "vehicleType.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	// TODO: enum Category
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.DisplayName, "vehicleType.displayName", "", "Display name for the vehicle type.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.HeightMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.HeightMetres.Value, "vehicleType.heightMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.LengthMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.LengthMetres.Value, "vehicleType.lengthMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.WidthMetres = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.WidthMetres.Value, "vehicleType.widthMetres.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.AxleCount = new(wrapperspb.Int32Value)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Int32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.AxleCount.Value, "vehicleType.axleCount.value", 10, "The int32 value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.CurbWeightKilograms = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.CurbWeightKilograms.Value, "vehicleType.curbWeightKilograms.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.GrossWeightKilograms = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.GrossWeightKilograms.Value, "vehicleType.grossWeightKilograms.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.TrailerCount = new(wrapperspb.Int32Value)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Int32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.TrailerCount.Value, "vehicleType.trailerCount.value", 10, "The int32 value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.TotalBatteryCapacityJoules = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.TotalBatteryCapacityJoules.Value, "vehicleType.totalBatteryCapacityJoules.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.AverageFuelConsumptionLitresPerKilometre.Value, "vehicleType.averageFuelConsumptionLitresPerKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.LoadCapacityEurPallets = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.LoadCapacityEurPallets.Value, "vehicleType.loadCapacityEurPallets.value", 10, "The float value.")
	// TODO: enum VehicleClass
	// TODO: enum Drivetrain
	// TODO: enum DriveType
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionEmptyJoulesPerKilometre.Value, "vehicleType.energyConsumptionEmptyJoulesPerKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionPayloadJoulesPerKilogramKilometre.Value, "vehicleType.energyConsumptionPayloadJoulesPerKilogramKilometre.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.EnergyConsumptionAuxJoulesPerSecond.Value, "vehicleType.energyConsumptionAuxJoulesPerSecond.value", 10, "The float value.")
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond = new(wrapperspb.FloatValue)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().Float32Var(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.VehicleType.BatteryMaxDesiredChargingPowerJoulesPerSecond.Value, "vehicleType.batteryMaxDesiredChargingPowerJoulesPerSecond.value", 10, "The float value.")
	// TODO: list: LoadCapacities message

	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType.Flags().StringSliceVar(&einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes)

	einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes_Request.Parent, "parent", "", "Resource name of the parent carrier.")

	einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes.Flags().Int32Var(&einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType)

	einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType_Request.Name, "name", "", "Resource name of the vehicle type to delete.")
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType)

	einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType.Flags().StringVar(&einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType_Request.Name, "name", "", "Resource name of the vehicle type to undelete.")
}
