[Home](../Readme.md)

##### Enums with const
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

##### Collections:
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