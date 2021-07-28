package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	type Skill struct {
		Name  string
		Level int
	}

	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}

	a := Actor{
		Name: "cow body",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll adn roll", Level: 1},
			{Name: "Flash your doy eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	if result, err := MarshalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func MarshalJson(v interface{}) (string, error) {
	var b bytes.Buffer
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {

		return b.String(), nil

	} else {
		return "", err
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupport kind: " + value.Kind().String())
	}
	return nil
}

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")
	for i := 0; i < value.Len(); i++ {
		sliceValue := value.Index(i)
		writeAny(buff, sliceValue)
		if i < value.Len()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}
func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()
	buff.WriteString("{")

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := valueType.Field(i)
		buff.WriteString("\"")
		buff.WriteString(fieldType.Name)
		buff.WriteString("\":")
		writeAny(buff, fieldValue)
		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("}")

	return nil
}
