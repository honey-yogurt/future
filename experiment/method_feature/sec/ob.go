package sec

import "strconv"

type SecSyncHarderPoint struct {
	id     string
	number int
}

func NewSecSyncHarderPoint() *SecSyncHarderPoint {
	return &SecSyncHarderPoint{
		id:     "SecSyncHarderPoint",
		number: 77,
	}
}

func (s *SecSyncHarderPoint) GetId() string {
	return s.id
}

func (s *SecSyncHarderPoint) getNumber() string {
	return strconv.Itoa(s.number)
}

type sSecSyncHarderPoint struct {
	id     string
	number int
}

func newsSecSyncHarderPoint() *sSecSyncHarderPoint {
	return &sSecSyncHarderPoint{
		id:     "sSecSyncHarderPoint",
		number: 77,
	}
}

func (s *sSecSyncHarderPoint) GetId() string {
	return s.id
}

func (s *sSecSyncHarderPoint) getNumber() string {
	return strconv.Itoa(s.number)
}

type Sec struct {
	S *sSecSyncHarderPoint
}

func NewSec() *Sec {
	return &Sec{S: newsSecSyncHarderPoint()}
}
