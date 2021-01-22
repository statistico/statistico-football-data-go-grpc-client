package statisticodata

import (
	"context"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type TeamStatClient interface {
	Stats(ctx context.Context, req *statisticoproto.TeamStatRequest) ([]*statisticoproto.TeamStat, error)
}

type teamStatClient struct {
	client statisticoproto.TeamStatsServiceClient
}

func (t *teamStatClient) Stats(ctx context.Context, req *statisticoproto.TeamStatRequest) ([]*statisticoproto.TeamStat, error) {
	stats := []*statisticoproto.TeamStat{}

	stream, err := t.client.GetStatForTeam(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				return nil, ErrorInvalidArgument{err}
			case codes.Internal:
				return nil, ErrorInternalServerError{err}
			default:
				return nil, ErrorBadGateway{err}
			}
		}

		return nil, err
	}

	for {
		st, err := stream.Recv()

		if err == io.EOF {
			return stats, nil
		}

		if err != nil {
			return stats, ErrorInternalServerError{err}
		}

		stats = append(stats, st)
	}
}

func NewTeamStatClient(p statisticoproto.TeamStatsServiceClient) TeamStatClient {
	return &teamStatClient{client: p}
}
