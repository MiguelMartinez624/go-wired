package factorytest

import (
	"testing"

	"github.com/miguelmartinez624/go-wired/factory"
)

type ScannerTestFunction = func(f *factory.Scanner, t *testing.T)

type ScannerTestCase struct {
	Name     string
	CaseFunc ScannerTestFunction
}

func ScannerSuite(f *factory.Scanner, t *testing.T) {
	ttcc := []ScannerTestCase{
		{Name: "Scan Test", CaseFunc: ScanShouldSuccedTest},
		{Name: "Package Path imformation", CaseFunc: ScanShouldGetPackageTest},
	}

	// Run all test cases
	for _, tc := range ttcc {
		t.Run(tc.Name, func(t *testing.T) { tc.CaseFunc(f, t) })
	}
}

func ScanShouldSuccedTest(s *factory.Scanner, t *testing.T) {
	var result factory.ObjectSchema
	s.Scan(GrandChild{}, &result)

	if &result == nil {
		t.Error("Object schema no created")
	}
}

func ScanShouldGetPackageTest(s *factory.Scanner, t *testing.T) {
	var result factory.ObjectSchema
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
	scanner := &factory.Scanner{}
	ScannerSuite(scanner, t)
}
