package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type xField struct {
	ID             int
	Name           string
	FullName       string
	Type           string
	LongType       string
	Tag            string
	Data           string
	Count          int
	Size           int
	Layer          int
	Level          int
	Template       *xField
	Fields         []*xField
	ParentField    *xField
	HasDefaultData bool
}

// 是否自定义类型
func (f *xField) IsCustomType() bool {
	return strings.HasPrefix(f.Type, "PB") || strings.HasPrefix(f.Type, "E")
}

func (f *xField) IsEnum() bool {
	return strings.HasPrefix(f.Type, "E")
}

// 初始化单个字段定义
func (f *xField) Init(name string, definition string, tag string) (bool, int) {
	if name == "" && definition == "" && f.Name == "" && f.Type == "" {
		return false, -1
	}

	if name != "" {
		f.Name = name
		if f.ParentField == nil {
			f.FullName = name
		} else {
			f.FullName = f.ParentField.FullName + "." + name
		}
	}
	if definition != "" {
		f.Size = 1
		hasSubField, subFieldDef := f.parseDefinition(definition, tag)
		if hasSubField {
			f.Fields = make([]*xField, 0)
		}
		if subFieldDef != "" {
			f.Template = new(xField)
			f.Template.ParentField = f
			if ok, layer := f.Template.Init("", subFieldDef, tag); ok {
				f.Size = f.Template.Size*f.Count + 1
				f.Layer = layer + 1
			}
		}
	}
	return true, f.Layer
}

func (f *xField) ParseSubFieldsDefs(names []string, defs []string, tags []string) {
	subFieldIndex := 1
	for i := 0; i < len(names); {
		subFieldName := names[i]
		// # 作为结尾
		if strings.HasPrefix(strings.TrimSpace(subFieldName), "#") {
			break
		}
		if f.Template == nil {
			f.Template = new(xField)
			f.Template.Size = 1
		}
		f.Template.ParentField = f
		field := f.Template.Copy()

		if f.Type == "list" {
			subFieldName = strconv.Itoa(subFieldIndex)
		}
		if ok, _ := field.Init(subFieldName, defs[i], tags[i]); ok {
			num := field.Size
			if num > 1 {
				field.ParseSubFieldsDefs(names[i+1:i+num], defs[i+1:i+num], tags[i+1:i+num])
			}
			f.Fields = append(f.Fields, field)

			i += num
		} else {
			i++
		}
		subFieldIndex++
	}
}

// 解析一条数据
func (f *xField) ParseDatas(id int, datas []string) error {
	data := strings.TrimSpace(datas[0])
	if strings.ToLower(data) == "nil" || strings.ToLower(data) == "null" {
		return errors.New("this field is null")
	}
	if strings.HasPrefix(f.Type, "//") {
		return errors.New("this field is comment")
	}
	if f.Tag != "" && f.Tag != params.tag {
		return errors.New("unexported tag")
	}
	f.ID = id
	if f.ParentField != nil && f.ParentField.Type == "dict" && strings.TrimSpace(f.Name) == "" {
		nameData := splitName(data)
		f.Name = nameData[0]
		f.FullName = f.ParentField.FullName + "." + f.Name
		data = nameData[1]
	}

	if f.Count == 0 {
		// 单表格嵌套数据
		data = strings.TrimSpace(data)
		subDatas := splitSubData(f.Layer, data)

		f.setSubFieldsData(subDatas)
	} else if f.Count == 1 {
		// 单表格基础类型数据
		if err := handleData(f, data); err == nil {

		} else {
			log.Fatalln("[", err, "] in field", f.FullName, "of data id", f.ID)
		}
	} else {
		// 多表格嵌套数据
		f.setSubFieldsData(datas)
	}
	return nil
}

func (f *xField) setSubFieldsData(data []string) {
	fieldNum := 0
	var offset int
	if f.Count == -1 || f.Count == 0 {
		offset = 0
	} else {
		offset = 1
	}

	for i := offset; i < len(data); {
		if len(f.Fields) <= fieldNum {
			if data[i] == "" {
				i++
				continue
			}
			subField := f.Template.Copy()
			subField.ParentField = f
			f.Fields = append(f.Fields, subField)
		}
		size := f.Fields[fieldNum].Size
		subdata := data[i : i+size]
		if f.Type == "list" {
			f.Fields[fieldNum].Name = strconv.Itoa(fieldNum)
			f.Fields[fieldNum].FullName = f.FullName + "." + strconv.Itoa(fieldNum)
		}
		if err := f.Fields[fieldNum].ParseDatas(f.ID, subdata); err != nil {
			f.Fields = append(f.Fields[:fieldNum], f.Fields[fieldNum+1:]...)
			fieldNum--
		}
		i += size
		fieldNum++
	}
}

// 解析字段的类型定义
func (f *xField) parseDefinition(def string, tag string) (bool, string) {
	first := strings.Index(def, "<")
	last := strings.LastIndex(def, ">:")
	if first != -1 && last != -1 {
		// list,dict类型
		if count, err := strconv.Atoi(def[last+2:]); err == nil {
			f.Type = def[:first] //list,dcit
			f.LongType = def[:last+1] //list<XX>,dict<XX>
			f.Count = count
			f.Tag = tag
		}
		return true, def[first+1 : last] // XX
	}

	// 处理基础类型的默认值
	f.HasDefaultData = false
	first = strings.Index(def, ",")
	if first != -1 {
		// 配置了默认值
		f.Type = strings.TrimSpace(def[:first])
		// 读取默认值
		err := handleData(f, def[first+1:])
		if err != nil {
			log.Fatalln("[", err, "] in field", f.FullName, "of default value")
			return false, ""
		}
		log.Println("[", f.Data, "] in field", f.FullName, "of default value")
		f.HasDefaultData = true
		f.LongType = f.Type
	} else {
		f.Type = strings.TrimSpace(def)
		f.LongType = f.Type
	}

	if strings.HasPrefix(f.Type, "E") {
		f.LongType = "enum"
	}

	f.Count = 1
	f.Size = 1
	f.Tag = tag
	if def == "struct" {
		f.Count = -1
		f.Size = -1
		f.Tag = ""
		return true, ""
	}

	return false, ""
}

func (f *xField) Copy() *xField {
	field := new(xField)
	field.ID = f.ID
	field.Name = f.Name
	field.FullName = f.FullName
	field.Tag = f.Tag
	field.Type = f.Type
	field.LongType = f.LongType
	field.Data = f.Data
	field.Count = f.Count
	field.Size = f.Size
	field.Layer = f.Layer
	field.Level = f.Level
	field.ParentField = f.ParentField
	field.HasDefaultData = f.HasDefaultData
	if f.Template != nil {
		field.Template = f.Template.Copy()
	}
	if f.Fields != nil {
		field.Fields = make([]*xField, 0)
		for i := 0; i < len(f.Fields); i++ {
			field.Fields = append(field.Fields, f.Fields[i].Copy())
		}
	}
	return field
}

func (f *xField) SetLevel(level int) {
	f.Level = level
	if f.Fields != nil {
		for j := 0; j < len(f.Fields); j++ {
			f.Fields[j].SetLevel(level + 4)
		}
	}
}

func (f *xField) Print() {
	for i := 0; i < f.Level; i++ {
		fmt.Print(" ")
	}
	fmt.Print(" name = ", f.FullName, " | type = ", f.Type, " | long type = ", f.LongType)
	if f.Data != "" {
		fmt.Print(" | data = ", f.Data)
	}
	fmt.Print("\n")
	if f.Fields != nil {
		for j := 0; j < len(f.Fields); j++ {
			f.Fields[j].Print()
		}
	}
}
