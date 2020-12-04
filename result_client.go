package statisticodata

import (
	"context"
	"github.com/statistico/statistico-proto/data/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type ResultClient interface {
	ByID(ctx context.Context, fixtureID uint64) (*statisticodata.Result, error)
	ByTeam(ctx context.Context, req *statisticodata.TeamResultRequest) ([]*statisticodata.Result, error)
}

type resultClient struct {
	client statisticodata.ResultServiceClient
}

func (r resultClient) ByID(ctx context.Context, fixtureID uint64) (*statisticodata.Result, error) {
	request := statisticodata.ResultRequest{FixtureId: fixtureID}

	result, err := r.client.GetById(ctx, &request)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, ErrorNotFound
			case codes.Internal:
				return nil, ErrorInternalServerError
			default:
				return nil, ErrorBadGateway
			}
		}
	}

	return result, nil
}

func (r resultClient) ByTeam(ctx context.Context, req *statisticodata.TeamResultRequest) ([]*statisticodata.Result, error) {
	results := []*statisticodata.Result{}

	stream, err := r.client.GetResultsForTeam(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return results, err
			case codes.Internal:
				return results, ErrorInternalServerError
			default:
				return results, ErrorBadGateway
			}
		}
	}

	for {
		result, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return results, ErrorInternalServerError
		}

		results = append(results, result)
	}

	return results, nil
}

func NewResultClient(p statisticodata.ResultServiceClient) ResultClient {
	return resultClient{client: p}
}
