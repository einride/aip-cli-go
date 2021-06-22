package carrierv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	log "log"
)

// einride.carrier.v1beta1.VehicleService.
var (
	einride_carrier_v1beta1_VehicleServiceClient v1beta1.VehicleServiceClient
	einride_carrier_v1beta1_VehicleService       = &cobra.Command{
		Use: "einride.carrier.v1beta1.VehicleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_carrier_v1beta1_VehicleServiceClient = v1beta1.NewVehicleServiceClient(conn)
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.CreateVehicle.
var (
	einride_carrier_v1beta1_VehicleService_CreateVehicle_Request v1beta1.CreateVehicleRequest
	einride_carrier_v1beta1_VehicleService_CreateVehicle         = &cobra.Command{
		Use: "CreateVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.CreateVehicle")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.GetVehicle.
var (
	einride_carrier_v1beta1_VehicleService_GetVehicle_Request v1beta1.GetVehicleRequest
	einride_carrier_v1beta1_VehicleService_GetVehicle         = &cobra.Command{
		Use: "GetVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.GetVehicle")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.BatchGetVehicles.
var (
	einride_carrier_v1beta1_VehicleService_BatchGetVehicles_Request v1beta1.BatchGetVehiclesRequest
	einride_carrier_v1beta1_VehicleService_BatchGetVehicles         = &cobra.Command{
		Use: "BatchGetVehicles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.BatchGetVehicles")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.UpdateVehicle.
var (
	einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request v1beta1.UpdateVehicleRequest
	einride_carrier_v1beta1_VehicleService_UpdateVehicle         = &cobra.Command{
		Use: "UpdateVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.UpdateVehicle")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.ListVehicles.
var (
	einride_carrier_v1beta1_VehicleService_ListVehicles_Request v1beta1.ListVehiclesRequest
	einride_carrier_v1beta1_VehicleService_ListVehicles         = &cobra.Command{
		Use: "ListVehicles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.ListVehicles")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.DeleteVehicle.
var (
	einride_carrier_v1beta1_VehicleService_DeleteVehicle_Request v1beta1.DeleteVehicleRequest
	einride_carrier_v1beta1_VehicleService_DeleteVehicle         = &cobra.Command{
		Use: "DeleteVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.DeleteVehicle")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.UndeleteVehicle.
var (
	einride_carrier_v1beta1_VehicleService_UndeleteVehicle_Request v1beta1.UndeleteVehicleRequest
	einride_carrier_v1beta1_VehicleService_UndeleteVehicle         = &cobra.Command{
		Use: "UndeleteVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.UndeleteVehicle")
			return nil
		},
	}
)

// einride.carrier.v1beta1.VehicleService.SearchVehicles.
var (
	einride_carrier_v1beta1_VehicleService_SearchVehicles_Request v1beta1.SearchVehiclesRequest
	einride_carrier_v1beta1_VehicleService_SearchVehicles         = &cobra.Command{
		Use: "SearchVehicles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.carrier.v1beta1.VehicleService.SearchVehicles")
			return nil
		},
	}
)

func AddVehicleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_VehicleService)
}

func init() {
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_CreateVehicle)

	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Parent, "parent", "", "Resource name of the parent carrier where this vehicle will be created.")

	einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Vehicle = new(v1beta1.Vehicle)
	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Vehicle.Name, "vehicle.name", "", "The resource name of the vehicle.")
	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Vehicle.Etag, "vehicle.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Vehicle.VehicleType, "vehicle.vehicleType", "", "Resource name of the vehicle's type.")
	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.Vehicle.DisplayName, "vehicle.displayName", "", "Free-text display name for the vehicle.")

	einride_carrier_v1beta1_VehicleService_CreateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_CreateVehicle_Request.VehicleId, "vehicleId", "", "The ID to use for the vehicle, which will become the final component of\nthe vehicle's resource name.\n\nIf a vehicle ID is not provided, a vehicle ID will be automatically\ngenerated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_GetVehicle)

	einride_carrier_v1beta1_VehicleService_GetVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_GetVehicle_Request.Name, "name", "", "Resource name of the vehicle to retrieve.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_BatchGetVehicles)

	einride_carrier_v1beta1_VehicleService_BatchGetVehicles.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_BatchGetVehicles_Request.Parent, "parent", "", "Resource name of the parent carrier shared by all vehicles being retrieved.\nIf this is set, the parent of all of the vehicles specified in `names`\nmust match this field.")

	einride_carrier_v1beta1_VehicleService_BatchGetVehicles.Flags().StringSliceVar(&einride_carrier_v1beta1_VehicleService_BatchGetVehicles_Request.Names, "names", []string{}, "Resource names of the vehicles to retrieve.\nA maximum of 1000 vehicles can be retrieved in a batch.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UpdateVehicle)

	einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.Vehicle = new(v1beta1.Vehicle)
	einride_carrier_v1beta1_VehicleService_UpdateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.Vehicle.Name, "vehicle.name", "", "The resource name of the vehicle.")
	einride_carrier_v1beta1_VehicleService_UpdateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.Vehicle.Etag, "vehicle.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_carrier_v1beta1_VehicleService_UpdateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.Vehicle.VehicleType, "vehicle.vehicleType", "", "Resource name of the vehicle's type.")
	einride_carrier_v1beta1_VehicleService_UpdateVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.Vehicle.DisplayName, "vehicle.displayName", "", "Free-text display name for the vehicle.")

	einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_carrier_v1beta1_VehicleService_UpdateVehicle.Flags().StringSliceVar(&einride_carrier_v1beta1_VehicleService_UpdateVehicle_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_ListVehicles)

	einride_carrier_v1beta1_VehicleService_ListVehicles.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_ListVehicles_Request.Parent, "parent", "", "Resource name of the parent carrier.")

	einride_carrier_v1beta1_VehicleService_ListVehicles.Flags().Int32Var(&einride_carrier_v1beta1_VehicleService_ListVehicles_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_carrier_v1beta1_VehicleService_ListVehicles.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_ListVehicles_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_DeleteVehicle)

	einride_carrier_v1beta1_VehicleService_DeleteVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_DeleteVehicle_Request.Name, "name", "", "Resource name of the vehicle to delete.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UndeleteVehicle)

	einride_carrier_v1beta1_VehicleService_UndeleteVehicle.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_UndeleteVehicle_Request.Name, "name", "", "Resource name of the vehicle to undelete.")
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_SearchVehicles)

	einride_carrier_v1beta1_VehicleService_SearchVehicles.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_SearchVehicles_Request.Parent, "parent", "", "Resource name of the parent carrier.")

	einride_carrier_v1beta1_VehicleService_SearchVehicles.Flags().Int32Var(&einride_carrier_v1beta1_VehicleService_SearchVehicles_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_carrier_v1beta1_VehicleService_SearchVehicles.Flags().StringVar(&einride_carrier_v1beta1_VehicleService_SearchVehicles_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
