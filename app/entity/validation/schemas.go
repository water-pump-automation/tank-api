package validation

const GetTankSchemaLoader = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"additionalProperties": true,
	"properties": {
		"group": {
			"type": "string",
			"format": "name_match"
		},
		"tank_name": {
			"type": "string",
			"format": "name_match"
		}
	},
	"required": [
		"group",
		"tank_name"
	]
}
`

const GetGroupSchemaLoader = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"additionalProperties": true,
	"properties": {
		"group": {
			"type": "string",
			"format": "name_match"
		}
	},
	"required": [
		"group"
	]
}
`

const CreateTankSchemaLoader = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"additionalProperties": true,
	"properties": {
		"group": {
			"type": "string",
			"format": "name_match"
		},
		"tank_name": {
			"type": "string",
			"format": "name_match"
		},
		"maximum_capacity": {
			"type": "number",
			"minimum": 1,
            "exclusiveMinimum": false
		}
	},
	"required": [
		"group",
		"tank_name",
		"maximum_capacity"
	]
}
`

const UpdateTankSchemaLoader = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"additionalProperties": true,
	"properties": {
		"group": {
			"type": "string",
			"format": "name_match"
		},
		"tank_name": {
			"type": "string",
			"format": "name_match"
		},
		"water_level": {
			"type": "number",
			"minimum": 0,
            "exclusiveMinimum": true
		},

	},
	"required": [
		"group",
		"tank_name",
		"water_level"
	]
}
`
