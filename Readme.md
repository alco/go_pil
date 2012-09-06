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
    Image decode took 15.508 ms
    Creating an empty image took 0.027 ms
    Drawing the image took 43.289 ms
    Image save took 31.448 ms
    ---
    Total time: 90.27199999999999 ms

    Image decode took 17.162 ms
    Creating an empty image took 0.049 ms
    Converting pixels took 0.609 ms
    Image save took 35.479 ms
    ---
    Total time: 53.299 ms

    $ python --version
    Python 2.7.3

    $ python to_grayscale.py
    Image open took 2.65503 ms
    Converting the image took 3.75509 ms
    Image save took 6.58894 ms
    ---
    Total time: 12.9991 ms

  [1]: http://www.pythonware.com/products/pil/
  [2]: http://golang.org/pkg/image/
