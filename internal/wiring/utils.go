package wiring

import "strings"

func toPascalCase(s string) string {
	parts := strings.Split(s, "-")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func toCamelCase(s string) string {
	p := toPascalCase(s)
	return strings.ToLower(p[:1]) + p[1:]
}

func indent(s string, spaces int) string {
	prefix := strings.Repeat(" ", spaces)
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}
