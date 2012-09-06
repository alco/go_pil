try:
    import Image
except ImportError:
    from PIL import Image

import time


bench_total = 0

def benchmark(comment, fun):
    global bench_total
    t = time.time()
    result = fun()
    delta_ms = (time.time() - t) * 1000
    bench_total += delta_ms
    print '%s took %s ms' % (comment, delta_ms)
    return result

def print_summary():
    print '---'
    print 'Total time: %s ms' % bench_total


def open_image(path):
    return benchmark("Image open", lambda: Image.open(path))

def save_image(image, path):
    return benchmark("Image save", lambda: image.save(path))

def convert(image):
    """Convert the input image from RGBA to grayscale and return a new image"""
    assert image.mode == 'RGBA'
    return benchmark("Converting the image", lambda: image.convert('L'))

if __name__ == '__main__':
    image = open_image('image.png')
    grayscale = convert(image)
    save_image(grayscale, 'py_output.png')

    print_summary()
