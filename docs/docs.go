// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Slipneff",
            "url": "https://github.com/slipneff"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/history": {
            "post": {
                "description": "Generates a history of requests for a specified time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "History"
                ],
                "summary": "Generate a history",
                "parameters": [
                    {
                        "description": "Generate history",
                        "name": "History",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.HistoryStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/segment": {
            "post": {
                "description": "Create a new segment with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segments"
                ],
                "summary": "Create a new segment",
                "parameters": [
                    {
                        "description": "Create segment",
                        "name": "Segment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a segment with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Segments"
                ],
                "summary": "Delete a segment",
                "parameters": [
                    {
                        "description": "Delete segment",
                        "name": "Segment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Create a new user with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/user/addSegment": {
            "post": {
                "description": "Provides or deletes all user segments by ID and name of segments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Adds segments to the user",
                "parameters": [
                    {
                        "description": "Add or delete segments to user",
                        "name": "AddSegmentsToUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.AddSegmentsToUserStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "Get all users without data payload",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {}
                    }
                }
            }
        },
        "/user/segment": {
            "post": {
                "description": "Provides all user segments by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Provides all user segments",
                "parameters": [
                    {
                        "description": "Get user segments",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.GetUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Segment": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "percentage": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "segments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Segment"
                    }
                }
            }
        },
        "services.AddSegmentsToUserStruct": {
            "type": "object",
            "properties": {
                "add_name": {
                    "type": "string"
                },
                "delete_name": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "services.GetUser": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "services.HistoryStruct": {
            "type": "object",
            "properties": {
                "month": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito Test Task",
	Description:      "This is the swagger document for the Avito task.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
