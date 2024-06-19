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
        "/game/all/count": {
            "get": {
                "description": "Get the number of games in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Get number of games in the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/game.GetGamesCountResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    }
                }
            }
        },
        "/game/create": {
            "post": {
                "description": "Create a new game with the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Create a new game",
                "parameters": [
                    {
                        "description": "Game details",
                        "name": "createGame",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/game.CreateGame"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/game.Game"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    }
                }
            }
        },
        "/game/move": {
            "post": {
                "description": "Post a move to the game with the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Post a move",
                "parameters": [
                    {
                        "description": "Move details",
                        "name": "createMove",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/game.Move"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/game.Game"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    }
                }
            }
        },
        "/game/{gameName}": {
            "get": {
                "description": "Get the game with the given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Get a game",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game name",
                        "name": "gameName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/game.Game"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.FieldError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.ResponseData": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.FieldError"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "game.Cell": {
            "type": "object",
            "properties": {
                "col": {
                    "type": "integer"
                },
                "row": {
                    "type": "integer"
                },
                "value": {
                    "description": "\"X\", \"O\", or null",
                    "type": "string"
                },
                "winningCell": {
                    "type": "boolean"
                }
            }
        },
        "game.CreateGame": {
            "type": "object",
            "properties": {
                "boardSize": {
                    "type": "integer"
                },
                "currentPlayer": {
                    "type": "string"
                },
                "gameName": {
                    "type": "string"
                },
                "winningCondition": {
                    "type": "integer"
                }
            }
        },
        "game.Game": {
            "type": "object",
            "properties": {
                "board": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/game.Cell"
                        }
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "currentPlayer": {
                    "type": "string"
                },
                "gameName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isDraw": {
                    "type": "boolean"
                },
                "oPlayer": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "winner": {
                    "type": "string"
                },
                "winningCondition": {
                    "type": "integer"
                },
                "xPlayer": {
                    "type": "string"
                }
            }
        },
        "game.GetGamesCountResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "game.Move": {
            "type": "object",
            "properties": {
                "col": {
                    "type": "integer"
                },
                "gameName": {
                    "type": "string"
                },
                "player": {
                    "description": "\"X\" or \"O\"",
                    "type": "string"
                },
                "row": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Online TicTacToe API",
	Description:      "This is the REST API for Online TicTacToe.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
