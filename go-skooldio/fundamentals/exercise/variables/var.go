package variables

import "fmt"

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}

func GreetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %v. You are %v years old.", name, age)
}

func GreetingWithAgeAndDrink(name string, age int, drink string) string {
	return fmt.Sprintf("Hello, %v. You are %v years old and your favorite drink is %v.", name, age, drink)
}
