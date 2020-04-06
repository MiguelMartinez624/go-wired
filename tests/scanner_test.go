package gowiredtest

import (
	"testing"

	gowired "github.com/miguelmartinez624/go-wired"
	"github.com/miguelmartinez624/go-wired/models"
)

type ScannerTestFunction = func(f *gowired.Scanner, t *testing.T)

type ScannerTestCase struct {
	Name     string
	CaseFunc ScannerTestFunction
}

func ScannerSuite(f *gowired.Scanner, t *testing.T) {
	ttcc := []ScannerTestCase{
		{Name: "Scan Test", CaseFunc: ScanShouldSuccedTest},
		{Name: "Package Path imformation", CaseFunc: ScanShouldGetPackageTest},
	}

	// Run all test cases
	for _, tc := range ttcc {
		t.Run(tc.Name, func(t *testing.T) { tc.CaseFunc(f, t) })
	}
}

func ScanShouldSuccedTest(s *gowired.Scanner, t *testing.T) {
	var result models.ObjectSchema
	s.Scan(GrandChild{}, &result)

	if &result == nil {
		t.Error("Object schema no created")
	}
}

func ScanShouldGetPackageTest(s *gowired.Scanner, t *testing.T) {
	var result models.ObjectSchema
	s.Scan(GrandChild{}, &result)

	if &result == nil {
		t.Error("Object schema no created.")
	}

	if result.Package == "" {
		t.Error("Pakcage info not set.")
	}

	if result.Package != "github.com/miguelmartinez624/go-wired/tests" {
		t.Errorf("Invalid package name %v", result.Package)

	}
}

func TestScanner(t *testing.T) {
	scanner := &gowired.Scanner{}
	ScannerSuite(scanner, t)
}
