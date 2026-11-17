package main 
import (
	"fmt"
	"os"
	"bufio"
	
)

type log_entry struct {
	timestamp string
	GPU_id string
	XID_code int
	message string
}



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
	fmt.Println("reading the XID logs, hold on bro")
	fileName := "example.txt"
	data, err := os.Open(fileName)
	
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("the file being read: " + fileName)
}

func main() {
	
	ReadFile()
}

