package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	path := os.Args[1]
	if path == "" {
		fmt.Println("usage: store-sizer <desired store path>")
		os.Exit(1)
	}

	var fsStats syscall.Statfs_t
	if err := syscall.Statfs(path, &fsStats); err != nil {
		// TODO
		panic(err)
	}

	totalSize := fsStats.Bsize * int64(fsStats.Blocks)

	fmt.Println(totalSize)
}
