package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minullin/coq/controller"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/minullin/coq/docs"
)

// @title			User Services Platform - Controller API
// @version		1.0.0
//
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath		/v2
func main() {
	r := gin.Default()
	c := controller.NewController()
	v2 := r.Group("/v2")
	{
		agents := v2.Group("/agents")
		{
			agents.POST("", c.CreateAgent)
			agents.GET(":endpointID", c.GetAgentDetails)
			agents.PUT(":endpointID", c.UpdateAgent)
			agents.DELETE(":endpointID", c.DeleteAgent)

			agents.GET(":endpointID/status", c.GetAgentStatus)
			agents.GET(":endpointID/associatedAgents", c.SearchAgents)

			serviceElements := agents.Group(":endpointID/serviceElements")
			{
				serviceElements.POST("", c.CreateAgentServiceElement)
				serviceElements.GET(":name", c.GetAgentServiceElement)
				serviceElements.PATCH(":name", c.UpdateAgentServiceElements)
				serviceElements.DELETE(":name", c.DeleteAgentServiceElement)
			}

			commands := agents.Group(":endpointID/commands")
			{
				commands.POST("", c.ExecuteAgentServiceElementAction)
				commands.GET(":asyncRequestID", c.GetAgentServiceElementActionResults)
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
