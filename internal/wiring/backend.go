package wiring

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateBackendWiring(projectRoot string, modules []string) error {

	outputPath := filepath.Join(projectRoot, "backend/modules/modules.gen.go")

	var imports []string
	var registrations []string

	for _, m := range modules {

		importAlias := m
		importPath := fmt.Sprintf("modules/%s/backend", m)

		imports = append(imports, fmt.Sprintf(`%s "%s"`, importAlias, importPath))

		structName := toPascalCase(m) + "Module"

		registrations = append(registrations,
			fmt.Sprintf("&%s.%s{}", importAlias, structName),
		)
	}

	code := fmt.Sprintf(`package modules

import (
	"github.com/gin-gonic/gin"
	"backend/core"
	%s
)

func LoadModules() []core.Module {
	return []core.Module{
		%s,
	}
}
`,
		strings.Join(imports, "\n"),
		strings.Join(registrations, ",\n"),
	)

	return os.WriteFile(outputPath, []byte(code), 0644)
}
