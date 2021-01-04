package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type xlsx struct {
	Name     string
	FileName string
	ClassName string
	SheetName string
	Template *xField
	Data     []*xField
	keymap   map[int]*xField
}

func (x *xlsx) Init(fileName string, name string) {
	x.Name = name
	x.FileName = fileName
	x.Data = make([]*xField, 0)
	x.keymap = make(map[int]*xField)
}

func PrintStringArray (row []string) {
	print("\t")
	for i := 0; i < len(row); i++ {
		fmt.Printf("[%d]%s, ", i, row[i])
	}
	println()
}

func (x *xlsx) Parse(rows [][]string) bool {
	if len(rows) < 3 {
		log.Fatalln("Parse", x.Name, "header rows < 3")
		return false
	}
	if len(rows[2]) < 1 {
		log.Println("Parse", x.Name, "header need ClassName in row[2][0]")
		return false
	}

	x.ClassName = rows[2][0]
	rows[2][0] = ""

	fmt.Println()
	x.Template = new(xField)
	fmt.Printf("---------- FieldName:\n")
	PrintStringArray(rows[0])
	fmt.Printf("---------- FieldType:\n")
	PrintStringArray(rows[1])
	fmt.Printf("---------- Tag:\n")
	PrintStringArray(rows[2])
	fmt.Println()

	if ok, _ := x.Template.Init(x.Name, "struct", ""); ok {
		x.Template.ParseSubFieldsDefs(rows[0], rows[1], rows[2])
		for i := 4; i < len(rows); i++ {
			field := x.Template.Copy()
			// PrintStringArray(rows[i])
			// comment row
			if len(rows[i]) == 0 || strings.HasPrefix(rows[i][0], "//") || rows[i][0] == "" {
				continue
			}
			id, _ := strconv.Atoi(rows[i][0])
			if _, ok2 := x.keymap[id]; !ok2 {
				field.ParseDatas(id, rows[i])
				field.SetLevel(4)
				x.Data = append(x.Data, field)
				x.keymap[id] = field
			} else {
				log.Fatalln("Parse", x.Name, "failed, Id", id, "is duplicated")
			}

		}
		i := 0
		for i < len(x.Template.Fields) {
			v := x.Template.Fields[i]
			if strings.HasPrefix(v.Type, "//") || (v.Tag != "" && v.Tag != params.tag) {
				x.Template.Fields = append(x.Template.Fields[:i], x.Template.Fields[i+1:]...)
			} else {
				i++
			}
		}
	} else {
		log.Fatalln("Parse", x.Name, "head field")
	}

	return true
}

func (x *xlsx) Print() {
	for k, v := range x.Data {
		fmt.Print(k, " ")
		v.Print()
	}
}
