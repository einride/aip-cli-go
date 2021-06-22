package optimizerv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/optimizer/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.optimizer.v1.OptimizerService.
var (
	einride_optimizer_v1_OptimizerServiceClient v1.OptimizerServiceClient
	einride_optimizer_v1_OptimizerService       = &cobra.Command{
		Use: "einride.optimizer.v1.OptimizerService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_optimizer_v1_OptimizerServiceClient = v1.NewOptimizerServiceClient(conn)
			return nil
		},
	}
)

// einride.optimizer.v1.OptimizerService.OptimizeProblemParameters.
var (
	einride_optimizer_v1_OptimizerService_OptimizeProblemParameters_Request v1.OptimizeProblemParametersRequest
	einride_optimizer_v1_OptimizerService_OptimizeProblemParameters         = &cobra.Command{
		Use: "OptimizeProblemParameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.optimizer.v1.OptimizerService.OptimizeProblemParameters")
			return nil
		},
	}
)

func AddOptimizerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_optimizer_v1_OptimizerService)
}

func init() {
	einride_optimizer_v1_OptimizerService.AddCommand(einride_optimizer_v1_OptimizerService_OptimizeProblemParameters)
}
