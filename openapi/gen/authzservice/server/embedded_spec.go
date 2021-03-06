// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "authz service for istio demo\n",
    "title": "authz",
    "contact": {
      "name": "Joey Zhang",
      "email": "me@zhangzhoujian.com"
    },
    "version": "1.0.0"
  },
  "host": "authz",
  "basePath": "/authz/v1",
  "paths": {
    "/authorize": {
      "post": {
        "tags": [
          "authorization"
        ],
        "summary": "Authorize a request",
        "operationId": "Authorize",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthorizationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status 200",
            "schema": {
              "type": "object",
              "properties": {
                "allow": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Status 400",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "401": {
            "description": "Status 401",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "403": {
            "description": "Status 403",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "500": {
            "description": "Status 500",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "summary": "health",
        "responses": {
          "200": {
            "description": "Status 200"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthorizationRequest": {
      "type": "object",
      "properties": {
        "context": {
          "type": "object"
        },
        "method": {
          "type": "string"
        },
        "resource": {
          "type": "object"
        },
        "resource_url": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "AuthorizationResult": {
      "type": "object",
      "properties": {
        "allow": {
          "type": "boolean"
        },
        "permitted_companies": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "permitted_stores": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "company_id": {
                "type": "string"
              },
              "store_ids": {
                "type": "array",
                "items": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "ErrorCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "authz service for istio demo\n",
    "title": "authz",
    "contact": {
      "name": "Joey Zhang",
      "email": "me@zhangzhoujian.com"
    },
    "version": "1.0.0"
  },
  "host": "authz",
  "basePath": "/authz/v1",
  "paths": {
    "/authorize": {
      "post": {
        "tags": [
          "authorization"
        ],
        "summary": "Authorize a request",
        "operationId": "Authorize",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthorizationRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Status 200",
            "schema": {
              "type": "object",
              "properties": {
                "allow": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Status 400",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "401": {
            "description": "Status 401",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "403": {
            "description": "Status 403",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          },
          "500": {
            "description": "Status 500",
            "schema": {
              "$ref": "#/definitions/ErrorCode"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "summary": "health",
        "responses": {
          "200": {
            "description": "Status 200"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthorizationRequest": {
      "type": "object",
      "properties": {
        "context": {
          "type": "object"
        },
        "method": {
          "type": "string"
        },
        "resource": {
          "type": "object"
        },
        "resource_url": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "AuthorizationResult": {
      "type": "object",
      "properties": {
        "allow": {
          "type": "boolean"
        },
        "permitted_companies": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "permitted_stores": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "company_id": {
                "type": "string"
              },
              "store_ids": {
                "type": "array",
                "items": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "ErrorCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
