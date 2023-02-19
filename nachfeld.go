package alpinods

var (
	vpCat = map[string]bool{
		"inf":   true,
		"ti":    true,
		"ssub":  true,
		"oti":   true,
		"ppart": true,
	}
	rHead = map[string]bool{
		"hd":   true,
		"cmp":  true,
		"mwp":  true,
		"crd":  true,
		"rhd":  true,
		"whd":  true,
		"nucl": true,
		"dp":   true,
	}
)

func doNachfeld(node *Node, q *context) {
	findNachfeld(node, q)
	if node.Node != nil {
		for _, n := range node.Node {
			doNachfeld(n, q)
		}
	}
}

func findNachfeld(node *Node, q *context) {
	if !vpCat[node.Cat] {
		return
	}

	if node.Node == nil || len(node.Node) == 0 {
		return
	}

	// zoek begin van head
	headBegin := -1
	for _, n := range node.Node {
		n = idx(n, q)
		if n.Rel == "hd" {
			headBegin = n.Begin
			break
		}
	}
	if headBegin < 0 {
		return
	}
	// zoek nachfeld
	for _, n := range node.Node {
		n = idx(n, q)
		if n.Rel != "hd" {
			setNachfeld(n, headBegin, node, q)
		}
	}
}

func setNachfeld(node *Node, begin int, vp *Node, q *context) {
	if node.Rel != "hd" && node.Rel != "svp" && node.Cat != "inf" && node.Cat != "ppart" {

		var skip func(*Node, int) bool
		skip = func(current *Node, state int) bool {
			current = idx(current, q)
			if current == node {
				return state == 2
			}
			switch state {
			case 0:
				state = 1
			case 1:
				if vpCat[current.Cat] {
					state = 2
				}
			}
			if current.Node != nil {
				for _, n := range current.Node {
					if skip(n, state) {
						return true
					}
				}
			}
			return false
		}

		n2 := make([]*Node, 0)
		if node.Node != nil {
			for _, nn2 := range node.Node {
				nn2 = idx(nn2, q)
				if rHead[nn2.Rel] {
					n2 = append(n2, nn2)
				}
			}
		}
		if len(n2) == 0 {
			if begin < node.Begin {
				if !skip(vp, 0) {
					node.IsNachfeld = "true"
				}
				return
			}
		} else {
			ok := true
			for _, nn2 := range n2 {
				if begin >= nn2.Begin {
					ok = false
					break
				}
			}
			if ok {
				if !skip(vp, 0) {
					node.IsNachfeld = "true"
				}
				return
			}
		}
	}
	if vpCat[node.Cat] {
		return
	}
	if node.Node == nil {
		return
	}
	for _, n := range node.Node {
		n = idx(n, q)
		setNachfeld(n, begin, vp, q)
	}
}
