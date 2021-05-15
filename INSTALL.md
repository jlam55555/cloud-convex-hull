# Installing the application to the cloud
See the [top-level Makefile](Makefile). This contains some top-level commands to build and install the files, as well as general configuration settings. (See [BUILD.md](BUILD.md) for build details and prerequisites.)

You should not have to modify any of the parameters in the top-level Makefile except for the bucket name, because it has to be globally unique. (Note that this also affects the output URL). Change these as needed. The default values are shown below.
```bash
HOST_BUCKET_NAME?=$(APP_PREFIX)hostbucket
UPLOAD_BUCKET_NAME?=$(APP_PREFIX)uploadsbucket
```

Additional sub-install steps can be found in the [`scripts/`](scripts) directory. You should not have to modify any of the variable declarations in this directory.

This requires [`aws-cli` version 2][aws-cli-v2] and GNU `make`.

To build and install the whole app:
```bash
$ make all
```

This will build the app to `http://[HOST_BUCKET_NAME].s3-website-[AWS_REGION].amazonaws.com/`, where `HOST_BUCKET_NAME` and `AWS_REGION` are specified in the Makefile.

To tear down the whole app:
```bash
$ make clean
```

[aws-cli-v2]: https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html
