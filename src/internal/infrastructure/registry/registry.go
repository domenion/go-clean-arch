package registry

import "go-clean-arch/src/internal/interface/controllers"

type Registry interface {
	NewAppControllers() controllers.Controllers
}

type registry struct {
}

type appControllers struct {
	controllers.SampleController
	// ...controllers
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppControllers() controllers.Controllers {
	return &appControllers{
		r.NewSampleController(),
		// ...controller creators
	}
}
