# fAPI
Fake any API easily.

## Plan

1. Define routes and mock responses in JSON.
2. Run `fapi <input_file>`

## Example JSON

```json
{
    'port': 8080,
    'routes': [
        {
            'url': '/hello',
            'GET': {
                'status': 200,
                'response': {
                    'type': 'plaintext',
                    'value': 'Hello, World',
                }
            },
            'POST': {
                'status': 200,
                'response': {
                    'type': 'json',
                    'value': '{ "id": "1234567890" }',
                }
            }
        },
        {
            'url': '/users/:id',
            'GET': {
                'status': 200,
                'response': {
                    'type': 'json',
                    'value': '{ "id": "${url.id}", "name": "${random_name()}", "age": "${random_int(18, 50)}" }',
                }
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
