package environ

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const structTag = "environ"

// Inject will populate the fields of the given struct with values from the environment.
// the struct fields must be tagged with the "environ" tag.
func Inject(holder interface{}) error {
	hPtr := reflect.TypeOf(holder)

	if hPtr.Kind() != reflect.Ptr {
		return fmt.Errorf(`expected Inject to be called on a pointer type, not "%v"`, hPtr)
	}

	v := reflect.ValueOf(holder).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag, ok := f.Tag.Lookup(structTag)
		if !ok {
			continue
		}

		tokens := strings.Split(tag, ",")

		if len(tokens) < 1 || len(tokens) > 2 {
			continue
		}

		varName := strings.TrimSpace(tokens[0])
		varType := f.Type

		varValue, ok := os.LookupEnv(varName)
		if !ok {
			if len(tokens) < 2 {
				return fmt.Errorf("environment variable %s not found", varName)
			}

			varValue = strings.TrimSpace(tokens[1])
		}

		if varType.Kind() == reflect.String {
			v.Field(i).Set(reflect.ValueOf(varValue))
			continue
		}

		parser, ok := parsers[varType]
		if !ok {
			return fmt.Errorf("no parser for type %s", varType)
		}

		result := parser.Call([]reflect.Value{reflect.ValueOf(varValue)})
		if len(result) == 2 && !result[1].IsNil() {
			err := result[1].Interface().(error)
			return fmt.Errorf("error parsing %s: %w", varName, err)
		}

		v.Field(i).Set(result[0])
	}

	return nil
}
