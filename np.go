package alpinods

func doNp(node *Node, q *context) {
	q.rels = make(map[int][]string)
	prepareNp1(node, q)
	prepareNp2(node, q)
}

func prepareNp1(node *Node, q *context) {
	if q.rels[node.ID] == nil {
		q.rels[node.ID] = make([]string, 0)
	}
	q.rels[node.ID] = append(q.rels[node.ID], node.Rel)
	if n := idx(node, q); n != node {
		if q.rels[n.ID] == nil {
			q.rels[n.ID] = make([]string, 0)
		}
		q.rels[n.ID] = append(q.rels[n.ID], node.Rel)
	}
	if node.Node != nil {
		for _, n := range node.Node {
			prepareNp1(n, q)
		}
	}
}

func prepareNp2(node *Node, q *context) {
	if isNP(node, q) {
		node.IsNp = "true"
	}
	if node.Node != nil {
		for _, n := range node.Node {
			doNp(n, q)
		}
	}
}

func isNP(node *Node, q *context) bool {

	if node.Cat == "" && node.Pt == "" {
		return false
	}

	if node.Cat == "np" {
		return true
	}

	if node.Lcat == "np" && otherString(q.rels[node.ID], "hd", "mwp") {
		return true
	}

	if node.Pt == "n" && otherString(q.rels[node.ID], "hd") {
		return true
	}

	if node.Pt == "vnw" && node.Pdtype == "pron" && otherString(q.rels[node.ID], "hd") {
		return true
	}

	if node.Cat == "mwu" && hasString(q.rels[node.ID], "su", "obj1", "obj2", "app") {
		return true
	}

	if node.Node != nil {
		for _, n := range node.Node {
			if n.Rel != "cnj" {
				continue
			}
			n = idx(n, q)
			if isNP(n, q) {
				return true
			}
		}
	}

	return false
}

// is er een string in ss die gelijk is aan een string in s ?
func hasString(ss []string, s ...string) bool {
	for _, s1 := range ss {
		for _, s2 := range s {
			if s1 == s2 {
				return true
			}
		}
	}
	return false
}

// is er een string in ss die ongelijk is aan alle strings in s ?
func otherString(ss []string, s ...string) bool {
LOOP:
	for _, s1 := range ss {
		for _, s2 := range s {
			if s1 == s2 {
				continue LOOP
			}
		}
		return true
	}
	return false
}
