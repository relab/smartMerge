package rpc

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	lat "github.com/relab/smartMerge/directCombineLattice"
	pb "github.com/relab/smartMerge/proto"
	"github.com/relab/smartMerge/regserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	silentLogger := log.New(ioutil.Discard, "", log.LstdFlags)
	grpclog.SetLogger(silentLogger)
}

func TestRegisterCall(t *testing.T) {
	addr1 := "127.0.0.1:8080"
	addr2 := "127.0.0.1:9090"

	serv1, err1 := regserver.StartTest(8080)
	serv2, err2 := regserver.StartTest(9090)
	defer serv1.Stop()
	defer serv2.Stop()

	if err1 != nil || err2 != nil {
		t.Fatalf("Error creating server")
	}

	time.Sleep(50 * time.Millisecond)

	mgr, err := NewManager(
		[]string{addr1, addr2},
		WithGrpcDialOptions(
			grpc.WithBlock(),
			grpc.WithTimeout(50*time.Millisecond),
		),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// All ids (2 machines)
	ids := mgr.IDs()

	// Example: Write configuration, quorum=2, n=2
	myWriteConfig, err := mgr.NewConfiguration(ids, 2)
	if err != nil {
		t.Fatalf("%v", err)
	}
	i := myWriteConfig.QuorumSize()
	if i != 2 {
		t.Errorf("Quorum Size was %d", i)
	}
	i = myWriteConfig.ReadQuorumSize()
	if i != 1 {
		t.Errorf("Read Quorum Size was %d", i)
	}

	s, newCur, err := myWriteConfig.ReadS(nil)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}
	fmt.Print(s)
	s = pb.State{nil, 2, 0}
	newCur, err = myWriteConfig.WriteS(&s, nil)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}

	sr, newCur, err := myWriteConfig.ReadS(nil)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}
	if sr.Timestamp != s.Timestamp {
		t.Errorf("return was %v, expected %v.", sr, s)
	}

	blps, newCur, err := myWriteConfig.ReadN(nil)
	if err != nil {
		t.Fatalf("readN: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}

	fmt.Printf("ReadN before Write: %v", blps)

	bluep1 := lat.GetBlueprint(bp1)
	bluep2 := lat.GetBlueprint(bp2)

	newCur, err = myWriteConfig.WriteN(&bluep1, nil)
	if err != nil {
		t.Fatalf("writeN: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}

	_, err = myWriteConfig.WriteN(&bluep2, nil)
	if err != nil {
		t.Fatalf("writeN: %v", err)
	}

	blps, newCur, err = myWriteConfig.ReadN(nil)
	if err != nil {
		t.Fatalf("readN: %v", err)
	}
	if newCur != nil {
		t.Errorf("new Cur: %v", newCur)
	}

	if len(blps) != 2 || !blps[0].Equals(bluep1) || !blps[1].Equals(bluep2) {
		t.Errorf("readN returned: %v", blps)
	}

}
