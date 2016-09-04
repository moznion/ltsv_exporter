ltsv_exporter
==

An [Prometheus](https://prometheus.io/) exporter that scrapes LTSV.

Example
--

### Supplied LTSV

```
size:123	duration:4.097	time:2016-09-04T22:22:14+09:00
```

### Exporter's output

```
# HELP ltsv_value LTSV value
# TYPE ltsv_value gauge
ltsv_value{key="duration"} 4.097
ltsv_value{key="size"} 123
```

This exporter only collects numeric values. Other type values are ignored.

Description
--

This exporter can scrape LSTV string and export contents of the LTSV as Prometheus format.

Usage
--

```
$ ltsv_exporter --url https://example.com/sample.ltsv
```

### Options

```
-h, --help          display help information
-v, --version       display version and revision
-p, --port[=6666]   set the port number to listen
-u, --url          *set a URL of the LTSV
```

How to build
--

```
$ make VERSION=1.0.0
```

Author
--

moznion(<moznion@gmail.com>)

License
--

```
The MIT License (MIT)
Copyright © 2016 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

