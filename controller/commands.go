package controller

import "github.com/gin-gonic/gin"

// ExecuteAgentServiceElementAction godoc
//
//	@Summary		Execute Service Element Action on USP Agent
//	@Description	Execute a Service Element Action on a USP Agent
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID	path		string						true	"USP Agent Endpoint ID"
//	@Param			body		body		model.ServiceElementAction	true	"Service Element action to execute on the USP Agent"
//	@Success		201			{object}	model.AsyncActionResults	"Asynchronous Action Invoked"
//	@Success		202			{object}	model.ActionResults			"Synchronous Action Invoked"
//	@Failure		400			"Service Element Action invocation failed"
//	@Failure		404			"USP Agent not found"
//	@Router			/agents/{endpointID}/commands [post]
func (c *Controller) ExecuteAgentServiceElementAction(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
}

// GetAgentServiceElementActionResults godoc
//
//	@Summary		Get Asynchronous Service Element Action results
//	@Description	Retrieve the asynchronous results of a Service Element Action
//	@Tags			agents
//	@Accept			json
//	@Produce		json
//	@Param			endpointID		path		string				true	"USP Agent Endpoint ID"
//	@Param			asyncRequestID	path		string				true	"Service Element Asynchronous Action Request ID"
//	@Success		200				{object}	model.ActionResults	"Successful operation"
//	@Failure		400				"Service Element Asynchronous Action Request ID not found"
//	@Failure		404				"USP Agent not found"
//	@Router			/agents/{endpointID}/commands/{asyncRequestID} [get]
func (c *Controller) GetAgentServiceElementActionResults(ctx *gin.Context) {
	// eid := ctx.Param("endpointID")
	// arid := ctx.Param("asyncRequestID")
}
