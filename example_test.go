package main

import "testing"

func TestPart1(t *testing.T) {
	data := `
	part 
	1 
	data
	`
	expected := !
	actual := part1(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := `
	part 
	2 
	data
	`
	expected := !
	actual := part2(data)
	if expected != actual {
		t.Fatalf(`Error: Wanted %v but got %v`, expected, actual)
	}
}
