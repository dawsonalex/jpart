# golang-cli

![example workflow](https://github.com/dawsonalex/jutil/workflows/Build/badge.svg)

Get the value of a JSON element via it's path name.

## Usage

```bash
=======
jutil
=======
Usage: jutil [-p <path>]
Options:
  -p string
        A path in the format 'first.second.third' which defines the data you want. Leave empty to see the whole JSON input.
```

## Building

Jutil requires `go 1.18 beta1`. To install it follow the instructions on the _prerequisites_
section [here](https://go.dev/doc/tutorial/generics).

After that you can build binaries by running `make release`.

## GitHub Workflow

The GitHub workflow defined in `base.yml` attempts to do some common things in a simple way. Currently, it does the
following steps under a single job called `Build`:

    - Set up Go environment.
    - Run Go Tests (ignoring if there are none).
    - Runs `make release` to create all binaries.
    - Bump the version based on the commit message.
        - Use `#major`, `#minor`, or `#patch` in your commit message to bump the version and create a new release.
        - Leaving out the above tags will not create a new tag or release version.
    - Generate release logs from the commits between this tag and the last.
    - Create a GitHub release and upload the content of `bin`.