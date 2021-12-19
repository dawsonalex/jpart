package jpart

import (
	"encoding/json"
	"fmt"
	"testing"
)

func ExampleSelect() {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Bob",
			"address": map[string]interface{}{
				"house":  3,
				"street": "Some Street",
				"place":  "Some Place",
			},
		},
	}

	path := "user.address"

	res, _ := Select(Path(path), data)
	fmt.Println(res)
	// Output:
	//{
	//   "house": 3,
	//   "place": "Some Place",
	//   "street": "Some Street"
	//}
}

func TestSelect(t *testing.T) {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Bob",
			"address": map[string]interface{}{
				"house":  3,
				"street": "Some Street",
				"place":  "Some Place",
			},
		},
	}

	tests := []struct {
		Name      string
		path      Path
		shouldErr bool
		expects   interface{}
	}{
		{"user.address", Path("user.address"), false, map[string]interface{}{
			"house":  3,
			"street": "Some Street",
			"place":  "Some Place",
		}},
		{
			"not a path", Path("user.errorpath"), true, nil,
		},
		{
			"no path", Path(""), false, data,
		},
		{
			"just whitespace", Path(" "), false, data,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			res, err := Select(test.path, data)
			if test.shouldErr {
				if err == nil {
					tt.Error("expected error, got nil")
				} else {
					return
				}
			}
			if err != nil {
				tt.Error(err)
			}

			var result interface{}
			err = json.Unmarshal([]byte(res), &result)
			if err != nil {
				tt.Error(err)
			}

			// going on the string representation of the type here
			if fmt.Sprintf("%#v", result) != fmt.Sprintf("%#v", test.expects) {
				tt.Error(fmt.Sprintf("expected %#v, got %#v", test.expects, result))
			}
		})
	}
}
