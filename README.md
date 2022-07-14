# Open Graph to JSON server

Parses meta tags of given URL and returns as a JSON.
It will be useful to generate URL cards with [Hugo](https://gohugo.io), or other static site generator.

## Get started

Once you started the server, you can get JSON response by sending a request to `localhost:8080` with `url` parameter.

Example using `curl`:

```bash
curl -s http://localhost:8080/?url=https://ogp.me | jq .
```

<details>
  <summary>Output</summary>

```json
{
  "Policy": {
    "TrustedTags": [
      "meta",
      "link",
      "title"
    ]
  },
  "Title": "Open Graph protocol",
  "Type": "website",
  "URL": {
    "Source": "https://ogp.me",
    "Scheme": "https",
    "Opaque": "",
    "User": null,
    "Host": "ogp.me",
    "Path": "",
    "RawPath": "",
    "ForceQuery": false,
    "RawQuery": "",
    "Fragment": "",
    "RawFragment": "",
    "Value": "https://ogp.me/"
  },
  "SiteName": "",
  "Image": [
    {
      "URL": "https://ogp.me/logo.png",
      "SURL": "",
      "Type": "",
      "Width": 300,
      "Height": 300,
      "Alt": ""
    }
  ],
  "Video": [],
  "Audio": [],
  "Description": "The Open Graph protocol enables any web page to become a rich object in a social graph.",
  "Determiner": "",
  "Locale": "",
  "LocaleAlt": [],
  "Favicon": "/favicon.ico"
}
```
</details>

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

## License
[MIT](./LICENSE)

## Author
Akimo ([@akimon658](https://github.com/Akimon658))
