package parser

import (
	"github.com/markoczy/goutil/cli/command"
	"sort"
)

type SimpleParser struct {
	commands []command.Command
}

type Parser interface {
	AddCommand(aCommand command.Command)
	Exec(aArgs []string) (interface{}, error)
}

func New() Parser {
	return &SimpleParser{}
}

func (p *SimpleParser) AddCommand(aCommand command.Command) {
	p.commands = append(p.commands, aCommand)
	sort.Sort(command.CommandSorter(p.commands))
}

func (p *SimpleParser) Exec(aArgs []string) (interface{}, error) {
	for _, cmd := range p.commands {
		match, err := cmd.Match(aArgs[0])
		if err != nil {
			return nil, err
		}
		if match {
			err := cmd.Validate(aArgs)
			if err != nil {
				return nil, err
			}
			return cmd.Exec(aArgs)
		}
	}

	return nil, nil // TODO error..
}
