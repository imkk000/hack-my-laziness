package conv

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/urfave/cli/v3"
)

func BuildCommands(baseCmds []*cli.Command) []*cli.Command {
	return append(baseCmds, commands...)
}

var commands = []*cli.Command{
	{
		Name: "conv",
		Commands: []*cli.Command{
			{
				Name:  "json2struct",
				Usage: "JSON to Go Struct",
				Action: func(_ context.Context, c *cli.Command) error {
					input := strings.Join(c.Args().Slice(), "")
					var data map[string]any
					if err := json.Unmarshal([]byte(input), &data); err != nil {
						panic(err)
					}

					fmt.Println(generateStruct("Root", data))

					return nil
				},
			},
		},
	},
}

func generateStruct(name string, data map[string]any) string {
	var sb strings.Builder
	var nested strings.Builder

	fmt.Fprintf(&sb, "type %s struct {\n", name)

	for k, v := range data {
		fieldName := toPascalCase(k)
		goType, nestedStruct := inferType(fieldName, v)
		fmt.Fprintf(&sb, "\t%s %s `json:\"%s\"`\n", fieldName, goType, k)
		if nestedStruct != "" {
			nested.WriteString("\n" + nestedStruct)
		}
	}

	sb.WriteString("}")
	return sb.String() + nested.String()
}

func inferType(fieldName string, v any) (string, string) {
	switch val := v.(type) {
	case map[string]any:
		return fieldName, generateStruct(fieldName, val)
	case []any:
		if len(val) > 0 {
			elemType, nested := inferType(fieldName+"Item", val[0])
			return "[]" + elemType, nested
		}
		return "[]any", ""
	case float64:
		if val == float64(int(val)) {
			return "int", ""
		}
		return "float64", ""
	case string:
		return "string", ""
	case bool:
		return "bool", ""
	default:
		return "any", ""
	}
}

func toPascalCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	for i, p := range parts {
		if len(p) > 0 {
			r := []rune(p)
			r[0] = unicode.ToUpper(r[0])
			parts[i] = string(r)
		}
	}
	return strings.Join(parts, "")
}
