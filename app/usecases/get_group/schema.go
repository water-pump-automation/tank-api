package get_group

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
