package alpinods

import (
	"encoding/xml"
	"strings"
)

// A document in alpino_ds XML format
type AlpinoT struct {
	XMLName  xml.Name   `xml:"alpino_ds"`
	Version  string     `xml:"version,attr,omitempty"`
	Metadata *MetadataT `xml:"metadata,omitempty"`
	Parser   *ParserT   `xml:"parser,omitempty"`
	Node     *NodeT     `xml:"node,omitempty"`
	Sentence *SentenceT `xml:"sentence,omitempty"`
	Comments *CommentsT `xml:"comments,omitempty"`
	Root     []*DeprelT `xml:"root,omitempty"`
	Conllu   *ConlluT   `xml:"conllu,omitempty"`

	UserData interface{} `xml:"-"`
}

type MetadataT struct {
	Meta []MetaT `xml:"meta,omitempty"`
}

type MetaT struct {
	Type  string `xml:"type,attr,omitempty"`
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`

	UserData interface{} `xml:"-"`
}

type ParserT struct {
	Build string `xml:"build,attr,omitempty"`
	Date  string `xml:"date,attr,omitempty"`
	Cats  string `xml:"cats,attr,omitempty"`
	Skips string `xml:"skips,attr,omitempty"`
}

type CommentsT struct {
	Comment []string `xml:"comment,omitempty"`
}

type SentenceT struct {
	Sentence string `xml:",chardata"`
	SentId   string `xml:"sentid,attr,omitempty"`
}

type ConlluT struct {
	Conllu string `xml:",cdata"`
	Status string `xml:"status,attr,omitempty"`
	Error  string `xml:"error,attr,omitempty"`
	Auto   string `xml:"auto,attr,omitempty"`
}

type NodeT struct {
	AttributesT
	Ud   *UdT     `xml:"ud,omitempty"`
	Node []*NodeT `xml:"node"`

	UserData interface{} `xml:"-"`
}

type UdT struct {
	Id    string `xml:"id,attr,omitempty"`
	Form  string `xml:"form,attr,omitempty"`
	Lemma string `xml:"lemma,attr,omitempty"`
	Upos  string `xml:"upos,attr,omitempty"`
	FeatsT
	Head       string `xml:"head,attr,omitempty"`
	Deprel     string `xml:"deprel,attr,omitempty"`
	DeprelMain string `xml:"deprel_main,attr,omitempty"`
	DeprelAux  string `xml:"deprel_aux,attr,omitempty"`
	Dep        []DepT `xml:"dep,omitempty"`

	UserData interface{} `xml:"-"`
}

type DepT struct {
	Id         string `xml:"id,attr,omitempty"`
	Head       string `xml:"head,attr,omitempty"`
	Deprel     string `xml:"deprel,attr,omitempty"`
	DeprelMain string `xml:"deprel_main,attr,omitempty"`
	DeprelAux  string `xml:"deprel_aux,attr,omitempty"`
	Elided     bool   `xml:"elided,attr,omitempty"`

	UserData interface{} `xml:"-"`
}

type FeatsT struct {
	Abbr     string `xml:"Abbr,attr,omitempty"`
	Case     string `xml:"Case,attr,omitempty"`
	Definite string `xml:"Definite,attr,omitempty"`
	Degree   string `xml:"Degree,attr,omitempty"`
	Foreign  string `xml:"Foreign,attr,omitempty"`
	Gender   string `xml:"Gender,attr,omitempty"`
	Number   string `xml:"Number,attr,omitempty"`
	Person   string `xml:"Person,attr,omitempty"`
	PronType string `xml:"PronType,attr,omitempty"`
	Reflex   string `xml:"Reflex,attr,omitempty"`
	Tense    string `xml:"Tense,attr,omitempty"`
	VerbForm string `xml:"VerbForm,attr,omitempty"`
}

type DeprelT struct {
	XMLName xml.Name

	RecursionLimit string `xml:"recursion_limit,attr,omitempty"`

	Ud    string `xml:"ud,attr,omitempty"`
	Id    string `xml:"id,attr,omitempty"`
	Eid   string `xml:"eid,attr,omitempty"`
	Form  string `xml:"form,attr,omitempty"`
	Lemma string `xml:"lemma,attr,omitempty"`
	Upos  string `xml:"upos,attr,omitempty"`
	FeatsT
	Head      string `xml:"head,attr,omitempty"`
	Deprel    string `xml:"deprel,attr,omitempty"`
	DeprelAux string `xml:"deprel_aux,attr,omitempty"`

	Buiging  string `xml:"buiging,attr,omitempty"`
	Conjtype string `xml:"conjtype,attr,omitempty"`
	Dial     string `xml:"dial,attr,omitempty"`
	Genus    string `xml:"genus,attr,omitempty"`
	Getal    string `xml:"getal,attr,omitempty"`
	GetalN   string `xml:"getal-n,attr,omitempty"`
	Graad    string `xml:"graad,attr,omitempty"`
	Lwtype   string `xml:"lwtype,attr,omitempty"`
	Naamval  string `xml:"naamval,attr,omitempty"`
	Npagr    string `xml:"npagr,attr,omitempty"`
	Ntype    string `xml:"ntype,attr,omitempty"`
	Numtype  string `xml:"numtype,attr,omitempty"`
	Pdtype   string `xml:"pdtype,attr,omitempty"`
	Persoon  string `xml:"persoon,attr,omitempty"`
	Positie  string `xml:"positie,attr,omitempty"`
	Pt       string `xml:"pt,attr,omitempty"`
	Pvagr    string `xml:"pvagr,attr,omitempty"`
	Pvtijd   string `xml:"pvtijd,attr,omitempty"`
	Spectype string `xml:"spectype,attr,omitempty"`
	Status   string `xml:"status,attr,omitempty"`
	Vwtype   string `xml:"vwtype,attr,omitempty"`
	Vztype   string `xml:"vztype,attr,omitempty"`
	Wvorm    string `xml:"wvorm,attr,omitempty"`

	Deps []*DeprelT `xml:",any"`

	UserData interface{} `xml:"-"`
}

func (a AlpinoT) String() string {
	b, err := xml.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err) // Shouldn't happen
	}
	s := string(b)
	for _, a := range []string{"parser", "meta", "node", "dep" /* , "nattr", "rattr" */} {
		s = strings.Replace(s, "></"+a+">", "/>", -1)
	}
	return "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + s + "\n"
}
