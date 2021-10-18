package model

import "github.com/alecthomas/jsonschema"

type TemplateID string

func ToTemplateID(s string) TemplateID {
	return TemplateID(s)
}

func (id TemplateID) String() string {
	return string(id)
}

func (TemplateID) JSONSchemaType() *jsonschema.Type {
	return &jsonschema.Type{
		Type:        "string",
		Title:       "id of the template",
		Description: "The Unique ID for the template",
		Examples:    []interface{}{"cve-2021-19520"},
		Pattern:     "^(?:[a-zA-Z0-9][a-zA-Z0-9]*-?[a-zA-Z0-9]+?)+$",
	}
}
