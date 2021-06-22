package prospectiveshipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/prospect/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.prospect.shipper.v1beta1.ElectrificationPlannerService.
var (
	einride_prospect_shipper_v1beta1_ElectrificationPlannerServiceClient v1beta1.ElectrificationPlannerServiceClient
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService       = &cobra.Command{
		Use: "einride.prospect.shipper.v1beta1.ElectrificationPlannerService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_prospect_shipper_v1beta1_ElectrificationPlannerServiceClient = v1beta1.NewElectrificationPlannerServiceClient(conn)
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ElectrificationPlannerService.GetDatasetNetwork.
var (
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork_Request v1beta1.GetDatasetNetworkRequest
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork         = &cobra.Command{
		Use: "GetDatasetNetwork",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.prospect.shipper.v1beta1.ElectrificationPlannerService.GetDatasetNetwork")
			return nil
		},
	}
)

// einride.prospect.shipper.v1beta1.ElectrificationPlannerService.ComputeNetworkInsights.
var (
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights_Request v1beta1.ComputeNetworkInsightsRequest
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights         = &cobra.Command{
		Use: "ComputeNetworkInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.prospect.shipper.v1beta1.ElectrificationPlannerService.ComputeNetworkInsights")
			return nil
		},
	}
)

func AddElectrificationPlannerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService)
}

func init() {
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork)
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights)
}
