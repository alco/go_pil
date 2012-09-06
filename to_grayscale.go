package main

import (
    "fmt"
    "image"
    "image/png"
    "image/draw"
    "log"
    "os"
    "time"
)

var bench_total int64

func benchmark(comment string, fun func() interface{}) interface{} {
    t := time.Now()
    result := fun()
    delta_ms := time.Now().Sub(t).Nanoseconds() / 1000 / 1000
    bench_total += delta_ms
    fmt.Printf("%s took %d ms\n", comment, delta_ms)
    return result
}

func printSummary() {
    fmt.Println("---")
    fmt.Printf("Total time: %d ms\n", bench_total)
}


func openImage(path string) image.Image {
    // Open the file.
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Decode the image.
    return benchmark("Image decode", func() (img interface{}) {
        img, err := png.Decode(file)
        if err != nil {
            log.Fatal(err)
        }
        return
    }).(image.Image)
}

func saveImage(image image.Image, path string) {
    // Create or truncate the file.
    out_file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer out_file.Close()

    // Save the image
    benchmark("Image save", func() (_ interface{}) {
        err = png.Encode(out_file, image)
        if err != nil {
            log.Fatal(err)
        }
        return
    })
}

func newGray(bounds image.Rectangle) *image.Gray {
    return benchmark("Creating an empty image", func() interface{} {
        return image.NewGray(bounds)
    }).(*image.Gray)
}

func convertDraw(img image.Image, dst *image.Gray) {
    benchmark("Drawing the image", func() interface{} {
        draw.Draw(dst, dst.Bounds(), img, image.ZP, draw.Src)
        return nil
    })
}

func convertLoop(img image.Image, dst *image.Gray) {
    benchmark("Converting pixels", func() interface{} {
        for i := 0; i < len(dst.Pix); i++ {
            _, _, _, a := img.At(i % dst.Stride, i / dst.Stride).RGBA()
            dst.Pix[i] = uint8(a / 256)
        }
        return nil
    })
}


func main() {
    for i := 0; i < 2; i++ {
        bench_total = 0

        image := openImage("image.png")
        grayscale := newGray(image.Bounds())
        if i == 0 {
            convertDraw(image, grayscale)
        } else {
            convertLoop(image, grayscale)
        }
        saveImage(grayscale, "go_output.png")

        printSummary()
        fmt.Println()
    }
}
