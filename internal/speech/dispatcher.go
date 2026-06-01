package speech

import (
	"context"
	"os/exec"
)

// Speak sends text to speech-dispatcher via spd-say.
// It's a fire-and-forget operation that ignores errors.
func Speak(ctx context.Context, text string) {
	go func() {
		_ = exec.Command("spd-say", text).Run()
	}()
}
