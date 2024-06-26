// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/agents": {
            "post": {
                "description": "Create a new USP Agent with a Given Endpoint ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Create a USP Agent with a Given Endpoint ID",
                "parameters": [
                    {
                        "description": "Details about the new USP Agent",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAgentDetails"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "USP Agent created"
                    },
                    "400": {
                        "description": "USP Agent creation failed"
                    }
                }
            }
        },
        "/agents/{endpointID}": {
            "get": {
                "description": "Retrieve details related to a specific USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get USP Agent Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.AgentDetails"
                        }
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            },
            "put": {
                "description": "Update the details of a USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Update a USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "USP Agent details that need to be updated",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AgentDetails"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "USP Agent updates applied"
                    },
                    "400": {
                        "description": "USP Agent update failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            },
            "delete": {
                "description": "Remove an existing USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Delete a USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "USP Agent removed"
                    },
                    "400": {
                        "description": "USP Agent removal failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/associatedAgents": {
            "get": {
                "description": "Search for USP Agents associated to this USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Search for USP Agents associated to this USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.AgentSearchResults"
                        }
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/commands": {
            "post": {
                "description": "Execute a Service Element Action on a USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Execute Service Element Action on USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service Element action to execute on the USP Agent",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ServiceElementAction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Asynchronous Action Invoked",
                        "schema": {
                            "$ref": "#/definitions/model.AsyncActionResults"
                        }
                    },
                    "202": {
                        "description": "Synchronous Action Invoked",
                        "schema": {
                            "$ref": "#/definitions/model.ActionResults"
                        }
                    },
                    "400": {
                        "description": "Service Element Action invocation failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/commands/{asyncRequestID}": {
            "get": {
                "description": "Retrieve the asynchronous results of a Service Element Action",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get Asynchronous Service Element Action results",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Element Asynchronous Action Request ID",
                        "name": "asyncRequestID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.ActionResults"
                        }
                    },
                    "400": {
                        "description": "Service Element Asynchronous Action Request ID not found"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/serviceElements": {
            "post": {
                "description": "Update the USP Agent's configuration by creating a new instance of the specified Service Element",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Create Service Element instance on a USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service Element object that needs to be created on the USP Agent",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateServiceElementDetail"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Service Element instance created"
                    },
                    "400": {
                        "description": "Service Element instance creation failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/serviceElements/{name}": {
            "get": {
                "description": "Retrieve an instance of the specified Service Element from a specific USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get Service Element instance by Name from USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Element Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The identifier of the Service Element to retrieve",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ServiceElementIdentifier"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.ServiceElementDetail"
                        }
                    },
                    "400": {
                        "description": "Service Element not found"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            },
            "delete": {
                "description": "Update the USP Agent's configuration by removing an existing instance of the specified Service Element",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Delete Service Element instance by Name from USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Element Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The identifier of the Service Element to remove",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ServiceElementIdentifier"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Service Element instance removed"
                    },
                    "400": {
                        "description": "Service Element instance removal failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            },
            "patch": {
                "description": "Update the USP Agent's configuration with the given instance of the specified Service Element",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Update Service Element instance by Name on a USP Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service Element Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service Element objects that need to be configured on the USP Agent",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ServiceElementDetail"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Service Element updates applied"
                    },
                    "400": {
                        "description": "Service Element update failed"
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        },
        "/agents/{endpointID}/status": {
            "get": {
                "description": "Retrieve communications status of a specific USP Agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agents"
                ],
                "summary": "Get USP Agent Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USP Agent Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.AgentStatus"
                        }
                    },
                    "404": {
                        "description": "USP Agent not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ActionResults": {
            "type": "object",
            "properties": {
                "outputArgs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Arg"
                    }
                }
            }
        },
        "model.AgentDetails": {
            "type": "object",
            "properties": {
                "communicationsAddress": {
                    "type": "string"
                },
                "mtpToUse": {
                    "type": "string"
                },
                "supportedServiceElements": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.AgentSearchResults": {
            "type": "object",
            "properties": {
                "outputArgs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.AgentStatus": {
            "type": "object",
            "properties": {
                "communicationsStatus": {
                    "type": "string"
                }
            }
        },
        "model.Arg": {
            "type": "object",
            "properties": {
                "argName": {
                    "type": "string"
                },
                "argValue": {
                    "type": "string"
                }
            }
        },
        "model.AsyncActionResults": {
            "type": "object",
            "properties": {
                "asyncRequestID": {
                    "type": "string"
                }
            }
        },
        "model.CreateAgentDetails": {
            "type": "object",
            "properties": {
                "communicationsAddress": {
                    "type": "string"
                },
                "endpointID": {
                    "type": "string"
                },
                "mtpToUse": {
                    "type": "string"
                },
                "supportedServiceElements": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.CreateServiceElementDetail": {
            "type": "object",
            "properties": {
                "identifier": {
                    "$ref": "#/definitions/model.ServiceElementIdentifier"
                },
                "name": {
                    "type": "string"
                },
                "propertyList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Element"
                    }
                }
            }
        },
        "model.Element": {
            "type": "object",
            "properties": {
                "elementProperty": {
                    "type": "string"
                },
                "elementValue": {
                    "type": "string"
                }
            }
        },
        "model.ServiceElementAction": {
            "type": "object",
            "properties": {
                "inputArgs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Arg"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ServiceElementDetail": {
            "type": "object",
            "properties": {
                "identifier": {
                    "$ref": "#/definitions/model.ServiceElementIdentifier"
                },
                "propertyList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Element"
                    }
                }
            }
        },
        "model.ServiceElementIdentifier": {
            "type": "object",
            "properties": {
                "uniqueKeyList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Element"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "User Services Platform - Controller API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
