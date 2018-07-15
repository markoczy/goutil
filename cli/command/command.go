package command

import (
	"github.com/markoczy/goutil/cli/clierror"
	"regexp"
)

// Command executable entity
type Command interface {
	Name() string
	Priority() int
	Match(aVal string) (bool, error)
	Validate(aArgs []string) error
	Exec(aArgs []string) (interface{}, error)
}

// SimpleCommand std impl of Command
type SimpleCommand struct {
	name      string
	priority  int
	regex     *regexp.Regexp
	params    int
	operation Operation
}

// Name gets name
func (cmd *SimpleCommand) Name() string {
	return cmd.name
}

// Priority gets Priority
func (cmd *SimpleCommand) Priority() int {
	return cmd.priority
}

// Match matches to command regex
func (cmd *SimpleCommand) Match(aVal string) (bool, error) {
	return cmd.regex.MatchString(aVal), nil
}

// Validate checks command count
func (cmd *SimpleCommand) Validate(aArgs []string) error {
	if cmd.params == -1 {
		return nil
	}
	// the first argument matches regex so it isn't counted
	if cmd.params != len(aArgs)-1 {
		return clierror.ErrorArgsCountMismatch
	}
	return nil
}

// Exec processes command
func (cmd *SimpleCommand) Exec(aArgs []string) (interface{}, error) {
	return cmd.operation(aArgs)
}

// Operation function definition
type Operation func(aArgs []string) (interface{}, error)

// New creates new simple command
func New(aName string, aPriority int, aRegex string, aParams int,
	aOperation Operation) (Command, error) {

	// Compile regex
	rx, err := regexp.Compile(aRegex)
	if err != nil {
		return nil, err
	}
	return &SimpleCommand{aName, aPriority, rx, aParams, aOperation}, nil
}

func (cmd *SimpleCommand) checkParams(aArgs []string) bool {
	if cmd.params == -1 {
		return true
	}
	return cmd.params == len(aArgs)
}

// --- Sorting mechanism ---

// CommandSorter ...
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
