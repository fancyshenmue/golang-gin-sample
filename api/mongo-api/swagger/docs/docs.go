// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Charles Hsu",
            "email": "fancyshenmue@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/deleteSingleDocument": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete Data"
                ],
                "summary": "Delete Single Document",
                "operationId": "1",
                "parameters": [
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoDeleteSingleDocumentBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/getDocument": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "Get Document",
                "operationId": "1",
                "parameters": [
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoAPIGetBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/getDocumentRegex": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "Get Document Regex",
                "operationId": "1",
                "parameters": [
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoAPIGetRegexBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/insertManyDocument": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Insert Data"
                ],
                "summary": "Insert Many Document",
                "operationId": "1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "collection",
                        "name": "collection",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoAPIInsertManyDocumentBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/insertSingleDocument": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Insert Data"
                ],
                "summary": "Insert Single Document",
                "operationId": "1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "collection",
                        "name": "collection",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoAPIInsertSingleDocumentBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/updateSingleDocument": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Data"
                ],
                "summary": "Update Single Document",
                "operationId": "1",
                "parameters": [
                    {
                        "description": "body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.MongoUpdateSingleDocumentBodyTop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.MongoAPIGetBodyTop": {
            "type": "object",
            "required": [
                "body",
                "collection"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "collection": {
                    "type": "string"
                }
            }
        },
        "api.MongoAPIGetRegexBodyTop": {
            "type": "object",
            "required": [
                "collection",
                "data",
                "field"
            ],
            "properties": {
                "collection": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                }
            }
        },
        "api.MongoAPIInsertManyDocumentBodyTop": {
            "type": "object",
            "required": [
                "collection"
            ],
            "properties": {
                "body": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "collection": {
                    "type": "string"
                }
            }
        },
        "api.MongoAPIInsertSingleDocumentBodyTop": {
            "type": "object",
            "required": [
                "body",
                "collection"
            ],
            "properties": {
                "body": {
                    "type": "object",
                    "additionalProperties": true
                },
                "collection": {
                    "type": "string"
                }
            }
        },
        "api.MongoDeleteSingleDocumentBodyTop": {
            "type": "object",
            "required": [
                "collection",
                "filter"
            ],
            "properties": {
                "collection": {
                    "type": "string"
                },
                "filter": {
                    "type": "string"
                }
            }
        },
        "api.MongoUpdateSingleDocumentBodyTop": {
            "type": "object",
            "required": [
                "collection",
                "data",
                "field",
                "filter"
            ],
            "properties": {
                "collection": {
                    "type": "string"
                },
                "data": {
                    "type": "object"
                },
                "field": {
                    "type": "string"
                },
                "filter": {
                    "type": "string"
                },
                "upsert": {
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Mongo RESTful API",
	Description: "Mongo RESTful API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
