# Golang

- 🚀 **Guía Go:** [Aprende a programar](https://docs.google.com/document/d/1yHK6Aqzd-KHMJL1pZBQXvdZURJFIkfaUxYbQNnLTVl4/edit)

- 🚀 **Golang book:** [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)

# Learn How To Code: Google’s Go (golang) Programming Language

## Section 2: Variables, Values, & Type

- [2.01 Playground](section_02/2.01.md)
- [2.02 Hello World](section_02/2.02.md)
- [2.03 Short Declaration Operator](section_02/2.03.md)
- [2.04 The Var Keyword](section_02/2.04.md)
- [2.05 Exploring Type](section_02/2.05.md)
- [2.06 Zero Value](section_02/2.06.md)
- [2.07 The fmt Package](section_02/2.07.md)
- [2.08 Creating Your Own Type](section_02/2.08.md)
- [2.09 Conversion, Not Casting](section_02/2.09.md)

## Section 3: Exercises - Level 1

- [3.01 Hands-on Exercise #1](section_03/3.01.md)
- [3.02 Hands-on Exercise #2](section_03/3.02.md)
- [3.03 Hands-on Exercise #3](section_03/3.03.md)
- [3.04 Hands-on Exercise #4](section_03/3.04.md)
- [3.05 Hands-on Exercise #5](section_03/3.05.md)
- [3.06 Hands-on Exercise #6](section_03/3.06.md)

## Section 4: Programming Fundamentals

- [4.01 Bool Type](section_04/4.01.md)
- [4.02 How Computers Work](section_04/4.02.md)
- [4.03 Numeric Types](section_04/4.03.md)
- [4.04 String Type](section_04/4.04.md)
- [4.05 Numeral Systems](section_04/4.05.md)
- [4.06 Constants](section_04/4.06.md)
- [4.07 Iota](section_04/4.07.md)
- [4.08 Bit Shifting](section_04/4.08.md)

## Section 5: Exercises - Level 2

- [5.01 Hands-on Exercise #1](section_05/5.01.md)
- [5.02 Hands-on Exercise #2](section_05/5.02.md)
- [5.03 Hands-on Exercise #3](section_05/5.03.md)
- [5.04 Hands-on Exercise #4](section_05/5.04.md)
- [5.05 Hands-on Exercise #5](section_05/5.05.md)
- [5.06 Hands-on Exercise #6](section_05/5.06.md)
- [5.07 Hands-on Exercise #7](section_05/5.07.md)
- [5.07a Hands-on Exercise #7 Answers](section_05/5.07a.md)

## Section 6: Control Flow

- [6.01 Understanding Control Flow](section_06/6.01.md)
- [6.02 Loop - Init, Condition, Post](section_06/6.02.md)
- [6.03 Loop - Nested Loops](section_06/6.03.md)
- [6.04 Loop - For Statement](section_06/6.04.md)
- [6.05 Loop - Break & Continue](section_06/6.05.md)
- [6.06 Loop - Printing ASCII](section_06/6.06.md)
- [6.07 Conditional - If Statement](section_06/6.07.md)
- [6.08 Conditional - If, Else if, Else](section_06/6.08.md)
- [6.09 Loop, Conditional, Modulus](section_06/6.09.md)
- [6.10 Conditional - Switch Statement](section_06/6.10.md)
- [6.11 Conditional - Switch Statement Documentation](section_6/6.11.md)
- [6.12 Conditional Logic Operators ](section_06/6.12.md)

## Section 7: Exercises - Level 3

- [7.01 Hands-on Exercise #1](section_07/7.01.md)
- [7.02 Hands-on Exercise #2](section_07/7.02.md)
- [7.03 Hands-on Exercise #3](section_07/7.03.md)
- [7.04 Hands-on Exercise #4](section_07/7.04.md)
- [7.05 Hands-on Exercise #5](section_07/7.05.md)
- [7.06 Hands-on Exercise #6](section_07/7.06.md)
- [7.07 Hands-on Exercise #7](section_07/7.07.md)
- [7.08 Hands-on Exercise #8](section_07/7.08.md)
- [7.09 Hands-on Exercise #9](section_07/7.09.md)
- [7.10 Hands-on Exercise #10](section_07/7.10.md)

## Importing Multiple Packages

Previously, we imported a single package, fmt. But, we can import so many more! Go has an extensive list of packages that we can take advantage of. [Here’s a list of Go’s standard packages](https://pkg.go.dev/std)

To import multiple packages we can add multiple import statements, like so:

```go
import "package1"
import "package2"
```

Or we can use a single import with a pair of parentheses that contain our packages:

```go
import (
  "package1"
  "package2"
)
```

Notice, when using an import with parentheses, we’re not separating each package with a comma. Instead, each package is on a different line.

We can also provide an alias to a package by including the alias name before the file. Including an alias will make it easier to refer to the package without typing out the full package name:

```go
import (
  p1 "package1"
  "package2"
)
```

```go
import p1 "package1"
import p2 "package2"
```

In the example above we’ve aliased package1 as p1 and now we can call functions from package1 by using p1 like:

```go
p1.SampleFunc()
```

Instead of:

```go
package1.SampleFunc()
```

Example:

```go
package main

import "fmt"
import t "time"

func main() {
	fmt.Println(t.Now())
}
```

## Comments

We can’t always be there in person to explain to the next developer (or even our future self) what our code does or what our intentions were when writing it. That’s where comments come in.

Comments are ignored by the compiler and helpful for many things. We can use comments to explain our code, or thought process, and even to comment out code to debug (fix errors).

Go also encourages the use of comments for describing what functions do, it gets used by Go’s documentation tool. In fact, it started on their Github Wiki page!

There are two types of comments in Go. A line comment is created using //:

```go
// This entire line is ignored by the compiler
// fmt.Println("Does NOT print")
fmt.Println("This gets printed!") // This part gets ignored
```

Notice that line comments out the content to the right of //.

There are also block comments that can span multiple lines— it starts with /_, ends with a _/ and envelopes everything inside (including new lines):

```go
/*
This is ignored.
This is also ignored.
fmt.Println("This WON'T print!")
*/
```

Example:

```go
package main

import "fmt"

func main() {
  // Are we racing or coding?
	/*
  fmt.Println("Ready")
  fmt.Println("Set")
  */
  fmt.Println("Gooooo!")
}
```

## Go Resources

Learning a new language like Go involves learning the accompanying rules and syntax. But, we don’t have to commit everything to memory! It’s ok to search things up, in fact, that’s what all good programmers do!

If you’re ever stuck on something, check out:

- [Codecademy’s Forums](https://discuss.codecademy.com/)

- [Stack Overflow's Go questions](https://stackoverflow.com/collectives/go)

- [Go's official site](https://golang.org/)

- [Google](https://google.com)

In addition to online resources, Go also has it’s own built-in documentation system. To use it, in the command line, use the command go doc followed by a package name. For instance, to find out more information on the fmt package, you can use the command:

```
go doc fmt
```

In the terminal, you’ll see at the top:

```
package fmt // import "fmt"

Package fmt implements formatted I/O with functions analogous to C's printf
and scanf. The format 'verbs' are derived from C's but are simpler.
...
```

The information returned actually spans quite a few lines, the example above is truncated. To get more specific information about a function in a package (like fmt‘s Println function) append .Println (or .println, the capitalization of the function doesn’t matter to go doc) to the command:

```
go doc fmt.Println
```
