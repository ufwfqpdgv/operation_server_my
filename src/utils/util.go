package utils

import (
	"fmt"

	"utils/log"

	"github.com/json-iterator/go"
	"github.com/viant/toolbox"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Struct2Map(in interface{}) (out map[string]string) {
	m := make(map[string]interface{})
	err := toolbox.NewConverter("", "json").AssignConverted(&m, in)
	if err != nil {
		log.Error(err)
	}
	out = make(map[string]string)
	for k, v := range m {
		out[k] = fmt.Sprintf("%v", v)
	}
	return
}
