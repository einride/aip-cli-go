package gridv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/grid/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.grid.v1beta1.ChargerService.
var (
	einride_grid_v1beta1_ChargerServiceClient v1beta1.ChargerServiceClient
	einride_grid_v1beta1_ChargerService       = &cobra.Command{
		Use: "einride.grid.v1beta1.ChargerService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_grid_v1beta1_ChargerServiceClient = v1beta1.NewChargerServiceClient(conn)
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.CreateChargingPool.
var (
	einride_grid_v1beta1_ChargerService_CreateChargingPool_Request v1beta1.CreateChargingPoolRequest
	einride_grid_v1beta1_ChargerService_CreateChargingPool         = &cobra.Command{
		Use: "CreateChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.CreateChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ListChargingPools.
var (
	einride_grid_v1beta1_ChargerService_ListChargingPools_Request v1beta1.ListChargingPoolsRequest
	einride_grid_v1beta1_ChargerService_ListChargingPools         = &cobra.Command{
		Use: "ListChargingPools",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.ListChargingPools")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.GetChargingPool.
var (
	einride_grid_v1beta1_ChargerService_GetChargingPool_Request v1beta1.GetChargingPoolRequest
	einride_grid_v1beta1_ChargerService_GetChargingPool         = &cobra.Command{
		Use: "GetChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.GetChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ReferenceChargingPool.
var (
	einride_grid_v1beta1_ChargerService_ReferenceChargingPool_Request v1beta1.ReferenceChargingPoolRequest
	einride_grid_v1beta1_ChargerService_ReferenceChargingPool         = &cobra.Command{
		Use: "ReferenceChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.ReferenceChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.BatchGetChargingPools.
var (
	einride_grid_v1beta1_ChargerService_BatchGetChargingPools_Request v1beta1.BatchGetChargingPoolsRequest
	einride_grid_v1beta1_ChargerService_BatchGetChargingPools         = &cobra.Command{
		Use: "BatchGetChargingPools",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.BatchGetChargingPools")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UpdateChargingPool.
var (
	einride_grid_v1beta1_ChargerService_UpdateChargingPool_Request v1beta1.UpdateChargingPoolRequest
	einride_grid_v1beta1_ChargerService_UpdateChargingPool         = &cobra.Command{
		Use: "UpdateChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.UpdateChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.DeleteChargingPool.
var (
	einride_grid_v1beta1_ChargerService_DeleteChargingPool_Request v1beta1.DeleteChargingPoolRequest
	einride_grid_v1beta1_ChargerService_DeleteChargingPool         = &cobra.Command{
		Use: "DeleteChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.DeleteChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UndeleteChargingPool.
var (
	einride_grid_v1beta1_ChargerService_UndeleteChargingPool_Request v1beta1.UndeleteChargingPoolRequest
	einride_grid_v1beta1_ChargerService_UndeleteChargingPool         = &cobra.Command{
		Use: "UndeleteChargingPool",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.UndeleteChargingPool")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.GetCharger.
var (
	einride_grid_v1beta1_ChargerService_GetCharger_Request v1beta1.GetChargerRequest
	einride_grid_v1beta1_ChargerService_GetCharger         = &cobra.Command{
		Use: "GetCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.GetCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ReferenceCharger.
var (
	einride_grid_v1beta1_ChargerService_ReferenceCharger_Request v1beta1.ReferenceChargerRequest
	einride_grid_v1beta1_ChargerService_ReferenceCharger         = &cobra.Command{
		Use: "ReferenceCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.ReferenceCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.CreateCharger.
var (
	einride_grid_v1beta1_ChargerService_CreateCharger_Request v1beta1.CreateChargerRequest
	einride_grid_v1beta1_ChargerService_CreateCharger         = &cobra.Command{
		Use: "CreateCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.CreateCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UpdateCharger.
var (
	einride_grid_v1beta1_ChargerService_UpdateCharger_Request v1beta1.UpdateChargerRequest
	einride_grid_v1beta1_ChargerService_UpdateCharger         = &cobra.Command{
		Use: "UpdateCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.UpdateCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.DeleteCharger.
var (
	einride_grid_v1beta1_ChargerService_DeleteCharger_Request v1beta1.DeleteChargerRequest
	einride_grid_v1beta1_ChargerService_DeleteCharger         = &cobra.Command{
		Use: "DeleteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.DeleteCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.UndeleteCharger.
var (
	einride_grid_v1beta1_ChargerService_UndeleteCharger_Request v1beta1.UndeleteChargerRequest
	einride_grid_v1beta1_ChargerService_UndeleteCharger         = &cobra.Command{
		Use: "UndeleteCharger",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.UndeleteCharger")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.ListChargers.
var (
	einride_grid_v1beta1_ChargerService_ListChargers_Request v1beta1.ListChargersRequest
	einride_grid_v1beta1_ChargerService_ListChargers         = &cobra.Command{
		Use: "ListChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.ListChargers")
			return nil
		},
	}
)

// einride.grid.v1beta1.ChargerService.BatchGetChargers.
var (
	einride_grid_v1beta1_ChargerService_BatchGetChargers_Request v1beta1.BatchGetChargersRequest
	einride_grid_v1beta1_ChargerService_BatchGetChargers         = &cobra.Command{
		Use: "BatchGetChargers",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.grid.v1beta1.ChargerService.BatchGetChargers")
			return nil
		},
	}
)

func AddChargerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_grid_v1beta1_ChargerService)
}

func init() {
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_CreateChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ListChargingPools)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_GetChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ReferenceChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_BatchGetChargingPools)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UpdateChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_DeleteChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UndeleteChargingPool)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_GetCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ReferenceCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_CreateCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UpdateCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_DeleteCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_UndeleteCharger)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_ListChargers)
	einride_grid_v1beta1_ChargerService.AddCommand(einride_grid_v1beta1_ChargerService_BatchGetChargers)
}
