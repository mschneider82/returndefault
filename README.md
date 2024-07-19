# defaultvalue

A golang library that returns the given default value if the value is set to
golangs default.

It uses golangs generics with type comparable.

## Usage

```go
import "schneider.vip/returndefault"

...

returndefault.Value(123, 8) # returns 123
returndefault.Value(0, 8) # returns 8

returndefault.Value("x", "unset string") # returns "x"
returndefault.Value("", "unset string") # returns unset string
```