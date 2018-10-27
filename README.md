# `jsonassert`

`jsonassert` is a Go test assertion library for asserting JSON payloads.

## Installation

```
go get github.com/kinbiko/jsonassert
```

## Usage

Create a new `jsonassert.Asserter` in your test and use this to make assertions against your JSON payloads:

```go
func TestWhatever(t *testing.T) {
    ja := jsonassert.New(t)
    // find some sort of payload
    ja.Assert(payload, `
    {
        "name": "%s",
        "age": %d,
        "skills": [
            { "name": "martial arts", "level": 99 },
            { "name": "intelligence", "level": 100 },
            { "name": "mental fortitude", "level": 4 }
        ]
    }`, "River Tam", 16)
}
```

Notice that you can pass in `fmt.Sprintf` arguments after the expected JSON structure.

`jsonassert.Assert` supports assertions against the following payload data types:

- `string`
- `*json.RawJSON`
- `*simplejson.Json`
- `*http.Request`
- Any `struct` with `json:` tags

### Advanced usage

Some times you do not care about the value of a key, but only its presence. Take timestamps and UUIDs for example:

```json
{
    "uuid": "cb5230fc-f98f-4c63-abb7-d0588295983b",
    "timestamp": "2018-10-26T23:43:50+00:00"
}
```

The properties here may be non-deterministic and difficult to test against.
In this case you may replace the expected value with `"<PRESENCE>"` to assert its presence, but not its value:

```json
{
    "uuid": "<PRESENCE>",
    "timestamp": "<PRESENCE>"
}
```

`"<PRESENCE>"` works for any key, regardless of type.

## Docs

You can find the [GoDocs for this package here](https://godoc.org/github.com/kinbiko/jsonassert).

## Contributing and appreciation

Contributions are welcome. Please discuss feature requests in an issue before opening a PR.

If you use this project and would like to show your appreciation, I do accept donations in the form of code reviews. Feel free to have a look at the [eternal PR](https://github.com/kinbiko/jsonassert/pulls/1).