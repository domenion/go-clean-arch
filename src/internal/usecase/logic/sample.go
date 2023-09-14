package logic

import (
	"go-clean-arch/src/internal/domain/entities"
	"go-clean-arch/src/internal/usecase/repositories"
)

type Sample interface {
	GetSample() (*entities.Sample, error)
}

type sample struct {
	sampleRepo repositories.SampleRepository
}

func NewSample(sr repositories.SampleRepository) Sample {
	return &sample{sr}
}

func (s *sample) GetSample() (*entities.Sample, error) {
	sample := &entities.Sample{}
	sample, err := s.sampleRepo.Get(sample)
	if err != nil {
		return nil, err
	}

	return sample, nil
}
