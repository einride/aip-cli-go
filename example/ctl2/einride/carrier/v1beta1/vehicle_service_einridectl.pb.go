package carrierv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
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
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_GetVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_BatchGetVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UpdateVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_ListVehicles)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_DeleteVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_UndeleteVehicle)
	einride_carrier_v1beta1_VehicleService.AddCommand(einride_carrier_v1beta1_VehicleService_SearchVehicles)
}
