# geproxy

geproxy is a simple gemini server which accepts gemini urls (as url parameters) to proxy.

eg. `gemini://localhost?gemini://gemini.circumlunar.space/docs/specification.gmi` will return the response body from `gemini://gemini.circumlunar.space/docs/specification.gmi`

## Usage

```sh
$ geproxy --help
Usage of ./geproxy:
  -cert string
        certificate PEM file (default "cert.pem")
  -key string
        key PEM file (default "key.pem")
  -port string
        listening port (default "1965")
```