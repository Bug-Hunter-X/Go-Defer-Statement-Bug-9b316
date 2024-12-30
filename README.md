# Go Defer Statement Bug

This repository demonstrates a common mistake related to the `defer` statement in Go. Many developers are unaware that `defer` statements execute after the function has finished, even if it returns early.

## The Bug

The `bug.go` file contains a program with a `defer` statement inside an `if` block. Because the if statement returns early, the deferred statement won't be executed at the moment you might expect.

## The Solution

The `bugSolution.go` file provides a solution that demonstrates how to avoid the unexpected behavior.  

## How to Reproduce

1. Clone this repository
2. Run `go run bug.go` to see the unexpected output.
3. Run `go run bugSolution.go` to see the intended behavior. 