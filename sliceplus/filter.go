package sliceplus

type intPredicate struct {
	f func(int) bool
}

type intPredicateMatcher struct {
	f    func(int) bool
	data []int
}

func (intPredicateMatcher) Match(i int) {

}

type Matcher interface {
	Match(i int) bool
}

func GetFilter() {}
