package client

import (
	"github.com/heroku/silvia-runtime-university/spec"
	"golang.org/x/net/context"
)

type RouteGuide struct {
	Client spec.RouteGuideClient
}

func (rg *RouteGuide) GetFeatures(ctx context.Context, points []spec.Point) ([]spec.Feature, error) {
	var features = []spec.Feature{}
	for _, point := range points {
		feature, err := rg.Client.GetFeature(ctx, &point)
		if err != nil {
			return nil, err
		}
		features = append(features, *feature)
	}

	return features, nil
}
