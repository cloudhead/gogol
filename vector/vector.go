package vector

type Vector struct {
	X, Y float32
}

func New(x, y float32) Vector {
	return Vector{x, y}
}

func (v *Vector) Add(v1 Vector) Vector {
	return Vector{v.X + v1.X, v.Y + v1.Y}
}

func (v *Vector) Sub(v1 Vector) Vector {
	return Vector{v.X - v1.X, v.Y - v1.Y}
}

func (v *Vector) Mul(k float32) Vector {
	return Vector{v.X * k, v.Y * k}
}
