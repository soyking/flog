flog
====

Golang logger:
- Support children loggers / sub loggers, just like [Python3 logging](https://docs.python.org/3/library/logging.html)
- Based on [logrus](https://github.com/sirupsen/logrus)

Case 1: using different level logger in different package/struct/instance, which is convenient to debug

Case 2: update logger's level dynamically, which is useful for debugging online service. For example:

```
http.HandleFunc("/debug/log/level", func(rw http.ResponseWriter, r *http.Request) {
    loggerNames, level := parseLogLevel(r)
    Setup(func(l *Logger) {
        l.SetLevel(level)
    }, loggerNames)
})
```
