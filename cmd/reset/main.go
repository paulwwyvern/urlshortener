package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const MagicWord = "generate:reset"

type StructData struct {
	ShortName     string
	Name          string
	NullifyFields []string
}

type PackageData struct {
	Exec string
	Name string

	Structs []*StructData
}

func getAllFiles(root string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") &&
			!strings.HasSuffix(path, ".gen.go") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	exec := os.Args[0]

	files, err := getAllFiles(".")
	if err != nil {
		logger.Error("Failed to get all files", slog.String("error", err.Error()))
		panic(err)
	}

	g := NewGenerator()

	packages := make(map[string]*PackageData)

	for _, file := range files {

		fileDir := filepath.Dir(file)

		// парсим файл
		packageName, structs, err := parse(g, file)
		if err != nil {
			logger.Error("Failed to parse file", slog.String("file", file), slog.String("error", err.Error()))
			continue
		}

		// файлы без структур скипаем
		if len(structs) == 0 {
			continue
		}

		// если встречаем пакет, который до этого не видели
		if _, ok := packages[fileDir]; !ok {
			packages[fileDir] = &PackageData{
				Exec:    exec,
				Name:    packageName,
				Structs: make([]*StructData, 0, len(structs)),
			}
		}

		packages[fileDir].Structs = append(packages[fileDir].Structs, structs...)
	}

	for dir, pkg := range packages {
		file := filepath.Join(dir, "reset.gen.go")
		buf := g.GenerateFile(pkg)

		err = os.WriteFile(file, buf, 0644)
		if err != nil {
			logger.Error("Failed to write file", slog.String("file", file), slog.String("error", err.Error()))
			continue
		}
		logger.Info("Generated file", slog.String("file", file))
	}
}

// parse перебирает конкретный файл по пути path
func parse(g *Generator, path string) (string, []*StructData, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)

	if err != nil {
		return "", nil, err
	}

	structs := make([]*StructData, 0)

	//перебираем декларации ищем с type
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if gd.Tok != token.TYPE {
			continue
		}

		// ищем метку
		tag := hasTag(gd.Doc)

		// внутри полей type ищем structы
		for _, spec := range gd.Specs {
			tspec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// ищем метку
			if !hasTag(tspec.Doc) && !tag {
				continue
			}

			structType, ok := tspec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			shortName := string(unicode.ToLower(rune(tspec.Name.Name[0])))

			structData := &StructData{
				Name:          tspec.Name.Name,
				ShortName:     shortName,
				NullifyFields: make([]string, 0),
			}

			// перебираем поля структуры, смотрим как их занулять
			for _, field := range structType.Fields.List {

				for _, fname := range field.Names {
					structData.NullifyFields = append(structData.NullifyFields,
						g.GenerateNullifyField(shortName, fname.Name, field.Type))
				}
			}

			structs = append(structs, structData)
		}
	}

	return f.Name.Name, structs, nil
}

func hasTag(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}
	for _, c := range doc.List {
		if strings.Contains(c.Text, MagicWord) {
			return true
		}
	}
	return false
}
