[Home](../Readme.md)

#### Run tests and benches

uses regex 
> ^ => starts with
> $ => ends with

````
go test .          # Run all tests in the current directory
go test ./...      # Run all tests in the current directory and sub-directories
go test ./foo/bar  # Run all tests in the ./foo/bar directory
````
  


````
go test -v -run=^TestFooBar$ .          # Run the test with the exact name TestFooBar
go test -v -run=^TestFoo .              # Run tests whose names start with TestFoo
go test -v -run=^TestFooBar$/^Baz$ .    # Run the Baz subtest of the TestFooBar test only

go test -bench=. ./...                        # Run all benchmarks and tests
go test -run=^$ -bench=. ./...                # Run all benchmarks (and no tests)
go test -run=^$ -bench=^BenchmarkFoo$ ./...   # Run only the BenchmarkFoo benchmark (and no tests)
````
> go test -race ./...
        
More info:
        
        
        
        
| Common links |
| ----------- |
|https://dave.cheney.net/2019/05/07/prefer-table-driven-tests|
|https://www.alexedwards.net/blog/an-overview-of-go-tooling#testing|
|https://www.alexedwards.net/blog/an-overview-of-go-tooling#running-and-comparing-benchmarks|
|https://pkg.go.dev/cmd/go#hdr-Testing_flags|
|[Assertions for tests](https://github.com/stretchr/testify) <br/>|
|[HTTP mocking](https://github.com/h2non/gock)    <br/>   |

<hr/>

### Profiling

Ways to get profile information

1. Through tests and benchmarks
>using the **-cpuprofile & -memprofile** flags

````
    go test . -bench . -cpuprofile prof.cpu
    go test -bench . -benchmem -cpuprofile prof.cpu -memprofile prof.mem
````
To analyse:    
 `go tool pprof [binary] prof.cpu`
    *binary will be generated while running test/bench

2. Endpoints in  service
    > import _ net/http/pprof

    `go tool pprof -seconds 5 http://localhost:9090/debug/pprof/profile`


3. From code
    by calling `runtime.StartCPUProfile` or `runtime.WriteHeapProfile`
    
    
[pprof, Profiling Go Program](https://pkg.go.dev/runtime/pprof#hdr-Profiling_a_Go_program)

[pprof Readme](https://github.com/google/pprof/blob/master/doc/README.md)

Analyse:
`go tool pprof stats.test prof.cpu`
or
`go tool pprof -alloc_objects stats.test prof.mem`

Commands:
````
top10           #show top 10 funcs by time spent only in that function
top10 -cum      #top 10 funcs by time spent in that func or a func it called.
list regex      #show code for functions matching regex
disasm regex    #show the disassembly for functions matching regex
web             #Visualize graph through web browser
````
<hr/>

[A tutorial on profiling & flamegraphs](https://www.matoski.com/article/golang-profiling-flamegraphs/)
