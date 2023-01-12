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
        "contact": {
            "name": "@a48zhang \u0026 @Cg1028",
            "email": "3557695455@qq.com 2194028175@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "login and return id\u0026token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "native login",
                "parameters": [
                    {
                        "description": "the User who is logging in",
                        "name": "loginAuth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "username or password incorrect",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "403": {
                        "description": "param not satisfied",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "token generation failed",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/newproject": {
            "post": {
                "description": "Create user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "newproject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project": {
            "get": {
                "description": "Get a project with its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the project",
                        "name": "info_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user's project(login required)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update one's project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "operating project",
                        "name": "newproject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/project/template": {
            "get": {
                "description": "Get a template with its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a templte",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the template",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Template"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "create a new account in certain way",
                "produces": [
                    "application/json"
                ],
                "summary": "register account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the type of the register",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "description": "register data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get User's info with its userID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get User's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            },
            "put": {
                "description": "upload sth with its UserID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "upload profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
                        "description": "update failed",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/users/myproject": {
            "get": {
                "description": "Get all the projects from user(login required)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dev"
                ],
                "summary": "Get one's projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.ProposalInfo"
                        }
                    }
                }
            }
        },
        "/users/photo": {
            "put": {
                "description": "upload photo with its UserID",
                "consumes": [
                    "image/jpeg"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "upload photo",
                "parameters": [
                    {
                        "type": "object",
                        "description": "the photo of the user",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
                        "description": "file not received",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "failed to upload file",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "db.ProposalInfo": {
            "type": "object",
            "properties": {
                "aim": {
                    "description": "活动目的",
                    "type": "string"
                },
                "budget": {
                    "description": "活动预算",
                    "type": "string"
                },
                "department": {
                    "description": "部门",
                    "type": "string"
                },
                "game": {
                    "description": "游戏项目",
                    "type": "string"
                },
                "info_id": {
                    "description": "活动序号",
                    "type": "integer"
                },
                "name": {
                    "description": "活动",
                    "type": "string"
                },
                "nodes": {
                    "description": "项目环节",
                    "type": "string"
                },
                "optional_tables": {
                    "description": "可选标签",
                    "type": "string"
                },
                "place": {
                    "description": "活动位置",
                    "type": "string"
                },
                "time": {
                    "description": "活动时间",
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "db.Template": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "temid": {
                    "type": "integer"
                }
            }
        },
        "db.User": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "muxipass": {
                    "type": "integer"
                },
                "nick_name": {
                    "description": "最多七个汉字",
                    "type": "string"
                },
                "photo": {
                    "description": "头像",
                    "type": "string"
                },
                "position": {
                    "description": "职位",
                    "type": "string"
                },
                "qq": {
                    "type": "string"
                },
                "studentid": {
                    "type": "integer"
                },
                "tag": {
                    "description": "标签",
                    "type": "string"
                },
                "uid": {
                    "description": "序号",
                    "type": "integer"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "properties": {
                "auth": {
                    "type": "string"
                },
                "nick_name": {
                    "description": "最多七个汉字",
                    "type": "string"
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "under developing",
            "name": "dev"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Melting API",
	Description:      "Backend system of Muxi_Melting",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
