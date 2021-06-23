package prospectiveshipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/prospect/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.prospect.shipper.v1beta1.ElectrificationPlannerService.
var (
	einride_prospect_shipper_v1beta1_ElectrificationPlannerServiceClient v1beta1.ElectrificationPlannerServiceClient
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService       = &cobra.Command{
		Use:   "einride.prospect.shipper.v1beta1.ElectrificationPlannerService",
		Short: "Electrification planner service.",
		Long:  "Electrification planner service.",
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
			response, err := einride_prospect_shipper_v1beta1_ElectrificationPlannerServiceClient.GetDatasetNetwork(cmd.Context(), &einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_prospect_shipper_v1beta1_ElectrificationPlannerServiceClient.ComputeNetworkInsights(cmd.Context(), &einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddElectrificationPlannerServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService)
}

func init() {
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork)

	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork.Flags().StringVar(&einride_prospect_shipper_v1beta1_ElectrificationPlannerService_GetDatasetNetwork_Request.Name, "name", "", "Resource name of the network to get.")
	einride_prospect_shipper_v1beta1_ElectrificationPlannerService.AddCommand(einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights)

	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights.Flags().StringVar(&einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights_Request.Dataset, "dataset", "", "Resource name of the network to get insights for.")

	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights.Flags().StringVar(&einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights_Request.Hub, "hub", "", "id of a site that also has to be a hub.\nSpecifying hub will return insights in the context of the specified hub.")

	einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights.Flags().StringSliceVar(&einride_prospect_shipper_v1beta1_ElectrificationPlannerService_ComputeNetworkInsights_Request.Sites, "sites", []string{}, "A list of site ids to include in the calculation of the insights.\nAn empty list includes all sites in the scope (hub or network).")
}
