# api
Command line HTTP client

### Usage
Execute an HTTP GET request against the given URL

Usage:
  `api GET [URL] [flags]`

### Session Cookies
api supports session cookies out-of-the-box

### Request Response Cycle
```
+---+  api.Call  +---+ http.Request  +-----+ http.Request  +------+ HTTP request  +---+
|   |----------->|   |-------------->|     |-------------->|      |-------------->|   |
|cli|            |api|               |plugs|               |client|               | ? |
|   |<-----------|   |<--------------|     |<--------------|      |<--------------|   |
+---+ api.Answer +---+ http.Response +-----+ http.Response +------+ HTTP response +---+
```
