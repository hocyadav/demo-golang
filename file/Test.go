package main

import (
	"fmt"
	//"github.com/pelletier/go-toml" //working
	//toml "github.com/pelletier/go-toml" //working
	toml1 "github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	//open file

	file, err := os.Open("/Users/hariomyadav/go/src/awesomeProject1/file/file1.toml")
	if err != nil {
		log.Fatalf("error occur when opening file - %s", err)
	}
	defer file.Close()
	//close file using defer

	// create a new config obj
	configObj := &Config{}

	decoder := toml1.NewDecoder(file)// create a decoder obj using file , and using decode obj decode our class obj i.e. COnfig

	if err := decoder.Decode(configObj); err != nil {
		log.Fatalf("error .. %s", err)
	}
	fmt.Printf("config %+v \n", configObj)
}
