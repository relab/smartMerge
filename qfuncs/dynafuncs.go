package qfuncs

import (

	pr "github.com/relab/smartMerge/proto"

)

var DReadSQF = func(c *pr.Configuration, replies []*pr.AdvReadReply) (*pr.AdvReadReply, bool){
	
	// Stop RPC if new current configuration reported. 
	lastrep := replies[len(replies)-1]
	if lastrep.GetCur() != nil {
		return lastrep, true
	}
	
	// Return false, if not enough replies yet.
	if len(replies) < c.MaxQuorum() {
		return nil, false
	}
	
	lastrep = new(pr.AdvReadReply)
	for _,rep := range replies {
		if lastrep.GetState().Compare(rep.GetState()) == 1 {
			lastrep.State = rep.GetState()
		}
	}
	
	next := make([]*pr.Blueprint,0,1)
	for _, rep := range replies {
		next = DGetBlueprintSlice(next, rep)
	}
	
	lastrep.Next = next
	
	return lastrep, true	
}

var DWriteSQF = func(c *pr.Configuration, replies []*pr.AdvWriteSReply) (*pr.AdvWriteSReply, bool) {
	
	// Stop RPC if new current configuration reported. 
	lastrep := replies[len(replies)-1]
	if lastrep.GetCur() != nil {
		return lastrep, true
	}
	
	// Return false, if not enough replies yet.
	// This rpc is both reading and writing.
	if len(replies) < c.MaxQuorum() {
		return nil, false
	}
		
	next := make([]*pr.Blueprint,0,1)
	for _, rep := range replies {
		next = DGetBlueprintSlice(next, rep)
	}
	
	lastrep.Next = next
	
	
	return lastrep, true
}


var DWriteNSetQF = func(c *pr.Configuration, replies []*pr.DWriteNReply) (*pr.DWriteNReply, bool) {
	
	// Stop RPC if new current configuration reported. 
	lastrep := replies[len(replies)-1]
	if lastrep.GetCur() != nil {
		return lastrep, true
	}
	
	// Return false, if not enough replies yet.
	if len(replies) < c.WriteQuorum() {
		return nil, false
	}
	
	return nil, true
}

var GetOneNQF = func(c *pr.Configuration, replies []*pr.GetOneReply) (*pr.GetOneReply, bool) {
	return replies[0], true
}

var DSetCurQF = SetCurQF

func DGetBlueprintSlice(next []*pr.Blueprint, rep NextReport) []*pr.Blueprint {
	for _, blp := range rep.GetNext() {
		next = add(next, blp)
	}

	return next
}

func add(bls []*pr.Blueprint, bp *pr.Blueprint) []*pr.Blueprint {
	place := 0
	
	findplacefor:
	for _, blpr := range bls {
		switch blpr.Compare(bp) {
		case 1: 
			if bp.Compare(blpr) == 1 {
				//New blueprint already present
				return bls
			}
			continue
		case 0:
			continue
		case -1:
			break findplacefor
		}
		place += 1
	}
	
	bls = append(bls, nil)
	
	for i := len(bls)-1; i > place; i-- {
		bls[i] = bls[i-1]
	}
	bls[place] = bp
	
	return bls	
}
