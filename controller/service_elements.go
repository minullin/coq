package controller

import "github.com/gin-gonic/gin"

// CreateAgentServiceElement godoc
//
//	@Summary		Create Service Element instance on a USP Agent
//	@Description	Update the USP Agent's configuration by creating a new instance of the specified Service Element
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path	string								true	"USP Agent Endpoint ID"
//	@Param			body		body	model.CreateServiceElementDetail	true	"Service Element object that needs to be created on the USP Agent"
//	@Success		201			"Service Element instance created"
//	@Failure		400			"Service Element instance creation failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/serviceElements [post]
func (c *Controller) CreateAgentServiceElement(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
}

// GetAgentServiceElement godoc
//
//	@Summary		Get Service Element instance by Name from USP Agent
//	@Description	Retrieve an instance of the specified Service Element from a specific USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path		string							true	"USP Agent Endpoint ID"
//	@Param			name		path		string							true	"Service Element Name"
//	@Param			body		body		model.ServiceElementIdentifier	true	"The identifier of the Service Element to retrieve"
//	@Success		200			{object}	model.ServiceElementDetail		"Successful operation"
//	@Failure		400			"Service Element not found"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/serviceElements/{name} [get]
func (c *Controller) GetAgentServiceElement(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
	// name := ctx.Param("name")
}

// UpdateAgentServiceElements godoc
//
//	@Summary		Update Service Element instance by Name on a USP Agent
//	@Description	Update the USP Agent's configuration with the given instance of the specified Service Element
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path	string						true	"USP Agent Endpoint ID"
//	@Param			name		path	string						true	"Service Element Name"
//	@Param			body		body	model.ServiceElementDetail	true	"Service Element objects that need to be configured on the USP Agent"
//	@Success		204			"Service Element updates applied"
//	@Failure		400			"Service Element update failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/serviceElements/{name} [patch]
func (c *Controller) UpdateAgentServiceElements(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
	// name := ctx.Param("name")
}

// DeleteAgentServiceElement godoc
//
//	@Summary		Delete Service Element instance by Name from USP Agent
//	@Description	Update the USP Agent's configuration by removing an existing instance of the specified Service Element
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path	string							true	"USP Agent Endpoint ID"
//	@Param			name		path	string							true	"Service Element Name"
//	@Param			body		body	model.ServiceElementIdentifier	true	"The identifier of the Service Element to remove"
//	@Success		204			"Service Element instance removed"
//	@Failure		400			"Service Element instance removal failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/serviceElements/{name} [delete]
func (c *Controller) DeleteAgentServiceElement(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
	// name := ctx.Param("name")
}
