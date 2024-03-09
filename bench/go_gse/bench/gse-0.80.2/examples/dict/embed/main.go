//go:build go1.16
// +build go1.16

package main

import (
	_ "embed"
	"fmt"

	"github.com/go-ego/gse"
)

//go:embed test_dict3.txt
var testDict string

var (
	text  = "他在命运的沉浮中随波逐流, 扮演着受害与加害者的双重角色"
	text1 = "沙漠的那边还是沙漠, hello world"
	seg   gse.Segmenter
)

func main() {
	var err error
	seg, err = gse.NewEmbed("zh, word 20 n"+testDict, "en")
	// err := seg.LoadDictEmbed()
	// seg.LoadDictStr(testDict)
	fmt.Println("gse NewEmbed error: ", err)

	freq, pos, ok := seg.Find("hello")
	fmt.Println(freq, pos, ok)
	freq, pos, ok = seg.Find("world")
	fmt.Println(freq, pos, ok)

	freq, pos, ok = seg.Find("1号店")
	fmt.Println(freq, pos, ok)

	s := seg.Cut(text, true)
	fmt.Println("cut: ", s)

	s = seg.Cut(text1)
	fmt.Println("cut: ", s, len(s))
}

func load1() {
	err := seg.LoadDictEmbed()
	fmt.Println(err)
	err = seg.LoadDictStr(testDict)
	fmt.Println(err)

	err = seg.LoadStopEmbed()
	fmt.Println(err)
}
