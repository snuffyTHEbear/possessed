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
// 😱 Possessed run at 2025-06-18 22:57:47.987652381 +0100 BST m=+0.000065320

// 😱 Possessed run at 2025-06-18 22:57:50.008642849 +0100 BST m=+0.000063289

	lines = append(lines[:insertionPoint], append([]string{incantation}, lines[insertionPoint:]...)...)
// 😱 Possessed run at 2025-06-18 22:57:13.043804063 +0100 BST m=+0.000041480


// 😱 Possessed run at 2025-06-18 22:56:00.654983305 +0100 BST m=+0.000043229

	// Overwrite the file
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
// 😱 Possessed run at 2025-06-18 22:57:48.997727174 +0100 BST m=+0.000063949

	}

	// Commit the mutation
	exec.Command("git", "add", filename).Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("🧛 Possessed rewrite %v", time.Now())).Run()

	fmt.Println("💀 Code rewritten. It’s learning...")
}

