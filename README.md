[![Build Status](https://travis-ci.com/dimw/cidrls.svg?branch=master)](https://travis-ci.com/dimw/cidrls)

# CIDRLS

This command-line tool converts the list of IP addresses in CIDR format to a list of IP addresses.

## Usage

Print the list of IPs for the given addresses in CIDR format:

```
$ go run main.go --clean 192.168.11.0/30 127.0.0.1/32
192.168.11.1
192.168.11.2
127.0.0.1
```
