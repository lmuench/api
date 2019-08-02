# api
Command line HTTP client

### Request Response Cycle
```
+---+  api.Call  +---+ http.Request  +-----+ http.Request  +------+ HTTP request  +---+
|   |----------->|   |-------------->|     |-------------->|      |-------------->|   |
|cli|            |api|               |plugs|               |client|               | ? |
|   |<-----------|   |<--------------|     |<--------------|      |<--------------|   |
+---+ api.Answer +---+ http.Response +-----+ http.Response +------+ HTTP response +---+
```
