package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
)

type MovieServer struct {
	uc MovieUseCase
	svc.UnimplementedMovieServiceServer
}

func NewMovieServer(uc MovieUseCase) *MovieServer {
	return &MovieServer{
		uc: uc,
	}
}

func (s *MovieServer) Create(ctx context.Context, req *svc.CreateRequest) (*svc.CreateResponse, error) {
	movie := dto.ToMovieFromCreateRequest(req)

	createdMovie, err := s.uc.Create(ctx, movie)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.CreateResponse{
		Movie: dto.FromMovieToPb(createdMovie),
	}, nil
}

func (s *MovieServer) Get(ctx context.Context, req *svc.GetRequest) (*svc.GetResponse, error) {
	id := req.ID

	movie, err := s.uc.GetByID(ctx, id)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.GetResponse{
		Movie: dto.FromMovieToPb(movie),
	}, nil
}

func (s *MovieServer) GetAll(ctx context.Context, req *svc.GetAllRequest) (*svc.GetAllResponse, error) {
	movies, err := s.uc.GetAll(ctx)
	if err != nil {
		return nil, dto.FromError(err)
	}

	var moviesPb []*base.Movie

	for _, movie := range movies {
		moviesPb = append(moviesPb, dto.FromMovieToPb(movie))
	}

	return &svc.GetAllResponse{
		Movies: moviesPb,
	}, nil
}

func (s *MovieServer) Update(ctx context.Context, req *svc.UpdateRequest) (*svc.UpdateResponse, error) {
	id, update := dto.ToMovieUpdateFromUpdateRequest(req)

	updatedMovie, err := s.uc.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.UpdateResponse{
		Movie: dto.FromMovieToPb(updatedMovie),
	}, nil
}

func (s *MovieServer) Delete(ctx context.Context, req *svc.DeleteRequest) (*svc.DeleteResponse, error) {
	id := req.ID

	movie, err := s.uc.DeleteByID(ctx, id)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.DeleteResponse{
		Movie: dto.FromMovieToPb(movie),
	}, nil
}
