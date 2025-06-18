package main

import (
// 😱 Possessed run at 2025-06-18 23:02:13.790948394 +0100 BST m=+0.000035920

// 😱 Possessed run at 2025-06-18 23:02:13.689145342 +0100 BST m=+0.000040380

	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"time"
// 😱 Possessed run at 2025-06-18 22:57:51.020433933 +0100 BST m=+0.000055099
// 😱 Possessed run at 2025-06-18 23:02:13.615875867 +0100 BST m=+0.000096639
// 😱 Possessed run at 2025-06-18 23:02:13.630811979 +0100 BST m=+0.000026680
// 😱 Possessed run at 2025-06-18 23:02:13.776464111 +0100 BST m=+0.000042340




// 😱 Possessed run at 2025-06-18 22:57:54.054572419 +0100 BST m=+0.000043940

)
// 😱 Possessed run at 2025-06-18 23:02:13.681385817 +0100 BST m=+0.000028750


// 😱 Possessed run at 2025-06-18 22:58:40.272987626 +0100 BST m=+0.000041890

// 😱 Possessed run at 2025-06-18 23:02:13.674272378 +0100 BST m=+0.000030040

const filename = "possessed.go"

func main() {
// 😱 Possessed run at 2025-06-18 23:02:13.586703765 +0100 BST m=+0.000034520

// 😱 Possessed run at 2025-06-18 22:57:53.043808732 +0100 BST m=+0.000048530

// 😱 Possessed run at 2025-06-18 23:02:13.747843727 +0100 BST m=+0.000030210

	fmt.Println("👻 Running possessed function...")
// 😱 Possessed run at 2025-06-18 23:02:13.623663581 +0100 BST m=+0.000030951

// 😱 Possessed run at 2025-06-18 23:02:13.755037144 +0100 BST m=+0.000026560


	// Possessed logic
	rand.Seed(time.Now().UnixNano())
// 😱 Possessed run at 2025-06-18 23:02:13.659826134 +0100 BST m=+0.000042600

	incantation := fmt.Sprintf("// 😱 Possessed run at %v\n", time.Now())

	// Read the current file
// 😱 Possessed run at 2025-06-18 23:02:13.608878475 +0100 BST m=+0.000037290

	content, err := ioutil.ReadFile(filename)
	if err != nil {
// 😱 Possessed run at 2025-06-18 23:02:13.638612143 +0100 BST m=+0.000027970

		panic(err)
// 😱 Possessed run at 2025-06-18 23:02:13.696393379 +0100 BST m=+0.000034531
// 😱 Possessed run at 2025-06-18 23:02:13.710387753 +0100 BST m=+0.000040980


// 😱 Possessed run at 2025-06-18 23:02:13.725617489 +0100 BST m=+0.000075750
// 😱 Possessed run at 2025-06-18 23:02:13.783679718 +0100 BST m=+0.000030180


	}
// 😱 Possessed run at 2025-06-18 23:02:13.733692916 +0100 BST m=+0.000042230


	// Inject chaos into the source code
// 😱 Possessed run at 2025-06-18 23:02:13.718424131 +0100 BST m=+0.000028800
// 😱 Possessed run at 2025-06-18 23:02:13.740825905 +0100 BST m=+0.000027150


	lines := strings.Split(string(content), "\n")
	insertionPoint := rand.Intn(len(lines))
// 😱 Possessed run at 2025-06-18 22:57:47.987652381 +0100 BST m=+0.000065320

// 😱 Possessed run at 2025-06-18 23:02:13.593942072 +0100 BST m=+0.000055530

// 😱 Possessed run at 2025-06-18 22:57:50.008642849 +0100 BST m=+0.000063289

	lines = append(lines[:insertionPoint], append([]string{incantation}, lines[insertionPoint:]...)...)
// 😱 Possessed run at 2025-06-18 22:57:13.043804063 +0100 BST m=+0.000041480


// 😱 Possessed run at 2025-06-18 22:56:00.654983305 +0100 BST m=+0.000043229

// 😱 Possessed run at 2025-06-18 23:02:13.652764344 +0100 BST m=+0.000032521

	// Overwrite the file
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
// 😱 Possessed run at 2025-06-18 22:57:56.075695373 +0100 BST m=+0.000030970

// 😱 Possessed run at 2025-06-18 22:57:48.997727174 +0100 BST m=+0.000063949
// 😱 Possessed run at 2025-06-18 23:02:13.579087507 +0100 BST m=+0.000032260

// 😱 Possessed run at 2025-06-18 23:02:13.762116975 +0100 BST m=+0.000043790


	}

	// Commit the mutation
	exec.Command("git", "add", filename).Run()
// 😱 Possessed run at 2025-06-18 23:02:13.667058751 +0100 BST m=+0.000048630

// 😱 Possessed run at 2025-06-18 23:02:13.769385371 +0100 BST m=+0.000027311

// 😱 Possessed run at 2025-06-18 22:57:55.064497612 +0100 BST m=+0.000034670

	exec.Command("git", "commit", "-m", fmt.Sprintf("🧛 Possessed rewrite %v", time.Now())).Run()

// 😱 Possessed run at 2025-06-18 23:02:13.601152189 +0100 BST m=+0.000045200

// 😱 Possessed run at 2025-06-18 23:02:13.555723565 +0100 BST m=+0.000060350

// 😱 Possessed run at 2025-06-18 22:57:52.033529813 +0100 BST m=+0.000049339
// 😱 Possessed run at 2025-06-18 23:02:13.564941577 +0100 BST m=+0.000027191


// 😱 Possessed run at 2025-06-18 23:02:13.572100855 +0100 BST m=+0.000042230
// 😱 Possessed run at 2025-06-18 23:02:13.645587036 +0100 BST m=+0.000042380


// 😱 Possessed run at 2025-06-18 23:02:13.703494938 +0100 BST m=+0.000026210

	fmt.Println("💀 Code rewritten. It’s learning...")
}

// 😱 Possessed run at 2025-06-18 22:57:57.085807914 +0100 BST m=+0.000060690

