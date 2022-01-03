package v1

import "testing"

func TestSubjectNotify(t *testing.T) {
	s := &Subject{}
	s.Register(&Observer1{})
	s.Register(&Observer2{})
	s.Notify("Hi")
}