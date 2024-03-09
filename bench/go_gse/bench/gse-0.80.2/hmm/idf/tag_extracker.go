package idf

import (
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/go-ego/gse"
)

// Segment type a word with weight.
type Segment struct {
	text   string
	weight float64
}

// Text return the segment's text.
func (s Segment) Text() string {
	return s.text
}

// Weight return the segment's weight.
func (s Segment) Weight() float64 {
	return s.weight
}

// Segments type a slice of Segment.
type Segments []Segment

func (ss Segments) Len() int {
	return len(ss)
}

func (ss Segments) Less(i, j int) bool {
	if ss[i].weight == ss[j].weight {
		return ss[i].text < ss[j].text
	}

	return ss[i].weight < ss[j].weight
}

func (ss Segments) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

// TagExtracter is extract tags struct.
type TagExtracter struct {
	seg gse.Segmenter

	Idf      *Idf
	stopWord *StopWord
}

// WithGse register the gse segmenter
func (t *TagExtracter) WithGse(segs gse.Segmenter) {
	t.stopWord = NewStopWord()
	t.seg = segs
}

// LoadDict load and create a new dictionary from the file
func (t *TagExtracter) LoadDict(fileName ...string) error {
	t.stopWord = NewStopWord()
	return t.seg.LoadDict(fileName...)
}

// LoadIdf load and create a new Idf dictionary from the file.
func (t *TagExtracter) LoadIdf(fileName ...string) error {
	t.Idf = NewIdf()
	return t.Idf.LoadDict(fileName...)
}

// LoadIdfStr load and create a new Idf dictionary from the string.
func (t *TagExtracter) LoadIdfStr(str string) error {
	t.Idf = NewIdf()
	return t.Idf.seg.LoadDictStr(str)
}

// LoadStopWords load and create a new StopWord dictionary from the file.
func (t *TagExtracter) LoadStopWords(fileName ...string) error {
	t.stopWord = NewStopWord()
	return t.stopWord.LoadDict(fileName...)
}

// ExtractTags extract the topK key words from text.
func (t *TagExtracter) ExtractTags(text string, topK int) (tags Segments) {
	freqMap := make(map[string]float64)

	for _, w := range t.seg.Cut(text, true) {
		w = strings.TrimSpace(w)
		if utf8.RuneCountInString(w) < 2 {
			continue
		}
		if t.stopWord.IsStopWord(w) {
			continue
		}

		if f, ok := freqMap[w]; ok {
			freqMap[w] = f + 1.0
		} else {
			freqMap[w] = 1.0
		}
	}

	total := 0.0
	for _, freq := range freqMap {
		total += freq
	}

	for k, v := range freqMap {
		freqMap[k] = v / total
	}

	ws := make(Segments, 0)
	var s Segment
	for k, v := range freqMap {
		if freq, _, ok := t.Idf.Freq(k); ok {
			s = Segment{text: k, weight: freq * v}
		} else {
			s = Segment{text: k, weight: t.Idf.median * v}
		}
		ws = append(ws, s)
	}

	sort.Sort(sort.Reverse(ws))

	if len(ws) > topK {
		tags = ws[:topK]
		return
	}

	tags = ws
	return
}
