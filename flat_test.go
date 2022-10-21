package eflat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  map[string]interface{}
	}{
		/////////////////// string
		{
			name:  "string value",
			given: `{"hello": "world"}`,
			want:  map[string]interface{}{"hello": "world"},
		},
		{
			name:  "nested string value",
			given: `{"hello":{"world":"good morning"}}`,
			want:  map[string]interface{}{"hello.world": "good morning"},
		},
		{
			name:  "double nested string value",
			given: `{"hello":{"world":{"again":"good morning"}}}`,
			want:  map[string]interface{}{"hello.world.again": "good morning"},
		},

		/////////////////// float
		{
			name:  "float",
			given: `{"hello": 1234.99}`,
			want:  map[string]interface{}{"hello": 1234.99},
		},
		{
			name:  "nested float value",
			given: `{"hello":{"world":1234.99}}`,
			want:  map[string]interface{}{"hello.world": 1234.99},
		},

		/////////////////// boolean
		{
			name:  "boolean value",
			given: `{"hello": true}`,
			want:  map[string]interface{}{"hello": true},
		},
		{
			name:  "nested boolean",
			given: `{"hello":{"world":true}}`,
			want:  map[string]interface{}{"hello.world": true},
		},

		/////////////////// nil
		{
			name:  "nil value",
			given: `{"hello": null}`,
			want:  map[string]interface{}{"hello": nil},
		},
		{
			name:  "nested nil value",
			given: `{"hello":{"world":null}}`,
			want:  map[string]interface{}{"hello.world": nil},
		},

		/////////////////// map
		{
			name:  "empty value",
			given: `{"hello":{}}`,
			want:  map[string]interface{}{"hello": map[string]interface{}{}},
		},
		{
			name:  "empty object",
			given: `{"hello":{"empty":{"nested":{}}}}`,
			want:  map[string]interface{}{"hello.empty.nested": map[string]interface{}{}},
		},

		/////////////////// slice
		{
			name:  "empty slice",
			given: `{"hello":[]}`,
			want:  map[string]interface{}{"hello": []interface{}{}},
		},
		{
			name:  "nested empty slice",
			given: `{"hello":{"world":[]}}`,
			want:  map[string]interface{}{"hello.world": []interface{}{}},
		},
		{
			name:  "nested slice",
			given: `{"hello":{"world":["one","two"]}}`,
			want: map[string]interface{}{
				"hello.world": []interface{}{"one", "two"},
			},
		},

		/////////////////// combos
		{
			name: "multiple keys",
			given: `{
				"hello": {
					"lorem": {
						"ipsum": "again",
						"dolor": "sit"
					}
				},
				"world": {
					"lorem": {
						"ipsum": "again",
						"dolor": "sit"
					}
				}
			}`,
			want: map[string]interface{}{
				"hello.lorem.ipsum": "again",
				"hello.lorem.dolor": "sit",
				"world.lorem.ipsum": "again",
				"world.lorem.dolor": "sit",
			},
		},

		/////////////////// nested slices
		{
			name: "array of strings",
			given: `{
				"hallo": {
					"lorem": ["10", "1"],
					"ipsum": {
						"dolor": ["1", "10"]
					}
				}
			}`,
			want: map[string]interface{}{
				"hallo.lorem":       []interface{}{"10", "1"},
				"hallo.ipsum.dolor": []interface{}{"1", "10"},
			},
		},
		{
			name: "array of integers",
			given: `{
				"hallo": {
					"lorem": [10, 1],
					"ipsum": {
						"dolor": [1, 10]
					}
				}
			}`,
			want: map[string]interface{}{
				"hallo.lorem":       []interface{}{float64(10), float64(1)},
				"hallo.ipsum.dolor": []interface{}{float64(1), float64(10)},
			},
		},

		/////////////////// slice combos
		{
			name: "array of numbers and strings",
			given: `{
				"hallo": {
					"lorem": [10, 1],
					"ipsum": {
						"dolor": ["1", "10"]
					}
				}
			}`,
			want: map[string]interface{}{
				"hallo.lorem":       []interface{}{float64(10), float64(1)},
				"hallo.ipsum.dolor": []interface{}{"1", "10"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var given map[string]interface{}
			err := json.Unmarshal([]byte(test.given), &given)
			if err != nil {
				t.Errorf("failed to unmarshal JSON: %v", err)
			}

			got, err := Flatten(given)
			if err != nil {
				t.Errorf("failed to flatten: %+v", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("mismatch:\ngot:  %+v\nwant: %+v", got, test.want)
			}
		})
	}
}
