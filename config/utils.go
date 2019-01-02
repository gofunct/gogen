package config

import (
	"github.com/Masterminds/sprig"
	"log"
	"strings"
)

//returns sprig funcs
func List() {
	list := []string{}
	funcs := sprig.GenericFuncMap()
	for k := range funcs {
		list = append(list, k)
	}
	log.Printf("(%s)", strings.Join(list, "|"))
}
