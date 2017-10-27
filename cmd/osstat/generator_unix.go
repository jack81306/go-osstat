// +build !windows,!linux,!darwin

package main

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
