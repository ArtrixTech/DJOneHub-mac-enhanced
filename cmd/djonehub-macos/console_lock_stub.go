//go:build !darwin || !cgo

package main

func consoleLocked() (locked bool, known bool) {
	return false, false
}
