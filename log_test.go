package logger

import (
	"testing"
)


func TestGet(t *testing.T){
	t.Parallel()
	l := Get("testing")

	l.Println("this is a test")
}