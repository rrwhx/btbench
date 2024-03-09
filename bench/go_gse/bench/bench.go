package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-ego/gse"
)

var (
	seg gse.Segmenter

	b, _ = os.ReadFile("./testdata/zh/bailuyuan_zh.txt")
	text = strings.Repeat(string(b), 2)
)

func main() {
	flag.Parse()

	// Loading the default dictionary
	// seg.LoadDict()
	// Loading the default dictionary with embed
	// seg.LoadDictEmbed()
	//
	// Loading the simple chinese dictionary
	// seg.LoadDict ("zh_s")
	seg.LoadDictEmbed("zh_s")
	//
	// Loading the traditional chinese dictionary
	// seg.LoadDict("zh_t")
	//
	// Loading the japanese dictionary
	// seg.LoadDict("jp")
	//
	// seg.LoadDict("../data/dict/dictionary.txt")
	//
	// Loading the custom dictionary
	// seg.LoadDict("zh,../../testdata/zh/test_dict.txt,../../testdata/zh/test_dict1.txt")

	cut()
	cut()
	cut()
}

// 使用 DAG 或 HMM 模式分词
func cut() {
	// "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片."

	// use DAG and HMM
	hmm := seg.Cut(text, true)
	fmt.Println("cut use hmm: ", len(hmm))
	// cut use hmm:  [《复仇者联盟3：无限战争》 是 全片 使用 imax 摄影机 拍摄 制作 的 的 科幻片 .]

	cut := seg.Cut(text)
	fmt.Println("cut: ", len(cut))
	// // cut:  [《 复仇者 联盟 3 ： 无限 战争 》 是 全片 使用 imax 摄影机 拍摄 制作 的 的 科幻片 .]

	hmm = seg.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", len(hmm))
	// //cut search use hmm:  [复仇 仇者 联盟 无限 战争 复仇者 《复仇者联盟3：无限战争》 是 全片 使用 imax 摄影 摄影机 拍摄 制作 的 的 科幻 科幻片 .]
	// // fmt.Println("analyze: ", seg.Analyze(hmm, text))

	cut = seg.CutSearch(text)
	fmt.Println("cut search: ", len(cut))
	// // cut search:  [《 复仇 者 复仇者 联盟 3 ： 无限 战争 》 是 全片 使用 imax 摄影 机 摄影机 拍摄 制作 的 的 科幻 片 科幻片 .]

	cut = seg.CutAll(text)
	fmt.Println("cut all: ", len(cut))
	// cut all:  [《复仇者联盟3：无限战争》 复仇 复仇者 仇者 联盟 3 ： 无限 战争 》 是 全片 使用 i m a x 摄影 摄影机 拍摄 摄制 制作 的 的 科幻 科幻片 .]

	// s := seg.CutStr(cut, ", ")
	// fmt.Println("cut all to string: ", s)
	// cut all to string:  《复仇者联盟3：无限战争》, 复仇, 复仇者, 仇者, 联盟, 3, ：, 无限, 战争, 》, 是, 全片, 使用, i, m, a, x, 摄影, 摄影机, 拍摄, 摄制, 制作, 的, 的, 科幻, 科幻片, .
}
