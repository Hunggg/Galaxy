package server

import (
	"context"
	"fmt"

	"github.com/metrogalaxy-io/metrogalaxy-api/grpc/proto/wearables/v1"
	"github.com/metrogalaxy-io/metrogalaxy-api/inject"
	"go.uber.org/zap"
)

type WearablesController struct {
	l *zap.SugaredLogger
}

func NewWearablesController(injector *inject.Injector) *WearablesController {
	return &WearablesController{
		l: zap.S(),
	}
}

func (c *WearablesController) GetWearables(ctx context.Context, request *wearables.GetWearablesRequest) (*wearables.GetWearablesResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetListWearables(ctx context.Context, request *wearables.GetListWearablesRequest) (*wearables.GetListWearablesResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetListWearablesActivity(ctx context.Context, request *wearables.GetListWearablesActivityRequest) (*wearables.GetListWearablesActivityResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetWearablesListing(ctx context.Context, request *wearables.GetWearablesListingRequest) (*wearables.GetWearablesListingResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetWearablesOffer(ctx context.Context, request *wearables.GetWearablesOfferRequest) (*wearables.GetWearablesOfferResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetListWearablesListing(ctx context.Context, request *wearables.GetListWearablesListingRequest) (*wearables.GetListWearablesListingResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetListWearablesOffer(ctx context.Context, request *wearables.GetListWearablesOfferRequest) (*wearables.GetListWearablesOfferResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetListWearablesInformation(ctx context.Context, request *wearables.GetListWearablesInformationRequest) (*wearables.GetListWearablesInformationResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *WearablesController) GetWearablesInformation(ctx context.Context, request *wearables.GetWearablesInformationRequest) (*wearables.GetWearablesInformationResponse, error) {
	return nil, fmt.Errorf("not implemented")
}