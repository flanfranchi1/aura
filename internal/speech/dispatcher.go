// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (c) 2026 Fernando Lanfranchi

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
