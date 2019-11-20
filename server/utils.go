package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func prettyPrintJSON(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func jsonCodeBlock(in string) string {
	return fmt.Sprintf("``` json\n%s\n```", in)
}

func codeBlock(in string) string {
	return fmt.Sprintf("```\n%s\n```", in)
}

func inlineCode(in string) string {
	return fmt.Sprintf("`%s`", in)
}

func standardizeName(name string) string {
	return strings.ToLower(name)
}

func validLicenseOption(license string) bool {
	return license == licenseOptionE10 || license == licenseOptionE20 || license == licenseOptionTE
}

func validDatabaseOption(databaseChoice string) bool {
	return databaseChoice == databaseOptionRDS || databaseChoice == databaseOptionOperator
}

func validInstallationName(name string) bool {
	return installationNameMatcher.MatchString(name)
}

func validFilestoreOption(filestoreChoice string) bool {
	return filestoreChoice == filestoreOptionS3 || filestoreChoice == filestoreOptionOperator
}

// NewBool returns a pointer to a given bool.
func NewBool(b bool) *bool { return &b }

// NewInt returns a pointer to a given int.
func NewInt(n int) *int { return &n }

// NewInt32 returns a pointer to a given int32.
func NewInt32(n int32) *int32 { return &n }

// NewInt64 returns a pointer to a given int64.
func NewInt64(n int64) *int64 { return &n }

// NewString returns a pointer to a given string.
func NewString(s string) *string { return &s }
