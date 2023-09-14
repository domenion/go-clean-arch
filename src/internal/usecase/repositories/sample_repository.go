package repositories

import "go-clean-arch/src/internal/domain/entities"

type SampleRepository interface {
	Get(s *entities.Sample) (*entities.Sample, error)
}
