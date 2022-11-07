package raft

import "6.824/utils"

type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []Entry
	LeaderCommit int
}

type AppendEntriesReply struct {
	Term    int
	Success bool
	XTerm   int // for fast backup
	XIndex  int
	XLen    int
}

type InstallSnapshotArgs struct {
	// Your data here (2A, 2B).
	Term              int
	LeaderId          int
	LastIncludedIndex int
	LastIncludedTerm  int
	Data              []byte
}

type InstallSnapshotReply struct {
	Term int
}

func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	utils.Debug(utils.DInfo, "S%d send AppendEntries request to %d {%+v}", rf.me, server, args)
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	if !ok {
		utils.Debug(utils.DWarn, "S%d call (AppendEntries)rpc to C%d error", rf.me, server)
		return ok
	}
	utils.Debug(utils.DInfo, "S%d get AppendEntries response from %d {%+v}", rf.me, server, reply)
	return ok
}

func (rf *Raft) sendInstallSnapshot(server int, args *InstallSnapshotArgs, reply *InstallSnapshotReply) bool {
	utils.Debug(utils.DInfo, "S%d send InstallSnapshot request to %d {%+v}", rf.me, server, args)
	ok := rf.peers[server].Call("Raft.InstallSnapshot", args, reply)
	if !ok {
		utils.Debug(utils.DWarn, "S%d call (InstallSnapshot)rpc to C%d error", rf.me, server)
		return ok
	}
	utils.Debug(utils.DInfo, "S%d get InstallSnapshot response from %d {%+v}", rf.me, server, reply)
	return ok
}
