package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	for {
		var d Data
		d.File = name
		var der string
		n, e := fmt.Fscanln(f, d.StorePath, d.Hash, der)
		if e != nil || n != 3 {
			break
		}
		d.Prebuilt = (len(der) > 0)
		gdata = append(gdata, d)
	}
}

func CheckHashes() {
	for i := 0; i < len(gdata); i += 1 {
		path := gdata[i].StorePath
		hash := gdata[i].Hash
		for j := i + 1; j < len(gdata) && gdata[j].StorePath != path; j += 1 {
			if gdata[j].Hash != hash {
				fmt.Println("Discrepancy for ", path, " between ",
					gdata[i].File, " and ", gdata[j].File)
			}
		}
	}
}

func main() {
	for _, a := range os.Args[1:] {
		load(a)
	}
	gdata.Sort()
	CheckHashes()
}
