package tests

import "github.com/go-openapi/strfmt"

type PerfSession struct {
}

func (self *PerfSession) CheckInt(expected int, actual int, message string) {
	if expected != actual {
		panic(message)
	}
}
func (self *PerfSession) CheckStr(expected string, actual string, message string) {
	if expected != actual {
		panic(message)
	}
}
func (self *PerfSession) CheckHash(expected PHash, actual string, message string) {
	if expected != Hash(actual) {
		panic(message)
	}
}
func (self *PerfSession) CheckDate(expected *strfmt.DateTime, actual *strfmt.DateTime, message string) {
	if *expected != *actual {
		panic(message)
	}
}
func (self *PerfSession) Finish(before PVersion, after PVersion) {

}
