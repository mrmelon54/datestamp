package main

import (
	"github.com/djherbis/times"
	"io/fs"
	"time"
)

type TimeType uint8

const (
	ATime TimeType = iota
	BTime
	CTime
	MTime
)

func (t TimeType) String() string {
	switch t {
	case ATime:
		return "atime"
	case BTime:
		return "btime"
	case CTime:
		return "ctime"
	case MTime:
		return "mtime"
	default:
		return "unknown"
	}
}

func getEarliestFileTime(fi fs.FileInfo) (time.Time, TimeType) {
	spec := times.Get(fi)
	t := spec.AccessTime()
	tType := ATime
	mTime := spec.ModTime()
	if mTime.Before(t) {
		t = mTime
		tType = MTime
	}
	if spec.HasBirthTime() {
		bTime := spec.BirthTime()
		if bTime.Before(t) {
			t = bTime
			tType = BTime
		}
	}
	if spec.HasChangeTime() {
		cTime := spec.ChangeTime()
		if cTime.Before(t) {
			t = cTime
			tType = CTime
		}
	}
	return t, tType
}
