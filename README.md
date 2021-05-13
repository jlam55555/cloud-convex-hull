# Distributed Convex Hull
Distributed convex hull cloud service implemented on AWS

<!-- TODO: summary here -->

<!-- TODO: report incoming -->

---

### Architecture
![Architecture diagram](assets/arch.png)

<small>

1. Access to uploads bucket by presigned URLs only.
2. Appopriate CORS headers for website use.
3. Presign lambda has read access to uploads bucket.
4. Convex hull lambda has GET/PUT access to uploads bucket.
5. Presign endpoint calls presign lambda.
6. Convex hull endpoint calls convex hull lambda.
7. API Gateway stage has logging set up.
8 Each lambda has logging set up (to different loggroups).
9. Lambdas need access to uploads bucket.
10. Route 53 acts as a DNS server.
11. Cloudfront acts as a CDN for the website, uploads bucket, API Gateway, and authorization page. It also supports HTTPS (necessary for Cognito).
12. Cognito authorizes API requests using the users access token.
13. Authorization lambda verifies the user's access token.
14. Listmodels endpoint calls listmodels lambda.
15. Listmodels lambda retrieves models entries from database.
16. New model lambda creates new model entry in database.
17. Listmodels/Newmodel lambdas require permissions for RDS.
18. On new upload (PUT request) or creation of a hull (a new model), newmodel lambda will trigger.

</small>

---

### Build instructions

<!-- TODO: these build instructions have to be updated for
	use with the new AWS Makefile -->

##### Prerequisites
Update your `GOPATH` environment variable to include the current directory:
```bash
export GOPATH=$GOPATH:$(pwd)
```
Download the following packages:
```bash
go get gonum.org/v1/plot/                       # plotting for 2D convex hull
go get github.com/aws/aws-lambda-go/lambda      # lambda sdk
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

<!-- TODO: describe all the packages here -->