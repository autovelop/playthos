// +build autovelop_playthos_collision !play

package collision

import (
	"github.com/autovelop/playthos/std"
	"math"
)

// CheckCollisionAABB tests whether two colliders are intersecting
func CheckCollisionAABB(one *Collider, two *Collider) bool {
	onePosition, _, oneSize := one.Get()
	twoPosition, _, twoSize := two.Get()

	onePositionMax := std.Vector2{
		onePosition.X + (oneSize.X / 2),
		onePosition.Y + (oneSize.Y / 2),
	}
	onePositionMin := std.Vector2{
		onePosition.X - (oneSize.X / 2),
		onePosition.Y - (oneSize.Y / 2),
	}
	twoPositionMax := std.Vector2{
		twoPosition.X + (twoSize.X / 2),
		twoPosition.Y + (twoSize.Y / 2),
	}
	twoPositionMin := std.Vector2{
		twoPosition.X - (twoSize.X / 2),
		twoPosition.Y - (twoSize.Y / 2),
	}

	collisionX := onePositionMax.X > twoPositionMin.X && twoPositionMax.X > onePositionMin.X
	collisionY := onePositionMax.Y > twoPositionMin.Y && twoPositionMax.Y > onePositionMin.Y

	return collisionX && collisionY
}

// FindRectAxis returns two axis representing all four sides of the rect
func FindRectAxis(rect *std.Rect) (std.Vector2, std.Vector2) {
	halfW := rect.W / 2
	halfH := rect.H / 2
	axisX := std.Vector2{rect.X + halfW, rect.X - halfW}
	axisY := std.Vector2{rect.Y + halfH, rect.Y - halfH}
	return axisX, axisY
}

// Dot returns dot product of points and axis (Vector2)
func Dot(point std.Vector2, axis std.Vector2) float32 {
	return point.X*axis.X + point.Y*axis.Y
}

// Distance3 returns distance between two Vector3 points
func Distance3(p1 std.Vector3, p2 std.Vector3) float32 {
	return float32(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}

// Distance2 returns distance between two Vector2 points
func Distance2(p1 std.Vector2, p2 std.Vector2) float32 {
	return float32(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}

// ProjectRectOnAxis returns max and min of a single rect projected onto the axis
func ProjectRectOnAxis(rect *std.Rect, axis *std.Vector2) (float32, float32) {
	halfW := rect.W / 2
	halfH := rect.H / 2

	ul := std.Vector2{rect.X - halfW, rect.Y + halfH}
	ur := std.Vector2{rect.X + halfW, rect.Y + halfH}
	ll := std.Vector2{rect.X - halfW, rect.Y - halfH}
	lr := std.Vector2{rect.X + halfW, rect.Y - halfH}

	min := Dot(ul, *axis)
	max := min

	p := Dot(ur, *axis)
	if p < min {
		min = p
	} else if p > max {
		max = p
	}

	p = Dot(ll, *axis)
	if p < min {
		min = p
	} else if p > max {
		max = p
	}

	p = Dot(lr, *axis)
	if p < min {
		min = p
	} else if p > max {
		max = p
	}
	// double max = min;
	// for (int i = 1; i < shape.vertices.length; i++) {
	// // NOTE: the axis must be normalized to get accurate projections
	// double p = axis.dot(shape.vertices[i]);
	// if (p < min) {
	// min = p;
	// } else if p > max if (p > max) {
	// max = p;
	// }
	// }
	// Projection proj = new Projection(min, max);
	// return proj;

	// ul_proj_x := (((ul.X*axis.X)+(ul.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.X
	// ul_proj_y := (((ul.X*axis.X)+(ul.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.Y

	// ur_proj_x := (((ur.X*axis.X)+(ur.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.X
	// ur_proj_y := (((ur.X*axis.X)+(ur.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.Y

	// ll_proj_x := (((ll.X*axis.X)+(ll.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.X
	// ll_proj_y := (((ll.X*axis.X)+(ll.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.Y

	// lr_proj_x := (((lr.X*axis.X)+(lr.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.X
	// lr_proj_y := (((lr.X*axis.X)+(lr.Y*axis.Y))/(axis.X*axis.X) + (axis.Y * axis.Y)) * axis.Y

	// only get min and max
	return min, max
	// return engine.Vector2{ul_proj_x, ul_proj_y}, engine.Vector2{ur_proj_x, ur_proj_y}, engine.Vector2{ll_proj_x, ll_proj_y}, engine.Vector2{lr_proj_x, lr_proj_y}
}
