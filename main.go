package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const filePerm = 0666

func main() {
	// Function to generate SHA-256 hash
	generateHash := func(input string) string {
		hash := sha256.New()
		hash.Write([]byte(input))
		hashBytes := hash.Sum(nil)
		return hex.EncodeToString(hashBytes)
	}

	// check for arguments
	if len(os.Args) < 2 {
		if runtime.GOOS == "windows" {
			fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [FILE | \\PATH\\TO\\FILE]")
			os.Exit(0)
		} else {
			fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [FILE | /PATH/TO/FILE]")
			os.Exit(0)
		}

	}
	// take filepath as an argument
	filePath := os.Args[1]
	filename := filepath.Base(filePath)

	// reading file
	file, err := os.ReadFile(filePath)

	// error handling
	if os.IsNotExist(err) {
		fmt.Println("Error reading file: File does not exist!")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error reading file: ", err)
		fmt.Println("This is unhandled error, so if you see this, report error to the developer")
		os.Exit(1)
	}

	// generating provided file's SHA-256 hash
	hash := generateHash(string(file))

	// save file if --save or -s argument provided
	if len(os.Args) >= 3 && (os.Args[2] == "--save" || os.Args[2] == "-s") {
		sha256file, err := os.OpenFile(filename+".sha256", os.O_CREATE|os.O_WRONLY, filePerm)
		if err != nil {
			panic(err)
		}

		sha256file.WriteString(hash)
		defer sha256file.Close()
		fmt.Printf("SHA-256 hash of %s has been written to %s.sha256\n", filename, filename)
		os.Exit(0)
	}

	fmt.Println(filename, "SHA-256 hash:", hash)
}
