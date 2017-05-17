goamz
=====

Go package to interact with Amazon Web Services

[![Build Status](https://travis-ci.org/go-amz/amz.svg?branch=v1)](https://travis-ci.org/go-amz/amz)
[![GoDoc](https://godoc.org/gopkg.in/amz.v1?status.png)](http://godoc.org/gopkg.in/amz.v1)

**NOTE**: This is a stable branch, which is compatible with the original goamz source from [Launchpad](https://launchpad.net/goamz/trunk), except for the changed import paths. It is not actively developed, unlike the other branches (`v2`, `v3`).

Instructions
------------

Install the package with:

    go get github.com/veritone/amz/...

Import it with:

    import "github.com/veritone/amz/<package>"

Example:

    import "github.com/veritone/amz/ec2"

and use _ec2_ as the package name inside the code.
The same applies to the other sub-packages: _aws_, _s3_, etc.

For more details, visit the project page:

* http://wiki.ubuntu.com/goamz

and the API documentation:

* https://gopkg.in/amz.v1

Users Group
-----------

In addition to contacting directly one of the [maintainers](https://github.com/orgs/go-amz/people), users are encouraged to discuss goamz topics on the project's [Google Group](https://groups.google.com/forum/#!forum/goamz).

Issues
------

Please report bugs by opening an [issue](https://github.com/go-amz/amz/issues).

Contributing
------------

Contributors are most welcome!
Please have a look at [CONTRIBUTING.md](CONTRIBUTING.md) for details.

Authors
-------

List of people who made relevant contributions to goamz can be found in [AUTHORS.md](AUTHORS.md).

License
-------

goamz is licensed under LGPLv3, but includes an exception that allows
linking code statically. See the [LICENSE](LICENSE) for details.
