package xml

const XML10ProcessingInstruction = `xml version="1.0" encoding="UTF-8"`

// Document represents an XML document consisting with an optional list of
// namespaces, and ProcessingInstruction.
// Use SetRoot to set the document's root node.
type Document struct {
	Namespaces            []*Namespace
	ProcessingInstruction string
	root                  *Node
}

// Root returns the document's root node.
func (d Document) Root() *Node {
	return d.root
}

// SetRoot sets the document's root node.
func (d *Document) SetRoot(n *Node) {
	if n != nil {
		n.Parent = nil
		n.parentDoc = d
	}
	d.root = n
}

// String returns the Document's XML representation.
func (d Document) String() string {
	return printXML(d)
}

// AddNamespace registers a new namespace to be used within nodes of this
// document.
func (d *Document) AddNamespace(name string, url string) *Namespace {
	ns := &Namespace{
		Name: name,
		URL:  url,
	}
	d.Namespaces = append(d.Namespaces, ns)
	return ns
}

// Namespace represents a single namespace to be used within nodes of a
// document.
type Namespace struct {
	Name string
	URL  string
}

// NodeMutatorFn represents a function that received and mutates a given Node
type NodeMutatorFn func(n *Node)

// WithAttribute adds an attribute to a node. To add an attribute associated
// with a namespace, use WithNSAttribute
func WithAttribute(name, value string) NodeMutatorFn {
	return func(n *Node) {
		n.Attributes = append(n.Attributes, Attribute{
			Name:      name,
			Value:     value,
			Namespace: nil,
		})
	}
}

// WithNSAttribute adds an attribute from a Namespace with a given name and
// value.
func WithNSAttribute(ns *Namespace, name, value string) NodeMutatorFn {
	return func(n *Node) {
		n.Attributes = append(n.Attributes, Attribute{
			Name:      name,
			Value:     value,
			Namespace: ns,
		})
	}
}

// WithNamespace indicates the element being added belongs to the given
// namespace, and will receive the namespace's name as its prefix.
func WithNamespace(ns *Namespace) NodeMutatorFn {
	return func(n *Node) {
		n.Namespace = ns
	}
}

// WithInnerText associates a given string to the node's content.
func WithInnerText(inner string) NodeMutatorFn {
	return func(n *Node) {
		n.InnerText = inner
	}
}

// Node represents a single node on a Document
type Node struct {
	Attributes []Attribute
	Name       string
	Namespace  *Namespace
	InnerText  string
	Children   []*Node
	Parent     *Node
	parentDoc  *Document
}

// AddChild creates, appends, and returns a node using the receiver as its
// parent.
func (n *Node) AddChild(name string, mutators ...NodeMutatorFn) *Node {
	newNode := &Node{
		Parent: n,
		Name:   name,
	}
	for _, fn := range mutators {
		fn(newNode)
	}
	n.Children = append(n.Children, newNode)

	return newNode
}

// Attribute represents a single XML attribute for a node
type Attribute struct {
	Name      string
	Value     string
	Namespace *Namespace
}

// MakeNode creates a new Node with no parent, applies the provided mutators,
// and returns it.
func MakeNode(name string, mutators ...NodeMutatorFn) *Node {
	newNode := &Node{
		Name: name,
	}

	for _, fn := range mutators {
		fn(newNode)
	}

	return newNode
}
