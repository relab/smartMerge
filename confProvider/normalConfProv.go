package confProvider

import (
	bp "github.com/relab/smartmerge/blueprints"
	pb "github.com/relab/smartmerge/proto"
)

// This Config provider always return a full configuration.
type NormalConfP struct {
	Provider
}

func (cp *NormalConfP) ReadC(blp *bp.Blueprint, rids []uint32) *pb.Configuration {
	return cp.Provider.FullC(blp)
}

func (cp *NormalConfP) WriteC(blp *bp.Blueprint, rids []uint32) *pb.Configuration {
	return cp.Provider.FullC(blp)
}

func (cp *NormalConfP) WriteCNoS(blp *bp.Blueprint, rids []uint32) *pb.Configuration {
	return cp.Provider.FullC(blp)
}

/*
func (cp *NormalConfP) SingleC(blp *pb.Blueprint) *pb.Configuration {
	return cp.Provider.ReadC(blp, nil)
}*/
