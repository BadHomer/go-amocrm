package amocrm

import (
	"github.com/BadHomer/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Get) Pipelines(id string, params *Params) (pipeline []models.Pipeline, err error) {
	c.api.log("GetLead request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: pipelineURL,
		In:      nil,
		Out:     &models.Pipeline{},
		Params:  params,
	}

	if id != "" {
		// ID exists
		options.BaseURL += "/" + id
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		pipeline = []models.Pipeline{*options.Out.(*models.Pipeline)}
		c.api.log("returning the struct...")
		return
	} else {
		// All pipelines
		options.Out = &models.RequestResponse{}
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		pipeline = options.Out.(*models.RequestResponse).Embedded.Pipelines
		c.api.log("returning the struct...")
		return
	}
}
