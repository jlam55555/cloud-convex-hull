# Distributed Convex Hull
Distributed convex hull cloud service implemented on AWS

<!-- TODO: summary here -->

<!-- TODO: report incoming -->

---

### Architecture

<!-- TODO: include diagram of architecture -->

---

### Build instructions

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
go run src/convexhull/main.go
```

<!-- TODO: include more detailed build instructions here -->

##### Tests
```bash
go test src/convexhull/main_test.go
```

<!-- TODO: include a list of tests -->

[gonumplot]: https://github.com/gonum/plot