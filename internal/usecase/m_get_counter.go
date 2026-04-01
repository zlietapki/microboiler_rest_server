package usecase

import (
	"context"

	"github.com/zlietapki/gena/internal/domain"
)

func (uc *Usecase) GetCounter(_ context.Context) (domain.Counter, error) {
	return uc.counterRepo.GetNextCounter()
}
