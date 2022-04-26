package variables

import "testing"

func TestGreeting(t *testing.T) {
	name := "karoon"
	want := "Hello, " + name
	output := Greeting(name)
	if want != output {
		t.Errorf("%q is expected but got %q\n", want, output)
	}
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func TestGreetingWithAge(t *testing.T) {
	name := "Pallat"
	age := uint(30)

	want := "Hello, Pallat. You are 30 years old."
	op := GreetingWithAge(name, age)

	if want != op {
		t.Errorf("%q is expected but got %q\n", want, op)
	}

}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func TestGreetingWithAgeAndDrink(t *testing.T) {
	name := "Pallt"
	age := 30
	drink := "Cola"

	want := "Hello, Pallt. You are 30 years old and your favorite drink is Cola."
	op := GreetingWithAgeAndDrink(name, age, drink)

	if want != op {
		t.Errorf("%q is expected but got %q\n", want, op)
	}

}
