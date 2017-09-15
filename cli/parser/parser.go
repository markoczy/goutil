package parser

import 
(
	"sort"
	"github.com/markoczy/goutil/cli/command"
)

type Parser struct {
	commands []command.Command
}

func New() *Parser {
	return &Parser{}
}

func AddCommand(aParser *Parser, aCommand command.Command) error {
	aParser.commands = append(aParser.commands, aCommand)
	sort.Sort(command.CommandSorter(aParser.commands))
	return nil
}

func Exec(aParser *Parser, aArgs []string) (interface{}, error) {
	
	for _, cmd := range aParser.commands {
		if command.Match(&cmd, aArgs[0]) {
			return command.Exec(&cmd, aArgs)
		}
	}

	return nil, nil // TODO error..
}