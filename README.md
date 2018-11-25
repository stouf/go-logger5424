# Go Logger 5424

Simple logger that lets you write logs with contextual information and
severities such as defined by the
[RFC-5424](https://tools.ietf.org/html/rfc5424). The current time is also
automatically added to the generated log for you.

## How to use

```golang
import (
  logger5424 "github.com/stouf/go-logger5424"
)

wc := makeWriter()
myLogger := logger5424.New(wc)
info := map[string]string{"foo": "bar"}
err := myLogger.Json(
  logger5424.Info,
  "Service started",
  &info,
)
if err != nil {
  panic(err)
}
```

The available severities are the following:

- `Emergency`
- `Alert`
- `Critical`
- `Error`
- `Warning`
- `Notice`
- `Info`
- `Debug`
