Go vs PIL
=========

This is a performance comparison between [PIL][1] and Go's [image package][2]. See related discussion at [go-nuts][3] mailing list.


Run the Python script with:

    python to_grayscale.py


Run the Go program with:

    go run to_grayscale.go

## Sample timings ##

Here are sample timings from my MacBook Pro Core i7 2.2 GHz. These are timings from one run, not averages, but they're representative enough.

    +---------------------------------------------+------------+
    |                   | Python, ms | Go, ms     | Go tip, ms |
    +===================+============+============+============+
    | Image decode      |       3.13 |      15.61 |       6.81 |
    | RGBA -> Grayscale |       0.88 |       2.70 |       1.54 |
    | Image save        |       6.55 |      33.89 |      24.26 |
    +-------------------+------------+------------+------------+
    | Total             |      10.56 |      57.50 |      32.65 |
    +---------------------------------------------+------------+

    $ go version
    go version go1.0.2

    go tip version
    14026:4a9c3b3e39c6

    $ python --version
    Python 2.7.3

    PIL version 1.1.7


  [1]: http://www.pythonware.com/products/pil/
  [2]: http://golang.org/pkg/image/
  [3]: https://groups.google.com/group/golang-nuts/browse_thread/thread/47143dea57243d0e/5686f7aadae4fa06#5686f7aadae4fa06
