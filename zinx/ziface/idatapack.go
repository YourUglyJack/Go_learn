package ziface

// 处理tcp粘包问题
// 在收到数据的时候分两次进行读取，先读取固定长度的head部分，得到后续Data的长度，再根据DataLen读取之后的body

type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
