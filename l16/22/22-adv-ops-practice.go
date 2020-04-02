package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {

	s.PrintBacking = true
	names := make([]string, 5)
	s.Show("1st step", names)

	names = append(names, "einstein", "tesla", "aristo")
	s.Show("2nd step", names)

	// names = names[:0]
	names = append(names[:0], "einstein", "tesla", "aristo")
	s.Show("3rd step", names)

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}

	copy(names[3:5], moreNames[:2])
	s.Show("4th step", names[:10])
	s.Show("4th step #2", names)

	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	s.Show("5th step (before append)", clone)

	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	sliced := clone[1:4]
	sliced = append(sliced, "hypatia")
	clone[2] = "elder"

	s.Show("6th step", clone, sliced)
}
