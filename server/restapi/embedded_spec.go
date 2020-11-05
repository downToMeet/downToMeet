// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
    "title": "Internal DownToMeet API",
    "version": "1.0.0"
  },
  "paths": {
    "/hello": {
      "get": {
        "description": "If id is \"error\", an error response is returned.\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "summary": "Get a hello world message",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "successful hello world response",
            "schema": {
              "type": "object",
              "properties": {
                "hello": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "A sample parameter.",
          "name": "id",
          "in": "query"
        }
      ]
    },
    "/meetup": {
      "get": {
        "description": "If the required parameters were not specified correctly, an error is returned",
        "summary": "Get the list of meetups",
        "parameters": [
          {
            "type": "number",
            "description": "The latitude of the center of search",
            "name": "lat",
            "in": "query",
            "required": true
          },
          {
            "type": "number",
            "description": "The longitude of the center of search",
            "name": "lon",
            "in": "query",
            "required": true
          },
          {
            "type": "number",
            "description": "Desired search radius (kilometers)",
            "name": "radius",
            "in": "query",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "The longitude of the center of search",
            "name": "tags",
            "in": "query",
            "required": true
          }
		],
		"responses": {
			"200": {
			  "description": "OK",
			  "schema": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/meetupID"
				}
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"post": {
		  "consumes": [
			"application/json"
		  ],
		  "summary": "Post a new meetup",
		  "parameters": [
			{
			  "description": "The meetup to create",
			  "name": "meetup",
			  "in": "body",
			  "schema": {
				"$ref": "#/definitions/meetupRequestBody"
			  }
			}
		  ],
		  "responses": {
			"201": {
			  "description": "Created",
			  "schema": {
				"$ref": "#/definitions/meetup"
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
    "/restricted": {
      "get": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "This is a sample endpoint that is restricted only to users who are\n\"logged in\".\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "produces": [
          "text/plain"
        ],
        "summary": "Restricted endpoint",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Not authenticated",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/meetup/{id}": {
      "get": {
        "description": "If the specified meetup does not exist, an error is returned",
        "summary": "Get the specified meetup's information",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/meetup"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "If the specified meetup does not exist, an error is returned",
        "summary": "Delete the specified meetup",
        "responses": {
          "204": {
            "description": "No Content"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "consumes": [
          "application/json"
        ],
        "summary": "Patch the specified meetup",
        "parameters": [
          {
            "description": "The meetup to update",
            "name": "meetup",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/meetupRequestBody"
			}
		}
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/meetup"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "ID of the desired meetup",
          "name": "id",
          "in": "path",
          "required": true
        }
      ],
    "/set-cookie": {
      "get": {
        "description": "This is a sample endpoint that simulates the action of logging in. After\na successful call to this endpoint, one should then be able to go to\n/restricted and receive a message about who they are logged in as.\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "produces": [
          "text/plain"
        ],
        "summary": "Set cookie session",
        "deprecated": true,
        "parameters": [
          {
            "type": "string",
            "description": "User ID to set",
            "name": "user_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            },
            "headers": {
              "Set-Cookie": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/user/facebook/redirect": {
      "get": {
        "description": "If authentication fails, the user is not logged in.",
        "summary": "Facebook OAuth redirect",
        "parameters": [
          {
            "type": "string",
            "description": "Authorization code from Facebook",
            "name": "code",
            "in": "query"
          }
        ],
        "responses": {
          "303": {
            "description": "Redirect to home page.",
            "headers": {
              "Location": {
                "type": "string",
                "description": "Redirect URL"
              }
            }
          }
        }
      }
    },
    "/user/me": {
      "get": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "If user is not logged in, an error response is returned.",
        "summary": "Get the current user's information",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "description": "If specified user does not exist, an error is returned.",
        "summary": "Get the specified user's information",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "If specified user does not exist or current user is not the specified\nuser, an error is returned.\n",
        "summary": "Patch the specified user",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "ID of the desired user",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "coordinates": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number",
          "maximum": 90,
          "minimum": -90
        },
        "lon": {
          "type": "number",
          "maximum": 180,
          "minimum": -180
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "location": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number",
          "maximum": 90,
          "minimum": -90
        },
        "lon": {
          "type": "number",
          "maximum": 180,
          "minimum": -180
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "meetup": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "attendees": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/userID"
          }
        },
        "canceled": {
          "type": "boolean"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "$ref": "#/definitions/meetupID"
        },
        "location": {
          "$ref": "#/definitions/location"
        },
        "maxCapacity": {
          "type": "number"
        },
        "minCapacity": {
          "type": "number"
        },
        "owner": {
          "$ref": "#/definitions/userID"
        },
        "pendingAttendees": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/userID"
          }
        },
        "rejected": {
          "type": "boolean"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "time": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "meetupID": {
      "type": "string"
    },
    "meetupRequestBody": {
      "type": "object",
      "properties": {
        "capacity": {
          "type": "number"
        },
        "location": {
          "$ref": "#/definitions/location"
        },
        "name": {
          "type": "string"
        },
        "taggedInterests": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "time": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "attending": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        },
        "contactInfo": {
          "type": "string"
        },
        "coordinates": {
          "$ref": "#/definitions/coordinates"
        },
        "id": {
          "$ref": "#/definitions/userID"
        },
        "interests": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "ownedMeetups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        },
        "pendingApproval": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        }
      }
    },
    "userID": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "cookieSession": {
      "description": "Session stored in a cookie.\n\n(If you're reading this in the API documentation, ignore the\n\"query parameter name\" below. It is ignored at runtime.)\n",
      "type": "apiKey",
      "name": "COOKIE",
      "in": "query"
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
    "title": "Internal DownToMeet API",
    "version": "1.0.0"
  },
  "paths": {
    "/hello": {
      "get": {
        "description": "If id is \"error\", an error response is returned.\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "summary": "Get a hello world message",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "successful hello world response",
            "schema": {
              "type": "object",
              "properties": {
                "hello": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "A sample parameter.",
          "name": "id",
          "in": "query"
        }
      ]
    },
    "/meetup": {
      "get": {
        "description": "If the required parameters were not specified correctly, an error is returned",
        "summary": "Get the list of meetups",
        "parameters": [
          {
            "type": "number",
            "description": "The latitude of the center of search",
            "name": "lat",
            "in": "query",
            "required": true
          },
          {
            "type": "number",
            "description": "The longitude of the center of search",
            "name": "lon",
            "in": "query",
            "required": true
          },
          {
            "type": "number",
            "description": "Desired search radius (kilometers)",
            "name": "radius",
            "in": "query",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "The longitude of the center of search",
            "name": "tags",
            "in": "query",
            "required": true
          }
		],
		"responses": {
			"200": {
			  "description": "OK",
			  "schema": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/meetupID"
				}
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"post": {
		  "consumes": [
			"application/json"
		  ],
		  "summary": "Post a new meetup",
		  "parameters": [
			{
			  "description": "The meetup to create",
			  "name": "meetup",
			  "in": "body",
			  "schema": {
				"$ref": "#/definitions/meetupRequestBody"
			  }
			}
		  ],
		  "responses": {
			"201": {
			  "description": "Created",
			  "schema": {
				"$ref": "#/definitions/meetup"
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/meetup/{id}": {
		"get": {
		  "description": "If the specified meetup does not exist, an error is returned",
		  "summary": "Get the specified meetup's information",
		  "responses": {
			"200": {
			  "description": "OK",
			  "schema": {
				"$ref": "#/definitions/meetup"
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"404": {
			  "description": "Not found",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"delete": {
		  "description": "If the specified meetup does not exist, an error is returned",
		  "summary": "Delete the specified meetup",
		  "responses": {
			"204": {
			  "description": "No Content"
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"403": {
			  "description": "Forbidden",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"404": {
			  "description": "Not found",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			},
			"500": {
			  "description": "Internal server error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"patch": {
		  "consumes": [
			"application/json"
		  ],
		  "summary": "Patch the specified meetup",
		  "parameters": [
			{
			  "description": "The meetup to update",
			  "name": "meetup",
			  "in": "body",
			  "schema": {
				"$ref": "#/definitions/meetupRequestBody"
			  }
			}
			],
			"responses": {
			  "200": {
				"description": "OK",
				"schema": {
				  "$ref": "#/definitions/meetup"
				}
			  },
			  "400": {
				"description": "Bad Request",
				"schema": {
				  "$ref": "#/definitions/error"
				}
			  },
			  "403": {
				"description": "Forbidden",
				"schema": {
				  "$ref": "#/definitions/error"
				}
			  },
			  "404": {
				"description": "Not found",
				"schema": {
				  "$ref": "#/definitions/error"
				}
			  },
			  "500": {
				"description": "Internal server error",
				"schema": {
				  "$ref": "#/definitions/error"
				}
			  }
			}
		  },
		  "parameters": [
			{
			  "type": "string",
			  "description": "ID of the desired meetup",
			  "name": "id",
			  "in": "path",
			  "required": true
			}
		  ]
    "/restricted": {
      "get": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "This is a sample endpoint that is restricted only to users who are\n\"logged in\".\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "produces": [
          "text/plain"
        ],
        "summary": "Restricted endpoint",
        "deprecated": true,
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Not authenticated",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/set-cookie": {
      "get": {
        "description": "This is a sample endpoint that simulates the action of logging in. After\na successful call to this endpoint, one should then be able to go to\n/restricted and receive a message about who they are logged in as.\n\nThis is a dummy endpoint for testing purposes. It should be removed soon.\n",
        "produces": [
          "text/plain"
        ],
        "summary": "Set cookie session",
        "deprecated": true,
        "parameters": [
          {
            "type": "string",
            "description": "User ID to set",
            "name": "user_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            },
            "headers": {
              "Set-Cookie": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/user/facebook/redirect": {
      "get": {
        "description": "If authentication fails, the user is not logged in.",
        "summary": "Facebook OAuth redirect",
        "parameters": [
          {
            "type": "string",
            "description": "Authorization code from Facebook",
            "name": "code",
            "in": "query"
          }
        ],
        "responses": {
          "303": {
            "description": "Redirect to home page.",
            "headers": {
              "Location": {
                "type": "string",
                "description": "Redirect URL"
              }
            }
          }
        }
      }
    },
    "/user/me": {
      "get": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "If user is not logged in, an error response is returned.",
        "summary": "Get the current user's information",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "description": "If specified user does not exist, an error is returned.",
        "summary": "Get the specified user's information",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "cookieSession": []
          }
        ],
        "description": "If specified user does not exist or current user is not the specified\nuser, an error is returned.\n",
        "summary": "Patch the specified user",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "ID of the desired user",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "coordinates": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number",
          "maximum": 90,
          "minimum": -90
        },
        "lon": {
          "type": "number",
          "maximum": 180,
          "minimum": -180
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "location": {
      "type": "object",
      "properties": {
        "lat": {
          "type": "number",
          "maximum": 90,
          "minimum": -90
        },
        "lon": {
          "type": "number",
          "maximum": 180,
          "minimum": -180
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "meetup": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "attendees": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/userID"
          }
        },
        "canceled": {
          "type": "boolean"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "$ref": "#/definitions/meetupID"
        },
        "location": {
          "$ref": "#/definitions/location"
        },
        "maxCapacity": {
          "type": "number"
        },
        "minCapacity": {
          "type": "number"
        },
        "owner": {
          "$ref": "#/definitions/userID"
        },
        "pendingAttendees": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/userID"
          }
        },
        "rejected": {
          "type": "boolean"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "time": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "meetupID": {
      "type": "string"
    },
    "meetupRequestBody": {
      "type": "object",
      "properties": {
        "capacity": {
          "type": "number"
        },
        "location": {
          "$ref": "#/definitions/location"
        },
        "name": {
          "type": "string"
        },
        "taggedInterests": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "time": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "attending": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        },
        "contactInfo": {
          "type": "string"
        },
        "coordinates": {
          "$ref": "#/definitions/coordinates"
        },
        "id": {
          "$ref": "#/definitions/userID"
        },
        "interests": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string"
        },
        "ownedMeetups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        },
        "pendingApproval": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/meetupID"
          }
        }
      }
    },
    "userID": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "cookieSession": {
      "description": "Session stored in a cookie.\n\n(If you're reading this in the API documentation, ignore the\n\"query parameter name\" below. It is ignored at runtime.)\n",
      "type": "apiKey",
      "name": "COOKIE",
      "in": "query"
    }
  }
}`))
}
