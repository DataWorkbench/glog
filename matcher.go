package glog

// Filter used to control the behavior in glog
type Filter interface {
	Match(level Level) bool
}

// MatchFunc is wrappers for match the specified level
type MatchFunc func(level Level) bool

func (f MatchFunc) Match(level Level) bool {
	return f(level)
}

// MatchGTLevel used to match an level is granter than the level(`lvl`).
func MatchGTLevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level > lvl
	})
}

// MatchGELevel used to match an level is granter than or equal the specified level(`lvl`).
func MatchGELevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level >= lvl
	})
}

// MatchLTLevel used to match an level is less than the level(`lvl`).
func MatchLTLevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level < lvl
	})
}

// MatchLELevel used to match an level is less than or equal the level(`lvl`).
func MatchLELevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level <= lvl
	})
}

// MatchEQLevel used to match an level is equal the level(`lvl`).
func MatchEQLevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level == lvl
	})
}

// MatchNELevel used to match an level is not equal the level(`lvl`).
func MatchNELevel(lvl Level) Filter {
	return MatchFunc(func(level Level) bool {
		return level != lvl
	})
}
