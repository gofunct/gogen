package fmap

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"strings"
)

func List() {
	list := []string{}
	funcs := sprig.GenericFuncMap()
	for k := range funcs {
		list = append(list, k)
	}
	fmt.Printf("(%s)", strings.Join(list, "|"))
}
