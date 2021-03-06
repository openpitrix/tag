// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated by "makestatic"; DO NOT EDIT.

package static

var Files = map[string]string{
	"api.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "Tag Project",
    "version": "0.0.1",
    "contact": {
      "name": "Tag Project",
      "url": "https://github.com/openpitrix/tag"
    }
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/tag/attach": {
      "post": {
        "summary": "Attach tags",
        "operationId": "AttachTags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbAttachTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbAttachTagsRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    },
    "/v1/tag/create": {
      "post": {
        "summary": "Create tag",
        "operationId": "CreateTag",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbCreateTagResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateTagRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    },
    "/v1/tag/delete": {
      "post": {
        "summary": "Delete tags",
        "operationId": "DeleteTags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbDeleteTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDeleteTagsRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    },
    "/v1/tag/detach": {
      "post": {
        "summary": "Detach tags",
        "operationId": "DetachTags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbDetachTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDetachTagsRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    },
    "/v1/tag/modify": {
      "post": {
        "summary": "Modify tag",
        "operationId": "ModifyTag",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbModifyTagResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbModifyTagRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    },
    "/v1/tags": {
      "post": {
        "summary": "Describe tag",
        "operationId": "DescribeTags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbDescribeTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDescribeTagsRequest"
            }
          }
        ],
        "tags": [
          "tag"
        ]
      }
    }
  },
  "definitions": {
    "pbAttachTagsRequest": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resource_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbAttachTagsResponse": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resource_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbCreateTagRequest": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        },
        "creator": {
          "type": "string"
        },
        "create_time": {
          "type": "string",
          "format": "date-time"
        },
        "update_time": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "pbCreateTagResponse": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "string"
        }
      }
    },
    "pbDeleteTagsRequest": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbDeleteTagsResponse": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbDescribeTagsRequest": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "key": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "value": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "create_time": {
          "type": "string",
          "format": "date-time"
        },
        "update_time": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "pbDescribeTagsResponse": {
      "type": "object",
      "properties": {
        "total_count": {
          "type": "integer",
          "format": "int64"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbTag"
          }
        }
      }
    },
    "pbDetachTagsRequest": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resource_id": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbDetachTagsResponse": {
      "type": "object",
      "properties": {
        "tag_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resource_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "pbModifyTagRequest": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "pbModifyTagResponse": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "string"
        }
      }
    },
    "pbTag": {
      "type": "object",
      "properties": {
        "tag_id": {
          "type": "string"
        },
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        },
        "creator": {
          "type": "string"
        },
        "create_time": {
          "type": "string",
          "format": "date-time"
        },
        "update_time": {
          "type": "string",
          "format": "date-time"
        }
      }
    }
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "description": "The Authorization header must be set to Bearer followed by a space and a token. For example, 'Bearer vHUabiBEIKi8n1RdvWOjGFulGSM6zunb'.",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "BearerAuth": []
    }
  ]
}
`,
}
