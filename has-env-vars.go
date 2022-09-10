package tutils

import "os"

func HasEnvVars(envsToCheck []string) ([]string, bool) {
	var notPresent []string

	for _, value := range envsToCheck {
		if _, exists := os.LookupEnv(value); !exists {
			notPresent = append(notPresent, value)
		}
	}

	return notPresent, len(notPresent) != 0
}

