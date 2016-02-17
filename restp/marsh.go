package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func main() {
	type Table struct {
		// XMLName xml.Name `xml: "Table"`
		Country string `xml:"Country"`
		City    string `xml: City"`
	}

	type NewDataSet struct {
		XMLName xml.Name `xml: "NewDataSet"`
		Table   []Table  `xml: "Table"`
	}

	type Root struct {
		XMLName    xml.Name   `xml: "string"`
		NewDataSet NewDataSet `xml: "NewDataSet"`
	}

	v := NewDataSet{}
	d := `
		<string xmlns="http://www.webserviceX.NET">
			<NewDataSet> 
			    <Table> 
			        <Country>Ukraine</Country>
			        <City>Boryspil</City> 
			    </Table>
			    <Table>
			        <Country>Ukraine</Country> 
			        <City>Simferopol</City>
			    </Table>
			    <Table>
			        <Country>Ukraine</Country> 
			        <City>Kharkiv</City>
			    </Table>
			    <Table>
			        <Country>Ukraine</Country>
			        <City>Kyiv</City>
			    </Table>
			    <Table>
			        <Country>Ukraine</Country>
			        <City>L'Viv</City>
			    </Table>
			    <Table>
			        <Country>Ukraine</Country>
			        <City>Odesa</City>
			    </Table> 
			</NewDataSet>
		</string>
	`

	err := xml.Unmarshal([]byte(d), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("XMLName: ", v.XMLName)
	/*fmt.Printf("{country: %s, cities:[", v.NewDataSet.Table[0].Country)
	for _, element := range v.NewDataSet.Table {
		// element is the element from someSlice for where we are
		fmt.Print(element.City + ", ")
	}
	fmt.Print("]}\n")
	fmt.Println(v.NewDataSet.Table)*/
	str := "hello"
	b := []byte(str)
	runes := bytes.Runes(b)
	fmt.Println(runes)
	fmt.Println(string(runes))
}
