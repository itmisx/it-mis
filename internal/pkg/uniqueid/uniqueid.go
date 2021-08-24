package uniqueid

import (
	"it-mis/init/config"
	"it-mis/internal/pkg/debugx"

	"github.com/bwmarrin/snowflake"
)

// NewID 获取全局唯一id
func NewID() (int64, error) {
	nodeID := config.Config.APP.NodeID
	if nodeID >= 1024 {
		nodeID = 0
	}
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		debugx.PrintError(err)
		return -1, err
	}
	id := node.Generate()
	return id.Int64(), nil
}
