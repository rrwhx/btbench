package gse

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/vcaesar/tt"
)

var (
	prodSeg = Segmenter{}

	testH = []byte("こんにちは世界")
)

func TestGetVer(t *testing.T) {
	fmt.Println("go version: ", runtime.Version())
	ver := GetVersion()

	tt.Expect(t, Version, ver)
	tt.Expect(t, Version, ver)
	tt.Equal(t, Version, ver)
}

func TestSplit(t *testing.T) {
	var seg1 Segmenter
	tt.Expect(t, "世/界/有/七/十/亿/人/口/",
		bytesToString(seg1.SplitTextToWords([]byte("世界有七十亿人口"))))

	tt.Expect(t, "github/ /is/ /a/ /web/-/based/ /hosting/ /service/,/ /for/ /software/ /development/ /projects/./",
		bytesToString(SplitWords([]byte(
			"GitHub is a web-based hosting service, for software development projects."))))

	tt.Expect(t, "雅/虎/yahoo/!/ /致/力/于/，/领/先/的/门/户/网/站/。/",
		bytesToString(seg1.SplitTextToWords([]byte(
			"雅虎Yahoo! 致力于，领先的门户网站。"))))

	tt.Expect(t, "こ/ん/に/ち/は/",
		bytesToString(seg1.SplitTextToWords([]byte("こんにちは"))))

	tt.Expect(t, "안/녕/하/세/요/",
		bytesToString(seg1.SplitTextToWords([]byte("안녕하세요"))))

	tt.Expect(t, "Я/ /тоже/ /рада/ /Вас/ /видеть/",
		bytesToString(seg1.SplitTextToWords([]byte("Я тоже рада Вас видеть"))))

	tt.Expect(t, "¿/cómo/ /van/ /las/ /cosas/",
		bytesToString(seg1.SplitTextToWords([]byte("¿Cómo van las cosas"))))

	tt.Expect(t, "wie/ /geht/ /es/ /ihnen/",
		bytesToString(seg1.SplitTextToWords([]byte("Wie geht es Ihnen"))))

	tt.Expect(t, "je/ /suis/ /enchanté/ /de/ /cette/ /pièce/",
		bytesToString(seg1.SplitTextToWords([]byte("Je suis enchanté de cette pièce"))))

	tt.Expect(t, "[[116 111 32 119 111 114 100 115]]", toWords("to words"))
}

func TestSegment(t *testing.T) {
	var seg Segmenter
	seg.LoadDict("testdata/zh/test_dict1.txt,testdata/zh/test_zh_dict2.txt")
	// seg.LoadDict("testdata/zh/test_dict1.txt", "testdata/zh/test_zh_dict2.txt")
	tt.Expect(t, "16", seg.Dict.NumTokens())
	// tt.Expect(t, "5", seg.Dict.NumTokens())
	segments := seg.Segment([]byte("世界有七十亿人口"))
	tt.Expect(t, "世界/ 有/p3 七十亿/ 人口/p12 ", ToString(segments, false))
	// tt.Expect(t, "世界/ 有/x 七十亿/ 人口/p12 ", ToString(segments, false))

	tt.Expect(t, "4", len(segments))
	tt.Expect(t, "0", segments[0].start)
	tt.Expect(t, "6", segments[0].end)
	tt.Expect(t, "6", segments[1].start)
	tt.Expect(t, "9", segments[1].end)

	tt.Expect(t, "9", segments[2].start)
	tt.Expect(t, "18", segments[2].end)
	tt.Expect(t, "18", segments[3].start)
	tt.Expect(t, "24", segments[3].end)
}

func TestSegmentJp(t *testing.T) {
	var seg Segmenter
	// SkipLog = true
	err := seg.LoadDict("data/dict/jp/dict.txt")
	tt.Nil(t, err)
	tt.Equal(t, 794146, len(seg.Dict.Tokens))
	tt.Equal(t, 4.784183005e+09, seg.Dict.totalFreq)

	f, pos, ok := seg.Find("自由")
	tt.Bool(t, ok)
	tt.Equal(t, "名詞", pos)
	tt.Equal(t, 3636, f)

	f, pos, ok = seg.Find("此の度")
	tt.Bool(t, ok)
	tt.Equal(t, "名詞", pos)
	tt.Equal(t, 5257, f)

	segments := seg.Segment(testH)

	tt.Expect(t, "こんにちは/感動詞 世界/名詞 ", ToString(segments, false))
	tt.Expect(t, "こん/名詞 こんにちは/感動詞 世界/名詞 ", ToString(segments, true))
	tt.Expect(t, "[こん こんにちは 世界]", ToSlice(segments, true))
	tt.Expect(t, "[こんにちは 世界]", ToSlice(segments, false))
	tt.True(t, IsJp(ToSlice(segments)[0]))

	tt.Expect(t, "2", len(segments))
	tt.Expect(t, "0", segments[0].start)
	tt.Expect(t, "15", segments[0].end)

	token := segments[0].Token()
	tt.Expect(t, "こんにちは", token.Text())
	tt.Expect(t, "5704", token.Freq())
	tt.Expect(t, "感動詞", token.Pos())

	var tokenArr []*Token
	for i := 0; i < len(segments); i++ {
		tokenArr = append(tokenArr, segments[i].Token())
	}
	tt.Expect(t, "こんにちは 世界 ", printTokens(tokenArr, 2))

	tseg := token.Segments()
	tt.Expect(t, "0", tseg[0].Start())
	tt.Expect(t, "6", tseg[0].End())
}

func TestLoadDictionary(t *testing.T) {
	var seg, seg1 Segmenter

	err := seg.LoadDict("zh")
	tt.Nil(t, err)

	err = seg1.LoadDict("jp")
	tt.Nil(t, err)
	seg1.Load = false

	err = seg1.LoadDict()
	tt.Nil(t, err)

	err = seg.LoadDict("en")
	tt.Nil(t, err)
}

func TestToken(t *testing.T) {
	ToLower = false
	defer func() { ToLower = true }()

	var seg = prodSeg
	seg.Load = false

	seg.LoadNoFreq = true
	seg.TextFreq = "2.0"
	seg.MinTokenFreq = 1.0
	seg.MoreLog = true
	seg.SkipLog = true

	tt.Expect(t, "世界/n 人口/n ", ToString(prodSeg.Segment(
		[]byte("世界人口")), false))

	dict := seg.Dictionary()
	tt.Expect(t, "16", dict.MaxTokenLen())
	tt.Equal(t, 5.3226765e+07, dict.TotalFreq())

	freq, pos, ok := dict.Find([]byte("世界"))
	tt.Equal(t, 34387, freq)
	tt.Equal(t, "n", pos)
	tt.True(t, ok)

	freq, _, ok = dict.Find([]byte("帝国大"))
	tt.Equal(t, 0, freq)
	tt.True(t, ok)

	freq, _, ok = dict.Find([]byte("帝国大厦"))
	tt.Equal(t, 3, freq)
	tt.True(t, ok)

	freq, _, ok = seg.Find("帝国大厦大")
	tt.Equal(t, 0, freq)
	tt.False(t, ok)

	val, id, err := seg.Value("帝国")
	tt.Equal(t, 96493, val)
	tt.Equal(t, 597213, id)
	tt.Nil(t, err)

	err = seg.AddToken("伦敦摘星塔", 100)
	tt.Nil(t, err)
	err = seg.AddToken("Winter is coming", 100)
	tt.Nil(t, err)
	err = seg.AddToken("Winter is coming", 200)
	tt.Nil(t, err)

	freq, _, ok = seg.Find("Winter is coming")
	tt.Equal(t, 100, freq)
	tt.True(t, ok)

	freq, _, ok = prodSeg.Find("伦敦摘星塔")
	tt.Equal(t, 100, freq)
	tt.True(t, ok)

	err = prodSeg.AddToken("西雅图中心", 100)
	tt.Nil(t, err)

	err = prodSeg.AddToken("西雅图太空针", 100)
	tt.Nil(t, err)
	freq, pos, ok = prodSeg.Find("西雅图太空针")
	tt.Equal(t, 100, freq)
	tt.Equal(t, "", pos)
	tt.True(t, ok)

	err = prodSeg.ReAddToken("西雅图太空针", 200, "n")
	tt.Nil(t, err)
	freq, pos, ok = prodSeg.Find("西雅图太空针")
	tt.Equal(t, 200, freq)
	tt.Equal(t, "n", pos)
	tt.True(t, ok)

	prodSeg.AddTokenForce("Space Needle", 100, "n")
	err = prodSeg.RemoveToken("西雅图太空针")
	tt.Nil(t, err)
	freq, _, ok = dict.Find([]byte("西雅图太空针"))
	tt.Equal(t, 0, freq)
	tt.False(t, ok)
}

func TestDictPaths(t *testing.T) {
	var seg1 Segmenter
	// seg.SkipLog = true
	paths := DictPaths("./dictDir", "zh, jp")
	tt.Expect(t, "3", len(paths))

	tt.Expect(t, "dictDir/dict/zh/t_1.txt", paths[0])
	tt.Expect(t, "dictDir/dict/zh/s_1.txt", paths[1])
	tt.Expect(t, "dictDir/dict/jp/dict.txt", paths[2])

	paths1 := DictPaths("./dictDir", "zh, jp")
	tt.Expect(t, "3", len(paths))
	tt.Equal(t, paths, paths1)

	p := strings.ReplaceAll(seg1.GetCurrentFilePath(), "/segmenter_test.go", "") +
		`/data/dict/zh/idf.txt`
	tt.Equal(t, "["+p+"]", seg1.GetIdfPath([]string{}...))

	seg1.DictPath = "testdata/zh"
	tt.Equal(t, "testdata/zh", seg1.GetCurrentFilePath())
	tt.Equal(t, "[testdata/data/dict/zh/idf.txt]", seg1.GetIdfPath([]string{}...))
}

func TestInAlphaNum(t *testing.T) {
	// var seg Segmenter
	// AlphaNum = true
	// seg.LoadDict("zh,./testdata/test_en_dict3.txt")
	//
	// AlphaNum = true
	// ToLower = true
	seg, err := New("zh,./testdata/test_en_dict3.txt", "alpha")
	tt.Nil(t, err)

	freq, _, ok := seg.Find("hello")
	tt.Equal(t, 20, freq)
	tt.True(t, ok)

	freq, _, ok = seg.Find("world")
	tt.Equal(t, 20, freq)
	tt.True(t, ok)

	text := "helloworld! 你好世界, Helloworld."
	tx := seg.Cut(text)
	tt.Equal(t, 11, len(tx))
	tt.Equal(t, "[hello world !   你好 世界 ,   hello world .]", tx)

	tx = seg.Cut(text, false)
	tt.Equal(t, 11, len(tx))
	tt.Equal(t, "[hello world !   你好 世界 ,   hello world .]", tx)

	tx = seg.Cut(text, true)
	tt.Equal(t, 9, len(tx))
	tt.Equal(t, "[hello world !  你好 世界 ,  hello world .]", tx)
}
