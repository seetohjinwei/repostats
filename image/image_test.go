package image

import (
	"log"
	"reflect"
	"testing"

	"github.com/seetohjinwei/repostats/models"
)

func TestCreateAllValues(t *testing.T) {
	tests := []struct {
		name     string
		typeData map[string]models.TypeData
		want     []item
	}{
		{
			name:     "empty",
			typeData: map[string]models.TypeData{},
			want:     []item{},
		},
		{
			name: "6",
			typeData: map[string]models.TypeData{
				"a": {Type: "a", FileCount: 100, Bytes: 1},
				"b": {Type: "b", FileCount: 10, Bytes: 1},
				"c": {Type: "c", FileCount: 1, Bytes: 1},
				"d": {Type: "d", FileCount: 1, Bytes: 1},
				"e": {Type: "e", FileCount: 1, Bytes: 1},
				"f": {Type: "f", FileCount: 1, Bytes: 1},
			},
			want: []item{{name: "a", ratio: 0.8771929824561403, colour: "#F47A1F"}, {name: "b", ratio: 0.08771929824561403, colour: "#FDBB2F"}, {name: "f", ratio: 0.008771929824561403, colour: "#377B2B"}, {name: "e", ratio: 0.008771929824561403, colour: "#7AC142"}, {name: "d", ratio: 0.008771929824561403, colour: "#007CC3"}, {name: "c", ratio: 0.008771929824561403, colour: "#00529B"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := createAllValues(test.typeData)
			assertEquals(t, got, test.want)
		})
	}
}

func TestCreateValues(t *testing.T) {
	tests := []struct {
		name     string
		typeData map[string]models.TypeData
		want     []item
	}{
		{
			name:     "empty",
			typeData: map[string]models.TypeData{},
			want:     []item{},
		},
		{
			name: "7",
			typeData: map[string]models.TypeData{
				"a": {Type: "a", FileCount: 2, Bytes: 1},
				"b": {Type: "b", FileCount: 2, Bytes: 1},
				"c": {Type: "c", FileCount: 2, Bytes: 1},
				"d": {Type: "d", FileCount: 2, Bytes: 1},
				"e": {Type: "e", FileCount: 2, Bytes: 1},
				"f": {Type: "f", FileCount: 1, Bytes: 1},
				"g": {Type: "g", FileCount: 1, Bytes: 1},
			},
			want: []item{{name: "e", ratio: 0.16666666666666666, colour: "#F47A1F"}, {name: "d", ratio: 0.16666666666666666, colour: "#FDBB2F"}, {name: "c", ratio: 0.16666666666666666, colour: "#377B2B"}, {name: "b", ratio: 0.16666666666666666, colour: "#7AC142"}, {name: "a", ratio: 0.16666666666666666, colour: "#007CC3"}, {name: "others", ratio: 0.16666666666666682, colour: "#00529B"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := createValues(test.typeData)
			assertEquals(t, got, test.want)
		})
	}
}

func assertEquals(t testing.TB, got, want []item) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		log.Fatalf("got %#v want %#v", got, want)
	}
}
