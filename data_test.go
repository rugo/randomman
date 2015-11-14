package main
import (
	"testing"
	"fmt"
)

func TestLoadAvailableManpages(t *testing.T) {
	a := LoadAvailableManpages("man_pages")
	if true {
		t.Fail()
	}
	fmt.Printf(a) // Won't run ever, muhaha
}

func TestGetRandomManpageFilename(t *testing.T) {
	a := LoadAvailableManpages("man_pages")
	// Do not need to pass by reference (&), because slice and
	// map are referenced anyways (only ref gets copied in func call)
	fmt.Println(GetRandomManpageFilename([]int{1,2,3}, a))
	t.Fail()
}