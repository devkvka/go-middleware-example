# Middleware example

This repository contains a simple showcase of a very easy way to add middleware
to any route in Go, using the functional options pattern for chaining.

### Examples of the server running for reference

('this' refers to the ip:port the project is running on obviously)

```
$ curl this
-> Root called
$ curl this/secure
-> unauthorized
$ curl this/something
-> unauthorized

$ curl -H "X-Auth-Token: secure" this
-> Root called
$ curl -H "X-Auth-Token: secure" this/secure
-> Secure called
$ curl -H "X-Auth-Token: secure" this/something
-> Something called
```
