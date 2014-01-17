package parse

import ()

type ruleset func(s *tReader) (Tree, bool)

func once() {

}

func maybe() {

}

func or(rules ...ruleset) ruleset {

}

// series means "in the order of"
// x y z
func series(rules ...ruleset) ruleset {

}
