package main

import "strconv"

// 目的：
// 1.比较结构体和指针所有者是否有区别
// 2.方法的可见性是否有区别
// 3.还有结构体的可见性

type SyncHarder struct {
	id     string
	number int
}

func NewSyncHarder() SyncHarder {
	return SyncHarder{
		id:     "SyncHarder",
		number: 66,
	}
}
func NewSyncHarderPp() *SyncHarder {
	return &SyncHarder{
		id:     "SyncHarder",
		number: 66,
	}
}

func (s SyncHarder) GetId() string {
	return s.id
}

func (s SyncHarder) getNumber() string {
	return strconv.Itoa(s.number)
}

type SyncHarderPoint struct {
	id     string
	number int
}

func NewSyncHarderPoint() *SyncHarderPoint {
	return &SyncHarderPoint{
		id:     "SyncHarderPoint",
		number: 77,
	}
}

func NewSyncHarderPoint2() SyncHarderPoint {
	return SyncHarderPoint{
		id:     "SyncHarderNoPoint",
		number: 77,
	}
}

func (s *SyncHarderPoint) GetId() string {
	return s.id
}

func (s *SyncHarderPoint) getNumber() string {
	return strconv.Itoa(s.number)
}

type SyncSofter struct {
	id     string
	number int
}

func NewSyncSofter() SyncSofter {
	return SyncSofter{
		id:     "SyncSofter",
		number: 88,
	}
}

func (s SyncSofter) GetId() string {
	return s.id
}

func (s SyncSofter) getNumber() string {
	return strconv.Itoa(s.number)
}

type SyncSofterPoint struct {
	id     string
	number int
}

func NewSyncSofterPoint() *SyncSofterPoint {
	return &SyncSofterPoint{
		id:     "SyncSofterPoint",
		number: 99,
	}
}

func (s *SyncSofterPoint) GetId() string {
	return s.id
}

func (s *SyncSofterPoint) getNumber() string {
	return strconv.Itoa(s.number)
}
