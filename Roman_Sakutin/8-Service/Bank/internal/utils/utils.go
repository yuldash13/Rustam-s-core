package utils

func PrtTo[V any](v *V) V {
	if v == nil {
		return *new(V)
	}
	return *v
}
