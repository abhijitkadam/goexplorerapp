
### 1. Intro
   > main package, func, println
   > go key word with sleep
   > runtime.GoProc 
   
   > - Bank Balance example with 2 go routines. One adding 10. Another deducting 10. use sync.waitgroup. This will be incorrect
   > - Change runtime.GOMAXPROCS(1) to sync however single proc
   > - Revert MAX procs and then use sync Mutex to make it correct
   > - Use channel way to update balance

   > Generate thousands of go routines without any effort. Use counter over channel to tally count. Each routine will signal the channel and collector will add one. Measure time

### 2. util: Doubler 
> Double of uint8
> Do with normal operator then binary
> Should fail if input greater than half of MAX::uint8. 
        `const MaxUint = ^uint8(0)` 
        or use `math.MaxUint`
        `half is Max/2 ( i.e Max >> 1)`
> Add guard for overflow and return error
> Include tests and bench
> Include Docs


### 3. Email Validation

### 4. Invoice sample
Format vs text template
text template you need not bother about rearranging text and values
````
type InvoiceItem struct {
	Name  string
	Price float32
	Units int32
}

type Invoice struct {
	OrderId int32
	Name    string
	Address string
	Date    string
	//Items   []InvoiceItem
}

const templateformat = `
Date: %s

Invoice No: %d is confirmed.
Following are the details

Name : %s,

Address:
%s
`

const textTemplate = `
Date: {{.Date}}

Invoice No: {{.OrderId}} is confirmed.
Following are the details

Name : {{.Name}},

Address:
{{.Address}}
`
````

