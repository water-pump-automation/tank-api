package update_tank_state

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
		"level": {
			"type": "number",
			"minimum": 0,
            "exclusiveMinimum": true
		},

	},
	"required": [
		"group",
		"tank_name",
		"level"
	]
}
`
