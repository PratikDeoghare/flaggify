## Flaggify your structs

It is common practice in golang programs to put all the command line options in a struct. Create an object of this
struct and pass pointers of each of the fields to
`flag.Int`, `flag.String`, etc. Like this:

``` 
type Options struct {
	Hostname string
	Verbose  bool
	Retry    int
}

func main() {
	o := Options{}

	flag.StringVar(&o.Hostname, "hostname", "localhost", "Hostname to connect to")
	flag.BoolVar(&o.Verbose, "verbose", false, "Verbose output")
	flag.IntVar(&o.Retry, "retry", 3, "Number of retries")

	flag.Parse()

	// more stuff ...
}

```

Some programs take huge number of command line arguments and this becomes cumbersome.
See [this](https://github.com/kubernetes/kubernetes/blob/555623c07eabf22864f6147736fa191e020cca25/cmd/kubelet/app/options/options.go#L370)
. `flaggify` allows you to avoid all that.

Flaggify creates command line flag with `json` struct tag as name, default
value from provided default values struct and help text from the `x` struct tag (x for eXplanation).
With flaggify above code can be written as:

```
type Options struct {
	Hostname string `json:"hostname" x:"Hostname to connect to"`
	Verbose  bool   `json:"verbose" x:"Verbose output"`
	Retry    int    `json:"retry" x:"Number of retries"`
}

func main() {
	o := Options{}
	defaultOptions := Options{
		Hostname: "localhost",
		Verbose:  false,
		Retry:    3,
	}
	
	// Options.Retry will be ignored by flaggify
	// since it is already flaggified.
	flag.IntVar(&o.Retry, "retry", 4, "Number of retries") 
	
	flaggify.Flaggify(&o, defaultOptions)
	flag.Parse()
	// more stuff ...
}

```

Since the struct now has json struct tags you can also pass configuration from a json file instead of providing command
line arguments. This is useful when number of options is huge.

## TODO

- Add support for more field types. Currently only fields of type string, int, bool are supported.