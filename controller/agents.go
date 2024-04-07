package controller

import "github.com/gin-gonic/gin"

// CreateAgent godoc
//
//	@Summary		Create a USP Agent with a Given Endpoint ID
//	@Description	Create a new USP Agent with a Given Endpoint ID
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			body	body	model.CreateAgentDetails	true	"Details about the new USP Agent"
//	@Success		201		"USP Agent created"
//	@Failure		400		"USP Agent creation failed"
//	@Router			/agents [post]
func (c *Controller) CreateAgent(ctx *gin.Context) {

}

// GetAgentDetails godoc
//
//	@Summary		Get USP Agent Details
//	@Description	Retrieve details related to a specific USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path		string				true	"USP Agent Endpoint ID"
//	@Success		200			{object}	model.AgentDetails	"Successful operation"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID} [get]
func (c *Controller) GetAgentDetails(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
}

// UpdateAgent godoc
//
//	@Summary		Update a USP Agent
//	@Description	Update the details of a USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path	string				true	"USP Agent Endpoint ID"
//	@Param			body		body	model.AgentDetails	true	"USP Agent details that need to be updated"
//	@Success		204			"USP Agent updates applied"
//	@Failure		400			"USP Agent update failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID} [put]
func (c *Controller) UpdateAgent(ctx *gin.Context) {

}

// DeleteAgent godoc
//
//	@Summary		Delete a USP Agent
//	@Description	Remove an existing USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path	string	true	"USP Agent Endpoint ID"
//	@Success		204			"USP Agent removed"
//	@Failure		400			"USP Agent removal failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID} [delete]
func (c *Controller) DeleteAgent(ctx *gin.Context) {

}

// GetAgentStatus godoc
//
//	@Summary		Get USP Agent Status
//	@Description	Retrieve communications status of a specific USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path		string				true	"USP Agent Endpoint ID"
//	@Success		200			{object}	model.AgentStatus	"Successful operation"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/status [get]
func (c *Controller) GetAgentStatus(ctx *gin.Context) {

}

// SearchAgents godoc
//
//	@Summary		Search for USP Agents associated to this USP Agent
//	@Description	Search for USP Agents associated to this USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path		string						true	"USP Agent Endpoint ID"
//	@Success		200			{object}	model.AgentSearchResults	"Successful operation"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/associatedAgents [get]
func (c *Controller) SearchAgents(ctx *gin.Context) {

}
