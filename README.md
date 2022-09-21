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

To encode or decode a string with [ROT13](https://en.wikipedia.org/wiki/ROT13), just ``POST`` it (as "input" inside the common JSON object) to ``/rot13``.


## Equations & Regula Falsi

To evaluate a given equation for a given ``x``, just ``POST`` the equation (as "input" inside the common JSON object) to ``/math/evaluate-equation``.

The format for the equation is e.g. ``cos(x) - x`` for the equation _f(y) = cos(x) - x_.

To find a zero for a given equation inside a given interval using [regula falsi](https://en.wikipedia.org/wiki/Regula_falsi), just ``POST`` to ``/math/regula-falsi``.

The format for the input string is: ``[<lowerBound>;<upperBound>] <epsilon> <numberOfIterations> <equation>``. ``<lowerBound>``, ``<upperBound>`` and ``<epsilon>`` are floating point numbers, ``<numberOfIterations>`` is an positive integer ``> 0`` and the ``<equation>`` has the already established format.

So for evaluating the equation _f(y) = cos(x) - x_ inside the interval ``[0;1]`` with an epsilon of ``0.000000000000001`` and ``100`` iterations, use: ``[0;1] 0.000000000000001 100 cos(x) - x``.


# Addendum
This repository is for learning purposes only. The algorithms are not tested, have minimal to none error checking and generally are not suitable for production. Use at your own risk!
