# Building the application
The application consists of two parts: the front-end (Vue/Vite application), and the back-end (Go lambdas).

---

### The Go lambdas backend
There are (currently) two Go lambdas in the backend: the presign (`src/chpresign`) and convex hull (`src/chhull`) lambdas.

##### Prerequisites
Make sure you have golang (1.15+) installed. You should append the current directory (top-level of the repo) to your `GOPATH`. (Note that order is important; the dependencies will be installed to the first entry of `$GOPATH`.)
```bash
$ export GOPATH=$GOPATH:$(pwd)
```
Install dependencies.
```bash
$ go get gonum.org/v1/plot/                       # plotting for 2D convex hull
$ go get github.com/aws/aws-lambda-go/lambda      # lambda sdk
$ go get github.com/aws/aws-sdk-go/service/s3     # s3 sdk
$ go get github.com/markus-wa/quickhull-go        # quickhull 3D library
```

##### Build instructions
This will be called automatically when calling `make all`. See [`scripts/go_compile.mk`](scripts/go_compile.mk) and [`scripts/lambda_deploy.mk`](scripts/lambda_deploy.mk) for more details.

For the first time build, assuming that the lambda is not yet created (you probably don't need this; call `make all` to provision AWS resources in the correct order):
```bash
$ make lambda-deploy
```
To update the existing lambdas:
```bash
$ make lambda-update
```
You may also simply build the files without updating the lambda if needed. See the aforementioned Makefiles for details.

To clean temporary build files:
```bash
$ make target-clean
```

---

### The Vute/Vite frontend
The front-end (`src/chfrontend`) relies on the API Gateway already deployed (which depends on the lambdas being deployed). If the API Gateway is relaunched, make sure to rebuild the app (this is because the API url will likely change, and thus all the Vue endpoints will be broken).

##### Prerequisites
Make sure you have Node.JS and `npm` installed.
```bash
$ cd src/chfrontend
$ npm i
```

##### Build instructions
This will be called automatically when calling `make all`. This not only calls the Vite NPM build command, but also gets the API URL using `aws-cli` and writes it to `src/chfrontend/.env` so that it will be compiled into the application.
```bash
$ make build-website
```
If you want to update the website S3 bucket:
```bash
$ make host-bucket-sync
```

##### Testing instructions
You can test the application locally by using the local Vite testing server. Make sure to set the API CORS policy to allow `*` origins. By default, this will open a local testing session at `localhost:3000`.
```bash
$ cd src/chfrontend
$ npm run dev
```