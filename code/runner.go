package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// go run code-user/main.go
	var bm runtime.MemStats
	runtime.ReadMemStats(&bm)
	// Alloc 已申请，且仍在使用的字节
	fmt.Printf("KB: %v\n", bm.Alloc/1024)
	now := time.Now().UnixMilli()
	println("当前时间(毫秒) ==> ", now)
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe, "23 11\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	println("Err:", string(stderr.Bytes()))
	fmt.Println(out.String())

	println(out.String() == "34\n")
	var em runtime.MemStats
	runtime.ReadMemStats(&em)
	fmt.Printf("KB: %v\n", em.Alloc/1024)
	end := time.Now().UnixMilli()
	println("当前时间 ==> ", end)
	println("耗时 ==> ", end-now)
}
