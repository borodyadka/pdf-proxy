# HTML to PDF rendering proxy

Render any page into PDF document.

## Usage

As docker image: `docker run --rm -p 8080:8080 borodyadka/pdf-proxy`

PDF proxy can be used behind reverse proxy such as Nginx:

```
set $pdfproxy_upstream 'http://pdf-proxy-service:8080';
location /api/v1/pdf {
    # to restrict rendering pages by mydomain.tld
    # and query this as /api/v1/pdf?page=/foo/bar.html
    # will render mydomain.tld/foo/bar.html
    rewrite ^/api/v1/pdf /render?url=https://mydomain.tld$arg_page break;
    proxy_pass $pdfproxy_upstream;
}
```

## Configuration

Logging level configurable with `LOG_LEVEL` environment variable. Possible values is: error, warning, info and debug. 

Address to listen configurable with `ADDRESS` environment variable. Default value is `:8080`.

## TODO

* [ ] edit exif data
* [ ] more tests (more than 0)

## License

[MIT](LICENSE)
