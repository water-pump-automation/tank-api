package get_tank

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
