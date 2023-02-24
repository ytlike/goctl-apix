package swgger

import (
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
)

var strColon = []byte(":")

const (
	defaultOption   = "default"
	stringOption    = "string"
	optionalOption  = "optional"
	omitemptyOption = "omitempty"
	optionsOption   = "options"
	rangeOption     = "range"
	exampleOption   = "example"
	optionSeparator = "|"
	equalToken      = "="
	atRespDoc       = "@respdoc-"
)

// Generator defines the environment needs of rpc service generation
type Generator struct {
	log console.Console
}

// NewGenerator returns an instance of Generator
func NewGenerator() *Generator {
	log := console.NewColorConsole(true)
	return &Generator{
		log: log,
	}
}
