package command

import 
(
	"regexp"
)

type Command struct {
	name string
	regex *regexp.Regexp
	priority int
	operation Operation
}

type Operation func(aArgs []string) (interface{}, error)

func New(aName string, aRegex string, aPriority int, 
	aOperation Operation) (*Command, error) {

	// Compile regex
	rx, err := regexp.Compile(aRegex)
	if err !=nil {
		return nil, err;
	}
	return &Command {aName, rx, aPriority, aOperation}, nil
}

func Match(aCommand *Command, aCall string) bool {
	return aCommand.regex.MatchString(aCall)
}

func Exec(aCommand *Command, aArgs []string) (interface{}, error) {
	return aCommand.operation(aArgs)
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
    return s[i].priority < s[j].priority
}