// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/students": {
            "get": {
                "description": "Get Students",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Get Students",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "FirstName",
                        "name": "firstname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "LastName",
                        "name": "lastname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UserName",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetStudentResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Create a student",
                "parameters": [
                    {
                        "description": "Student",
                        "name": "blog",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateStudentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/students/{id}": {
            "get": {
                "description": "Get student by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Get Student by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a blog",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Update a blog",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Student",
                        "name": "blog",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateStudentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "Delete a Student",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseOK"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateStudentRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.GetStudentResult": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "student": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Student"
                    }
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseOK": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Student": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger for Student project api",
	Description:      "This is a student project api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}