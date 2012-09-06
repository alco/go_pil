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

bench_total := 0

func benchmark(comment string, fun func) interface{} {
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


func openImage(path) image.Image {
    // Open the file.
    file, err := os.Open("font_vinque.png")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Decode the image.
    return benchmark("Image decode", func() {
        img, err := png.Decode(file)
        if err != nil {
            log.Fatal(err)
            return nil
        }
        return img
    })
}

func saveImage(image, path) {
    // Create or truncate the file.
    out_file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer out_file.Close()

    // Save the image
    return benchmark("Image save", func() {
        err = png.Encode(out_file, image)
        if err != nil {
            log.Fatal(err)
            return nil
        }
    })
}

func newGray(bounds image.Rect) image.Gray {
    return benchmark("Creating an empty image", func() {
        return image.NewGray(bounds)
    })
}

func convertDraw(img image.Image, dst image.Gray) {
    return benchmark("Drawing the image", func() {
        draw.Draw(dst, dst.Bounds(), img, image.ZP, draw.Src)
    })
}

func convertLoop(img image.Image, dst image.Gray) {
    return benchmark("Converting pixels", func() {
        for i := 0; i < len(dst.Pix); i++ {
            _, _, _, a := img.At(i % dst.Stride, i / dst.Stride).RGBA()
            grayscale.Pix[i] = uint8(a / 256)
        }
    })
}


func main() {
    image := openImage("image.png")
    grayscale := newGray(image.Bounds())
    convertDraw(image, grayscale)
    convertLoop(image, grayscale)
    saveImage(grayscale, "go_output.png")

    printSummary()
}
