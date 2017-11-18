package command

import (
	"regexp"
)

type Command interface {
	Name() string
	Priority() int
	Match(aVal string) (bool, error)
	// Validate(aArgs []string) bool
	Exec(aArgs []string) (interface{}, error)
}

type SimpleCommand struct {
	name       string
	priority   int
	regex      *regexp.Regexp
	params	   int
	operation  Operation
}

func (cmd *SimpleCommand) Name() string {
	return cmd.name
}

func (cmd *SimpleCommand) Priority() int {
	return cmd.priority
}

func (cmd *SimpleCommand) Match(aVal string) (bool, error) {
	return cmd.regex.MatchString(aVal), nil
}


// func (cmd *SimpleCommand) Validate(aArgs []string) bool {
// 	if cmd.params == -1 {return true}
// 	return cmd.params == len(aArgs)
// }

func (cmd *SimpleCommand) Exec(aArgs []string) (interface{}, error) {
	
	return cmd.operation(aArgs)
}

type Operation func(aArgs []string) (interface{}, error)

func New(aName string, aPriority int, aRegex string, aParams int,
	aOperation Operation) (Command, error) {

	// Compile regex
	rx, err := regexp.Compile(aRegex)
	if err != nil {
		return nil, err
	}
	return &SimpleCommand{aName, aPriority, rx, aParams,  aOperation}, nil
}

func (cmd *SimpleCommand) checkParams(aArgs []string) bool {
	if cmd.params == -1 {return true}
	return cmd.params == len(aArgs)
}

// --- Sorting mechanism ---
type CommandSorter []Command

func (s CommandSorter) Len() int {
	return len(s)
}
func (s CommandSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CommandSorter) Less(i, j int) bool {
	return s[i].Priority() < s[j].Priority()
}
