# Terrago

This git repository contains one Go module with:
- One public Go package at the root directory.
- Two internal Go packages generated by the `make dag-code` target.

The module path is `github.com/h0tbird/terrago` which is also the import path used for the root directory.
It is intended to be consumed as a library by adding the `import "github.com/h0tbird/terrago"` statement to
your code.

### Architecture
A `Manifest` is a collection of `Resource`s (stored in a hash table) and its corresponding dependency DAG.
When a `Manifest` is applied, the DAG is generated and walked and that's when `Resource`s are reconciled.

### Development
Upgrade all dependencies at once:
```
go mod edit -go=1.16
go get -u ./...
go mod tidy
```