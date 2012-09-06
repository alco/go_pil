Go vs PIL
=========

This is a performance comparison between [PIL][1] and Go's [image package][2].


Run the Python script with:

    python to_grayscale.py


Run the Go program with:

    go run to_grayscale.go

## Sample timings ##

Here's benchmark results from my Mac Mini Core i5 2.5 GHz:

    $ go version
    go version go1.0.2

    $ go run to_grayscale.go
    Image decode took 16 ms
    Creating an empty image took 0 ms
    Drawing the image took 45 ms
    Image save took 39 ms
    ---
    Total time: 100 ms

    Image decode took 20 ms
    Creating an empty image took 0 ms
    Converting pixels took 28 ms
    Image save took 45 ms
    ---
    Total time: 93 ms

    $ python --version
    Python 2.7.3

    $ python to_grayscale.py
    Image open took 5.29599189758 ms
    Converting the image took 5.20396232605 ms
    Image save took 25.6631374359 ms
    ---
    Total time: 36.1630916595 ms

  [1]: http://www.pythonware.com/products/pil/
  [2]: http://golang.org/pkg/image/
