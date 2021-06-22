package carrierv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/carrier/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.carrier.v1beta1.VehicleTypeService.
var (
	einride_carrier_v1beta1_VehicleTypeServiceClient v1beta1.VehicleTypeServiceClient
	einride_carrier_v1beta1_VehicleTypeService       = &cobra.Command{
		Use: "einride.carrier.v1beta1.VehicleTypeService",
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.CreateVehicleType")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.GetVehicleType")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.BatchGetVehicleTypes")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.UpdateVehicleType")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.ListVehicleTypes")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.DeleteVehicleType")
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
			log.Println("einride.carrier.v1beta1.VehicleTypeService.UndeleteVehicleType")
			return nil
		},
	}
)

func AddVehicleTypeServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_carrier_v1beta1_VehicleTypeService)
}

func init() {
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_CreateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_GetVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_BatchGetVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UpdateVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_ListVehicleTypes)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_DeleteVehicleType)
	einride_carrier_v1beta1_VehicleTypeService.AddCommand(einride_carrier_v1beta1_VehicleTypeService_UndeleteVehicleType)
}
