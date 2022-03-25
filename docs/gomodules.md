[Home](../Readme.md)
# Go Modules
Go to the directory where you want to create Module and run following:
> `go mod init <moduleAppName>` 

`go get <module>` will get the module. However running `go build` and `go test` commands add new dependencies to go.mod as needed.  

To list all the dependencies:
> `go list -m all`

To clean unsed dependencies:
> `go mod tidy`

#### Module versioning

A semantic version has three parts: `major, minor, and patch` 

For example, for `v0.1.2`, the major version is 0, the minor version is 1, and the patch version is 2.

Semantic versioning follows the `import compatibility rule` by including the `major version` in the `import path`. *MAJOR version is incremented when a backward incompatible change to the public API of module is made*

A module `my/thing` is imported as `my/thing` for `v0`, the **incompatibility** period when breaking changes are expected and not protected against.

`v1`, will be first stable major version. It will be include as `my/thing/v1`

**Incompatible packages (those with different major versions) have different names**



##### [Reference link](https://go.dev/blog/using-go-modules)






