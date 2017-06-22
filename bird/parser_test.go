package bird

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func readSampleData(filename string) ([]byte, error) {
	sample := "../test/" + filename
	return ioutil.ReadFile(sample)
}

func TestParseBgpRoutes(t *testing.T) {

	inputs := []string{
		"1 imported, 0 filtered, 2 exported, 1 preferred",
		"0 imported, 2846 exported", // Bird 1.4.x
	}

	expected := []Parsed{
		Parsed{
			"imported":  int64(1),
			"filtered":  int64(0),
			"exported":  int64(2),
			"preferred": int64(1),
		},
		Parsed{
			"imported": int64(0),
			"exported": int64(2846),
		},
	}

	for i, in := range inputs {
		routes := parseBgpRoutes(in)
		if !reflect.DeepEqual(routes, expected[i]) {
			t.Error("Parse bgpRoutes:", routes, "expected:", expected[i])
		}
	}

}

func TestParseRoutesAll(t *testing.T) {
	sample, err := readSampleData("routes_all.sample")
	if err != nil {
		t.Error(err)
	}

	// Parse routes
	result := parseRoutes(sample)
	routes, ok := result["routes"].([]Parsed)
	if !ok {
		t.Error("Error getting routes")
	}

	if len(routes) != 4 {
		t.Error("Expected number of routes to be 3")
	}

	expectedNetworks := []string{"16.0.0.0/24", "200.0.0.0/24", "200.0.0.0/24", "16.0.0.0/24"}
	for i, route := range routes {
		net := route["network"].(string)
		if net != expectedNetworks[i] {
			t.Error("Expected network to be:", expectedNetworks[i], "not", net)
		}
	}
}
