package envriron

import (
	"fmt"
	"reflect"
)

const (
	NilParser         = "cannot pass nil as arg to UseParser"
	InvalidArgument   = "parser must take a single string as an argument"
	InvalidReturnType = "parser return type must be T or (T, error)"
)

var (
	parsers = map[reflect.Type]reflect.Value{}

	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

func UseParser(parser interface{}) {
	if parser == nil {
		panic(NilParser)
	}

	t := reflect.TypeOf(parser)

	if t.Kind() != reflect.Func {
		panic(fmt.Sprintf(`cannot use "%v" as a parser`, t))
	}

	if t.NumIn() != 1 || t.In(0).Kind() != reflect.String {
		panic(InvalidArgument)
	}

	if t.NumOut() < 1 || t.NumOut() > 2 {
		panic(InvalidReturnType)
	}

	if t.NumOut() == 2 && t.Out(1) != errorType {
		panic(InvalidReturnType)
	}

	if t.Out(0) == errorType {
		panic(InvalidReturnType)
	}

	parsers[t.Out(0)] = reflect.ValueOf(parser)
}
