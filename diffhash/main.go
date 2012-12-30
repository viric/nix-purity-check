package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Data struct {
	File      string
	Hash      string
	StorePath string
	Prebuilt  bool
}

type DataSlice []Data

var gdata DataSlice = make([]Data, 0)

func (d DataSlice) Len() int {
	return len(d)
}

func (d DataSlice) Less(i, j int) bool {
	return d[i].StorePath < d[j].StorePath
}

func (d DataSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DataSlice) Sort() {
	sort.Sort(d)
}

func load(name string) {
	f, e := os.Open(name)
	if e != nil {
		log.Fatal("Couldn't open ", name, ": ", e)
	}
	defer f.Close()
	fmt.Println("Loading ", name)

	bf := bufio.NewReader(f)

	for {
		var d Data
		d.File = name
		var der string

		line, e := bf.ReadString('\n')
		if e == io.EOF {
			break
		}
		line = line[:len(line)-1]
		if len(line) > 1 && line[len(line)-1] == '\r' {
			line = line[:len(line)-2]
		}
		line = strings.Replace(line, "|", " ", -1)
		n, e := fmt.Sscanf(line, "%s %s %s", &d.StorePath, &d.Hash, &der)
		if e == nil {
			d.Prebuilt = true
		} else if e == io.EOF && n == 2 {
			d.Prebuilt = false
		} else {
			log.Fatal("Failed scanf. n=", n, ", err = ", e)
		}
		gdata = append(gdata, d)
	}
}

func CheckHashes() {
	i := 0
	for i < len(gdata) {
		path := gdata[i].StorePath
		hash := gdata[i].Hash
		j := i + 1
		for ; j < len(gdata) && gdata[j].StorePath == path; j += 1 {
			if gdata[j].Hash != hash {
				fmt.Println(path, "between", gdata[i].File, "and", gdata[j].File)
			}
		}
		i = j
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Usage: diffhash [file1] [file2] ...")
	}
	for _, a := range os.Args[1:] {
		load(a)
	}
	fmt.Println("Loaded: ", len(gdata))
	gdata.Sort()
	CheckHashes()
}
