package main

import (
	"fmt"
	"os"
//	"path/filepath"
	"log"
	"strings"
)

func main (){
	envPath := "../.env"
	content, err := os.ReadFile(envPath)
	if (err != nil){
		log.Fatal(err)
		return 
	}
	envContent := strings.Split(string(content), "\n")
	envMap := make(map[string]string)
	for i := range(len(envContent)) - 1  {
		if strings.Contains(envContent[i], "=") {
			keyValue := strings.Split(envContent[i], "=")
			envMap[keyValue[0]] = keyValue[1]
		}
	}
	fmt.Println(envMap)
}
