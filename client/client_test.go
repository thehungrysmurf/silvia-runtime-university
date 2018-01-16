package client

import (
	"testing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/heroku/silvia-runtime-university/spec"
	"reflect"
	"fmt"
	"strings"
)

var testCases = []struct{
	name string
	inputPoints []spec.Point
	wantFeatures []spec.Feature
	wantErr	bool
}{
	{
		name:         "With zero points",
		inputPoints:  []spec.Point{},
		wantFeatures: []spec.Feature{},
		wantErr: false,
	},
	{
		name:        "With one point",
		inputPoints: []spec.Point{{ Latitude: 100001, Longitude: 100002},
		},
		wantFeatures: []spec.Feature{
			{ Name: "somefeature", Location: &spec.Point{Latitude: 100001, Longitude: 100002} },
		},
		wantErr: false,
	},
	{
		name:        "With two points",
		inputPoints: []spec.Point{{ Latitude: 100001, Longitude: 100002}, {Latitude: 100003, Longitude: 100004 }},
		wantFeatures: []spec.Feature{
			{ Name: "somefeature", Location: &spec.Point{Latitude: 100001, Longitude: 100002} },
			{ Name: "anotherfeature", Location: &spec.Point{Latitude: 100003, Longitude: 100004} },
		},
		wantErr: false,
	},
	{
		name:        "With point unassociated to feature",
		inputPoints: []spec.Point{{ Latitude: 100008, Longitude: 100009}},
		wantFeatures: []spec.Feature{
			{ Name: "somefeature", Location: &spec.Point{Latitude: 100001, Longitude: 100002} },
		},
		wantErr: true,
	},
}

func TestGetFeatures(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var routeGuide = RouteGuide{ Client: fakeRouteGuideClient{ features: test.wantFeatures }}
			ctx := context.Background()

			gotFeatures, err := routeGuide.GetFeatures(ctx, test.inputPoints)
			if err != nil {
				if test.wantErr && strings.Contains(err.Error(), "No feature associated with this point") {
					return
				} else {
					t.Fatalf("GetFeatures failed miserably: %v", err)
				}
			}
			if !reflect.DeepEqual(gotFeatures, test.wantFeatures) {
				t.Fatalf("Got: %#v\nWant: %#v", gotFeatures, test.wantFeatures)
				}
		})
	}
}

type fakeRouteGuideClient struct {
	features []spec.Feature
}

// Mock must implement the same methods as interface RouteGuideClient
func (f fakeRouteGuideClient) GetFeature(ctx context.Context, in *spec.Point, opts ...grpc.CallOption) (*spec.Feature, error) {
	for _, feature := range f.features {
		if reflect.DeepEqual(feature.Location, in) {
			return &feature, nil
		}
	}
	return nil, fmt.Errorf("No feature associated with this point: %v", in)
}

func (f fakeRouteGuideClient) ListFeatures(ctx context.Context, in *spec.Rectangle, opts ...grpc.CallOption) (spec.RouteGuide_ListFeaturesClient, error) {
	return nil, nil
}

func (f fakeRouteGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (spec.RouteGuide_RecordRouteClient, error) {
	return nil, nil
}

func (f fakeRouteGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (spec.RouteGuide_RouteChatClient, error) {
	return nil, nil
}
