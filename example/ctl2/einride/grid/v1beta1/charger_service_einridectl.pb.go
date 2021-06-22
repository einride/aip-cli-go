package gridv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/grid/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_grid_v1beta1_ChargerService = &cobra.Command{
	Use: "einride.grid.v1beta1.ChargerService",
}

var einride_grid_v1beta1_CreateChargingPoolRequest v1beta1.CreateChargingPoolRequest
var einride_grid_v1beta1_ChargerService_CreateChargingPool = &cobra.Command{
	Use: "CreateChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.CreateChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_ListChargingPoolsRequest v1beta1.ListChargingPoolsRequest
var einride_grid_v1beta1_ChargerService_ListChargingPools = &cobra.Command{
	Use: "ListChargingPools",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.ListChargingPools")
		return nil
	},
}

var einride_grid_v1beta1_GetChargingPoolRequest v1beta1.GetChargingPoolRequest
var einride_grid_v1beta1_ChargerService_GetChargingPool = &cobra.Command{
	Use: "GetChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.GetChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_ReferenceChargingPoolRequest v1beta1.ReferenceChargingPoolRequest
var einride_grid_v1beta1_ChargerService_ReferenceChargingPool = &cobra.Command{
	Use: "ReferenceChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.ReferenceChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_BatchGetChargingPoolsRequest v1beta1.BatchGetChargingPoolsRequest
var einride_grid_v1beta1_ChargerService_BatchGetChargingPools = &cobra.Command{
	Use: "BatchGetChargingPools",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.BatchGetChargingPools")
		return nil
	},
}

var einride_grid_v1beta1_UpdateChargingPoolRequest v1beta1.UpdateChargingPoolRequest
var einride_grid_v1beta1_ChargerService_UpdateChargingPool = &cobra.Command{
	Use: "UpdateChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.UpdateChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_DeleteChargingPoolRequest v1beta1.DeleteChargingPoolRequest
var einride_grid_v1beta1_ChargerService_DeleteChargingPool = &cobra.Command{
	Use: "DeleteChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.DeleteChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_UndeleteChargingPoolRequest v1beta1.UndeleteChargingPoolRequest
var einride_grid_v1beta1_ChargerService_UndeleteChargingPool = &cobra.Command{
	Use: "UndeleteChargingPool",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.UndeleteChargingPool")
		return nil
	},
}

var einride_grid_v1beta1_GetChargerRequest v1beta1.GetChargerRequest
var einride_grid_v1beta1_ChargerService_GetCharger = &cobra.Command{
	Use: "GetCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.GetCharger")
		return nil
	},
}

var einride_grid_v1beta1_ReferenceChargerRequest v1beta1.ReferenceChargerRequest
var einride_grid_v1beta1_ChargerService_ReferenceCharger = &cobra.Command{
	Use: "ReferenceCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.ReferenceCharger")
		return nil
	},
}

var einride_grid_v1beta1_CreateChargerRequest v1beta1.CreateChargerRequest
var einride_grid_v1beta1_ChargerService_CreateCharger = &cobra.Command{
	Use: "CreateCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.CreateCharger")
		return nil
	},
}

var einride_grid_v1beta1_UpdateChargerRequest v1beta1.UpdateChargerRequest
var einride_grid_v1beta1_ChargerService_UpdateCharger = &cobra.Command{
	Use: "UpdateCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.UpdateCharger")
		return nil
	},
}

var einride_grid_v1beta1_DeleteChargerRequest v1beta1.DeleteChargerRequest
var einride_grid_v1beta1_ChargerService_DeleteCharger = &cobra.Command{
	Use: "DeleteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.DeleteCharger")
		return nil
	},
}

var einride_grid_v1beta1_UndeleteChargerRequest v1beta1.UndeleteChargerRequest
var einride_grid_v1beta1_ChargerService_UndeleteCharger = &cobra.Command{
	Use: "UndeleteCharger",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.UndeleteCharger")
		return nil
	},
}

var einride_grid_v1beta1_ListChargersRequest v1beta1.ListChargersRequest
var einride_grid_v1beta1_ChargerService_ListChargers = &cobra.Command{
	Use: "ListChargers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.ListChargers")
		return nil
	},
}

var einride_grid_v1beta1_BatchGetChargersRequest v1beta1.BatchGetChargersRequest
var einride_grid_v1beta1_ChargerService_BatchGetChargers = &cobra.Command{
	Use: "BatchGetChargers",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.grid.v1beta1.ChargerService.BatchGetChargers")
		return nil
	},
}

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
