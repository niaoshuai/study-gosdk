package xml

import (
	"encoding/xml"
	"log"
	"os"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestEncode(t *testing.T) {
	user01 := User{
		Name: "niaoshuai",
		Age:  29,
	}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(user01); err != nil {
		log.Fatal(err)
	}
}

func TestDecode(t *testing.T) {
	data := `
		<User>
		  <Name>niaoshuai</Name>
		  <Age>29</Age>
		</User>
    `
	v := User{}
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
