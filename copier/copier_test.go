package copier_test

import (
	"testing"

	"github.com/goadesign/gorma/copier"
)

type Simple struct {
	Name  string
	Phone string
	Age   int
}

func (s *Simple) Reset() {
	s.Name = ""
	s.Phone = ""
	s.Age = 0
}

type LessSimple struct {
	Name  *string
	Phone *string
	Age   *int
}

type Complex struct {
	Name    string
	Phone   *string
	Age     *int32
	Balance *float64
}

func TestBadParameters(t *testing.T) {

	bad := "Bad"
	badder := "Badder"
	err := copier.Copy(bad, badder)
	if err == nil {
		t.Error("Expected error with bad parameters")
	}

}
func TestSimpleCopy(t *testing.T) {

	left := Simple{
		Name:  "Brian",
		Phone: "813-555-1212",
		Age:   31,
	}
	right := &Simple{}
	err := copier.Copy(left, right)
	if err != nil {
		t.Error(err)
	}

	if right.Age != left.Age {
		t.Errorf("Expected right age to be %d, but got %d", left.Age, right.Age)
	}
	if right.Phone != left.Phone {
		t.Errorf("Expected right phone to be %s, but got %s", left.Phone, right.Phone)
	}
	if right.Name != left.Name {
		t.Errorf("Expected right name to be %s, but got %s", left.Name, right.Name)
	}
}

func TestLessSimpleCopy(t *testing.T) {

	name := "Brian"
	phone := "813-368-3425"
	age := 31
	left := LessSimple{
		Name:  &name,
		Phone: &phone,
		Age:   &age,
	}
	right := &Simple{}
	err := copier.Copy(left, right)
	if err != nil {
		t.Error(err)
	}

	if right.Age != *left.Age {
		t.Errorf("Expected right age to be %d, but got %d", *left.Age, right.Age)
	}
	if right.Phone != *left.Phone {
		t.Errorf("Expected right phone to be %s, but got %s", *left.Phone, right.Phone)
	}
	if right.Name != *left.Name {
		t.Errorf("Expected right name to be %s, but got %s", *left.Name, right.Name)
	}

	nleft := Simple{
		Name:  "Brian",
		Phone: "813-555-1212",
		Age:   31,
	}
	nright := &LessSimple{}
	err = copier.Copy(nleft, nright)
	if err != nil {
		t.Error(err)
	}

	if *nright.Age != nleft.Age {
		t.Errorf("Expected right age to be %d, but got %d", nleft.Age, *nright.Age)
	}
	if *nright.Phone != nleft.Phone {
		t.Errorf("Expected right phone to be %s, but got %s", nleft.Phone, *nright.Phone)
	}
	if *nright.Name != nleft.Name {
		t.Errorf("Expected right name to be %s, but got %s", nleft.Name, *nright.Name)
	}
}

func TestLessLessCopy(t *testing.T) {

	name := "Brian"
	phone := "813-368-3425"
	age := 31
	left := LessSimple{
		Name:  &name,
		Phone: &phone,
		Age:   &age,
	}
	right := &LessSimple{}
	err := copier.Copy(left, right)
	if err != nil {
		t.Error(err)
	}

	if *right.Age != *left.Age {
		t.Errorf("Expected right age to be %d, but got %d", *left.Age, *right.Age)
	}
	if *right.Phone != *left.Phone {
		t.Errorf("Expected right phone to be %s, but got %s", *left.Phone, *right.Phone)
	}
	if *right.Name != *left.Name {
		t.Errorf("Expected right name to be %s, but got %s", *left.Name, *right.Name)
	}
}
func TestComplexCopy(t *testing.T) {

	name := "Brian"
	phone := "813-368-3425"
	var age int32
	age = 30
	balance := 5.50

	left := Complex{
		Name:    name,
		Phone:   &phone,
		Age:     &age,
		Balance: &balance,
	}
	right := &Complex{}
	err := copier.Copy(left, right)
	if err != nil {
		t.Error(err)
	}

	if *right.Age != *left.Age {
		t.Errorf("Expected right age to be %d, but got %d", *left.Age, *right.Age)
	}
	if *right.Phone != *left.Phone {
		t.Errorf("Expected right phone to be %s, but got %s", *left.Phone, *right.Phone)
	}
	if right.Name != left.Name {
		t.Errorf("Expected right name to be %s, but got %s", left.Name, right.Name)
	}
	if *right.Balance != *left.Balance {
		t.Errorf("Expected right balance to be %f, but got %f", *left.Balance, *right.Balance)
	}
}

func BenchmarkSimple(b *testing.B) {
	left := Simple{
		Name:  "Brian",
		Phone: "813-555-1212",
		Age:   31,
	}
	right := &Simple{}
	// run the  function b.N times
	for n := 0; n < b.N; n++ {
		copier.Copy(left, right)
		right.Reset()
	}
}
