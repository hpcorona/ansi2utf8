package main

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"os"
	"strings"
	"flag"
	)
	
func OnlyParseXML(input string) {
	idx := strings.LastIndex(input, ".")
	if idx < 0 {
		return
	}
	
	ext := input[idx+1:]
	
	if (ext == "xml") {
		RevisarUTF8(input)
	}
}

func main() {
	flag.Parse()
	
	if flag.NArg() != 0 {
		for i := 0; i < flag.NArg(); i++ {
			input := flag.Arg(i)
			OnlyParseXML(input)
		}
		
		return
	}
	
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("%v: %s\n", err, "directorio actual")
		os.Exit(1)
	}
	
	for i := 0; i < len(files); i++ {
		input := files[i].Name
		OnlyParseXML(input)
	}	
}

func RevisarUTF8(file string) {
	fmt.Printf("Verificando\t%s... ", file)
	cont, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return
	}
	
	mcount := 0
	if len(cont) >= 3 {
		if cont[0] == 0xEF {
			mcount++
		}

		if cont[1] == 0xBB {
			mcount++
		}

		if cont[2] == 0xBF {
			mcount++
		}
	}
	
	if mcount == 3 {
		fmt.Printf("correcto\n")
		return
	}
	
	outbuff := bytes.NewBuffer([]byte(""))
	outbuff.WriteByte(0xEF)
	outbuff.WriteByte(0xBB)
	outbuff.WriteByte(0xBF)
	outbuff.Write(cont)
	
	ioutil.WriteFile(file, outbuff.Bytes(), 444)
	
	fmt.Printf("CONVERTIDO\n")
}
