package immo

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
)

var showSchemaCmd = &cobra.Command{
	Use:   "show-schema",
	Short: "Show the object schema for a real-estate offer in Vertesia",
	Run:   runShowSchema,
}

func runShowSchema(cmd *cobra.Command, args []string) {
	schema := jsonschema.Reflect(&Good{})
	schemaJSON, _ := json.MarshalIndent(schema, "", "  ")
	fmt.Println(string(schemaJSON))
}
