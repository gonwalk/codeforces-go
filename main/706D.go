package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

var x706D = uint(1)

func fastRand() uint {
	x706D ^= x706D << 13
	x706D ^= x706D >> 17
	x706D ^= x706D << 5
	return x706D
}

type node706D struct {
	lr       [2]*node706D
	priority uint
	key      int
	value    int
}

func (o *node706D) rotate(d int) *node706D {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap706D struct {
	root       *node706D
	comparator func(a, b int) int
}

func newTreap706D() *treap706D {
	return &treap706D{comparator: func(a, b int) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *treap706D) _put(o *node706D, key int) *node706D {
	if o == nil {
		return &node706D{priority: fastRand(), key: key, value: 1}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.value++
	}
	return o
}

func (t *treap706D) put(key int) { t.root = t._put(t.root, key) }

func (t *treap706D) _delete(o *node706D, key int) *node706D {
	if o == nil {
		return nil
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		if o.value > 1 {
			o.value--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			cmp2 := 0
			if o.lr[0].priority > o.lr[1].priority {
				cmp2 = 1
			}
			o = o.rotate(cmp2)
			o.lr[cmp2] = t._delete(o.lr[cmp2], key)
		}
	}
	return o
}

func (t *treap706D) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap706D) max() int {
	var max *node706D
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return max.key
}

func (t *treap706D) floor(key int) int {
	var floor *node706D
	for o := t.root; o != nil; {
		switch cmp := t.comparator(key, o.key); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			floor = o
			o = o.lr[1]
		default:
			return o.key
		}
	}
	return floor.key
}

// github.com/EndlessCheng/codeforces-go
func Sol706D(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	t := newTreap706D()
	t.put(0)

	nextFloor := func(bitPos int, floor, x int) (f int, i int) {
		for i = bitPos; i >= 0; i-- {
			ui := uint(i)
			bitFloor := floor >> ui & 1
			bitX := x >> ui & 1
			if bitFloor == 1 && bitX == 1 {
				// mask floor
				check := floor&^(1<<ui) | (1<<ui - 1)
				newFloor := t.floor(check)
				return newFloor, i - 1
			}
		}
		return -1, -1
	}

	var q, x int
	var op string
	for Fscan(in, &q); q > 0; q-- {
		switch Fscan(in, &op, &x); op[0] {
		case '+':
			t.put(x)
		case '-':
			t.delete(x)
		default:
			floor := t.max()
			if floor == 0 {
				Fprintln(out, x)
				continue
			}
			ans := floor ^ x
			var newFloor int
			for i := min(bits.Len(uint(floor)), bits.Len(uint(x))) - 1; i >= 0; {
				newFloor, i = nextFloor(i, floor, x)
				if newFloor == -1 {
					break
				}
				if newXor := newFloor ^ x; newXor > ans {
					ans = newXor
					floor = newFloor
				}
			}
			Fprintln(out, ans)
		}
	}
}

//func main() {
//	Sol706D(os.Stdin, os.Stdout)
//}
