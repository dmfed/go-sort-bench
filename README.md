# go-sort-bench

Code demonstrating a bit surprising results of Go **sort** package benchmarks.

Basically all of three methods: 
* sorting sort.Sort(sort.IntSlice(s))
* sorting with sort.Slice() with own less function
* sorting with sort.Sort() of custom type implementing sort.Interface 

demonstrate similar but not quite identical results.

**git clone --depth=1 https://github.com/dmfed/go-sort-bench** then **cd** into the repo and run **go test -bench=.** to see the results.

Looks like sort.Slice with custom lessfunc seems a bit quicker than other methods.
