package hash

import (
	"github.com/mitchellh/hashstructure/v2"
	"strconv"
	"strings"
)

func Hash(val interface{}) string {
	hash, err := hashstructure.Hash(val, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(hash))
}

func HashStrings(input ...string) string {
	val := strings.Join(input, ";")
	return Hash(val)
}
