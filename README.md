# unifynil, unify nil and empty slices and maps in Golang

Empty slices and maps can be `nil` or not `nil` in Go.
It may become a nightmare in tests and JSON processing.

This package provides a simple way to convert all empty slices
or maps to `nil` or non-`nil` versions.

## Example

To replace all `nil` slices and maps to not-`nil`:

```go
import "github.com/starius/unifynil"

type Response struct {
    Items []int             `json:"items"`
    Users map[string]string `json:"users"`
}

res := &Response{} // Leave the slice and the map nil.

buf, _ := json.Marshal(res)
fmt.Println(string(buf))
// {"items": null, "users": null}

unifynil.Unify(res, unifynil.SliceToEmpty(), unifynil.MapToEmpty())

// Now res.Items = []int{} and res.Users = map[string]string{}.

buf, _ = json.Marshal(res)
fmt.Println(string(buf))
// {"items": [], "users": {}}
```


Convert back to `nil`s:

```go
unifynil.Unify(res, unifynil.SliceToNil(), unifynil.MapToNil())

// Now res.Items = nil and res.Users = nil.

buf, _ = json.Marshal(res)
fmt.Println(string(buf))
// {"items": null, "users": null}
```
