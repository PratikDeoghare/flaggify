package flaggify_test

import (
	"flag"
	"github.com/pratikdeoghare/flaggify"
	"os"
	"testing"
)

func TestFlaggifyIt(t *testing.T) {
	type student struct {
		Name  string `json:"name" x:"name of the student"`
		Age   int    `json:"age" x:"age of the student"`
		City  string `json:"city" x:"city of the student"`
		Hobby string `json:"hobby" x:"hobby of the student"`
		Lazy  bool   `json:"lazy" x:"lazy or not"`
	}

	studentDefault := student{
		Name:  "John",
		Age:   20,
		City:  "Rivendell",
		Hobby: "reading",
		Lazy:  true,
	}

	s := new(student)

	flag.StringVar(&s.Hobby, "hobby", "", "HOBBY of the student")
	flaggify.Flaggify(s, studentDefault)

	os.Args = []string{"", "-name", "Jack", "-age", "44"}

	flag.Parse()

	if s.Name != "Jack" {
		t.Errorf("Expected name to be Jack, got %s", s.Name)
	}

	if s.Age != 44 {
		t.Errorf("Expected age to be 44, got %d", s.Age)
	}

	if s.City != "Rivendell" {
		t.Errorf("Expected city to be Rivendell, got %s", s.City)
	}

	if s.Hobby != "" {
		t.Errorf("Expected hobby to be reading, got %s", s.Hobby)
	}

	if s.Lazy != true {
		t.Errorf("Expected lazy to be true, got %t", s.Lazy)
	}
}
