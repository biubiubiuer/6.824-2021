package raft

import (
	"6.824/labgob"
	"6.824/utils"
	"bytes"
)

func (rf *Raft) raftState() []byte {
	w := new(bytes.Buffer)
	e := labgob.NewEncoder(w)

	if e.Encode(rf.log) != nil ||
		e.Encode(rf.currentTerm) != nil ||
		e.Encode(rf.votedFor) != nil {

		utils.Debug(utils.DError, "S%d encode fail", rf.me)
		panic("encode fail")
	}
	data := w.Bytes()
	return data
}

func (rf *Raft) persist() {
	// Your code here (2C).
	rf.persister.SaveRaftState(rf.raftState())
}
