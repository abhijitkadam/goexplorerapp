[Home](../Readme.md)
#### Static Analysis and format
````
go fmt $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
staticcheck $(go list ./... | grep -v /vendor/)
````
<hr/>

Notes: <br/>
`grep -v /vendor` will exclude the vendor directory
To install static check: https://staticcheck.io/docs/install


#### add vendoring
go get package
go mod vendor
cd vendor
git add .