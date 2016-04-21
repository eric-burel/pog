#!/bin/bash
# @see https://blog.golang.org/cover
# Generate the coverage report
instrument(){
	go test -covermode=count -coverprofile=validation/"$1".out "./$1"
}
# Print coverage report as a nice html
cover(){
	go tool cover -html=validation/"$1".out
}
instrument event
instrument person
instrument habit
instrument timetogo
# generate report
cover event
cover person
cover habit
cover timetogo
