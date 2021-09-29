package xml

import "strings"

type builder struct {
	strings.Builder
}

func newBuilder() *builder {
	return &builder{}
}

func (b *builder) indent(level int) {
	b.writeString(strings.Repeat(string(space), level*4))
}

func (b *builder) writeInnerText(indent int, str string) {
	lines := strings.Split(str, linebreak)
	multiLine := len(lines) > 1
	if multiLine {
		b.linebreak()
	}
	for _, v := range lines {
		if multiLine {
			b.indent(indent)
		}
		b.escapeText([]byte(v), false)
		if multiLine {
			b.linebreak()
		}
	}
}

func (b *builder) writeString(str string) {
	_, _ = b.WriteString(str)
}

func (b *builder) writeRune(r rune) {
	_, _ = b.WriteRune(r)
}

func (b *builder) beginTag(indent int, n *Node) {
	b.indent(indent)
	b.writeRune(tagStart)
	b.tagName(n)
}

func (b *builder) tagName(n *Node) {
	if n.Namespace != nil {
		b.writeString(n.Namespace.Name)
		b.writeRune(colon)
	}
	b.writeString(n.Name)
}

func (b *builder) endTag(indent int, n *Node) {
	b.indent(indent)
	b.writeRune(tagStart)
	b.writeRune(slash)
	b.tagName(n)
	b.writeRune(tagEnd)
}

func (b *builder) linebreak() {
	b.writeString(linebreak)
}

func (b *builder) writeAttr(ns, name, value string) {
	b.writeRune(space)
	if ns != "" {
		b.writeString(ns)
		b.writeRune(colon)
	}
	b.writeString(name)
	b.writeRune(equal)
	b.writeRune(quote)
	b.escapeString(value)
	b.writeRune(quote)
}

func (b *builder) writeAttrs(n *Node) {
	for _, v := range n.Attributes {
		ns := ""
		if v.Namespace != nil {
			ns = v.Namespace.Name
		}
		b.writeAttr(ns, v.Name, v.Value)
	}
}

func (b *builder) writeNode(indent int, n *Node) {
	b.beginTag(indent, n)
	b.writeAttrs(n)

	if n.InnerText != "" {
		multiline := strings.Contains(n.InnerText, "\n")
		b.writeRune(tagEnd)
		b.writeInnerText(indent+1, n.InnerText)
		idn := indent
		if !multiline {
			idn = 0
		}
		b.endTag(idn, n)
	} else if len(n.Children) > 0 {
		b.writeRune(tagEnd)
		b.linebreak()
		for _, n := range n.Children {
			b.writeNode(indent+1, n)
		}
		b.endTag(indent, n)
	} else {
		b.writeRune(slash)
		b.writeRune(tagEnd)
		return
	}
	b.linebreak()
}
