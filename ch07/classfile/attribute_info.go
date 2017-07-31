package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) [] AttributeInfo  {
	attributesCount := reader.readUint16()
	attributes := make([] AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo  {
	attributeNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attributeNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo  {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	//case "LocalVariableTable":
	//	return &LocalVariabletableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
