package registry

import (
	"go-clean-arch/src/internal/interface/controllers"
	"go-clean-arch/src/internal/interface/repositories"
	"go-clean-arch/src/internal/usecase/logic"
)

func (r *registry) NewSampleController() controllers.SampleController {
	sr := repositories.NewSampleRepository()
	s := logic.NewSample(sr)
	return controllers.NewSampleController(s)
}
