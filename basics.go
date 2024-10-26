package wiz

import (
	"iter"
)

func IFilter[V any](i iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for e := range i {
			if f(e) && !yield(e) {
				return
			}
		}
	}
}

func IReduce[V any, R any](i iter.Seq[V], f func(R, V) R, init R) R {
	r := init
	for e := range i {
		r = f(r, e)
	}
	return r
}

func Chain[V any](iters ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, it := range iters {
			for e := range it {
				if !yield(e) {
					return
				}
			}
		}
	}
}

func IMap[V any, R any](i iter.Seq[V], f func(V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for e := range i {
			if !yield(f(e)) {
				return
			}
		}
	}
}

func Zip[K any, V any](ki []K, vi []V) iter.Seq2[K, V] {
	if len(ki) != len(vi) {
		panic("zip: ki and vi must have the same length")
	}
	return func(yield func(K, V) bool) {
		for i := 0; i < len(ki); i++ {
			if !yield(ki[i], vi[i]) {
				return
			}
		}
	}
}
