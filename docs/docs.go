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
            "name": "GPL-3.0-or-later",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/audio/{audioId}/tag": {
            "post": {
                "description": "return status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "set audio Tag",
                "parameters": [
                    {
                        "description": "tag",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.setTagForm"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "audio id",
                        "name": "audioId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/clip/{clipId}/tag": {
            "post": {
                "description": "return status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "set clip Tag",
                "parameters": [
                    {
                        "description": "tag",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.setTagForm"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "clip id",
                        "name": "clipId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file": {
            "post": {
                "description": "return number of added files",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "add file or files",
                "parameters": [
                    {
                        "description": "files",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.filesForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/picture/{pictureId}/tag": {
            "post": {
                "description": "return status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "set picture Tag",
                "parameters": [
                    {
                        "description": "tag",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.setTagForm"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "picture id",
                        "name": "pictureId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tag": {
            "post": {
                "description": "return status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "add Tag",
                "parameters": [
                    {
                        "description": "tag",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.addTagForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/video/{videoId}/tag": {
            "post": {
                "description": "return status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "set video Tag",
                "parameters": [
                    {
                        "description": "tag",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.setTagForm"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "video id",
                        "name": "videoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "router.addTagForm": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "router.filesForm": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "router.setTagForm": {
            "type": "object",
            "properties": {
                "tag": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "alpha",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Slide Goose",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
