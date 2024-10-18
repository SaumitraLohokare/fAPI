# fAPI
Fake any API easily.

## Build from source

1. Run `go build -o fapi` or `go build -o fapi.exe` on windows.

## Usage

1. Define routes and mock responses in JSON.
2. Run `fapi <input_file>`

> Examples present in `examples` folder

## Example JSON

```js
{
    "port": 8080,
    "routes": [
        {
            "url": "/hello/",
            "GET": {
                "status": 200,
                "Content-Type": "text/plain",
                "response": "Hello, World"
            }
        }
    ]
}
```

## Future Plans

- [ ] Update routing to match variable URL (eg. `/users/*/`)
- [ ] Add ability to generate random values for responses
- [ ] Provide some useful functions for above
- [ ] Allow users to define their own functions (Maybe LUA? :3)
