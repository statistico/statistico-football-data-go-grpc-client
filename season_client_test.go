package statisticofootballdata_test

import (
	"context"
	"errors"
	"github.com/statistico/statistico-football-data-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"testing"
)

func TestSeasonClient_ByTeamID(t *testing.T) {
	t.Run("calls season client and returns a slice of season struct", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		stream := new(MockSeasonStream)

		request := statistico.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrapperspb.StringValue{Value: "name_desc"},
		}

		response := statistico.TeamSeasonsResponse{Seasons: []*statistico.Season{
			newProtoSeason(),
			newProtoSeason(),
		}}

		ctx := context.Background()

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(&response, nil)

		seasons, err := client.ByTeamID(ctx, 55, "name_desc")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(seasons))
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		request := statistico.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(&statistico.TeamSeasonsResponse{}, e)

		_, err := client.ByTeamID(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error returned from the data service: rpc error: code = Internal desc = internal error", err.Error())
		s.AssertExpectations(t)
	})

	t.Run("logs error and returns bad gateway error for non internal server error returned by client", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		request := statistico.TeamSeasonsRequest{
			TeamId: 55,
			Sort:   &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Unavailable, "service unavailable")

		s.On("GetSeasonsForTeam", ctx, &request, []grpc.CallOption(nil)).Return(&statistico.TeamSeasonsResponse{}, e)

		_, err := client.ByTeamID(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error connecting to the data service: rpc error: code = Unavailable desc = service unavailable", err.Error())
		s.AssertExpectations(t)
	})
}

func TestSeasonClient_ByCompetitionID(t *testing.T) {
	t.Run("calls season client and returns a slice of season struct", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		stream := new(MockSeasonStream)

		request := statistico.SeasonCompetitionRequest{
			CompetitionId: 55,
			Sort:          &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&statistico.Season{}, io.EOF)

		seasons, err := client.ByCompetitionID(ctx, 55, "name_desc")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 2, len(seasons))
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})

	t.Run("logs error and returns internal server error if internal server error returned by client", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		stream := new(MockSeasonStream)

		request := statistico.SeasonCompetitionRequest{
			CompetitionId: 55,
			Sort:          &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Internal, "internal error")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByCompetitionID(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error returned from the data service: rpc error: code = Internal desc = internal error", err.Error())
		s.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns bad gateway error for non internal server error returned by client", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		stream := new(MockSeasonStream)

		request := statistico.SeasonCompetitionRequest{
			CompetitionId: 55,
			Sort:          &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := status.Error(codes.Unavailable, "service unavailable")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, e)

		_, err := client.ByCompetitionID(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error connecting to the data service: rpc error: code = Unavailable desc = service unavailable", err.Error())
		s.AssertExpectations(t)
		stream.AssertNotCalled(t, "Recv")
	})

	t.Run("logs error and returns internal server error if error reading from stream", func(t *testing.T) {
		t.Helper()

		s := new(MockProtoSeasonClient)
		client := statisticofootballdata.NewSeasonClient(s)

		stream := new(MockSeasonStream)

		request := statistico.SeasonCompetitionRequest{
			CompetitionId: 55,
			Sort:          &wrapperspb.StringValue{Value: "name_desc"},
		}

		ctx := context.Background()

		e := errors.New("oh damn")

		s.On("GetSeasonsForCompetition", ctx, &request, []grpc.CallOption(nil)).Return(stream, nil)
		stream.On("Recv").Twice().Return(newProtoSeason(), nil)
		stream.On("Recv").Once().Return(&statistico.Season{}, e)

		_, err := client.ByCompetitionID(ctx, 55, "name_desc")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "internal server error returned from the data service: oh damn", err.Error())
		s.AssertExpectations(t)
		stream.AssertExpectations(t)
	})
}

func newProtoSeason() *statistico.Season {
	return &statistico.Season{}
}

type MockProtoSeasonClient struct {
	mock.Mock
}

func (s *MockProtoSeasonClient) GetSeasonsForCompetition(ctx context.Context, in *statistico.SeasonCompetitionRequest, opts ...grpc.CallOption) (statistico.SeasonService_GetSeasonsForCompetitionClient, error) {
	args := s.Called(ctx, in, opts)
	return args.Get(0).(statistico.SeasonService_GetSeasonsForCompetitionClient), args.Error(1)
}

func (s *MockProtoSeasonClient) GetSeasonsForTeam(ctx context.Context, in *statistico.TeamSeasonsRequest, opts ...grpc.CallOption) (*statistico.TeamSeasonsResponse, error) {
	args := s.Called(ctx, in, opts)
	return args.Get(0).(*statistico.TeamSeasonsResponse), args.Error(1)
}

type MockSeasonStream struct {
	mock.Mock
	grpc.ClientStream
}

func (s *MockSeasonStream) Recv() (*statistico.Season, error) {
	args := s.Called()
	return args.Get(0).(*statistico.Season), args.Error(1)
}
