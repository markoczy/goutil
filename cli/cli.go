package cli

import 
(
	"github.com/markoczy/goutil/cli/parser"
	"github.com/markoczy/goutil/cli/command"
)

// NewParser ...
func NewParser() parser.Parser {
	return parser.New()
}

// AddCommand ...
func AddCommand(aParser parser.Parser, aName string, aPriority int, 
	aRegex string, aArgCount int, aOperation command.Operation) error {

	cmd, err := command.New(aName, aPriority, aRegex, aArgCount, aOperation)
	if err != nil { return err }
	aParser.AddCommand(cmd)
	return nil
}