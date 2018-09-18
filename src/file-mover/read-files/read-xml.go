package read_files

import (
	"encoding/xml"
	"fmt"
	"os"
)

type configuration1 struct {
	Enabled bool   `xml:"enabled"`
	Path    string `xml:"path"`
}

func main() {
	xmlFile, err := os.Open("src/file-mover/conf/conf.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	var conf configuration1
	if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
		fmt.Println("Error Decode file:", err)
		return
	}

	fmt.Println(conf.Enabled)
	fmt.Println(conf.Path)

}
