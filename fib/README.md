How the Go Fibonnaci demo works

You may have seen the Fibonnaci example on https://play.golang.org/ - it produces Fibonnaci numbers by providing a function which, when called, will generate the *next* number in the sequence.

Here is the complete code:

```
package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
```

At first glance, this doesn't look like it should be working at all. And it was designed to show off a few traits of the language that are worth understanding.

It seems as though we have created a `struct` to hold the current state - but we haven't. We've only created a function.

### Closures

Closures refer to environment (or state) that belongs to a function. When a function is defined, the surrounding variables are captured and remain available to the function throughout its life. This can catch you out if you don't properly understand it, especially if you start writing concurrent code - where two goroutines should never try to access or modify the same data at the same time.

In the code above, the `fib` function returns another function which is defined inside the body of the `fib` function itself. When this happens, the two variables that have been defined (`a` and `b`) are _closured_ into the function.

The main function assigns this newly created function to a variable called `f` which is then called many times. Each time it is called, it is referring to the original `a` and `b` variables - even though these are nolonger accessible to us (in the `main` function).

### Multiple assignment

The other oddities here are these lines:

```
a, b := 0, 1
```

and

```
a, b = b, a+b
```

The first line is quite simple; it's a way of declaring and assigning values to two variables at the same time. It's equivalent to this:

```
a := 0
b := 1
```

The second line, however, is a bit cleverer. We are changing the values of `a` and `b` at the same time. Similar to the first line, except consider how you would write this on many lines:

```
a = b
b = a+b
```

It won't work. The top line changes `a`, interferring with the result of `a+b`. We'd need to store the original `a` value in an interim variable:

```
olda := a
a = b
b = olda+b
```

So you can see how combining it into one line simplifies our code - which is always worth doing.