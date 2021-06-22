package cemv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/cem/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.cem.v1beta1.CEMService.
var (
	einride_cem_v1beta1_CEMServiceClient v1beta1.CEMServiceClient
	einride_cem_v1beta1_CEMService       = &cobra.Command{
		Use: "einride.cem.v1beta1.CEMService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_cem_v1beta1_CEMServiceClient = v1beta1.NewCEMServiceClient(conn)
			return nil
		},
	}
)

// einride.cem.v1beta1.CEMService.CalculateResult.
var (
	einride_cem_v1beta1_CEMService_CalculateResult_Request v1beta1.CalculateResultRequest
	einride_cem_v1beta1_CEMService_CalculateResult         = &cobra.Command{
		Use: "CalculateResult",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.cem.v1beta1.CEMService.CalculateResult")
			return nil
		},
	}
)

// einride.cem.v1beta1.CEMService.BatchCalculateResults.
var (
	einride_cem_v1beta1_CEMService_BatchCalculateResults_Request v1beta1.BatchCalculateResultsRequest
	einride_cem_v1beta1_CEMService_BatchCalculateResults         = &cobra.Command{
		Use: "BatchCalculateResults",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.cem.v1beta1.CEMService.BatchCalculateResults")
			return nil
		},
	}
)

func AddCEMServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_cem_v1beta1_CEMService)
}

func init() {
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_CalculateResult)
	einride_cem_v1beta1_CEMService.AddCommand(einride_cem_v1beta1_CEMService_BatchCalculateResults)
}
