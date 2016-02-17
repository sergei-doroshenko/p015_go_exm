package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

type Table struct {
	XMLName xml.Name `xml: "Table"`
	Country string   `xml:"Country"`
	City    string   `xml: City"`
}

type NewDataSet struct {
	XMLName xml.Name `xml: "NewDataSet"`
	Table   []Table  `xml: "Table"`
}

type Root struct {
	XMLName    xml.Name   `xml: "string"`
	NewDataSet NewDataSet `xml: "NewDataSet"`
}

func main() {
	fmt.Println("Hello, start process...")
	resp, err := http.Get("http://www.webservicex.net/globalweather.asmx/GetCitiesByCountry?CountryName=Ukrain")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()
	/*k := Root{}
	xml.NewDecoder(resp.Body).Decode(&k)*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	str := html.UnescapeString(string(contents))
	fmt.Println(str)
	v := Root{}

	err1 := xml.Unmarshal([]byte(str), &v)

	if err1 != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("XMLName: ", v.XMLName)
	// fmt.Printf("{country: %s, cities:[", v.NewDataSet.Table[0].Country)
	/*for _, element := range v.NewDataSet.Table {
		// element is the element from someSlice for where we are
		fmt.Print(element.City + ", ")
	}
	fmt.Print("]}\n")*/
	fmt.Println(v.NewDataSet.Table)
}
