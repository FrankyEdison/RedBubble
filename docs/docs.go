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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Franky",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT License",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category/getAllCategory": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有帖子分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子分类"
                ],
                "summary": "获取所有帖子分类",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseAllCategory"
                        }
                    },
                    "500": {
                        "description": "服务繁忙",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/category/getCategoryById/{cateId}": {
            "get": {
                "description": "根据分类id获取分类详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子分类"
                ],
                "summary": "获取某个分类详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类id",
                        "name": "cateId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCategory"
                        }
                    },
                    "500": {
                        "description": "服务繁忙",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ResponseAllCategory": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.docAllCategory"
                    }
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ResponseCategory": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "$ref": "#/definitions/models.docCategory"
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.docAllCategory": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                }
            }
        },
        "models.docCategory": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
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
	Host:             "127.0.0.1:8081",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "RedBubble项目接口文档",
	Description:      "RedBubble后端接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}