// Анализатор NoExit проверяет что в функции main() пакета main нет прямого запуска os.Exit
//
// Анализатор допускает запуск os.Exit() вне пакета main или в прочих, отличных от main функциях,
// но не допускает любой запуск os.Exit() внутри main(в том числе в замыканиях)
package noexit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "noexit",
	Doc:  "check for call os.Exit in main function",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, f := range pass.Files {

		// Все файлы корме "main" скипаем
		if f.Name.Name != "main" {
			continue
		}

		ast.Inspect(f, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncDecl:
				// все функции кроме mainа скипаем
				if n.Name.Name != "main" {
					return false
				}
			case *ast.CallExpr:
				// Проверяем что вызванная функция это os.Exit
				if isExitCall(n) {
					pass.Reportf(n.Pos(), "call os.Exit function")
					return false
				}
			}
			return true
		})
	}

	return nil, nil
}

func isExitCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if sel.Sel.Name != "Exit" {
		return false
	}

	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name != "os" {
		return false
	}

	return true
}
