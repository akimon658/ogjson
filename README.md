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

## Advanced usage

### Change `User-Agent`

There are some websites that `ogjson` cannot access by default because of `User-Agent`.
To avoid it, you can use `user-agent` flag.

```bash
ogjson -user-agent "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"
```

```
$ curl "http://localhost:8080/?url=https://docs.github.com"
{"Policy":{"TrustedTags":["meta","link","title"]},"Title":"GitHub.com Help Documentation","Type":"article","URL":{"Source":"https://docs.github.com","Scheme":"https","Opaque":"","User":null,"Host":"docs.github.com","Path":"","RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":"","RawFragment":"","Value":"http://ghdocs-prod.azurewebsites.net:80/en"},"SiteName":"GitHub Docs","Image":[{"URL":"https://github.githubassets.com/images/modules/open_graph/github-logo.png","SURL":"","Type":"","Width":0,"Height":0,"Alt":""}],"Video":[],"Audio":[],"Description":"Get started, troubleshoot, and make the most of GitHub. Documentation for new users, developers, administrators, and all of GitHub's products.","Determiner":"","Locale":"","LocaleAlt":[],"Favicon":"/assets/cb-803/images/site/favicon.svg"}
```

## License
[MIT](./LICENSE)

## Author
Akimo ([@akimon658](https://github.com/Akimon658))
