// +build !windows

package main

// GetCommand to get ip configuration
func GetCommand() []string {
	return []string{"ifconfig", "-a"}
}
