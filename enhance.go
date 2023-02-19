package alpinods

// Features defines a set of flags to pass to the method AlpinoDS.Enhance()
type Features uint

const (
	Fnp       Features = 1 << iota                  // Add attribute `is_np`
	Fvorfeld                                        // Add attribute `is_vorfeld`
	Fnachfeld                                       // Add attribute `is_nachfeld`
	Fall      Features = Fnp | Fvorfeld | Fnachfeld // All of the above
)

type context struct {
	idxnodes    map[int]*Node
	rels        map[int][]string
	vorfeld     map[int]map[int]bool
	vorfeldSkip map[int]map[int]bool
}

// Enhance enriches an AlpinoDS document with structural information
// that is not provided by the Alpino parser.
//
// In addition, it sets the document version to DtdVersion.
//
// Old values for the chosen options are erased.
func (alpino *AlpinoDS) Enhance(f Features) {

	alpino.Version = DtdVersion

	q := &context{
		idxnodes: make(map[int]*Node),
	}
	prepare(alpino.Node, f, q)

	if f&Fnp != 0 {
		doNp(alpino.Node, q)
	}
	if f&Fvorfeld != 0 {
		doVorfeld(alpino.Node, q)
	}
	if f&Fnachfeld != 0 {
		doNachfeld(alpino.Node, q)
	}
}

func prepare(node *Node, f Features, q *context) {
	if f&Fnp != 0 {
		node.IsNp = ""
	}
	if f&Fvorfeld != 0 {
		node.IsVorfeld = ""
	}
	if f&Fnachfeld != 0 {
		node.IsNachfeld = ""
	}
	if node.Index != 0 && (node.Cat != "" || node.Pt != "") {
		q.idxnodes[node.Index] = node
	}
	if node.Node != nil {
		for _, n := range node.Node {
			prepare(n, f, q)
		}
	}
}

func idx(node *Node, q *context) *Node {
	if n, ok := q.idxnodes[node.Index]; ok {
		return n
	}
	return node
}
