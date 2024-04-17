package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jimpick/go-ethereum/accounts/abi/bind"
	"github.com/jimpick/go-ethereum/common"
)

func StringifyArg(arg interface{}) string {
	val := reflect.ValueOf(arg)
	typ := val.Type()

	// If the argument is a slice, iterate over its elements and stringify each one
	if typ.Kind() == reflect.Slice {
		elemCount := val.Len()
		elemStrings := make([]string, elemCount)
		for i := 0; i < elemCount; i++ {
			elem := val.Index(i).Interface()
			elemStrings[i] = StringifyArg(elem)
		}
		return fmt.Sprintf("[%s]", strings.Join(elemStrings, ", "))
	}

	// Handle other types
	switch arg.(type) {
	case *bind.TransactOpts:
		return fmt.Sprintf("TransactOpts{From: %s, Nonce: %s}", val.FieldByName("From").Interface(), val.FieldByName("Nonce").Interface())
	case common.Address:
		return fmt.Sprintf("Address: %s", arg.(common.Address).Hex())
	default:
		return fmt.Sprintf("%v", arg)
	}
}
