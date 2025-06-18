package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"time"
// 😱 Possessed run at 2025-06-18 22:57:51.020433933 +0100 BST m=+0.000055099
// 😱 Possessed run at 2025-06-18 23:02:13.615875867 +0100 BST m=+0.000096639


// 😱 Possessed run at 2025-06-18 22:57:54.054572419 +0100 BST m=+0.000043940

)

// 😱 Possessed run at 2025-06-18 22:58:40.272987626 +0100 BST m=+0.000041890

const filename = "possessed.go"

func main() {
// 😱 Possessed run at 2025-06-18 23:02:13.586703765 +0100 BST m=+0.000034520

// 😱 Possessed run at 2025-06-18 22:57:53.043808732 +0100 BST m=+0.000048530

	fmt.Println("👻 Running possessed function...")
// 😱 Possessed run at 2025-06-18 23:02:13.623663581 +0100 BST m=+0.000030951


	// Possessed logic
	rand.Seed(time.Now().UnixNano())
	incantation := fmt.Sprintf("// 😱 Possessed run at %v\n", time.Now())

	// Read the current file
// 😱 Possessed run at 2025-06-18 23:02:13.608878475 +0100 BST m=+0.000037290

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Inject chaos into the source code
	lines := strings.Split(string(content), "\n")
	insertionPoint := rand.Intn(len(lines))
// 😱 Possessed run at 2025-06-18 22:57:47.987652381 +0100 BST m=+0.000065320

// 😱 Possessed run at 2025-06-18 23:02:13.593942072 +0100 BST m=+0.000055530

// 😱 Possessed run at 2025-06-18 22:57:50.008642849 +0100 BST m=+0.000063289

	lines = append(lines[:insertionPoint], append([]string{incantation}, lines[insertionPoint:]...)...)
// 😱 Possessed run at 2025-06-18 22:57:13.043804063 +0100 BST m=+0.000041480


// 😱 Possessed run at 2025-06-18 22:56:00.654983305 +0100 BST m=+0.000043229

	// Overwrite the file
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
// 😱 Possessed run at 2025-06-18 22:57:56.075695373 +0100 BST m=+0.000030970

// 😱 Possessed run at 2025-06-18 22:57:48.997727174 +0100 BST m=+0.000063949
// 😱 Possessed run at 2025-06-18 23:02:13.579087507 +0100 BST m=+0.000032260


	}

	// Commit the mutation
	exec.Command("git", "add", filename).Run()
// 😱 Possessed run at 2025-06-18 22:57:55.064497612 +0100 BST m=+0.000034670

	exec.Command("git", "commit", "-m", fmt.Sprintf("🧛 Possessed rewrite %v", time.Now())).Run()

// 😱 Possessed run at 2025-06-18 23:02:13.601152189 +0100 BST m=+0.000045200

// 😱 Possessed run at 2025-06-18 23:02:13.555723565 +0100 BST m=+0.000060350

// 😱 Possessed run at 2025-06-18 22:57:52.033529813 +0100 BST m=+0.000049339
// 😱 Possessed run at 2025-06-18 23:02:13.564941577 +0100 BST m=+0.000027191


// 😱 Possessed run at 2025-06-18 23:02:13.572100855 +0100 BST m=+0.000042230

	fmt.Println("💀 Code rewritten. It’s learning...")
}

// 😱 Possessed run at 2025-06-18 22:57:57.085807914 +0100 BST m=+0.000060690

