package main

import (
	"crypto/sha256"
	"fmt"
	"image/jpeg"
	"os"
 	"github.com/corona10/goimagehash"
)

func main() {

	fmt.Printf("hello world \n")

	file, err := os.Open("hugo.jpg")
	if err != nil {
		fmt.Println(err)
	}
	img, _ := jpeg.Decode(file)

	compressed, err := os.Create("compressed.jpg")
	options := jpeg.Options{1}
	err = jpeg.Encode(compressed, img, &options)
	if err != nil {
		fmt.Println(err)
	}

	open_compressed, _ := os.Open("compressed.jpg")
	img_compressed , _ := jpeg.Decode(open_compressed)





	hash, _ := goimagehash.AverageHash(img)
	hash_c, _ := goimagehash.AverageHash(img_compressed)


	hash2, _ := goimagehash.PerceptionHash(img)
	hash2_c, _ := goimagehash.PerceptionHash(img_compressed)

	hash_hello := sha256.Sum256([]byte("hello"))
	fmt.Println("hash hello", hash_hello)

	fmt.Println("Average hashes \n")
	fmt.Println(hash.ToString())
	fmt.Println(hash_c.ToString())

	fmt.Println("Perceptual hashes \n")
	fmt.Println(hash2.ToString())
	fmt.Println(hash2_c.ToString())




}
