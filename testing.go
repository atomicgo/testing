// Package testing is a module which is only used for testing purposes.
package testing

// HelloWorld returns "Hello, World!".
func HelloWorld() string {
	return "Hello, World!"
}

// Hello returns "Hello, {Name}!".
func Hello(name string) string {
	return "Hello, " + name + "!"
}

// notExported is not exported.
func notExported() string {
	return "Not exported!"
}
