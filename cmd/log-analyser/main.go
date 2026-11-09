package main 
import (
	"fmt"
	"os"
	
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile() {
	fmt.Println("Writing")
	file, err := os.Create("example.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("testing my writing skills")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s", file.Name())
	fmt.Println("File length: ", length)
}

func ReadFile() {
	fmt.Println("reading the file, hold on bro")
	fileName := "example.txt"
	data, err := os.ReadFile(fileName)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("the file being read: " + fileName)
	fmt.Printf("whats inside of the file you're reading: %s", data)
}

func main() {
	
	ReadFile()
}

