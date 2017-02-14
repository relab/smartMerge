package confProvider

import (
	bp "github.com/relab/smartMerge/blueprints"
	pb "github.com/relab/smartMerge/proto"
)

// ThriftyConfP is a configuration provider that does not avoid recontacting servers.
// Oups is only thrifty, if underlying provider is also thrifty.
type ThriftyConfP struct {
	Provider
}

func (cp *ThriftyConfP) ReadC(blp *bp.Blueprint, rids []int) *pb.Configuration {
	return cp.Provider.ReadC(blp, nil)
}

func (cp *ThriftyConfP) WriteC(blp *bp.Blueprint, rids []int) *pb.Configuration {
	return cp.Provider.WriteC(blp, nil)
}

func (cp *ThriftyConfP) GIDs(in []int) []uint32 {
	return nil
}

func (cp *ThriftyConfP) WriteCNoS(blp *bp.Blueprint, rids []int) *pb.Configuration {
	return cp.Provider.WriteCNoS(blp, nil)
}
