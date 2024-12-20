package goutils

type Tuple []any

func Pair[T1, T2 any](t1 T1, t2 T2) Tuple {
	return []any{t1, t2}
}

func Triple[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) Tuple {
	return []any{t1, t2, t3}
}
