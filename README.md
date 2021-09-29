# exml

🧑‍💻 Go XML generator without Structs™

Package `exml` allows XML documents to be generated without the usage of structs or maps. It is not intended for
every-day usage, since package `xml` will probably be a better option.

## Installing

You probably already know the drill:

```
go get github.com/heyvito/exml
```

## Usage

```go
package main

import (
	"fmt"
	
	"github.com/heyvito/exml"
)

func main() {
    doc := exml.Document{
        ProcessingInstruction: exml.XML10ProcessingInstruction,
    }

    ns := doc.AddNamespace("foo", "https://tempuri.org")
    root := exml.MakeNode("root", exml.WithNamespace(ns), exml.WithAttribute("attr", "value"))
    root.AppendChild("child", exml.WithInnerText("Some value here"))
    child := root.AppendChild("other-child", exml.WithAttribute("len", "1"))
    child.AppendChild("other-tag", exml.WithInnerText("More text\nWith line breaks."))

    fmt.Println(doc.String())
}
```

The code above will print the following XML document:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<root xmlns:foo="https://tempuri.org">
    <child>Some value here</child>
    <other-child len="1">
        <other-tag>
            More text
            With line breaks.
        </other-tag>
    </other-child>
</root>
```

## License

```
MIT License

Copyright (c) 2021 Victor Gama de Oliveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
