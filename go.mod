module example.com/hello

go 1.21.8

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/example/hello v0.0.0-20240205180059-32022caedd6a
)

replace example.com/greetings => ../greetings
