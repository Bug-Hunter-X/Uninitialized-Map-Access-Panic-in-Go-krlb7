# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: accessing an uninitialized map.  Attempting to access a map that hasn't been created will result in a runtime panic.

The `bug.go` file contains code that produces this panic. The `bugSolution.go` file shows the corrected code.

## Problem
In Go, maps must be explicitly initialized before use.  Failure to do so leads to a runtime panic with the message `assignment to entry in nil map`.

## Solution
Initialize the map before using it.  This can be done using the `make()` function or a literal map.
