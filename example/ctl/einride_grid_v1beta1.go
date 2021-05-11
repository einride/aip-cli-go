package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_grid_v1beta1_ChargerService = &cobra.Command{
	Use: "einride.grid.v1beta1.Charger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.grid.v1beta1.Charger called")
	},
}

var einride_grid_v1beta1_ChargerService_CreateChargingPool = &cobra.Command{
	Use: "CreateChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_ListChargingPools = &cobra.Command{
	Use: "ListChargingPools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListChargingPools called")
	},
}

var einride_grid_v1beta1_ChargerService_GetChargingPool = &cobra.Command{
	Use: "GetChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_ReferenceChargingPool = &cobra.Command{
	Use: "ReferenceChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReferenceChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_BatchGetChargingPools = &cobra.Command{
	Use: "BatchGetChargingPools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetChargingPools called")
	},
}

var einride_grid_v1beta1_ChargerService_UpdateChargingPool = &cobra.Command{
	Use: "UpdateChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_DeleteChargingPool = &cobra.Command{
	Use: "DeleteChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_UndeleteChargingPool = &cobra.Command{
	Use: "UndeleteChargingPool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UndeleteChargingPool called")
	},
}

var einride_grid_v1beta1_ChargerService_GetCharger = &cobra.Command{
	Use: "GetCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_ReferenceCharger = &cobra.Command{
	Use: "ReferenceCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReferenceCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_CreateCharger = &cobra.Command{
	Use: "CreateCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_UpdateCharger = &cobra.Command{
	Use: "UpdateCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_DeleteCharger = &cobra.Command{
	Use: "DeleteCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_UndeleteCharger = &cobra.Command{
	Use: "UndeleteCharger",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UndeleteCharger called")
	},
}

var einride_grid_v1beta1_ChargerService_ListChargers = &cobra.Command{
	Use: "ListChargers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListChargers called")
	},
}

var einride_grid_v1beta1_ChargerService_BatchGetChargers = &cobra.Command{
	Use: "BatchGetChargers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetChargers called")
	},
}

func init() {
	rootCmd.AddCommand(einride_grid_v1beta1_ChargerService)
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
