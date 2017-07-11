// +build collision

package collision

import (
	"github.com/autovelop/playthos/std"
	// "log"
	"math"
)

// TODO: run tests through these functions

func CheckCollisionAABB(one *Collider, two *Collider) bool {
	// collisionX := one.Get().Position().X+one.Get().GetScale().X >= two.Get().Position().X &&
	// 	two.Get().Position().X+two.Get().GetScale().X >= one.Get().Position().X
	// collisionY := one.Get().Position().Y+one.Get().GetScale().Y >= two.Get().Position().Y &&
	// 	two.Get().Position().Y+two.Get().GetScale().Y >= one.Get().Position().Y
	one_transform, one_relative := one.Get()
	two_transform, two_relative := two.Get()
	collisionX := one_transform.Position().X+one_relative.X+one_relative.W >= two_transform.Position().X+two_relative.X &&
		two_transform.Position().X+two_relative.X+two_relative.W >= one_transform.Position().X+one_relative.X

	collisionY := one_transform.Position().Y+one_relative.Y+one_relative.H >= two_transform.Position().Y+two_relative.Y &&
		two_transform.Position().Y+two_relative.Y+two_relative.H >= one_transform.Position().Y+one_relative.Y

	// log.Println(one.Get().Position().Y)
	// log.Println(two.Get().Position().Y)
	return collisionX && collisionY
}

// returns two axis representing all four sides of the rect
// assumes always a rect
func FindRectAxis(rect *std.Rect) (std.Vector2, std.Vector2) {
	half_w := rect.W / 2
	half_h := rect.H / 2
	axis_x := std.Vector2{rect.X + half_w, rect.X - half_w}
	axis_y := std.Vector2{rect.Y + half_h, rect.Y - half_h}
	return axis_x, axis_y
}

func Dot(point std.Vector2, axis std.Vector2) float32 {
	return point.X*axis.X + point.Y*axis.Y
}

func Distance3(p1 std.Vector3, p2 std.Vector3) float32 {
	return float32(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}
func Distance2(p1 std.Vector2, p2 std.Vector2) float32 {
	return float32(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}

// returns max and min of a signle rect projected onto the axis
func ProjectRectOnAxis(rect *std.Rect, axis *std.Vector2) (float32, float32) {
	half_w := rect.W / 2
	half_h := rect.H / 2

	ul := std.Vector2{rect.X - half_w, rect.Y + half_h}
	ur := std.Vector2{rect.X + half_w, rect.Y + half_h}
	ll := std.Vector2{rect.X - half_w, rect.Y - half_h}
	lr := std.Vector2{rect.X + half_w, rect.Y - half_h}

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
