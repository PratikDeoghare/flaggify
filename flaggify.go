package flaggify

import (
	"flag"
	"reflect"
)

func It(s, sDefault interface{}) {
	sVal := reflect.ValueOf(s)
	sTyp := reflect.TypeOf(s)
	sValElemTyp := sVal.Elem().Type()

	sDefaultVal := reflect.ValueOf(sDefault)

	for i := 0; i < sValElemTyp.NumField(); i++ {
		field := sVal.Elem().Field(i)
		flagTag := sTyp.Elem().Field(i).Tag.Get("x") // x for explanation
		jsonTag := sTyp.Elem().Field(i).Tag.Get("json")

		if flag.CommandLine.Lookup(jsonTag) != nil {
			continue
		}

		switch field.Kind() {
		case reflect.Int:
			flag.IntVar(field.Addr().Interface().(*int), jsonTag, int(sDefaultVal.Field(i).Int()), flagTag)
		case reflect.String:
			flag.StringVar(field.Addr().Interface().(*string), jsonTag, sDefaultVal.Field(i).String(), flagTag)
		}
	}
}
