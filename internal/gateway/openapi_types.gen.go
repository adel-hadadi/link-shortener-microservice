// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package main

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreateLink defines model for CreateLink.
type CreateLink struct {
	OriginalLink string `json:"original_link"`
}

// Link defines model for Link.
type Link struct {
	OriginalLink string `json:"original_link"`
	ShortLink    string `json:"short_link"`
}

// CreateLinkJSONRequestBody defines body for CreateLink for application/json ContentType.
type CreateLinkJSONRequestBody = CreateLink
