// +build animation

package animation

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
	"log"
	"math"
)

// type Vector3AnimationClip struct {
// 	AnimationClip
// 	value *std.Vector3
// }

// type Vector2AnimationClip struct {
// 	AnimationClip
// 	value *std.Vector2
// }
type AnimationFrame struct {
	key       bool
	value     std.Animatable
	initValue std.Animatable
	// remember that this step is the diff to the next keyframe / 60
	// step std.Animatable
}

type AnimationClip struct {
	engine.Component
	duration float64
	frames   []*AnimationFrame
	// value  std.Animatable
	// currentStep   std.Animatable
	// keyframeCount int
	count         float64
	ticks         float64
	percCompleted float64
	// Value std.Animatable
	// From  std.Animatable
	// Diff  std.Animatable
	// Start std.Animatable
	// To     std.Animatable
	// Step   std.Animatable
	// Frames []std.Animatable
	// From      *std.Vector3
	// Start     std.Vector3
	// To        std.Vector3
	// step      *std.Vector3
	// progress  float64
	playing  bool
	autoplay bool
	// loop      bool
	// reverse   bool
	// reversing bool
}

func (a *AnimationClip) Set(s float64, au bool) {
	a.count = 0
	a.duration = s
	// a.autoplay = au
	a.frames = make([]*AnimationFrame, 0)
	// log.Fatal(len(a.frames))
	// a.frameCount = n
	if au {
		a.playing = true
	}
}

func (a *AnimationClip) AddFrame(i int, value std.Animatable) {
	a.count++
	// if a.keyframeCount < 1 {
	// 	a.value = value
	// if i == 1 {
	// diff := a.frames[1].value.Copy()
	// diff.Sub(a.frames[0].value)
	// diff.Mul(float32(a.percCompleted))
	// a.frames[0].value.Add(diff)
	a.frames = append(a.frames, &AnimationFrame{true, value, value.Copy()})
	// a.frames[i] = &AnimationFrame{true, value, value.Copy()}
	// } else {
	// 	a.frames[i] = &AnimationFrame{true, value, nil}
	// }
	// 	a.keyframeCount++
	// 	if i == 0 {
	// for i := range a.frames {
	// 	a.frames[i] = &AnimationFrame{false, value}
	// }
	// } else {
	// 	log.Fatal("The first keyframe must be at index 0")
	// }
	// } else {
	// if a.keyframeCount == 1 {
	// 	for b := 0; b < i-1; b++ {
	// 		a.frames = append(a.frames, nil)
	// 	}
	// 	a.frames = append(a.frames, &AnimationFrame{true, value, *new(std.Animatable)})

	// 	prevFrame := a.frames[0]
	// 	currFrame := a.frames[i]
	// 	prevFrame.step = currFrame.value.Copy().Sub(prevFrame.value)
	// 	prevFrame.step.Div(60)
	// 	currFrame.step = currFrame.value.Copy().Sub(prevFrame.value)
	// 	currFrame.step.Div(60)
	// log.Fatal(prevFrame.step)

	// check how many frames are between this and zero

	// } else {
	// 	log.Fatal("Only supporting 2 keyframes for now")
	// }
	// }
}

func (a *AnimationClip) Stop() {
	a.playing = false
	// a.progress = 0.0
}

func (a *AnimationClip) Play() {
	// a.value
	a.playing = true
}

// func (a *AnimationClip) AddFrame(v std.Animatable, t std.Animatable) {
// a.Value = v
// a.From = v.Copy()
// a.To = t.Copy()
// a.Diff = t.Copy().Sub(v.Copy())
// a.Step = t.Copy().Sub(v.Copy())
// if a.duration > 0 {
// 	a.Step.Div(100)
// 	a.Step.Div(float32(a.duration))
// }
// log.Printf("%v, %v, %v, %v, %v", a.Value, a.From, a.To, a.Diff, a.Step)
// }

func (a *AnimationClip) Reset() {
	frame := a.frames[0].value
	frame.Zero()
	frame.Add(a.frames[0].initValue)
}

func (a *AnimationClip) Update() {
	// zero pointer value
	// add init value of current frame to pointer value
	// get current frame with progress
	// get next frame
	// check difference
	// divide progress by count

	thisFrameIndex := int(math.Floor((a.count - 1) * a.percCompleted))
	log.Printf("this frame #: %v\n", thisFrameIndex)
	log.Printf("overall %%: %v\n", a.percCompleted)
	log.Printf("every frame is %%: %v\n", 1/(a.count-1))
	// log.Printf("speed %%: %v\n", a.percCompleted)

	frame := a.frames[0].value
	frame.Zero()
	frame.Add(a.frames[thisFrameIndex].initValue)

	targetValue := a.frames[thisFrameIndex+1].initValue.Copy()
	base := a.frames[thisFrameIndex].initValue.Copy()
	targetValue.Sub(base)
	mul := a.frames[thisFrameIndex].initValue.Copy()

	_, ok := frame.(*std.Integer)
	if ok {
		mul.Set(float32(math.Floor(float64(float32(a.percCompleted*(a.count-1))-float32(thisFrameIndex)) + 0.5)))
	} else {
		mul.Set(float32(a.percCompleted*(a.count-1)) - float32(thisFrameIndex))
	}

	log.Printf("mul: %v\n", math.Floor(float64(float32(a.percCompleted*(a.count-1))-float32(thisFrameIndex))+0.5))
	targetValue.Mul(mul)
	log.Printf("mul2: %v\n", mul)
	log.Printf("target value %%: %v\n", a.frames[thisFrameIndex+1].initValue)
	log.Printf("frame diff to target: %v\n", targetValue)
	// frame.Add(a.frames[0].initValue)

	// if thisFrameIndex == 2 {
	// 	log.Fatalf("%v %v", frame, a.frames[thisFrameIndex+1].value)
	// }

	// one := a.frames[thisFrameIndex+1].initValue.Copy()
	// diff := a.frames[thisFrameIndex+1].initValue.Copy()
	// diff.Sub(a.frames[thisFrameIndex].initValue)
	// log.Printf("diff1: %v\n", diff)

	// dummy := a.frames[0].value.Copy()
	// dummy.Zero()
	// dummy.Add(diff)
	// dummy.Mul(diff)
	// log.Printf("step: %v\n", diff)
	// one.Set(float32(a.percCompleted * (4)))
	// log.Printf("perc: %v\n", one)

	// diff.Mul(one)
	// log.Printf("diff2: %v\n", diff)
	// log.Fatal(a.frames[thisFrameIndex].initValue)
	// diff.Add(a.frames[thisFrameIndex].initValue)
	// log.Printf("base: %v\n", a.frames[thisFrameIndex].initValue)
	frame.Add(targetValue)
	// log.Printf("value: %v\n", frame)
	// if thisFrameIndex == 1 {
	// 	log.Fatal(diff)
	// }
	// log.Fatalf("%v  %v\n", diff, float32(a.percCompleted/(a.count-1)))
	// diff.Mul(float32(a.percCompleted / (a.count - 1)))
	// frame.Add(diff)
	// log.Printf("diff to next frame %%: %v\n", diff)
	// log.Printf("%v -> %v == %v\n%v | %v | %v",
	// 	a.frames[thisFrameIndex].initValue,
	// 	a.frames[thisFrameIndex+1].initValue,
	// 	frame,
	// 	diff,
	// 	a.percCompleted,
	// 	math.Floor((a.count-1)*a.percCompleted),
	// )

}

func (a *AnimationClip) Loop() {
	// a.loop = true
	// REMEMBER THAT FROM MUST BE A VALUE AND NOT POINTER
	// from := a.From
	// var b *std.Animatable
	// b = &a.From
	// from := new(std.Animatable)
	// log.Fatal(from.Get())
	// from.Add(a.From)
	// a.Start = a.From.Copy()
	// a.Start = *a.From
}

func (a *AnimationClip) Reverse() {
	// a.reverse = true
}
