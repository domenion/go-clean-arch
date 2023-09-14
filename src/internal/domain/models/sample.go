package models

import "go-clean-arch/src/internal/domain/entities"

type SampleRequest struct{}

type SampleResponse struct {
	Message string           `json:"message"`
	Sample  *entities.Sample `json:"sample"`
}
