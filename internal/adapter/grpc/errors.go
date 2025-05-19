package grpc

import (
	"errors"
	"fmt"
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
)

func (s *MovieServer) logError(op string, err error) {
	if !errors.Is(err, model.ErrNotFound) {
		s.log.Error(fmt.Sprintf("movie %s", op), logger.Err(err))
	}
}
