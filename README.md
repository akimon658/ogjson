# Open Graph to JSON server

Parses meta tags of given URL and returns as a JSON.
It will be useful to generate URL cards with [Hugo](https://gohugo.io), or other static site generator.

## Get started

Once you started the server, you can get a JSON by sending request to `localhost:8080` with a `url` parameter.

Example using `curl`:

```
$ curl http://localhost:8080/?url=https://example.com
{"Policy":{"TrustedTags":["meta","link","title"]},"Title":"Example Domain","Type":"","URL":{"Source":"https://example.com","Scheme":"https","Opaque":"","User":null,"Host":"example.com","Path":"","RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":"","RawFragment":"","Value":""},"SiteName":"","Image":[],"Video":[],"Audio":[],"Description":"","Determiner":"","Locale":"","LocaleAlt":[],"Favicon":"/favicon.ico"}
```

There are 2 ways to use the server on your environment.

### Use Docker

```bash
docker run --rm -p 8080:8080 akimon658/ogjson:latest
```

You may use the `-d` flag to start the container in [detached mode](https://docs.docker.com/engine/reference/run/#detached--d).

### Install using Go

```bash
go install github.com/Akimon658/ogjson@latest
```

Then you can use `ogjson` command.

## License
[MIT](./LICENSE)

## Author
Akimo ([@akimon658](https://github.com/Akimon658))
