package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

const filename = "possessed.go"

func main() {
	fmt.Println("👻 Running possessed function...")

	// Possessed logic
	rand.Seed(time.Now().UnixNano())
	incantation := fmt.Sprintf("// 😱 Possessed run at %v\n", time.Now())

	// Read the current file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Inject chaos into the source code
	lines := strings.Split(string(content), "\n")
	insertionPoint := rand.Intn(len(lines))
	lines = append(lines[:insertionPoint], append([]string{incantation}, lines[insertionPoint:]...)...)

	// Overwrite the file
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}

	// Commit the mutation
	exec.Command("git", "add", filename).Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("🧛 Possessed rewrite %v", time.Now())).Run()

	fmt.Println("💀 Code rewritten. It’s learning...")
}

