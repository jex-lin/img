# img

Draw a rectangle on an image in Golang

# Example

    package main

    import (
        "github.com/jex-lin/img"
        "bytes"
        "image"
        "image/color"
        "io/ioutil"
        "os"
    )

    func main() {
        // Load bytes from image file
        src, _ := ioutil.ReadFile("/tmp/gopher.jpg")

        // Create new image file
        t2, _ := os.Create("/tmp/gopher2.jpg")
        defer t2.Close()

        // New Draw struct
        d, _ := img.NewDraw(bytes.NewReader(src))

        // Draw a rectangle with thickness of 3
        b := d.SetColor(color.RGBA{255, 0, 0, 0}).DrawRect(image.Rect(210, 10, 380, 130), 3).OutputBytes()

        // Write bytes into file
        _, _ = t2.Write(b)
    }

Source:

![](https://github.com/jex-lin/img/blob/master/gopher.jpg)

Result:

![](https://github.com/jex-lin/img/blob/master/gopher2.jpg)

