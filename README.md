# Terrago

This git repository contains one Go module with one Go package at the root directory.
The module path is `github.com/h0tbird/terrago` which is also the import path used for the root directory.
It is intended to be consumed as a library by adding `import "github.com/h0tbird/terrago"` to your code.

### Architecture
A `Manifest` is a collection of `Resource`s stored in a hash table and its corresponding dependency DAG.
When a `Manifest` is applied, the DAG is generated and walked and that's when `Resource`s are reconciled.

### Development
Run `make dag-code` to fetch the DAG code from upstream and adjust it so it can be used by this project.
