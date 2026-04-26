package main

import (
	"io"
	"log"
	"os"
	"regexp"
)

func Must[T any](x T, err error) T {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return x
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

func example1() {
	r := Must(regexp.Compile("^123"))
	println(r.MatchString("te123st"))
}

func example2() {
	src := "./template.txt"
	dst := "./template_copy.txt"
	r := Must(os.Open(src))
	defer r.Close()

	w := Must(os.Create(dst))
	defer w.Close()

	Must(io.Copy(w, r))
	CheckErr(w.Close())
}

func main() {
	example1()
	example2()
}
