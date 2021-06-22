package shipperinsightsv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_insights_v1beta1_TransformationAssessmentService = &cobra.Command{
	Use: "einride.shipper.insights.v1beta1.TransformationAssessmentService",
}

var einride_shipper_insights_v1beta1_GetTransformationAssessmentRequest v1beta1.GetTransformationAssessmentRequest
var einride_shipper_insights_v1beta1_TransformationAssessmentService_GetTransformationAssessment = &cobra.Command{
	Use: "GetTransformationAssessment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.TransformationAssessmentService.GetTransformationAssessment")
		return nil
	},
}

var einride_shipper_insights_v1beta1_CreateTransformationAssessmentRequest v1beta1.CreateTransformationAssessmentRequest
var einride_shipper_insights_v1beta1_TransformationAssessmentService_CreateTransformationAssessment = &cobra.Command{
	Use: "CreateTransformationAssessment",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.TransformationAssessmentService.CreateTransformationAssessment")
		return nil
	},
}

var einride_shipper_insights_v1beta1_ListTransformationAssessmentsRequest v1beta1.ListTransformationAssessmentsRequest
var einride_shipper_insights_v1beta1_TransformationAssessmentService_ListTransformationAssessments = &cobra.Command{
	Use: "ListTransformationAssessments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.TransformationAssessmentService.ListTransformationAssessments")
		return nil
	},
}

var einride_shipper_insights_v1beta1_BatchGetTransformationAssessmentsRequest v1beta1.BatchGetTransformationAssessmentsRequest
var einride_shipper_insights_v1beta1_TransformationAssessmentService_BatchGetTransformationAssessments = &cobra.Command{
	Use: "BatchGetTransformationAssessments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.insights.v1beta1.TransformationAssessmentService.BatchGetTransformationAssessments")
		return nil
	},
}

func AddTransformationAssessmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService)
}

func init() {
	einride_shipper_insights_v1beta1_TransformationAssessmentService.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService_GetTransformationAssessment)
	einride_shipper_insights_v1beta1_TransformationAssessmentService.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService_CreateTransformationAssessment)
	einride_shipper_insights_v1beta1_TransformationAssessmentService.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService_ListTransformationAssessments)
	einride_shipper_insights_v1beta1_TransformationAssessmentService.AddCommand(einride_shipper_insights_v1beta1_TransformationAssessmentService_BatchGetTransformationAssessments)
}
