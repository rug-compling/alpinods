//
// THIS IS A GENERATED FILE. DO NOT EDIT.
//

package alpinods

import (
	"fmt"
)

type AttributesT struct {
	Aform        string `xml:"aform,attr,omitempty"`
	Begin        int    `xml:"begin,attr"`
	Buiging      string `xml:"buiging,attr,omitempty"`
	Case         string `xml:"case,attr,omitempty"`
	Cat          string `xml:"cat,attr,omitempty"`
	Comparative  string `xml:"comparative,attr,omitempty"`
	Conjtype     string `xml:"conjtype,attr,omitempty"`
	Def          string `xml:"def,attr,omitempty"`
	Dial         string `xml:"dial,attr,omitempty"`
	Dscmanual    string `xml:"dscmanual,attr,omitempty"`
	Dscsense     string `xml:"dscsense,attr,omitempty"`
	End          int    `xml:"end,attr"`
	Frame        string `xml:"frame,attr,omitempty"`
	Gen          string `xml:"gen,attr,omitempty"`
	Genus        string `xml:"genus,attr,omitempty"`
	Getal        string `xml:"getal,attr,omitempty"`
	GetalN       string `xml:"getal-n,attr,omitempty"`
	Graad        string `xml:"graad,attr,omitempty"`
	His          string `xml:"his,attr,omitempty"`
	His1         string `xml:"his_1,attr,omitempty"`
	His2         string `xml:"his_2,attr,omitempty"`
	His11        string `xml:"his_1_1,attr,omitempty"`
	His12        string `xml:"his_1_2,attr,omitempty"`
	His21        string `xml:"his_2_1,attr,omitempty"`
	His22        string `xml:"his_2_2,attr,omitempty"`
	His111       string `xml:"his_1_1_1,attr,omitempty"`
	His112       string `xml:"his_1_1_2,attr,omitempty"`
	His121       string `xml:"his_1_2_1,attr,omitempty"`
	His122       string `xml:"his_1_2_2,attr,omitempty"`
	His211       string `xml:"his_2_1_1,attr,omitempty"`
	His212       string `xml:"his_2_1_2,attr,omitempty"`
	His221       string `xml:"his_2_2_1,attr,omitempty"`
	His222       string `xml:"his_2_2_2,attr,omitempty"`
	Id           int    `xml:"id,attr"`
	Iets         string `xml:"iets,attr,omitempty"`
	Index        string `xml:"index,attr,omitempty"`
	Infl         string `xml:"infl,attr,omitempty"`
	Lcat         string `xml:"lcat,attr,omitempty"`
	Lemma        string `xml:"lemma,attr,omitempty"`
	Lwtype       string `xml:"lwtype,attr,omitempty"`
	MwuRoot      string `xml:"mwu_root,attr,omitempty"`
	MwuSense     string `xml:"mwu_sense,attr,omitempty"`
	Naamval      string `xml:"naamval,attr,omitempty"`
	Neclass      string `xml:"neclass,attr,omitempty"`
	Npagr        string `xml:"npagr,attr,omitempty"`
	Ntype        string `xml:"ntype,attr,omitempty"`
	Num          string `xml:"num,attr,omitempty"`
	Numtype      string `xml:"numtype,attr,omitempty"`
	Pb           string `xml:"pb,attr,omitempty"`
	Pdtype       string `xml:"pdtype,attr,omitempty"`
	Per          string `xml:"per,attr,omitempty"`
	Personalized string `xml:"personalized,attr,omitempty"`
	Persoon      string `xml:"persoon,attr,omitempty"`
	Pos          string `xml:"pos,attr,omitempty"`
	Positie      string `xml:"positie,attr,omitempty"`
	Postag       string `xml:"postag,attr,omitempty"`
	Pron         string `xml:"pron,attr,omitempty"`
	Pt           string `xml:"pt,attr,omitempty"`
	Pvagr        string `xml:"pvagr,attr,omitempty"`
	Pvtijd       string `xml:"pvtijd,attr,omitempty"`
	Refl         string `xml:"refl,attr,omitempty"`
	Rel          string `xml:"rel,attr,omitempty"`
	Rnum         string `xml:"rnum,attr,omitempty"`
	Root         string `xml:"root,attr,omitempty"`
	Sc           string `xml:"sc,attr,omitempty"`
	Sense        string `xml:"sense,attr,omitempty"`
	Special      string `xml:"special,attr,omitempty"`
	Spectype     string `xml:"spectype,attr,omitempty"`
	Status       string `xml:"status,attr,omitempty"`
	Stype        string `xml:"stype,attr,omitempty"`
	Tense        string `xml:"tense,attr,omitempty"`
	Vform        string `xml:"vform,attr,omitempty"`
	Vwtype       string `xml:"vwtype,attr,omitempty"`
	Vztype       string `xml:"vztype,attr,omitempty"`
	Wh           string `xml:"wh,attr,omitempty"`
	Wk           string `xml:"wk,attr,omitempty"`
	Word         string `xml:"word,attr,omitempty"`
	Wvorm        string `xml:"wvorm,attr,omitempty"`
}

var NodeTags = []string{
	"aform",
	"begin",
	"buiging",
	"case",
	"cat",
	"comparative",
	"conjtype",
	"def",
	"dial",
	"dscmanual",
	"dscsense",
	"end",
	"frame",
	"gen",
	"genus",
	"getal",
	"getal-n",
	"graad",
	"his",
	"his_1",
	"his_2",
	"his_1_1",
	"his_1_2",
	"his_2_1",
	"his_2_2",
	"his_1_1_1",
	"his_1_1_2",
	"his_1_2_1",
	"his_1_2_2",
	"his_2_1_1",
	"his_2_1_2",
	"his_2_2_1",
	"his_2_2_2",
	"id",
	"iets",
	"index",
	"infl",
	"lcat",
	"lemma",
	"lwtype",
	"mwu_root",
	"mwu_sense",
	"naamval",
	"neclass",
	"npagr",
	"ntype",
	"num",
	"numtype",
	"pb",
	"pdtype",
	"per",
	"personalized",
	"persoon",
	"pos",
	"positie",
	"postag",
	"pron",
	"pt",
	"pvagr",
	"pvtijd",
	"refl",
	"rel",
	"rnum",
	"root",
	"sc",
	"sense",
	"special",
	"spectype",
	"status",
	"stype",
	"tense",
	"vform",
	"vwtype",
	"vztype",
	"wh",
	"wk",
	"word",
	"wvorm",
}

func (n *NodeT) getAttr(attr string) string {
	switch attr {
	case "aform":
		return n.Aform
	case "begin":
		return fmt.Sprint(n.Begin)
	case "buiging":
		return n.Buiging
	case "case":
		return n.Case
	case "cat":
		return n.Cat
	case "comparative":
		return n.Comparative
	case "conjtype":
		return n.Conjtype
	case "def":
		return n.Def
	case "dial":
		return n.Dial
	case "dscmanual":
		return n.Dscmanual
	case "dscsense":
		return n.Dscsense
	case "end":
		return fmt.Sprint(n.End)
	case "frame":
		return n.Frame
	case "gen":
		return n.Gen
	case "genus":
		return n.Genus
	case "getal":
		return n.Getal
	case "getal-n":
		return n.GetalN
	case "graad":
		return n.Graad
	case "his":
		return n.His
	case "his_1":
		return n.His1
	case "his_2":
		return n.His2
	case "his_1_1":
		return n.His11
	case "his_1_2":
		return n.His12
	case "his_2_1":
		return n.His21
	case "his_2_2":
		return n.His22
	case "his_1_1_1":
		return n.His111
	case "his_1_1_2":
		return n.His112
	case "his_1_2_1":
		return n.His121
	case "his_1_2_2":
		return n.His122
	case "his_2_1_1":
		return n.His211
	case "his_2_1_2":
		return n.His212
	case "his_2_2_1":
		return n.His221
	case "his_2_2_2":
		return n.His222
	case "id":
		return fmt.Sprint(n.Id)
	case "iets":
		return n.Iets
	case "index":
		return n.Index
	case "infl":
		return n.Infl
	case "lcat":
		return n.Lcat
	case "lemma":
		return n.Lemma
	case "lwtype":
		return n.Lwtype
	case "mwu_root":
		return n.MwuRoot
	case "mwu_sense":
		return n.MwuSense
	case "naamval":
		return n.Naamval
	case "neclass":
		return n.Neclass
	case "npagr":
		return n.Npagr
	case "ntype":
		return n.Ntype
	case "num":
		return n.Num
	case "numtype":
		return n.Numtype
	case "pb":
		return n.Pb
	case "pdtype":
		return n.Pdtype
	case "per":
		return n.Per
	case "personalized":
		return n.Personalized
	case "persoon":
		return n.Persoon
	case "pos":
		return n.Pos
	case "positie":
		return n.Positie
	case "postag":
		return n.Postag
	case "pron":
		return n.Pron
	case "pt":
		return n.Pt
	case "pvagr":
		return n.Pvagr
	case "pvtijd":
		return n.Pvtijd
	case "refl":
		return n.Refl
	case "rel":
		return n.Rel
	case "rnum":
		return n.Rnum
	case "root":
		return n.Root
	case "sc":
		return n.Sc
	case "sense":
		return n.Sense
	case "special":
		return n.Special
	case "spectype":
		return n.Spectype
	case "status":
		return n.Status
	case "stype":
		return n.Stype
	case "tense":
		return n.Tense
	case "vform":
		return n.Vform
	case "vwtype":
		return n.Vwtype
	case "vztype":
		return n.Vztype
	case "wh":
		return n.Wh
	case "wk":
		return n.Wk
	case "word":
		return n.Word
	case "wvorm":
		return n.Wvorm
	}
	return ""
}
