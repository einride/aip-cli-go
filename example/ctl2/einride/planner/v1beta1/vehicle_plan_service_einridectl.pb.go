package plannerv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	protojson "google.golang.org/protobuf/encoding/protojson"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// einride.planner.v1beta1.VehiclePlanService.
var (
	einride_planner_v1beta1_VehiclePlanServiceClient v1beta1.VehiclePlanServiceClient
	einride_planner_v1beta1_VehiclePlanService       = &cobra.Command{
		Use:   "einride.planner.v1beta1.VehiclePlanService",
		Short: "Vehicle plan service.",
		Long:  "Vehicle plan service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_VehiclePlanServiceClient = v1beta1.NewVehiclePlanServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CreateVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request v1beta1.CreateVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan         = &cobra.Command{
		Use: "CreateVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.CreateVehiclePlan(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.GetVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan_Request v1beta1.GetVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan         = &cobra.Command{
		Use: "GetVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.GetVehiclePlan(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchGetVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans_Request v1beta1.BatchGetVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans         = &cobra.Command{
		Use: "BatchGetVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.BatchGetVehiclePlans(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ShipperSearchVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request v1beta1.ShipperSearchVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks         = &cobra.Command{
		Use: "ShipperSearchVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.ShipperSearchVehicleTasks(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.SearchVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request v1beta1.SearchVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans         = &cobra.Command{
		Use: "SearchVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.SearchVehiclePlans(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ListVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request v1beta1.ListVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans         = &cobra.Command{
		Use: "ListVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.ListVehiclePlans(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeleteVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan_Request v1beta1.DeleteVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan         = &cobra.Command{
		Use: "DeleteVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.DeleteVehiclePlan(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UndeleteVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan_Request v1beta1.UndeleteVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan         = &cobra.Command{
		Use: "UndeleteVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.UndeleteVehiclePlan(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CreateVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request v1beta1.CreateVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask         = &cobra.Command{
		Use: "CreateVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.CreateVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchCreateVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks_Request v1beta1.BatchCreateVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks         = &cobra.Command{
		Use: "BatchCreateVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.BatchCreateVehicleTasks(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UpdateVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request v1beta1.UpdateVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask         = &cobra.Command{
		Use: "UpdateVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.UpdateVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ListVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request v1beta1.ListVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks         = &cobra.Command{
		Use: "ListVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.ListVehicleTasks(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.GetVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_GetVehicleTask_Request v1beta1.GetVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_GetVehicleTask         = &cobra.Command{
		Use: "GetVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.GetVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_GetVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchGetVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks_Request v1beta1.BatchGetVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks         = &cobra.Command{
		Use: "BatchGetVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.BatchGetVehicleTasks(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask_Request v1beta1.DeleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask         = &cobra.Command{
		Use: "DeleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.DeleteVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UndeleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask_Request v1beta1.UndeleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask         = &cobra.Command{
		Use: "UndeleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.UndeleteVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.StartVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask_Request v1beta1.StartVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask         = &cobra.Command{
		Use: "StartVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.StartVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_StartVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CompleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask_Request v1beta1.CompleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask         = &cobra.Command{
		Use: "CompleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.CompleteVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.AssignShipmentToPickupShipmentVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request v1beta1.AssignShipmentToPickupShipmentVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask         = &cobra.Command{
		Use: "AssignShipmentToPickupShipmentVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.AssignShipmentToPickupShipmentVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeassignShipmentFromPickupShipmentVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask_Request v1beta1.DeassignShipmentFromPickupShipmentVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask         = &cobra.Command{
		Use: "DeassignShipmentFromPickupShipmentVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.DeassignShipmentFromPickupShipmentVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.RejectShipmentInPickupVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask_Request v1beta1.RejectShipmentInPickupVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask         = &cobra.Command{
		Use: "RejectShipmentInPickupVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.RejectShipmentInPickupVehicleTask(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ApproveDocking.
var (
	einride_planner_v1beta1_VehiclePlanService_ApproveDocking_Request v1beta1.ApproveDockingRequest
	einride_planner_v1beta1_VehiclePlanService_ApproveDocking         = &cobra.Command{
		Use: "ApproveDocking",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_VehiclePlanServiceClient.ApproveDocking(cmd.Context(), &einride_planner_v1beta1_VehiclePlanService_ApproveDocking_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddVehiclePlanServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_VehiclePlanService)
}

func init() {
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan)

	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request.Parent, "parent", "", "Resource name of the parent vehicle where this vehicle plan will be created.")

	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request.VehiclePlan = new(v1beta1.VehiclePlan)
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request.VehiclePlan.Name, "vehiclePlan.name", "", "The resource name of the vehicle plan.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request.VehiclePlan.Etag, "vehiclePlan.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")

	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request.VehiclePlanId, "vehiclePlanId", "", "The ID to use for the vehicle plan, which will become the final component\nof the vehicle plan's resource name.\n\nIf a vehicle plan ID is not provided, a vehicle plan ID will be\nautomatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan)

	einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan_Request.Name, "name", "", "Resource name of the vehicle plan to retrieve.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans)

	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans_Request.Parent, "parent", "", "Resource name of the parent vehicle shared by all vehicle plans being retrieved.\nIf this is set, the parent of all of the vehicle plans specified in `names`\nmust match this field.")

	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans_Request.Names, "names", []string{}, "Resource names of the vehicle plans to retrieve.\nA maximum of 1000 vehicle plans can be retrieved in a batch.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks)

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.Shipper, "shipper", "", "Resource name of the shipper shared by all of the vehicle tasks\nbeing searched for.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.Carrier, "carrier", "", "Resource name of the carrier to search vehicle tasks for.\nIt is valid to use a wildcard carrier resource name, \"carriers/-\",\nto search for vehicle tasks across carriers.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.Sites, "sites", []string{}, "Resource names of the sites to filter on.\nFilters vehicle plans based on pickup or delivery task site.\nIf multiple values are set for this field, they are joined with OR.\n\nTo get all vehicle plans with a pickup or deliver task from one specific site:\nsites = [\"shippers/shipper-1/sites/site-1\"]\n\nTo get all vehicle plans for two sites:\nsites = [\"shippers/shipper-1/sites/site-1\",\"shippers/shipper-1/sites/site-2\"]\n\nIf no value is set, no filter on sites are applied, and instead all tasks related\nto a shipper are returned.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartEarliestTime.Seconds, "startEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartEarliestTime.Nanos, "startEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartLatestTime.Seconds, "startLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.StartLatestTime.Nanos, "startLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteEarliestTime.Seconds, "completeEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteEarliestTime.Nanos, "completeEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteLatestTime.Seconds, "completeLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.CompleteLatestTime.Nanos, "completeLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().BoolVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.ShowDeleted, "showDeleted", false, "Flag indicating whether to include deleted results.")

	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request.Vehicle, "vehicle", "", "Resource name of the vehicle to search vehicle tasks for.\nIf set, must be a resource name without wilcards.")

	// TODO: enum State
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans)

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.Parent, "parent", "", "Resource name of the parent vehicle shared by all vehicle plans being searched for.\nIt is valid to use a wildcard vehicle resource name, \"carriers/{carrier}/vehicles/-\",\nto search for vehicle plans across all vehicles for a carrier\nor a wildcard carrier and vehicle resource name, \"carriers/-/vehicles/-\",\nto search for vehicle plans across all vehicles and carriers")

	// TODO: list: VehiclePlanState enum

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.ScheduleBooking, "scheduleBooking", []string{}, "Resource names of connected schedule bookings.\nIf not empty, search for vehicle plans generated from on of these schedule bookings")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.Sites, "sites", []string{}, "Resource names of the pickup or delivery sites to filter on.\nFilters vehicle plans based on pickup or delivery task site.\nWithin a filter the arguments are joined with AND, and multiple\nfilters are joined with OR.\n\nTo get all vehicle plans with a pickup or deliver task from one specific site:\nsites = [\"shippers/shipper-1/sites/site-1\"]\n\nTo get all vehicle plans for two sites:\nsites = [\"shippers/shipper-1/sites/site-1\",\"shippers/shipper-1/sites/site-2\"]\n\nThis search filter is no longer supported by the backend and entries will be ignored")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartEarliestTime.Seconds, "startEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartEarliestTime.Nanos, "startEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartLatestTime.Seconds, "startLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.StartLatestTime.Nanos, "startLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteEarliestTime.Seconds, "completeEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteEarliestTime.Nanos, "completeEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteLatestTime.Seconds, "completeLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.CompleteLatestTime.Nanos, "completeLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().BoolVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.ShowDeleted, "showDeleted", false, "Flag indicating whether to include deleted results.")

	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request.OrderBy, "orderBy", "", "Order of the returned vehicle plans.\nIf not specified, vehicle plans are ordered by planned_start_time ascending.\n\nFormat:\n```\norder_by: field [{ asc | desc }] [, ...]\nfield: { planned_start_time | planned_complete_time }\n```\n\nExample:\n```\norder_by = \"planned_complete_time asc, planned_start_time asc\"\n```")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans)

	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request.Parent, "parent", "", "Resource name of the parent vehicle.\nIt is possible to list vehicle plans across vehicles\nwith \"carriers/{carrier]/vehicles/-\".")

	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan)

	einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan_Request.Name, "name", "", "Resource name of the vehicle plan to delete.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan)

	einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan_Request.Name, "name", "", "Resource name of the vehicle plan to undelete.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.Parent, "parent", "", "Resource name of the parent vehicle plan where this vehicle task will be created.")

	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask = new(v1beta1.VehicleTask)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Name, "vehicleTask.name", "", "The resource name.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Etag, "vehicleTask.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task = new(v1beta1.Task)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.Duration = new(durationpb.Duration)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.Duration.Seconds, "vehicleTask.task.duration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.Duration.Nanos, "vehicleTask.task.duration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EnergyJ, "vehicleTask.task.energyJ", 10, "Energy delta of the task (J).\nNegative when driving or operating, and positive when charging.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.Description, "vehicleTask.task.description", "", "Free-text description of the task.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartTime = new(timeofday.TimeOfDay)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartTime.Hours, "vehicleTask.task.startTime.hours", 10, "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose\nto allow the value \"24:00:00\" for scenarios like business closing time.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartTime.Minutes, "vehicleTask.task.startTime.minutes", 10, "Minutes of hour of day. Must be from 0 to 59.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartTime.Seconds, "vehicleTask.task.startTime.seconds", 10, "Seconds of minutes of the time. Must normally be from 0 to 59. An API may\nallow the value 60 if it allows leap-seconds.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartTime.Nanos, "vehicleTask.task.startTime.nanos", 10, "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndTime = new(timeofday.TimeOfDay)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndTime.Hours, "vehicleTask.task.endTime.hours", 10, "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose\nto allow the value \"24:00:00\" for scenarios like business closing time.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndTime.Minutes, "vehicleTask.task.endTime.minutes", 10, "Minutes of hour of day. Must be from 0 to 59.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndTime.Seconds, "vehicleTask.task.endTime.seconds", 10, "Seconds of minutes of the time. Must normally be from 0 to 59. An API may\nallow the value 60 if it allows leap-seconds.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndTime.Nanos, "vehicleTask.task.endTime.nanos", 10, "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.StartEnergyJ, "vehicleTask.task.startEnergyJ", 10, "Assumed energy before performing the task.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.Task.EndEnergyJ, "vehicleTask.task.endEnergyJ", 10, "Assumed energy after performing the task. Must be equal to start_energy_j + energy_j.")
	// TODO: one-of: PickupShipment
	// TODO: one-of: DeliverShipment
	// TODO: one-of: DriveBetweenSites
	// TODO: one-of: WaitAtSite
	// TODO: one-of: ChargeAtSite
	// TODO: one-of: AwaitDockingApproval
	// TODO: one-of: DockAtSite
	// TODO: one-of: PickupShipmentWhileCharging
	// TODO: one-of: DeliverShipmentWhileCharging
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedStartTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedStartTime.Seconds, "vehicleTask.plannedStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedStartTime.Nanos, "vehicleTask.plannedStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedCompleteTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedCompleteTime.Seconds, "vehicleTask.plannedCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTask.PlannedCompleteTime.Nanos, "vehicleTask.plannedCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request.VehicleTaskId, "vehicleTaskId", "", "The ID to use for the resource, which will become the final component of\nthe vehicle plan's resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\n/[a-zA-Z0-9]/.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks)

	einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks_Request.Parent, "parent", "", "Resource name of the parent vehicle plan shared by all vehicle tasks being created.\n\nIf this is set, the parent field in the request messages must either be\nempty or match this field.")

	// TODO: list: Requests message
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask = new(v1beta1.VehicleTask)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Name, "vehicleTask.name", "", "The resource name.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Etag, "vehicleTask.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task = new(v1beta1.Task)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.Duration = new(durationpb.Duration)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.Duration.Seconds, "vehicleTask.task.duration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.Duration.Nanos, "vehicleTask.task.duration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EnergyJ, "vehicleTask.task.energyJ", 10, "Energy delta of the task (J).\nNegative when driving or operating, and positive when charging.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.Description, "vehicleTask.task.description", "", "Free-text description of the task.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartTime = new(timeofday.TimeOfDay)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartTime.Hours, "vehicleTask.task.startTime.hours", 10, "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose\nto allow the value \"24:00:00\" for scenarios like business closing time.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartTime.Minutes, "vehicleTask.task.startTime.minutes", 10, "Minutes of hour of day. Must be from 0 to 59.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartTime.Seconds, "vehicleTask.task.startTime.seconds", 10, "Seconds of minutes of the time. Must normally be from 0 to 59. An API may\nallow the value 60 if it allows leap-seconds.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartTime.Nanos, "vehicleTask.task.startTime.nanos", 10, "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndTime = new(timeofday.TimeOfDay)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndTime.Hours, "vehicleTask.task.endTime.hours", 10, "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose\nto allow the value \"24:00:00\" for scenarios like business closing time.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndTime.Minutes, "vehicleTask.task.endTime.minutes", 10, "Minutes of hour of day. Must be from 0 to 59.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndTime.Seconds, "vehicleTask.task.endTime.seconds", 10, "Seconds of minutes of the time. Must normally be from 0 to 59. An API may\nallow the value 60 if it allows leap-seconds.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndTime.Nanos, "vehicleTask.task.endTime.nanos", 10, "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.StartEnergyJ, "vehicleTask.task.startEnergyJ", 10, "Assumed energy before performing the task.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Float32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.Task.EndEnergyJ, "vehicleTask.task.endEnergyJ", 10, "Assumed energy after performing the task. Must be equal to start_energy_j + energy_j.")
	// TODO: one-of: PickupShipment
	// TODO: one-of: DeliverShipment
	// TODO: one-of: DriveBetweenSites
	// TODO: one-of: WaitAtSite
	// TODO: one-of: ChargeAtSite
	// TODO: one-of: AwaitDockingApproval
	// TODO: one-of: DockAtSite
	// TODO: one-of: PickupShipmentWhileCharging
	// TODO: one-of: DeliverShipmentWhileCharging
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedStartTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedStartTime.Seconds, "vehicleTask.plannedStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedStartTime.Nanos, "vehicleTask.plannedStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedCompleteTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int64Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedCompleteTime.Seconds, "vehicleTask.plannedCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.VehicleTask.PlannedCompleteTime.Nanos, "vehicleTask.plannedCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks)

	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request.Parent, "parent", "", "Resource name of the parent vehicle plan.\nIt is possible to list vehicle tasks across vehicle plans\nwith \"carriers/{carrier]/vehicles/{vehicle}/plans/-\".\nIt is possible to list vehicle plans across vehicles\nwith \"carriers/{carrier]/vehicles/-/plans/-\".")

	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks.Flags().Int32Var(&einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_GetVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_GetVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_GetVehicleTask_Request.Name, "name", "", "Resource name of the vehicle task to retrieve.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks)

	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks_Request.Parent, "parent", "", "Resource name of the parent vehicle plan shared by all resources being retrieved. If this is\nset, the parent of all of the resources specified in `names` must match\nthis field.")

	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks.Flags().StringSliceVar(&einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks_Request.Names, "names", []string{}, "Resource names of the vehicle tasks to retrieve.\nA maximum of 1000 resources can be retrieved in a batch.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask_Request.Name, "name", "", "Resource name of the vehicle plan to delete.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask_Request.Name, "name", "", "Resource name of the vehicle plan to undelete.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_StartVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_StartVehicleTask_Request.Name, "name", "", "Resource name of the vehicle plan to delete.")

	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_StartVehicleTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask_Request.Name, "name", "", "Resource name of the vehicle plan to complete.")

	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request.Name, "name", "", "Resource name of the vehicle task which the shipment should be assigned to.")

	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request.Shipment, "shipment", "", "Resource name of the shipment to assign to the pickup task.\nIf not specified, one will automatically be selected based on:\n- the origin site\n- the destination site\n- the vehicle capacity")

	// TODO: enum ShipmentAssignmentSource

	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask_Request.Name, "name", "", "Resource name of the vehicle task which the shipment should be deassigned from.")

	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask)

	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask_Request.Name, "name", "", "Resource name of the vehicle task which the shipment should be deassigned from.\nThe shipment will be rejected after deassigning.")

	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask_Request.Comment, "comment", "", "Comment containing a reason for rejecting the shipment.")
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ApproveDocking)

	einride_planner_v1beta1_VehiclePlanService_ApproveDocking.Flags().StringVar(&einride_planner_v1beta1_VehiclePlanService_ApproveDocking_Request.VehicleTask, "vehicleTask", "", "Resource name of the vehicle task to approve docking for.\nIf the vehicle task is not an AwaitDockingApproval task, FailedPrecondition is returned.\nIf the vehicle task is not STARTED, FailedPrecondition is returned.")
}
