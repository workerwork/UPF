package elem

type NodeIDType byte

const (
	NodeIDTypeIPv4Addr NodeIDType = 0
	NodeIDTypeIPv6Addr NodeIDType = 1
	NodeIDTypeFQDN     NodeIDType = 2
)

type NodeID struct {
	EType   IEType
	ELength uint16
	NodeIDType
	NodeID []byte
}

func NewIPv4NodeID(ipv4 []byte) *NodeID {
	return &NodeID{
		EType:      IETypeNodeID,
		ELength:    uint16(4),
		NodeIDType: NodeIDTypeIPv4Addr,
		NodeID:     ipv4,
	}
}

func DecodeNodeID(data []byte, len uint16) *NodeID {
	return &NodeID{
		EType:      IETypeNodeID,
		ELength:    len,
		NodeIDType: NodeIDType(getValue(data, 1)[0]),
		NodeID:     getValue(data, len-1),
	}
}

func EncodeNodeID(nodeID NodeID) []byte {
	return setValue(nodeID.EType, nodeID.ELength, nodeID.NodeIDType, nodeID.NodeID)
}

//判断是否含有NodeID
func HasNodeID(nodeID NodeID) bool {
	if nodeID.EType == 0 {
		return false
	}
	return true
}
