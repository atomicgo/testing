// Package testing is a module which is only used for testing purposes.
package testing

// HelloWorld returns "Hello, World Updated!".
func HelloWorld() string {
	return "Hello, World! Updated"
}

// Hello returns "Hello, {Name}!".
func Hello(name string) string {
	return "Hello, " + name + "!"
}

// notExported is not exported.
func notExported() string {
	return "Not exported!"
}

// notExported2 is not exported.
func notExported2() string {
	return "Not exported2!"
}
