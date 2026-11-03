//go:build darwin && cgo

package main

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>
#include <CoreFoundation/CoreFoundation.h>

// Returns 1 when the login session shows the lock screen, 0 when it does not,
// and -1 when there is no GUI session to ask.
static int djonehub_console_locked(void) {
	CFDictionaryRef session = CGSessionCopyCurrentDictionary();
	if (session == NULL) {
		return -1;
	}
	int locked = 0;
	CFTypeRef value = CFDictionaryGetValue(session, CFSTR("CGSSessionScreenIsLocked"));
	if (value != NULL && CFGetTypeID(value) == CFBooleanGetTypeID()) {
		locked = CFBooleanGetValue((CFBooleanRef)value) ? 1 : 0;
	}
	CFRelease(session);
	return locked;
}
*/
import "C"

// consoleLocked reports whether the user's screen is locked. known is false
// when the process has no GUI session, so callers keep their old behaviour.
func consoleLocked() (locked bool, known bool) {
	switch C.djonehub_console_locked() {
	case 1:
		return true, true
	case 0:
		return false, true
	}
	return false, false
}
