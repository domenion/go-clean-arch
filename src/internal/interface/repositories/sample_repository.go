package repositories

import (
	"go-clean-arch/src/internal/domain/entities"
	"go-clean-arch/src/internal/usecase/repositories"
)

type sampleRepository struct {
}

func NewSampleRepository() repositories.SampleRepository {
	return &sampleRepository{}
}

func (r *sampleRepository) Get(s *entities.Sample) (*entities.Sample, error) {
	return &entities.Sample{Name: "test", Value: 0, Description: "test"}, nil
}
