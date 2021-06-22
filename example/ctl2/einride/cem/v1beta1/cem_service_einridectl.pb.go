package cemv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/cem/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_cem_v1beta1_CEMService = &cobra.Command{
	Use: "einride.cem.v1beta1.CEMService",
}

var einride_cem_v1beta1_CalculateResultRequest v1beta1.CalculateResultRequest
var einride_cem_v1beta1_CEMService_CalculateResult = &cobra.Command{
	Use: "CalculateResult",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.cem.v1beta1.CEMService.CalculateResult")
		return nil
	},
}

var einride_cem_v1beta1_BatchCalculateResultsRequest v1beta1.BatchCalculateResultsRequest
var einride_cem_v1beta1_CEMService_BatchCalculateResults = &cobra.Command{
	Use: "BatchCalculateResults",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.cem.v1beta1.CEMService.BatchCalculateResults")
		return nil
	},
}

func AddCEMServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_cem_v1beta1_CEMService)
}

func init() {
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_CalculateResult)
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_BatchCalculateResults)
}
