package wz

type WZImage struct {
	*WZSimpleNode // heh
	Properties    WZProperty
}

func NewWZImage(name string, parent *WZSimpleNode) *WZImage {
	node := new(WZImage)
	node.WZSimpleNode = NewWZSimpleNode(name, parent)
	return node
}

func (m *WZImage) Parse(file *WZFileBlob, offset int64) {
	if file.Debug {
		m.debug(file, "> WZImage::Parse")
		defer func() { m.debug(file, "< WZImage::Parse") }()
	}

	file.seek(offset)
	typename := file.readDeDuplicatedWZString(m.GetPath(), offset, true)
	parsedObject := ParseObject(m.Name, typename, m.WZSimpleNode, file, offset)

	objResult, isOK := parsedObject.(WZProperty)
	if !isOK {
		panic("Expected object to be WZProperty")
	}

	m.Properties = objResult
}
