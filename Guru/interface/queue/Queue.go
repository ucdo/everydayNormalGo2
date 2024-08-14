package Qeueu

type q int

type Queue []q

func (q *Queue) Push(v q) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (v q) {
	if len(*q) == 0 {
		return 0
	}
	v, *q = (*q)[0], (*q)[1:]
	return v
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
