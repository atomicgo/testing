/**
Package template is a template for AtomicGo projects.

Here is the place for the description of the module.

You can use **Markdown** here.
*/
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

// notExported2 is not exported.
func notExported2() string {
	return "Not exported2!"
}
