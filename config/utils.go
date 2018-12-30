package config

import (
	"github.com/Masterminds/sprig"
	"strings"
	"log"
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
