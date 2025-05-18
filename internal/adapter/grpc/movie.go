package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
)

type MovieHandler struct {
	uc MovieUseCase
	svc.UnimplementedMovieServiceServer
}

func NewMovieHandler(uc MovieUseCase) *MovieHandler {
	return &MovieHandler{
		uc: uc,
	}
}

func (h *MovieHandler) Create(ctx context.Context, req *svc.CreateRequest) (*svc.CreateResponse, error) {
	movie := dto.ToMovieFromCreateRequest(req)

	createdMovie, err := h.uc.Create(ctx, movie)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.CreateResponse{
		Movie: dto.FromMovieToPb(createdMovie),
	}, nil
}

func (h *MovieHandler) Get(ctx context.Context, req *svc.GetRequest) (*svc.GetResponse, error) {
	id := req.ID

	movie, err := h.uc.GetByID(ctx, id)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.GetResponse{
		Movie: dto.FromMovieToPb(movie),
	}, nil
}

func (h *MovieHandler) GetAll(ctx context.Context, req *svc.GetAllRequest) (*svc.GetAllResponse, error) {
	movies, err := h.uc.GetAll(ctx)
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

func (h *MovieHandler) Update(ctx context.Context, req *svc.UpdateRequest) (*svc.UpdateResponse, error) {
	id, update := dto.ToMovieUpdateFromUpdateRequest(req)

	updatedMovie, err := h.uc.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.UpdateResponse{
		Movie: dto.FromMovieToPb(updatedMovie),
	}, nil
}

func (h *MovieHandler) Delete(ctx context.Context, req *svc.DeleteRequest) (*svc.DeleteResponse, error) {
	id := req.ID

	movie, err := h.uc.DeleteByID(ctx, id)
	if err != nil {
		return nil, dto.FromError(err)
	}

	return &svc.DeleteResponse{
		Movie: dto.FromMovieToPb(movie),
	}, nil
}
