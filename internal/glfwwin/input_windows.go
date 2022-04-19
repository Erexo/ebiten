// Copyright 2002-2006 Marcus Geelnard
// Copyright 2006-2019 Camilla Löwy
// Copyright 2022 The Ebiten Authors
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would
//    be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source
//    distribution.

package glfwwin

func (w *Window) inputKey(key Key, scancode int, action Action, mods ModifierKey) {
	panic("inputKey is not impleented")
}

func (w *Window) inputMouseClick(button MouseButton, action Action, mods ModifierKey) {
	panic("inputMouseClick is not implemented")
}

func (w *Window) centerCursorInContentArea() error {
	width, height, err := w.platformGetWindowSize()
	if err != nil {
		return err
	}
	if err := w.platformSetCursorPos(float64(width/2), float64(height/2)); err != nil {
		return err
	}
	return nil
}

func (w *Window) inputChar(codepoint uint32, mods ModifierKey, plain bool) {
	panic("inputChar is not impleented")
}

func (w *Window) inputCursorPos(xpos float64, ypos float64) {
	panic("inputCursorPos is not implemented")
}

func (w *Window) inputCursorEnter(entered bool) {
	panic("inputCursorEnter is not implemented")
}

func (w *Window) inputDrop(paths []string) {
	panic("inputDrop is not implemented")
}

func (w *Window) inputScroll(xoffset, yoffset float64) {
	panic("inputScroll is not implemented")
}
