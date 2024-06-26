// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Quang Dang",
            "email": "quangdangfit@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/address": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Get list Address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id_user",
                        "name": "id_user",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListAddressRes"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "create Address",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateAddressReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Address"
                        }
                    }
                }
            }
        },
        "/address/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Get Address by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Address"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Update Address",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAddressReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Address"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Delete Address",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteAddressReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Address"
                        }
                    }
                }
            }
        },
        "/auth//verfiy-code": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Verfiy Code for phone",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyResponse"
                        }
                    }
                }
            }
        },
        "/auth/change-password": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "changes the password",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePasswordReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRes"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get my profile",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRes"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "City of the address\nexample: \"San Francisco\"",
                    "type": "string"
                },
                "id_address": {
                    "description": "ID of the address\nexample: \"12345\"",
                    "type": "string"
                },
                "id_user": {
                    "description": "User ID associated with the address\nexample: \"67890\"",
                    "type": "string"
                },
                "lat": {
                    "description": "Latitude of the address\nexample: \"37.7749\"",
                    "type": "string"
                },
                "long": {
                    "description": "Longitude of the address\nexample: \"-122.4194\"",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the address\nexample: \"Home\"",
                    "type": "string"
                },
                "street": {
                    "description": "Street of the address\nexample: \"Market Street\"",
                    "type": "string"
                }
            }
        },
        "dto.ChangePasswordReq": {
            "type": "object",
            "required": [
                "new_password",
                "password"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAddressReq": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "City of the address\nexample: \"San Francisco\"",
                    "type": "string"
                },
                "id_user": {
                    "description": "User ID associated with the address\nexample: \"67890\"",
                    "type": "string"
                },
                "lat": {
                    "description": "Latitude of the address\nexample: \"37.7749\"",
                    "type": "string"
                },
                "long": {
                    "description": "Longitude of the address\nexample: \"-122.4194\"",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the address\nexample: \"Home\"",
                    "type": "string"
                },
                "street": {
                    "description": "Street of the address\nexample: \"Market Street\"",
                    "type": "string"
                }
            }
        },
        "dto.DeleteAddressReq": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the address\nexample: \"12345\"",
                    "type": "string"
                },
                "id_user": {
                    "description": "User ID associated with the address\nexample: \"67890\"",
                    "type": "string"
                }
            }
        },
        "dto.ListAddressRes": {
            "type": "object",
            "properties": {
                "addresses": {
                    "description": "List of addresses\nexample: [{\"id_address\":\"12345\",\"id_user\":\"67890\",\"name\":\"Home\",\"city\":\"San Francisco\",\"street\":\"Market Street\",\"lat\":\"37.7749\",\"long\":\"-122.4194\"}]",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Address"
                    }
                },
                "pagination": {
                    "description": "Pagination info",
                    "allOf": [
                        {
                            "$ref": "#/definitions/paging.Pagination"
                        }
                    ]
                }
            }
        },
        "dto.LoginReq": {
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
        "dto.LoginRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/dto.User"
                }
            }
        },
        "dto.RefreshTokenReq": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterReq": {
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
        "dto.RegisterRes": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/dto.User"
                }
            }
        },
        "dto.UpdateAddressReq": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "City of the address\nexample: \"San Francisco\"",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the address\nexample: \"12345\"",
                    "type": "string"
                },
                "id_user": {
                    "description": "User ID associated with the address\nexample: \"67890\"",
                    "type": "string"
                },
                "lat": {
                    "description": "Latitude of the address\nexample: \"37.7749\"",
                    "type": "string"
                },
                "long": {
                    "description": "Longitude of the address\nexample: \"-122.4194\"",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the address\nexample: \"Home\"",
                    "type": "string"
                },
                "street": {
                    "description": "Street of the address\nexample: \"Market Street\"",
                    "type": "string"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.VerifyRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "verify_code": {
                    "type": "string"
                }
            }
        },
        "dto.VerifyResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "paging.Pagination": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "skip": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "User API",
	Description:      "API for user management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
