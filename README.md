# Distributed Convex Hull
Distributed convex hull cloud service implemented on AWS

<!-- TODO: summary here -->

<!-- TODO: report incoming -->

---

### Architecture

<!-- TODO: include diagram of architecture -->

---

### Build instructions

<!-- TODO: these build instructions have to be updated for
	use with the new AWS Makefile -->

##### Prerequisites
Update your `GOPATH` environment variable to include the current directory:
```bash
export GOPATH=$GOPATH:$(pwd)
```
Download the [`gonum/plot`][gonumplot] library:
```bash
go get gonum.org/v1/plot/
```
(This will download to the first entry in your `GOPATH`.)

##### Compiling
```bash
go run convexhull
```

<!-- TODO: include more detailed build instructions here -->

##### Tests and Benchmarks
Using the `go test` tool will run the tests in `main_test.go`.
```bash
go test convexhull [-bench .] [-benchmem]
```

<!-- TODO: include a list of tests -->

---

### The List of TODO's
(In no particular order)
- Implement the 2D algorithm
- Write tests for the algorithm
- Implement the 3D algorithm
- Implement a visualizer for the 2D and/or 3D cases
- Add support for (at least one) common 3D CAD format
- Generate a mesh from a 3D convex hull
- Write the front-end
- Write an AWS Lambda handler
- Write an AWS S3 handler
- Learn how to use AWS Step Functions and/or AWS SNS
- Learn how to use AWS API Gateway

[gonumplot]: https://github.com/gonum/plot
