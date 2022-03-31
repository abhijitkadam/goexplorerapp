[Home](../Readme.md)

[fmt format](#fmt-format)<br/>
[time format layout](#golang-time-format)<br/>
[collections](#collections)<br/>
[enum with const](#enums-with-const)<br/>

## fmt format
common formats are below. default is %v if you use that then it covers most
````
%v      default
%+v     adds fields name 
%T      type of value
%d      number base 10
%f      decimal floating point
%.2f    default width, precision 2
%s      bytes if string or slice
%p      pointer address, for slice address of 0th element
%b      base 2
%x      hexadecimal
````

Most used fmt funcs (by me):

````
func Println(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)

func Sprintf(format string, a ...any) string
````
## Golang time format
The reference time used in the layouts is:
Mon Jan 2 15:04:05 MST 2006

It is the numbers 1 2 3 4 5 6 7

1: month (January, Jan, 01, etc)
2: day
3: hour (15 is 3pm on a 24 hour clock)
4: minute
5: second
6: year (2006)
7: timezone (GMT-7 is MST)
## Collections:
Arrays:
````
nums := []int32{2,4,64,9}
primes := [6]int{2, 3, 5, 7, 11, 13}
````
Slices:
`a[low : high]`
including low and *exluding* high

creating a slice with the underlying array:
`b := make([]int, 0, 5) // len(b)=0, cap(b)=5`

## Enums with const
````
type Level int16`


const (
    Bronze Level = iota + 1
    Silver 
    Gold
)
````
note:
> iota + 1 => so that we do not init default to Bronze as 0


## select
1. for time.After: https://pkg.go.dev/time#After
````

    select {
    case out, ok := <-ch1:
        // handle out
    case <-time.After(1 * time.Second):
        fmt.Println("timed out")
    }

````

2. for timeout context: https://pkg.go.dev/context#example-WithTimeout
````
    const shortDuration = 1 * time.Millisecond

    ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
    defer cancel()

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
````
3. [With context](https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html)

