package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type parameters struct {
	excelSrc   string
	luaPath    string
	jsonPath   string
	cshapPath  string
	golangPath string
	tag        string
}

var params *parameters

type fileXlsx struct {
	file *excelize.File
	xl   *xlsx
}

var loadedFiles map[string]*fileXlsx

func init() {
	params = new(parameters)
	flag.StringVar(&params.excelSrc, "p", "", "[Required] Relative `path` of excel files folder")

	flag.StringVar(&params.tag, "tag", "", "only field with this tag or empty string will be exported")
	flag.StringVar(&params.luaPath, "lua", "", "path to place exported .lua files, export no .lua files if parameter is missing")
	flag.StringVar(&params.jsonPath, "json", "", "path to place exported .json files, export no .json files if parameter is missing")
	flag.StringVar(&params.cshapPath, "csharp", "", "path to place exported .cs class files, export no .cs files if parameter is missing")
	flag.StringVar(&params.golangPath, "golang", "", "path to place exported .go struct files, export no .go files if parameter is missing")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "test" {
		params.excelSrc = "./excels/"
		params.tag = "client"
		params.luaPath = "./exports/"
		params.jsonPath = "./exports/"
		params.cshapPath = "./exports/"
		params.golangPath = "./exports/"
	} else if params.excelSrc == "" || (params.luaPath == "" && params.jsonPath == "" && params.cshapPath == "" && params.golangPath == "") || flag.Arg(0) == "help" {
		fmt.Println("Usage: xlsxporter -p <path> [-lua=<luaExportPath>] [-json=<jsonExportPath>] [-csharp=<csharpExportPath>] [-golang=<golangExportPath>] [-tag=<tag>] ")
		fmt.Println("       xlsxporter (help|test)")
		flag.PrintDefaults()
		return
	}

	loadedFiles = make(map[string]*fileXlsx)

	err := filepath.Walk(params.excelSrc, loadFile)
	if err != nil {
		log.Fatalln(err)
	}

	for key, value := range loadedFiles {
		fmt.Print("============================================================\n")
		fmt.Printf("Parse and export file [ %s ] ", key)

		value.xl = parseFile(key, value.file)
		exportFile(value.xl)
		fmt.Printf("Success! [ %s ] - < %s >\n\n", key, value.xl.SheetName)
	}

	//x.Print()
	return
}

func loadFile(path string, f os.FileInfo, err error) error {
	if f == nil {
		log.Fatalln(err)
	}
	if f.IsDir() {
		return nil
	}
	xl, err := excelize.OpenFile(path)
	if err != nil {
		return nil
	}

	loadedFiles[f.Name()] = &fileXlsx{xl, nil}

	return err
}

func parseFile(fileName string, file *excelize.File) *xlsx {
	sheetName := file.GetSheetName(1)

	data := file.GetRows(sheetName)
	if sheetName == "Vertical" {
		data = convertToVertical(data)
	}

	x := new(xlsx)
	x.SheetName = sheetName

	lower, camel := name2lower2Camel(fileName)
	x.Init(lower, camel)
	x.Parse(data)
	return x
}

func exportFile(x *xlsx) {
	e := new(exporter)
	e.Init()

	if params.luaPath != "" {
		e.ExportLua(params.luaPath, x)

	}
	if params.jsonPath != "" {
		e.ExportJSON(params.jsonPath, x)

	}
	if params.cshapPath != "" {
		e.ExportCSharp(params.cshapPath, x)

	}
	if params.golangPath != "" {
		e.ExportGolang(params.golangPath, x)

	}
}
