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

License / Copyright
-------------------

This software is released under the MIT License.

Copyright (c) 2014 Marc Falzon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
