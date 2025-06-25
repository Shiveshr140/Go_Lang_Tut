package main

import (
	"encoding/json"
	"fmt"
)

// Goal is encoding of json i.e converting any kind of data like array, slice to valid json format

// why small letter c because we are not exporting the type
type course struct {
	Name  string
	Price int
	Tags  []string
}

func main() {
	jsonEncodedWithIndentation()
	jsonEncodedWithModification()
	fmt.Println("=====================================")
	DecodeJson()
	fmt.Println("=====================================")
	DecodeJsonWithMap()
}

func jsonEncodedWithIndentation() {
	kodnest := []course{
		{"JavaScript", 25000, []string{`Full Stack`, `Frontend Development`}},
		{"Java", 45000, []string{`Full Stack`, `Web Development`}},
		{"Python", 35000, []string{`Full Stack`, `ML Al`}},
	}

	data, err := json.MarshalIndent(kodnest, "", "\t")

	if err != nil {
		panic(err)
	}

	// when you see []byte and you use string(data), it simply interprets the []byte as a string so that you can see the actual JSON text. It’s already JSON, and you’re just converting it to a string for display.
	fmt.Println("json data with indentation", string(data))
}

//// Most of the time user will be using it as pay load so instead of Name most of the you will see coursename

type product struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Tags     []string `json:"tags,omitempty"` // omit if empty
	Password string   `json:"-"`              // omit password so who ever consuming it won't see this field
}

func jsonEncodedWithModification() {
	kodnest := []product{
		{"JavaScript", 25000, []string{`Full Stack`, `Frontend Development`}, "abc"},
		{"Java", 45000, []string{`Full Stack`, `Web Development`}, "abc"},
		{"Python", 35000, nil, "abc"},
	}

	data, err := json.MarshalIndent(kodnest, "", "\t")

	if err != nil {
		panic(err)
	}

	// when you see []byte and you use string(data), it simply interprets the []byte as a string so that you can see the actual JSON text. It’s already JSON, and you’re just converting it to a string for display.
	fmt.Println("json data with indentation", string(data))
}

//// Lets consume json that may be comming from another backend service or from a file
//// This is a type conversion in Go.
//// The backticks `...` define a raw string literal (multiline string without escape sequences).
//// []byte(...) converts that string literal into a byte slice ([]byte), which is often required when dealing with JSON or binary data.
//// second arument must be pointer where you want to store

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "JavaScript",
                "price": 25000,
                "tags": [
                        "Full Stack",
                        "Frontend Development"
                ]
        }
	`)

	var kodnest2 product

	isValid := json.Valid(jsonDataFromWeb)

	if !isValid {
		panic("Invalid JSON data")
	}

	// Unmarshal is used to decode JSON data into a Go data structure.
	// The first argument is the JSON data (as a byte slice), and the second argument is a pointer to the variable where you want to store the decoded data.
	// The & operator is used to pass the address of kodnest, allowing Unmarshal to modify its contents.
	json.Unmarshal(jsonDataFromWeb, &kodnest2)

	// for printing interface we need to use fmt.Printf with %v or %#v
	// %v prints the value in a default format, while %#v prints the value in Go syntax.
	fmt.Printf("%#v \n", kodnest2)

}

// Why does printing show main.product? Because Go shows the package name + type name to uniquely identify the type of the value being printed.
// This is helpful especially when working with multiple packages to avoid confusion between types with the same name.
// // In this case, main.product indicates that the product type is defined in the main package as it is define outside of a function

// map[KeyType]ValueType
// map is a built-in data type in Go that represents a collection of key-value pairs.
// If you do not know
func DecodeJsonWithMap() {
	var myOnlineData map[string]interface{}

	jsonDataFromWeb := []byte(`
	 {
                "coursename": "JavaScript",
                "price": 25000,
                "tags": [
                        "Full Stack",
                        "Frontend Development"
                ]
        }
	`)

	isValid := json.Valid(jsonDataFromWeb)

	if !isValid {
		panic("Invalid JSON data")
	}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("Decoded JSON with map %#v: \n", myOnlineData)
	// fmt.Println("Decoded JSON with map:", myOnlineData[])

	for key, value := range myOnlineData{
		fmt.Printf("Key: %s, Value: %v and type: %T \n", key, value, value)
	}
}
