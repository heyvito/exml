package xml

import (
	"fmt"
)

const (
	quote     = '"'
	equal     = '='
	slash     = '/'
	tagStart  = '<'
	tagEnd    = '>'
	colon     = ':'
	xmlns     = "xmlns"
	space     = ' '
	linebreak = "\n"
)

func printXML(d Document) string {
	b := newBuilder()
	if d.ProcessingInstruction != "" {
		b.writeString(fmt.Sprintf("<?%s?>\n", d.ProcessingInstruction))
	}

	r := d.root
	if r == nil {
		return b.String()
	}

	b.beginTag(0, r)
	b.writeAttrs(r)

	if len(d.Namespaces) > 0 {
		for _, ns := range d.Namespaces {
			b.writeAttr(xmlns, ns.Name, ns.URL)
		}
	}

	if len(r.Children) == 0 && r.InnerText == "" {
		b.writeRune(space)
		b.writeRune(slash)
		b.writeRune(tagEnd)
		b.linebreak()
		return b.String()
	} else if r.InnerText != "" {
		b.writeRune(tagEnd)
		b.linebreak()
		b.writeInnerText(1, r.InnerText)
		b.linebreak()
	} else {
		b.writeRune(tagEnd)
		b.linebreak()
		for _, n := range r.Children {
			b.writeNode(1, n)
		}
	}

	b.endTag(0, r)
	return b.String()
}
