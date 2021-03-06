shim
====

shim is a super simple http proxy.

It is most useful for web development with third party APIs that do not expose JSONP or the appropriate origin controls.

Examples

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

shim was written to be used with jQuery, D3, or your ajax library of choice:

```
// run ./shim_osx -h "http://frustrating-api.example.com" in the terminal.
// then you can use code like this to transparently connect to a previously unusable API.

$.getJSON("localhost:8080/stats.json", function(data) {
  console.log(data);
  
  // etc.
}
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

NOTE: shim rewrites the Host header of your local request to match the host of the destination URL. Don't use shim if this violates the terms of the API you are using.


