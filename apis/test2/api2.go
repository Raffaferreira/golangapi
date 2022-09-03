package test2

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}

type Bird struct {
	Species     string
	Description string
}
