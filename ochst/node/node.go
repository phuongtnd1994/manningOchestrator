package node

type Node struct {
	Name            string
	Ip              string
	Api             string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int64
	DiskAllocated   int64
	Role            string
	TaskCount       int
}

func NewNode(name string, api string, role string) *Node {
	return &Node{
		Name: name,
		Api:  api,
		Role: role,
	}
}
