package exml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleRoot(t *testing.T) {
	doc := Document{
		ProcessingInstruction: XML10ProcessingInstruction,
	}

	root := MakeNode("root")
	doc.SetRoot(root)

	result := doc.String()
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root />\n", result)
}

func TestNS(t *testing.T) {
	doc := Document{
		ProcessingInstruction: XML10ProcessingInstruction,
	}

	root := MakeNode("root")
	doc.SetRoot(root)
	doc.AddNamespace("nanolog", "https://nanolog.app/ns")

	result := doc.String()
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root xmlns:nanolog=\"https://nanolog.app/ns\" />\n", result)
}

func TestChildren(t *testing.T) {
	doc := Document{
		ProcessingInstruction: XML10ProcessingInstruction,
	}

	root := MakeNode("root")
	doc.SetRoot(root)
	ns := doc.AddNamespace("nanolog", "https://nanolog.app/ns")

	child1 := root.AddChild("firstChildren", WithNamespace(ns))
	child2 := child1.AddChild("secondChildren", WithAttribute("foo", "bar"))
	child2.AddChild("thirdChildren", WithInnerText("Some text here. </>"))

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<root xmlns:nanolog="https://nanolog.app/ns">
    <nanolog:firstChildren>
        <secondChildren foo="bar">
            <thirdChildren>Some text here. &lt;/&gt;</thirdChildren>
        </secondChildren>
    </nanolog:firstChildren>
</root>`
	assert.Equal(t, expected, doc.String())
}

func TestSelfClosing(t *testing.T) {
	doc := Document{
		ProcessingInstruction: XML10ProcessingInstruction,
	}

	root := MakeNode("root")
	doc.SetRoot(root)
	ns := doc.AddNamespace("nanolog", "https://nanolog.app/ns")

	child1 := root.AddChild("firstChildren", WithNamespace(ns))
	child2 := child1.AddChild("secondChildren", WithAttribute("foo", "bar"))
	child2.AddChild("thirdChildren", WithInnerText("Some text here. </>"))
	root.AddChild("selfClosing")

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<root xmlns:nanolog="https://nanolog.app/ns">
    <nanolog:firstChildren>
        <secondChildren foo="bar">
            <thirdChildren>Some text here. &lt;/&gt;</thirdChildren>
        </secondChildren>
    </nanolog:firstChildren>
    <selfClosing/>
</root>`
	assert.Equal(t, expected, doc.String())
}
