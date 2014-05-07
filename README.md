send_riemann
============

Simple utility to send custom events to Riemann from command line.

Installation
------------

```
go get github.com/falzm/send_riemann
$GOPATH/bin/send_riemann -h
```

Usage
-----

```
send_riemann [options] <JSON-formatted Riemann event>
```

Options:

```
    -server="localhost:5555": Riemann server (host:port)
```

Riemann event JSON spec:

```javascript
 {
   "host": "hostname (default: local hostname)",
   "service": "service",
   "time": <UNIX epoch timestamp (default: now)>,
   "metric": <numeric value>,
   "state": "state (ok, warning, critical)",
   "ttl": <numeric value (seconds, default 10)>,
   "description": "some event description...",
   "tags": [ "tag1", "tag2", "tagN" ],
   "attributes": { "key1": "val1", "key2": "val2" }
 }
```

Example
-------

```
marc@cobalt:~% send_riemann -server=my.riemann.server:5555 '{"service": "blah", "metric": 42}'
```

Result in the Riemann console:

```
#riemann.codec.Event{:host "cobalt", :service "blah", :state "", :description "", :metric 42.0, :tags nil, :time 1399457657, :ttl 10.0}
```
