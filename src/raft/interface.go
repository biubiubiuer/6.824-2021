package raft

import "6.824/utils"

func (rf *Raft) Start(command interface{}) (int, int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if rf.status != leader {
		utils.Debug(utils.DClient, "S%d Not leader cmd: %+v", rf.me, command)
		return -1, -1, false
	}

	index := rf.lastLogIndex() + 1
	rf.log = append(rf.log, Entry{index, rf.currentTerm, command})
	rf.persist()

	// defer utils.Debug(utils.DLog2, "S%d append log: %+v", rf.me, rf.log)
	utils.Debug(utils.DClient, "S%d cmd: %+v, logIndex: %d", rf.me, command, rf.lastLogIndex())

	rf.doAppendEntries()

	return rf.lastLogIndex(), rf.currentTerm, true
}
