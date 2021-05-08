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
go get gonum.org/v1/plot/                       # plotting library
go get github.com/aws/aws-lambda-go/lambda      # lambda model
go get github.com/aws/aws-sdk-go/service/s3     # s3 sdk
go get github.com/markus-wa/quickhull-go        # quickhull 3D library
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
- Write tests for the algorithm
- Implement the 3D algorithm
- Implement a visualizer for the 2D and/or 3D cases
- Add support for (at least one) common 3D CAD format
- Generate a mesh from a 3D convex hull
- Document all the targets in the Makefile (and update this README in general)
- Document the project using draw.io
- Read the papers and lectures in http://www.cs.jhu.edu/~misha/Spring20/
- Each time you create a new API gateway it triggers a rebuild?

### AWS TODO's:
- AWS Step Functions and/or AWS SNS (note -- cannot use step functions?)
- AWS CloudFront and Route 53 for SSL & domain names (note -- both are denied)

[gonumplot]: https://github.com/gonum/plot
