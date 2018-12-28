package gocookie

import (
	"github.com/Masterminds/sprig"
	"strings"
	"log"
)

func List() {
	list := []string{}
	funcs := sprig.GenericFuncMap()
	for k := range funcs {
		list = append(list, k)
	}
	log.Printf("(%s)", strings.Join(list, "|"))
}
