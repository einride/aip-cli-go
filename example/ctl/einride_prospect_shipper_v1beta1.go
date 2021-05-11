package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_prospect_shipper_v1beta1_ProspectiveShipperService = &cobra.Command{
	Use: "einride.prospect.shipper.v1beta1.ProspectiveShipper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.prospect.shipper.v1beta1.ProspectiveShipper called")
	},
}

func init() {
	rootCmd.AddCommand(einride_prospect_shipper_v1beta1_ProspectiveShipperService)
}
