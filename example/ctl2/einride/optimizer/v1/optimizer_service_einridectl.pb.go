package optimizerv1

import (
	v1 "github.com/einride/proto/gen/go/einride/optimizer/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_optimizer_v1_OptimizerService = &cobra.Command{
	Use: "einride.optimizer.v1.OptimizerService",
}

var einride_optimizer_v1_OptimizeProblemParametersRequest v1.OptimizeProblemParametersRequest
var einride_optimizer_v1_OptimizerService_OptimizeProblemParameters = &cobra.Command{
	Use: "OptimizeProblemParameters",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.optimizer.v1.OptimizerService.OptimizeProblemParameters")
		return nil
	},
}

func AddOptimizerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_optimizer_v1_OptimizerService)
}

func init() {
	einride_optimizer_v1_OptimizerService.AddCommand(einride_optimizer_v1_OptimizerService_OptimizeProblemParameters)
}
