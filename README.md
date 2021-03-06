ruby-go-thrift
===

> A quick intro to thrift in Ruby and Go

*Uses Ruby 2.1.0, golang 1.4.2, thrift 0.9.3*
*The thrift-generated go has changed significantly between pre 0.9x versions - this was tested with 0.9.3*

This repo is meant as a gentle introduction to thrift in Ruby and Go. You should be familiar with both thrift, ruby and go beforehand:
- thrift.apache.org/tutorial/go

Installation
==
On Mac OS X
```
brew update
brew install thrift
sudo gem install thrift
```
If you're using a Ruby/gem version manager then substitute the appropriate statements in place of `sudo gem install thrift` above.


Note regarding python installation. It has [been noted](http://thrift-tutorial.readthedocs.org/en/latest/usage-example.html) that the homebrew package doesn't install all of the required python bindings. The stated workaround is to [download thrift-0.9.3.tar.gz](https://thrift.apache.org/download) uncompress it to a working directory and run `sudo python setup.py install` from within the `lib/py` directory. This didn't seem to be sufficient in my case and will revisit it at a later date.


This intro is split into 5 commits, each of which introduces a new concept.

Generating and running the Go server which must be done before the client.
```
thrift -r -gen go example.thrift
go run main.go
```

Any commit with a working server can be tested with a ruby client:
```
thrift -r -gen rb example.thrift
ruby RubyClient.py
```

If the python installation is suceessful, the the following can also be used:
```
thrift -r -gen py example.thrift
python client.py
```

### Commits

#### thriftfile

A simple thriftfile. One service extends another, and an exception is present.
The goal is to present as many esoteric cases in as little space as possible.

#### generate go thrift service

Notice the `//go:generate` statement at the top of `main.go`.
This is run using `go generate`, and will regenerate the go thrift files for our defined service.

Generated files are committed so that they can be consumed by others using `go get`.
It is the responsibility of the author to generate these files.

#### example/service handler

The handler is implemented, and plugged into a simple server.

#### multiple services

Another service is added to the thriftfile.

Each service lives on its own port. Services should be distinct enough that they should not share a port, barring other concerns. If you find two services needing to share a port, they should probably not be two services.

Server initialization is factored into its own file, so that each server can be restarted separately.

#### racy data access

The `count` function may be racy - it is refactored to use `chan`s.

If a service requires internal data access, it may become racy.
In that case, use `chan`s to avoid data corruption, adding a goroutine to the constructor to watch for requests.

In real scenarios, files with large implementations should also be split to maintain readability.

### NOTES

- Compiled thrift files go into a folder named after the thriftfile: example.thrift -> example/
- Service handlers live within a that services folder: example.Service -> example/service/handler.go

---

For a preconfigured environment to run through this tutorial, visit https://www.terminal.com/tiny/LGLivnbJox
