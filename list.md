- Introduction to Go Programming language
- Variables, Pointers & Types in Go
- Packages in go with examples inducing docs, tests and benchmarking
- Go modules
- Closures
- Json encoding & Httpclient
- Higher order funcs and methods
- Interfaces & composition
- Channels and concurrency basics


Intro:
    [
        Intro slides
        go.dev
        setup.md
        pkg.go.dev site
        sample intro code, main, package, go routine.             
            `with runtime.GOMAXPROCS(0)`          
        go tour
    ]

Types:
    - Types are collection of things that are similar in construction and have same methods, funcs operating on them
built in types
variables and pointers    
struct
initialization & constructor
    
    [
        word size 
        unsafe.sizeof
        bytes when dealing with network or disk storage 
        big endian little endian
        binary data
        binary operators
        encoding/binary package is there however not so 
    ]
slices & maps
fmt format options
Linklist
*traversal assignment*
strings blog: byte vs string (utf-8) & rune(int32)
    https://go.dev/blog/strings
- interface types are collection of things that have similar traits (via methods)

packages
Fibonacci https://www.mathsisfun.com/numbers/fibonacci-sequence.html
email validation with tests
email send service for invoice may be
*enhancement to above*

slice
map
*url shortner*
    url to unique number map in database (key value)
    number to string
    string to base64encode
    when redirect, do reverse
    base64 to string
    string to number 
    number lookup in db for long url and redirect


interface and composition
    Why inheritance is problem. 
        Animal inheritance example. Animal <- {Mamal, egg(duckbill), noegg(rest all)}, {reptiles}
        Person<-Employee<-{Manager, Developer, Architect, Lead, LeadPlusManager...}



