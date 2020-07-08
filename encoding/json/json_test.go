package json

import (
	"encoding/json"
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

	b, err := json.Marshal(user01)
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout.Write(b)
}
func TestDecode(t *testing.T) {
	var jsonBlob = []byte(`[
	{"name":"niaoshuai","age":29},
	{"name":"erkengzi","age":23}
]`)
	var users []User
	err := json.Unmarshal(jsonBlob, &users)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}
