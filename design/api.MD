### api design

## HTTP API

- get
```
GET https://example.com/api/v0/get?key=abc
```

- put 
```
POST https://example.com/api/v0/put

{key: "abc", value: "cba"}
```

- delete
```
DELETE https://example.com/api/v0/delete?key=abc
```


## JSON-RPC 2.0

- get
```
POST https://example.com/api/v0/


{"jsonrpc": "2.0", "method": "get", "params": {"key": 123}, "id": 1}
```

- put 
```
POST https://example.com/api/v0/

{"jsonrpc": "2.0", "method": "put", "params": {"key": 123, "value": 321}, "id": 1}
```

- put 
```
POST https://example.com/api/v0/

{"jsonrpc": "2.0", "method": "delete", "params": {"key": 123}, "id": 1}
```