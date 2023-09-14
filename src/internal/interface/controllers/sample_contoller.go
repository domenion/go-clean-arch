package controllers

import (
	"go-clean-arch/src/internal/domain/models"
	"go-clean-arch/src/internal/usecase/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleController interface {
	GetSample(ctx *gin.Context)
}

type sampleController struct {
	sample logic.Sample
}

func NewSampleController(s logic.Sample) SampleController {
	return &sampleController{s}
}

func (sc *sampleController) GetSample(ctx *gin.Context) {
	status := http.StatusOK
	resp := &models.SampleResponse{}
	defer func() {
		ctx.JSON(status, resp)
	}()
	sample, err := sc.sample.GetSample()
	if err != nil {
		status = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}

	resp.Sample = sample
	resp.Message = "success"
}
