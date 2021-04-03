package glog

import "runtime"

// getCaller returns the calling function. skipFrames indicates how far up the
// ladder we go when looking for the caller (2 = direct caller, 3 = caller of
// the caller, …)
func getCaller(skipFrames int) string {
	// Thanks to StackOverflow User Not_a_Golfer for providing this code
	targetFrameIndex := skipFrames + 2

	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame.Function
}
