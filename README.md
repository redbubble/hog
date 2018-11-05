# hog

*Takes all the connections. Doesn't give them back.*

`hog` is a testing tool for finding how many simultaneous TCP connections a
service will accept.

- - - - - -

## Installation

With `go get`:

```
go get github.com/redbubble/hog
```

## Usage

```
hog --target 192.168.0.1 --port 80 --limit 1000
```

Where:

* `target` is the hostname or IP of the target service. Default: 0.0.0.0.
* `port` is the port of the target service. Default: 80.
* `limit` is the maximum number of simultaneous connections to be attempted. Default: 100.

## Limitations

### Open files

`hog` is limited by the number of files that the executing system is allowed to
open. One connection consumes one file descriptor.

On Unix-y systems, you can run `ulimit -n` to find out how many file descriptors
your system is allowed to consume.

This limit can generally be increased, but the method will vary depending on
your OS. It's likely that you'll need admin access to increase the limit.

## Contributing

If you'd like to contribute to `hog`, please see our [contributing doc](CONTRIBUTING.md).

## Maintainers

    delivery-engineers@redbubble.com

## License

`hog` is provided under an MIT license. See the [LICENSE](https://github.com/redbubble/hog/blob/master/LICENSE) file for
details.
