Go vs PIL
=========

This is a performance comparison between [PIL][1] and Go's [image package][2]. See related discussion at [go-nuts][3] mailing list.


Run the Python script with:

    python to_grayscale.py


Run the Go program with:

    go run to_grayscale.go

## Sample timings ##

Here are sample timings from my Mac Mini Core i5 2.5 GHz. These are timings from one run, not averages, but they're representative enough.

    +---------------------------------------------+
    |                   | Python, ms | Go, ms     |
    +===================+============+============+
    | Image decode      | 3.12781    | 15.611     |
    | RGBA -> Grayscale | 0.876188   | 2.701      |
    | Image save        | 6.55293    | 33.886     |
    +-------------------+------------+------------+
    | Total             | 10.5569    | 57.498     |
    +---------------------------------------------+

    $ go version
    go version go1.0.2

    $ python --version
    Python 2.7.3

    PIL version 1.1.7


  [1]: http://www.pythonware.com/products/pil/
  [2]: http://golang.org/pkg/image/
  [3]: https://groups.google.com/group/golang-nuts/browse_thread/thread/47143dea57243d0e/5686f7aadae4fa06#5686f7aadae4fa06
