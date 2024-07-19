package options

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func GetDefaultServeOptionUint64(envName string, defaultValue uint64) uint64 {
	envValue := os.Getenv(envName)
	if envValue != "" {
		// convert envValue to int
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return uint64(i)
		}
		return 0
	}
	return defaultValue
}

func GetDefaultServeOptionStringArray(envName string, defaultValue []string) []string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return strings.Split(envValue, ",")
	}
	return defaultValue
}

func GetDefaultServeOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

// use the helper function here for allowlist

func GetDefaultServeOptionBool(envName string, defaultValue bool) bool {
	envValue := os.Getenv(envName)
	if envValue != "" {
		fmt.Printf(" %s envValue: %s\n", envName, envValue)
		i, err := strconv.ParseBool(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}
