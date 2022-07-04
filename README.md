## 0x90.st

HTTP POST files here:
    `curl -F 'file=@yourfile.png' https://0xg0.st`


### Shotout

This project is a simpler and minimal clone of [https://0x0.st](https://0x0.st).

Big thank's to <a>Mia Herkt</a> for the initiative.

This project is built totally in pure [Go](https://go.dev) only using the basic standard library.


### Operator notes
If you run a server and like this site, clone it! Centralization is bad.
If you have any problem, open up an issue in GitHub.

[https://github.com/joaoofreitas/0x90.st](https://github.com/joaoofreitas/0x90.st)

### Usage


` ./0xg0.st -h`

```
USAGE: ./0xg0.st -p=8080 -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]
  -alsologtostderr
        log to standard error as well as files
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -p uint
        port (default 8000)
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```

##### Example of run in server

`./0x0.st -p=80 -stderrthreshold=INFO -log_dir="/path/to/log"`

### LICENSE

```
Creative Commons Legal Code
CC0 1.0 Universal
```

Check LICENSE file for more information about this software license.
