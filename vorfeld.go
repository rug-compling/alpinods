package alpinods

func doVorfeld(node *Node, q *context) {
	q.vorfeld = make(map[int]map[int]bool)
	q.vorfeldSkip = make(map[int]map[int]bool)
	prepareVorfeld1(node, q)
	prepareVorfeld2(node, q)
	prepareVorfeld3(node, q)
}

func prepareVorfeld1(node *Node, q *context) {
	q.vorfeld[node.ID] = make(map[int]bool)
	q.vorfeldSkip[node.ID] = make(map[int]bool)
	if node.Node != nil {
		for _, n := range node.Node {
			prepareVorfeld1(n, q)
		}
	}
}

func prepareVorfeld2(node *Node, q *context) {
	if node.Cat == "smain" {
		smainVorfeld(node, q)
	} else if node.Cat == "whq" {
		whqVorfeld(node, q)
	}
	if node.Node != nil {
		for _, n := range node.Node {
			prepareVorfeld2(n, q)
		}
	}
}

func prepareVorfeld3(node *Node, q *context) {
	if node.Cat != "" || node.Pt != "" {
		for id := range q.vorfeld[node.ID] {
			if !q.vorfeldSkip[node.ID][id] {
				node.IsVorfeld = "true"
				break
			}
		}
	}
	if node.Node != nil {
		for _, n := range node.Node {
			prepareVorfeld3(n, q)
		}
	}
}

func smainVorfeld(node *Node, q *context) {
	if node.Node != nil {
		for _, n := range node.Node {
			if n.Rel == "hd" {
				// NIET alleen primary links
				n = idx(n, q)
				if n.Word != "" {
					for _, topic := range findTopic(node, n.Begin, false, q) {
						if checkTopic(topic, node, n.Begin, q) {
							q.vorfeld[topic.ID][node.ID] = true
						} else {
							q.vorfeldSkip[topic.ID][node.ID] = true
						}
					}
				}
			}
		}
	}
}

func whqVorfeld(node *Node, q *context) {
	if node.Node != nil {
		for _, n := range node.Node {
			// NIET alleen primary links
			rel := n.Rel
			n = idx(n, q)
			if rel == "body" && n.Cat == "sv1" {
				for _, n2 := range n.Node {
					if n2.Rel == "hd" {
						// NIET alleen primary links
						n2 = idx(n2, q)
						if n2.Word != "" {
							for _, topic := range findTopic(node, n2.Begin, true, q) {
								if checkTopic(topic, node, n2.Begin, q) {
									q.vorfeld[topic.ID][node.ID] = true
								} else {
									q.vorfeldSkip[topic.ID][node.ID] = true
								}
							}
						}
					}
				}
			}
		}
	}
}

func findTopic(node *Node, begin int, needWhd bool, q *context) []*Node {
	topics := make([]*Node, 0)

	// hier: inclusief topnode
	//if isTopic(node, begin) {
	//      topics = append(topics, node)
	//}

	if node.Node != nil {
		for _, n := range node.Node {

			if needWhd && n.Rel != "whd" {
				continue
			}

			// hier: exclusief topnode
			if isTopic(n, begin, q) {
				topics = append(topics, n)
			}

			// ALLEEN primary links
			for _, topic := range findTopic(n, begin, false, q) {
				topics = append(topics, topic)
			}
		}
	}
	return topics
}

func isTopic(node *Node, begin int, q *context) bool {
	if node.Begin < begin && node.End <= begin {
		return true
	}
	if node.Lemma != "" || node.Cat == "mwu" {
		if node.Begin < begin {
			return true
		}
		return false
	}

	if node.Node != nil {
		for _, n := range node.Node {
			if n.Rel == "hd" || n.Rel == "cmp" || n.Rel == "crd" {
				// NIET alleen primary links
				n = idx(n, q)
				if (n.Lemma != "" || n.Cat == "mwu") && n.Begin < begin {
					return true
				}
			}
		}
	}
	return false
}

func checkTopic(topic, node *Node, begin int, q *context) bool {
	// alle nodes tussen node (exclusief) en topic (exclusief)
	nodes := make(map[*Node]bool)
	nodePath(node, topic, nodes, q)

	for n := range nodes {
		if isTopic(n, begin, q) {
			return false
		}
	}

	return true
}

func nodePath(top, bottom *Node, nodes map[*Node]bool, q *context) bool {
	retval := false
	if top.Node != nil {
		for _, n := range top.Node {
			// NIET alleen primaire links (lijkt niet logisch, maar werkt toch beter)
			n = idx(n, q)
			if n == bottom {
				retval = true
			} else if nodePath(n, bottom, nodes, q) {
				nodes[n] = true
				retval = true
			}
		}
	}
	return retval
}
