shim
====

shim is a super simple http proxy.

It is most useful for interacting with third party APIs that do not expose JSONP or the appropriate origin controls.

Examples.

```
> go run src/shim.go -h "http://api.flickr.com/" &
2013/08/22 13:07:58 Ok, we are listening on port 8080 and connecting to http://api.flickr.com/ ...
> curl -I "localhost:8080/services/rest/?method=flickr.photos.getRecent&api_key=[API_KEY]&format=json&nojsoncallback=1"
HTTP/1.1 200 OK
Cache-Control: private
Connection: Keep-Alive
Content-Encoding: gzip
Content-Length: 20
Content-Type: application/json
```

There are a few command line options.

Listen on another port:
```
go run src/shim.go -l 3333 -h "http://example.com"
```

Print debugging information:
```
go run /src/shim.go -v -h "http://example.com"
```


