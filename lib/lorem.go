package lib

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var COMMON_P = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
	"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim " +
	"veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea " +
	"commodo consequat. Duis aute irure dolor in reprehenderit in voluptate " +
	"velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint " +
	"occaecat cupidatat non proident, sunt in culpa qui officia deserunt " +
	"mollit anim id est laborum."

var WORDS = [...]string{
	"exercitationem", "perferendis", "perspiciatis", "laborum", "eveniet",
	"sunt", "iure", "nam", "nobis", "eum", "cum", "officiis", "excepturi",
	"odio", "consectetur", "quasi", "aut", "quisquam", "vel", "eligendi",
	"itaque", "non", "odit", "tempore", "quaerat", "dignissimos",
	"facilis", "neque", "nihil", "expedita", "vitae", "vero", "ipsum",
	"nisi", "animi", "cumque", "pariatur", "velit", "modi", "natus",
	"iusto", "eaque", "sequi", "illo", "sed", "ex", "et", "voluptatibus",
	"tempora", "veritatis", "ratione", "assumenda", "incidunt", "nostrum",
	"placeat", "aliquid", "fuga", "provident", "praesentium", "rem",
	"necessitatibus", "suscipit", "adipisci", "quidem", "possimus",
	"voluptas", "debitis", "sint", "accusantium", "unde", "sapiente",
	"voluptate", "qui", "aspernatur", "laudantium", "soluta", "amet",
	"quo", "aliquam", "saepe", "culpa", "libero", "ipsa", "dicta",
	"reiciendis", "nesciunt", "doloribus", "autem", "impedit", "minima",
	"maiores", "repudiandae", "ipsam", "obcaecati", "ullam", "enim",
	"totam", "delectus", "ducimus", "quis", "voluptates", "dolores",
	"molestiae", "harum", "dolorem", "quia", "voluptatem", "molestias",
	"magni", "distinctio", "omnis", "illum", "dolorum", "voluptatum", "ea",
	"quas", "quam", "corporis", "quae", "blanditiis", "atque", "deserunt",
	"laboriosam", "earum", "consequuntur", "hic", "cupiditate",
	"quibusdam", "accusamus", "ut", "rerum", "error", "minus", "eius",
	"ab", "ad", "nemo", "fugit", "officia", "at", "in", "id", "quos",
	"reprehenderit", "numquam", "iste", "fugiat", "sit", "inventore",
	"beatae", "repellendus", "magnam", "recusandae", "quod", "explicabo",
	"doloremque", "aperiam", "consequatur", "asperiores", "commodi",
	"optio", "dolor", "labore", "temporibus", "repellat", "veniam",
	"architecto", "est", "esse", "mollitia", "nulla", "a", "similique",
	"eos", "alias", "dolore", "tenetur", "deleniti", "porro", "facere",
	"maxime", "corrupti",
}
var numWords int

var COMMON_WORDS = [...]string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur",
	"adipisicing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt",
	"ut", "labore", "et", "dolore", "magna", "aliqua",
}

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
	numWords = len(WORDS)
}

func sampleWords(n int) []string {
	perm := r.Perm(numWords)
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = WORDS[perm[i]]
	}
	return ret
}

func firstToUpper(s string) string {
	var buf bytes.Buffer

	first := true
	for _, r := range s {
		if first {
			buf.WriteString(string(unicode.ToUpper(r)))
			first = false
		} else {
			buf.WriteString(string(r))
		}
	}
	return buf.String()
}

func Sentence() string {
	sections := make([]string, r.Intn(4)+1)
	for i := 0; i < len(sections); i++ {
		sections[i] = strings.Join(sampleWords(r.Intn(9)+3), " ")
	}
	tmp := firstToUpper(strings.Join(sections, ", "))
	p := "."
	if (r.Intn(99) % 2) == 1 { // I can't get rand.Intn(1) to work ...
		p = "?"
	}
	return fmt.Sprintf("%s%s", tmp, p)
}

func Paragraph() string {
	ss := make([]string, r.Intn(4)+1)
	for i := 0; i < len(ss); i++ {
		ss[i] = Sentence()
	}
	return strings.Join(ss, " ")
}

func Paragraphs(count int, common bool) []string {
	ps := make([]string, count)
	for i := 0; i < count; i++ {
		if common && i == 0 {
			ps[i] = COMMON_P
		} else {
			ps[i] = Paragraph()
		}
	}
	return ps
}
func ParagraphsC(count int) []string {
	return Paragraphs(count, true)
}

func Words(count int, common bool) []string {
	var ws []string
	if common {
		ws = COMMON_WORDS[:]
	}

	permIdxs := rand.Perm(len(WORDS))
	len_ws := len(ws)
	if count > len_ws {
		for i := 0; i < count-len_ws; i++ {
			w := WORDS[permIdxs[i]]
			ws = append(ws, w)
		}
	}
	return ws[:count]
}
