package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matthew-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID *string) (*model.Shipment, error) {
	shipment := &model.Shipment{
		ID:           fmt.Sprintf("SHIPMENT-%d", len(r.Shipments)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: &deliveryDate,
		Truck:        &model.Truck{ID: *truckID},
	}
	r.Shipments = append(r.Shipments, shipment)
	return shipment, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, page int, first int) ([]*model.Shipment, error) {
	return r.Shipments, nil
}
