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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Log in an existing user with the provided email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login request object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorValidationResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Log out the currently logged-in user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "Successfully logged out",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user with the specified name, email, and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration request object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully registered a new user",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorValidationResponse"
                        }
                    }
                }
            }
        },
        "/note": {
            "get": {
                "description": "Retrieve all available notes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "note"
                ],
                "summary": "Get all notes",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved all notes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.NoteResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new note with the specified title, content, and visibility",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "note"
                ],
                "summary": "Create a new note",
                "parameters": [
                    {
                        "description": "Note request object",
                        "name": "note",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.NoteRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created a new note",
                        "schema": {
                            "$ref": "#/definitions/domain.NoteResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorValidationResponse"
                        }
                    }
                }
            }
        },
        "/note/{id}": {
            "get": {
                "description": "Retrieve a note based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "note"
                ],
                "summary": "Get a note by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved a note by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.NoteResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing note based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "note"
                ],
                "summary": "Update a note by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated note object",
                        "name": "note",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.NoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated a note by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.NoteResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing note based on the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "note"
                ],
                "summary": "Delete a note by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Note ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted a note by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Retrieve data of all registered users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved all user data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.UserResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorValidationResponse"
                        }
                    }
                }
            }
        },
        "/user/current": {
            "get": {
                "description": "Retrieve data of the currently logged-in user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get current user data",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved current user data",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Retrieve data of a user based on the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved user by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data of an existing user based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user data by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user data object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated user by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorValidationResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing user based on the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted user by ID",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ErrorValidationResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ValidationError"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.NoteRequest": {
            "type": "object",
            "required": [
                "content",
                "title",
                "userID"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "visibility": {
                    "type": "string",
                    "enum": [
                        "private",
                        "public"
                    ]
                }
            }
        },
        "domain.NoteResponse": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "visibility": {
                    "type": "string"
                }
            }
        },
        "domain.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.UserLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UserRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                }
            }
        },
        "domain.UserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                }
            }
        },
        "domain.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.ValidationError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "gocrud",
	Description:      "golang crud api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
