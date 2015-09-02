# imageDiff
A project for detecting birds eating stealing tomatoes from my veg garden. It compares two same size JPEG images and filters the new through the old, removing shades of green and outputs the differences to output.jpeg. The images are input as arguments into the program. I've supplied three sets of example images

## How to use
Install Golang on your system from golang.org, then go into the imageDiff directory:

    $ go run main.go [image.jpg] [olderImage.jpg]
    i.e. $ go run main.go test2.jpg testOld2.jpg
OR

    $ go build .
    $ ./imageDiff [image.jpg] [olderImage.jpg]
