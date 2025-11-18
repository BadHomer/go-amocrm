package amocrm

import (
	"github.com/BadHomer/go-amocrm/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Get) Pipelines(id string, params *Params) (out models.RequestResponse, err error) {
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

		out = models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Pipelines: []models.Pipeline{
					*options.Out.(*models.Pipeline),
				},
			},
		}
		c.api.log("returning the struct...")
		return
	}

	// All leads
	options.Out = &out
	err = c.api.makeRequest(options)
	if err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
