package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipmentService.
var (
	einride_shipper_v1beta1_ShipmentServiceClient v1beta1.ShipmentServiceClient
	einride_shipper_v1beta1_ShipmentService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipmentService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentServiceClient = v1beta1.NewShipmentServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.CreateShipment.
var (
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request v1beta1.CreateShipmentRequest
	einride_shipper_v1beta1_ShipmentService_CreateShipment         = &cobra.Command{
		Use: "CreateShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.CreateShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.GetShipment.
var (
	einride_shipper_v1beta1_ShipmentService_GetShipment_Request v1beta1.GetShipmentRequest
	einride_shipper_v1beta1_ShipmentService_GetShipment         = &cobra.Command{
		Use: "GetShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.GetShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.BatchGetShipments.
var (
	einride_shipper_v1beta1_ShipmentService_BatchGetShipments_Request v1beta1.BatchGetShipmentsRequest
	einride_shipper_v1beta1_ShipmentService_BatchGetShipments         = &cobra.Command{
		Use: "BatchGetShipments",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.BatchGetShipments")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.SearchShipments.
var (
	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request v1beta1.SearchShipmentsRequest
	einride_shipper_v1beta1_ShipmentService_SearchShipments         = &cobra.Command{
		Use: "SearchShipments",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.SearchShipments")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UpdateShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request v1beta1.UpdateShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UpdateShipment         = &cobra.Command{
		Use: "UpdateShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.UpdateShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.DeleteShipment.
var (
	einride_shipper_v1beta1_ShipmentService_DeleteShipment_Request v1beta1.DeleteShipmentRequest
	einride_shipper_v1beta1_ShipmentService_DeleteShipment         = &cobra.Command{
		Use: "DeleteShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.DeleteShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UndeleteShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UndeleteShipment_Request v1beta1.UndeleteShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UndeleteShipment         = &cobra.Command{
		Use: "UndeleteShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.UndeleteShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ReferenceShipment.
var (
	einride_shipper_v1beta1_ShipmentService_ReferenceShipment_Request v1beta1.ReferenceShipmentRequest
	einride_shipper_v1beta1_ShipmentService_ReferenceShipment         = &cobra.Command{
		Use: "ReferenceShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.ReferenceShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.PlanShipment.
var (
	einride_shipper_v1beta1_ShipmentService_PlanShipment_Request v1beta1.PlanShipmentRequest
	einride_shipper_v1beta1_ShipmentService_PlanShipment         = &cobra.Command{
		Use: "PlanShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.PlanShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnplanShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnplanShipment_Request v1beta1.UnplanShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnplanShipment         = &cobra.Command{
		Use: "UnplanShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.UnplanShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.LoadShipment.
var (
	einride_shipper_v1beta1_ShipmentService_LoadShipment_Request v1beta1.LoadShipmentRequest
	einride_shipper_v1beta1_ShipmentService_LoadShipment         = &cobra.Command{
		Use: "LoadShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.LoadShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.LoadShipmentUnknownVehicle.
var (
	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle_Request v1beta1.LoadShipmentUnknownVehicleRequest
	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle         = &cobra.Command{
		Use: "LoadShipmentUnknownVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.LoadShipmentUnknownVehicle")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.TransitShipment.
var (
	einride_shipper_v1beta1_ShipmentService_TransitShipment_Request v1beta1.TransitShipmentRequest
	einride_shipper_v1beta1_ShipmentService_TransitShipment         = &cobra.Command{
		Use: "TransitShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.TransitShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ArriveShipment.
var (
	einride_shipper_v1beta1_ShipmentService_ArriveShipment_Request v1beta1.ArriveShipmentRequest
	einride_shipper_v1beta1_ShipmentService_ArriveShipment         = &cobra.Command{
		Use: "ArriveShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.ArriveShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnloadShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnloadShipment_Request v1beta1.UnloadShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnloadShipment         = &cobra.Command{
		Use: "UnloadShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.UnloadShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.DeliverShipment.
var (
	einride_shipper_v1beta1_ShipmentService_DeliverShipment_Request v1beta1.DeliverShipmentRequest
	einride_shipper_v1beta1_ShipmentService_DeliverShipment         = &cobra.Command{
		Use: "DeliverShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.DeliverShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.HandleShipmentExternally.
var (
	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally_Request v1beta1.HandleShipmentExternallyRequest
	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally         = &cobra.Command{
		Use: "HandleShipmentExternally",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.HandleShipmentExternally")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.RejectShipment.
var (
	einride_shipper_v1beta1_ShipmentService_RejectShipment_Request v1beta1.RejectShipmentRequest
	einride_shipper_v1beta1_ShipmentService_RejectShipment         = &cobra.Command{
		Use: "RejectShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.RejectShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnrejectShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnrejectShipment_Request v1beta1.UnrejectShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnrejectShipment         = &cobra.Command{
		Use: "UnrejectShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.UnrejectShipment")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ConsumeShipmentChangedEvent.
var (
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request v1beta1.ConsumeShipmentChangedEventRequest
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent         = &cobra.Command{
		Use: "ConsumeShipmentChangedEvent",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipmentService.ConsumeShipmentChangedEvent")
			return nil
		},
	}
)

func AddShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_CreateShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_GetShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_BatchGetShipments)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_SearchShipments)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UpdateShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeleteShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UndeleteShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ReferenceShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_PlanShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnplanShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_TransitShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ArriveShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnloadShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeliverShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_RejectShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnrejectShipment)
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent)
}
