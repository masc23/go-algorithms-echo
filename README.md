# Go-Algorithms-Echo
Some arbitrary algorithms implemented with [Go](https://github.com/golang/go) and the [Echo framework](https://github.com/labstack/echo).

# Algorithms

Each endpoint uses the same JSON to specify the input and output:

```json
{
    "input":  "my input string",
    "output": "the generated output"
}
```

## Base64

To convert a string to [Base64](https://en.wikipedia.org/wiki/Base64), just ``POST`` it (as "input" inside the common JSON object) to ``/base64/encode``.

To convert a string in [Base64](https://en.wikipedia.org/wiki/Base64) back to normal, just ``POST`` it (as "input" inside the common JSON object) to ``/base64/decode``.


## ROT13

To encode or decode a string with [ROT13](https://en.wikipedia.org/wiki/ROT13) just ``POST`` it (as "input" inside the common JSON object) to ``/rot13``.
