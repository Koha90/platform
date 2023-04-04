// Package services ...
package services

type lifecycle int

// Transient ...
const (
	Transient lifecycle = iota
	Singleton
	Scoped
)
