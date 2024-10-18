# fAPI
Fake any API easily.

## Plan

1. Define routes and mock responses in JSON.
2. Run `fapi <input_file>`

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

- [ ] Add routing expressions to accept variables from URL
- [ ] Add ability to accept parameters from POST/PUT...
- [ ] Add ability to generate random values for responses
- [ ] Provide some useful functions for above
- [ ] Allow users to define their own functions (Maybe LUA? :3)
