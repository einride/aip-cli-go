// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package freightv1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func NewFreightServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_einride_example_freight_v1_freight_service_proto.
			Services().ByName("FreightService"),
		map[protoreflect.FullName]string{
			"einride.example.freight.v1.FreightService": " This API represents a simple freight service.\n\n It defines the following resource model:\n\n - The API has a collection of Shipper resources.\n\n - Each Shipper has a collection of Site resources.\n\n - Each Shipper has a collection of Shipment resources.\n",
		},
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("GetShipper"),
			&GetShipperRequest{},
			&Shipper{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.GetShipper": " Get a shipper.\n\n See: https://google.aip.dev/131 (Standard methods: Get).\n",
				"einride.example.freight.v1.GetShipperRequest.name":    " The resource name of the shipper to retrieve.\n Format: shippers/{shipper}\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("ListShippers"),
			&ListShippersRequest{},
			&ListShippersResponse{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.ListShippers":    " List shippers.\n\n See: https://google.aip.dev/132 (Standard methods: List).\n",
				"einride.example.freight.v1.ListShippersRequest.page_size":  " Requested page size. Server may return fewer shippers than requested.\n If unspecified, server will pick an appropriate default.\n",
				"einride.example.freight.v1.ListShippersRequest.page_token": " A token identifying a page of results the server should return.\n Typically, this is the value of\n [ListShippersResponse.next_page_token][einride.example.freight.v1.ListShippersResponse.next_page_token]\n returned from the previous call to `ListShippers` method.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("CreateShipper"),
			&CreateShipperRequest{},
			&Shipper{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.CreateShipperRequest.shipper": " The shipper to create.\n",
				"einride.example.freight.v1.FreightService.CreateShipper": " Create a shipper.\n\n See: https://google.aip.dev/133 (Standard methods: Create).\n",
				"einride.example.freight.v1.Shipper.create_time":          " The creation timestamp of the shipper.\n",
				"einride.example.freight.v1.Shipper.delete_time":          " The deletion timestamp of the shipper.\n",
				"einride.example.freight.v1.Shipper.display_name":         " The display name of the shipper.\n",
				"einride.example.freight.v1.Shipper.name":                 " The resource name of the shipper.\n",
				"einride.example.freight.v1.Shipper.update_time":          " The last update timestamp of the shipper.\n\n Updated when create/update/delete operation is performed.\n",
				"google.protobuf.Timestamp.nanos":                         " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                       " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("UpdateShipper"),
			&UpdateShipperRequest{},
			&Shipper{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.UpdateShipper":     " Update a shipper.\n\n See: https://google.aip.dev/134 (Standard methods: Update).\n",
				"einride.example.freight.v1.Shipper.create_time":              " The creation timestamp of the shipper.\n",
				"einride.example.freight.v1.Shipper.delete_time":              " The deletion timestamp of the shipper.\n",
				"einride.example.freight.v1.Shipper.display_name":             " The display name of the shipper.\n",
				"einride.example.freight.v1.Shipper.name":                     " The resource name of the shipper.\n",
				"einride.example.freight.v1.Shipper.update_time":              " The last update timestamp of the shipper.\n\n Updated when create/update/delete operation is performed.\n",
				"einride.example.freight.v1.UpdateShipperRequest.shipper":     " The shipper to update with. The name must match or be empty.\n The shipper's `name` field is used to identify the shipper to be updated.\n Format: shippers/{shipper}\n",
				"einride.example.freight.v1.UpdateShipperRequest.update_mask": " The list of fields to be updated.\n",
				"google.protobuf.FieldMask.paths":                             " The set of field mask paths.\n",
				"google.protobuf.Timestamp.nanos":                             " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                           " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("DeleteShipper"),
			&DeleteShipperRequest{},
			&Shipper{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.DeleteShipperRequest.name":    " The resource name of the shipper to delete.\n Format: shippers/{shipper}\n",
				"einride.example.freight.v1.FreightService.DeleteShipper": " Delete a shipper.\n\n See: https://google.aip.dev/135 (Standard methods: Delete).\n See: https://google.aip.dev/164 (Soft delete).\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("GetSite"),
			&GetSiteRequest{},
			&Site{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.GetSite": " Get a site.\n\n See: https://google.aip.dev/131 (Standard methods: Get).\n",
				"einride.example.freight.v1.GetSiteRequest.name":    " The resource name of the site to retrieve.\n Format: shippers/{shipper}/sites/{site}\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("ListSites"),
			&ListSitesRequest{},
			&ListSitesResponse{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.ListSites":    " List sites for a shipper.\n\n See: https://google.aip.dev/132 (Standard methods: List).\n",
				"einride.example.freight.v1.ListSitesRequest.page_size":  " Requested page size. Server may return fewer sites than requested.\n If unspecified, server will pick an appropriate default.\n",
				"einride.example.freight.v1.ListSitesRequest.page_token": " A token identifying a page of results the server should return.\n Typically, this is the value of\n [ListSitesResponse.next_page_token][einride.example.freight.v1.ListSitesResponse.next_page_token]\n returned from the previous call to `ListSites` method.\n",
				"einride.example.freight.v1.ListSitesRequest.parent":     " The resource name of the parent, which owns this collection of sites.\n Format: shippers/{shipper}\n",
				"einride.example.freight.v1.ListSitesRequest.skip":       " Number of resource to skip in the request.\n * A request with no page token and a skip value of 30 returns a single\n   page of results starting with the 31st result.\n * A request with a page token corresponding to the 51st result (because the\n   first 50 results were returned on the first page) and a skip value of 30\n   returns a single page of results starting with the 81st result.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("CreateSite"),
			&CreateSiteRequest{},
			&Site{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.CreateSiteRequest.parent":  " The resource name of the parent shipper for which this site will be created.\n Format: shippers/{shipper}\n",
				"einride.example.freight.v1.CreateSiteRequest.site":    " The site to create.\n",
				"einride.example.freight.v1.FreightService.CreateSite": " Create a site.\n\n See: https://google.aip.dev/133 (Standard methods: Create).\n",
				"einride.example.freight.v1.Site.create_time":          " The creation timestamp of the site.\n",
				"einride.example.freight.v1.Site.delete_time":          " The deletion timestamp of the site.\n",
				"einride.example.freight.v1.Site.display_name":         " The display name of the site.\n",
				"einride.example.freight.v1.Site.lat_lng":              " The geographic location of the site.\n",
				"einride.example.freight.v1.Site.name":                 " The resource name of the site.\n",
				"einride.example.freight.v1.Site.update_time":          " The last update timestamp of the site.\n\n Updated when create/update/delete operation is performed.\n",
				"google.protobuf.Timestamp.nanos":                      " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                    " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                          " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                         " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("UpdateSite"),
			&UpdateSiteRequest{},
			&Site{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.UpdateSite":     " Update a site.\n\n See: https://google.aip.dev/134 (Standard methods: Update).\n",
				"einride.example.freight.v1.Site.create_time":              " The creation timestamp of the site.\n",
				"einride.example.freight.v1.Site.delete_time":              " The deletion timestamp of the site.\n",
				"einride.example.freight.v1.Site.display_name":             " The display name of the site.\n",
				"einride.example.freight.v1.Site.lat_lng":                  " The geographic location of the site.\n",
				"einride.example.freight.v1.Site.name":                     " The resource name of the site.\n",
				"einride.example.freight.v1.Site.update_time":              " The last update timestamp of the site.\n\n Updated when create/update/delete operation is performed.\n",
				"einride.example.freight.v1.UpdateSiteRequest.site":        " The site to update with. The name must match or be empty.\n The site's `name` field is used to identify the site to be updated.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.UpdateSiteRequest.update_mask": " The list of fields to be updated.\n",
				"google.protobuf.FieldMask.paths":                          " The set of field mask paths.\n",
				"google.protobuf.Timestamp.nanos":                          " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                        " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                              " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                             " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("DeleteSite"),
			&DeleteSiteRequest{},
			&Site{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.DeleteSiteRequest.name":    " The resource name of the site to delete.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.FreightService.DeleteSite": " Delete a site.\n\n See: https://google.aip.dev/135 (Standard methods: Delete).\n See: https://google.aip.dev/164 (Soft delete).\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("BatchGetSites"),
			&BatchGetSitesRequest{},
			&BatchGetSitesResponse{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.BatchGetSitesRequest.names":   " The names of the sites to retrieve.\n A maximum of 1000 sites can be retrieved in a batch.\n Format: `shippers/{shipper}/sites/{site}`\n",
				"einride.example.freight.v1.BatchGetSitesRequest.parent":  " The parent resource shared by all sites being retrieved.\n If this is set, the parent of all of the sites specified in `names`\n must match this field.\n Format: `shippers/{shipper}`\n",
				"einride.example.freight.v1.FreightService.BatchGetSites": " Batch get sites.\n\n See: https://google.aip.dev/231 (Batch methods: Get).\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("GetShipment"),
			&GetShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.GetShipment": " Get a shipment.\n\n See: https://google.aip.dev/131 (Standard methods: Get).\n",
				"einride.example.freight.v1.GetShipmentRequest.name":    " The resource name of the shipment to retrieve.\n Format: shippers/{shipper}/shipments/{shipment}\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("ListShipments"),
			&ListShipmentsRequest{},
			&ListShipmentsResponse{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.ListShipments":    " List shipments for a shipper.\n\n See: https://google.aip.dev/132 (Standard methods: List).\n",
				"einride.example.freight.v1.ListShipmentsRequest.page_size":  " Requested page size. Server may return fewer shipments than requested.\n If unspecified, server will pick an appropriate default.\n",
				"einride.example.freight.v1.ListShipmentsRequest.page_token": " A token identifying a page of results the server should return.\n Typically, this is the value of\n [ListShipmentsResponse.next_page_token][einride.example.freight.v1.ListShipmentsResponse.next_page_token]\n returned from the previous call to `ListShipments` method.\n",
				"einride.example.freight.v1.ListShipmentsRequest.parent":     " The resource name of the parent, which owns this collection of shipments.\n Format: shippers/{shipper}\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("CreateShipment"),
			&CreateShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.CreateShipmentRequest.parent":    " The resource name of the parent shipper for which this shipment will be created.\n Format: shippers/{shipper}\n",
				"einride.example.freight.v1.CreateShipmentRequest.shipment":  " The shipment to create.\n",
				"einride.example.freight.v1.FreightService.CreateShipment":   " Create a shipment.\n\n See: https://google.aip.dev/133 (Standard methods: Create).\n",
				"einride.example.freight.v1.LineItem.quantity":               " The quantity of the line item.\n",
				"einride.example.freight.v1.LineItem.title":                  " The title of the line item.\n",
				"einride.example.freight.v1.LineItem.volume_m3":              " The volume of the line item in cubic meters.\n",
				"einride.example.freight.v1.LineItem.weight_kg":              " The weight of the line item in kilograms.\n",
				"einride.example.freight.v1.Shipment.AnnotationsEntry.key":   "",
				"einride.example.freight.v1.Shipment.AnnotationsEntry.value": "",
				"einride.example.freight.v1.Shipment.annotations":            " Annotations of the shipment.\n",
				"einride.example.freight.v1.Shipment.create_time":            " The creation timestamp of the shipment.\n",
				"einride.example.freight.v1.Shipment.delete_time":            " The deletion timestamp of the shipment.\n",
				"einride.example.freight.v1.Shipment.delivery_earliest_time": " The earliest delivery time of the shipment at the destination site.\n",
				"einride.example.freight.v1.Shipment.delivery_latest_time":   " The latest delivery time of the shipment at the destination site.\n",
				"einride.example.freight.v1.Shipment.destination_site":       " The resource name of the destination site of the shipment.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.Shipment.line_items":             " The line items of the shipment.\n",
				"einride.example.freight.v1.Shipment.name":                   " The resource name of the shipment.\n",
				"einride.example.freight.v1.Shipment.origin_site":            " The resource name of the origin site of the shipment.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.Shipment.pickup_earliest_time":   " The earliest pickup time of the shipment at the origin site.\n",
				"einride.example.freight.v1.Shipment.pickup_latest_time":     " The latest pickup time of the shipment at the origin site.\n",
				"einride.example.freight.v1.Shipment.update_time":            " The last update timestamp of the shipment.\n\n Updated when create/update/delete operation is shipment.\n",
				"google.protobuf.Timestamp.nanos":                            " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                          " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("UpdateShipment"),
			&UpdateShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.FreightService.UpdateShipment":     " Update a shipment.\n\n See: https://google.aip.dev/134 (Standard methods: Update).\n",
				"einride.example.freight.v1.LineItem.quantity":                 " The quantity of the line item.\n",
				"einride.example.freight.v1.LineItem.title":                    " The title of the line item.\n",
				"einride.example.freight.v1.LineItem.volume_m3":                " The volume of the line item in cubic meters.\n",
				"einride.example.freight.v1.LineItem.weight_kg":                " The weight of the line item in kilograms.\n",
				"einride.example.freight.v1.Shipment.AnnotationsEntry.key":     "",
				"einride.example.freight.v1.Shipment.AnnotationsEntry.value":   "",
				"einride.example.freight.v1.Shipment.annotations":              " Annotations of the shipment.\n",
				"einride.example.freight.v1.Shipment.create_time":              " The creation timestamp of the shipment.\n",
				"einride.example.freight.v1.Shipment.delete_time":              " The deletion timestamp of the shipment.\n",
				"einride.example.freight.v1.Shipment.delivery_earliest_time":   " The earliest delivery time of the shipment at the destination site.\n",
				"einride.example.freight.v1.Shipment.delivery_latest_time":     " The latest delivery time of the shipment at the destination site.\n",
				"einride.example.freight.v1.Shipment.destination_site":         " The resource name of the destination site of the shipment.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.Shipment.line_items":               " The line items of the shipment.\n",
				"einride.example.freight.v1.Shipment.name":                     " The resource name of the shipment.\n",
				"einride.example.freight.v1.Shipment.origin_site":              " The resource name of the origin site of the shipment.\n Format: shippers/{shipper}/sites/{site}\n",
				"einride.example.freight.v1.Shipment.pickup_earliest_time":     " The earliest pickup time of the shipment at the origin site.\n",
				"einride.example.freight.v1.Shipment.pickup_latest_time":       " The latest pickup time of the shipment at the origin site.\n",
				"einride.example.freight.v1.Shipment.update_time":              " The last update timestamp of the shipment.\n\n Updated when create/update/delete operation is shipment.\n",
				"einride.example.freight.v1.UpdateShipmentRequest.shipment":    " The shipment to update with. The name must match or be empty.\n The shipment's `name` field is used to identify the shipment to be updated.\n Format: shippers/{shipper}/shipments/{shipment}\n",
				"einride.example.freight.v1.UpdateShipmentRequest.update_mask": " The list of fields to be updated.\n",
				"google.protobuf.FieldMask.paths":                              " The set of field mask paths.\n",
				"google.protobuf.Timestamp.nanos":                              " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                            " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_example_freight_v1_freight_service_proto.
				Services().ByName("FreightService").Methods().ByName("DeleteShipment"),
			&DeleteShipmentRequest{},
			&Shipment{},
			map[protoreflect.FullName]string{
				"einride.example.freight.v1.DeleteShipmentRequest.name":    " The resource name of the shipment to delete.\n Format: shippers/{shipper}/shipments/{shipment}\n",
				"einride.example.freight.v1.FreightService.DeleteShipment": " Delete a shipment.\n\n See: https://google.aip.dev/135 (Standard methods: Delete).\n See: https://google.aip.dev/164 (Soft delete).\n",
			},
		),
	)
}
