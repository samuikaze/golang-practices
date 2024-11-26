package main

import "fmt"

func mapping() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Printf("m[\"Bell Labs\"] = %v\n", m["Bell Labs"])

	var m2 = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Printf("m2[\"Bell Labs\"] = %v\nm2[\"Google\"] = %v", m2["Bell Labs"], m2["Google"])

	m3 := make(map[string]int)
	m3["Answer"] = 42

	fmt.Printf("m3[\"Answer\"] = %v\n", m3["Answer"])

	m3["Answer"] = 48
	fmt.Printf("m3[\"Answer\"] = %v\n", m3["Answer"])

	delete(m3, "Answer")
	fmt.Printf("m3[\"Answer\"] = %v\n", m3["Answer"])

	value, ok := m3["Answer"]
	fmt.Printf("m3[\"Answer\"] = %v, exist? %v\n", value, ok)
}
