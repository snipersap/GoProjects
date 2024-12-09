# Memory Layout

In Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct. For example this struct:

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

Looks like this in memory:

[struct stats memory layout with 4 bytes](./stats-struct.png)
## Field ordering... matters?

the order of fields in a struct can have a big impact on memory usage. This is the same struct as above, but poorly designed:

type stats struct {
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}

It looks like this in memory:

[struct stats memory layout with 4 bytes](./stats-struct-wasted-mem.png)

Notice that Go has "aligned" the fields, meaning that it has added some padding (wasted space) to make up for the size difference between the `uint16` and `uint8` types. It's done for execution speed, but it can lead to increased memory usage.
## Should I panic?

To be honest, you should not stress about [memory layout](https://go101.org/article/memory-layout.html). However, if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help. You can also use the [reflect](https://pkg.go.dev/reflect) package to debug the memory layout of a struct:
```bash
typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
```
## Real story

I once had a server in production that held a lot of structs in memory. Like hundreds of thousands in a list. When I re-ordered the fields in the struct, the memory usage of the program dropped by over 2 gigabytes! It was a huge performance win.
## Assignment

Our over-engineering boss is at it again. He's heard about memory layout and wants to squeeze every last byte out of our structs.

Run the tests to see the current size of the structs, then update the struct definitions to minimize memory usage.

### Test results
```bash
Pass 1
$> go test
---------------------------------
0 passed, 2 failed
--- FAIL: Test (0.00s)
    main_test.go:37: ---------------------------------
        Inputs:     (contact)
        Expecting:  24 bytes
        Actual:     32 bytes
        Fail
    main_test.go:37: ---------------------------------
        Inputs:     (perms)
        Expecting:  16 bytes
        Actual:     24 bytes
        Fail
FAIL
exit status 1
FAIL    memory-layout-in-struct 0.200s

Pass 2
$> go test
---------------------------------
Inputs:     (contact)
Expecting:  24 bytes
Actual:     24 bytes
Pass
---------------------------------
Inputs:     (perms)
Expecting:  16 bytes
Actual:     16 bytes
Pass
---------------------------------
2 passed, 0 failed
PASS
ok      memory-layout-in-struct 0.271s
```