package main

import (
	"fmt"

	"github.com/go-ego/gse"
)

var (
	text   = "旧金山湾金门大桥"
	new, _ = gse.New("zh,../../testdata/zh/test_dict.txt")

	seg gse.Segmenter
)

func main() {
	cut()

	// loadDict()
	loadDictEmbed()
	// loadDictMap()
	segment()
}

// loadDictEmbed supported from go1.16
func loadDictEmbed() {
	seg.LoadDictEmbed()
	seg.LoadStopEmbed()
}

func loadDict() {
	// var seg gse.Segmenter
	seg.LoadDict("zh, ../../testdata/zh/test_dict.txt, ../../testdata/zh/test_dict1.txt")
	seg.LoadStop()
}

func loadDictMap() {
	m := []map[string]string{
		{
			"text": "一城山水",
			"freq": "10",
			"pos":  "n",
		},
		{
			"text": "山河日月",
			"freq": "13",
		},
	}

	seg.LoadDictMap(m)

	a := []string{"abc", "123"}
	seg.LoadStopArr(a)
}

func cut() {
	fmt.Println("cut: ", new.Cut(text, true))
	fmt.Println("cut all: ", new.CutAll(text))
	fmt.Println("cut for search: ", new.CutSearch(text, true))
}

func segment() {
	text1 := []byte(text)
	fmt.Println(seg.String(text, true))
	// 金山/nr 旧金山/ns 湾/zg 旧金山湾/ns 金门/n 大桥/ns 金门大桥/nz

	segments := seg.Segment(text1)
	// fmt.Println(gse.ToString(segments, false))
	fmt.Println(gse.ToString(segments))
	//"旧金山湾/n 金门大桥/nz "

	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	segs := seg.ModeSegment(text1, true)
	fmt.Println(gse.ToString(segs, true))
	// "金山/nr 旧金山/ns 湾/zg 旧金山湾/ns 金门/n 大桥/ns 金门大桥/nz "
}
