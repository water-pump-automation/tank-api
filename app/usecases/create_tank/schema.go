package create_tank

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
