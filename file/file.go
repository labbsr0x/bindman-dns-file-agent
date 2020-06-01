package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// File struct which contains
// an array of users
type File struct {
	Domain  string    `json:"domain"`
	Records []Records `json:"records"`
}

// Records defines what we understand as a DNSRecord
type Records struct {
	// Name the DNS host name
	Name string `json:"name"`

	// Value the value of this record
	Value string `json:"value"`

	// Type the record type
	Type string `json:"type"`
}

// FileBuilder defines the parametric information of a server instance
type FileBuilder struct {
	*File
}

func (t File) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

// GetFile receve the path with file json and load
func GetFile(path string) File {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var file File
	json.Unmarshal(raw, &file)
	return file
}

// Init initializes the web server builder with properties retrieved from Viper.
func (b *FileBuilder) Init(path string) *FileBuilder {
	file := GetFile(path)

	fmt.Println(file.Domain)

	// b.File := GetFile(path)

	return b
}
