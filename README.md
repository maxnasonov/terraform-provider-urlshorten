# Terraform Provider for Shorten URL

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

## Introduction

The Shorten URL provider uses a service and gives a short url for the given url.

## Using the provider

If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory, run `terraform init` to initialize it.

ex.

```hcl
  # Configure provider with your service details
  provider "urlshorten" {
    # Shorten URL service's password
    password = "admin123"
    # it can be set using the environment variable SHORTEN_PASSWORD

    # Service's name
    service = "tinyurl"
    # it can be set using the environment variable SHORTEN_SERVICE
  }

  # Retrieve project's templates
  data "shorten_url" "response" {
    provider = urlshorten
    url     = "https://www.iana.org/dnssec/archive/launch-status-updates"
  }
  output "shorten_url_response" {
    value = data.shorten_url.response
  }
```

In the examples directory you can find more.

## Building The Provider

Clone this repository to: `$GOPATH/src/github.com/wastorga/terraform-provider-urlshorten`

```sh
$ mkdir -p $GOPATH/src/github.com/wastorga
$ cd $GOPATH/src/github.com/wastorga
$ git clone https://github.com/wastorga/terraform-provider-urlshorten.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/wastorga/terraform-provider-urlshorten
$ make build
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.15+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary
in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-urlshorten
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources.

```sh
$ make testacc
```

## Documentation

In the docs directory you can find the documentation.

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Shorten URL repository](https://github.com/wastorga/terraform-provider-urlshorten/issues).

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
