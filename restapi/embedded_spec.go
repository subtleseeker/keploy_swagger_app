// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

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
    "description": "Example server uploading a file",
    "title": "File Upload",
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/upload": {
      "get": {
        "tags": [
          "uploads"
        ],
        "summary": "uploads",
        "operationId": "uploadFile",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  }
}`))
}
