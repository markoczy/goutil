package cli

import 
(
	"github.com/markoczy/goutil/cli/parser"
	"github.com/markoczy/goutil/cli/command"
)

func NewParser() *parser.Parser {
	return parser.New()
}

func AddCommand(aParser *parser.Parser, aName string, aRegex string, 
	aPriority int, aOperation command.Operation) error {

	cmd, err := command.New(aName, aRegex, aPriority, aOperation)
	if err != nil { return err }
	return parser.AddCommand(aParser, *cmd) 
}