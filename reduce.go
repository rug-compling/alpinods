package alpinods

// Reduce returns a minimal copy of an AlpinoDS document.
//
// Entities and atttributes that are not used in de reduced
// feature set of the minimal Alpino plug-in for TrEd are removed.
//
// # The document version is set to 1.5
//
// see: https://www.let.rug.nl/vannoord/alp/Alpino/tred/
func Reduce(alpino *AlpinoDS) *AlpinoDS {
	ds := &AlpinoDS{
		Version: "1.5",
		Sentence: &Sentence{
			SentID:   alpino.Sentence.SentID,
			Sentence: alpino.Sentence.Sentence,
		},
	}

	if alpino.Metadata != nil {
		ds.Metadata = &Metadata{
			Meta: make([]Meta, 0),
		}
		for _, meta := range alpino.Metadata.Meta {
			ds.Metadata.Meta = append(ds.Metadata.Meta, Meta{
				Type:  meta.Type,
				Name:  meta.Name,
				Value: meta.Value,
			})
		}
	}

	if alpino.Node != nil {
		ds.Node = reducedCopy(alpino.Node)
	}

	if alpino.Comments != nil {
		ds.Comments = &Comments{
			Comment: make([]string, 0),
		}
		ds.Comments.Comment = append(ds.Comments.Comment, alpino.Comments.Comment...)
	}

	return ds
}

func reducedCopy(node *Node) *Node {
	n := &Node{
		NodeAttributes: NodeAttributes{
			// minimal set recognized in TrEd
			Begin:  node.Begin,
			Cat:    node.Cat,
			End:    node.End,
			ID:     node.ID,
			Index:  node.Index,
			Lemma:  node.Lemma,
			Pos:    node.Pos,
			Postag: node.Postag,
			Rel:    node.Rel,
			Root:   node.Root,
			Sense:  node.Sense,
			Word:   node.Word,
			// attributes added by TrEd upon save
			Buiging:  node.Buiging,
			Conjtype: node.Conjtype,
			Dial:     node.Dial,
			Genus:    node.Genus,
			Getal:    node.Getal,
			GetalN:   node.GetalN,
			Graad:    node.Graad,
			Lwtype:   node.Lwtype,
			Naamval:  node.Naamval,
			Npagr:    node.Npagr,
			Ntype:    node.Ntype,
			Numtype:  node.Numtype,
			Pdtype:   node.Pdtype,
			Persoon:  node.Persoon,
			Positie:  node.Positie,
			Pt:       node.Pt,
			Pvagr:    node.Pvagr,
			Pvtijd:   node.Pvtijd,
			Spectype: node.Spectype,
			Status:   node.Status,
			Vwtype:   node.Vwtype,
			Vztype:   node.Vztype,
			Wvorm:    node.Wvorm,
		},
		Node: make([]*Node, 0),
	}
	if node.Node != nil {
		for _, nn := range node.Node {
			n.Node = append(n.Node, reducedCopy(nn))
		}
	}
	return n
}
