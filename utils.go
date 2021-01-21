package main

import (
	"errors"
	"strconv"
	"strings"
)

func convertToVertical(data [][]string) [][]string {
	ret := make([][]string, 0)
	for i := 0; i < len(data[0]); i++ {
		row := make([]string, 0)
		row = append(row, data[0][i])
		ret = append(ret, row)
	}
	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			ret[j] = append(ret[j], data[i][j])
		}
	}
	return ret
}

func trimData(data string) string {
	before := data
	for {
		data = strings.TrimPrefix(data, "{")
		data = strings.TrimSuffix(data, "}")
		if before == data {
			return data
		}
		before = data
	}

}

func splitName(data string) []string {
	subDatas := make([]string, 0)
	equal := strings.Index(data, "=")

	var n string
	var d string
	if equal == -1 {
		d = data
	} else {
		n = data[:equal]
		d = data[equal+1:]
	}
	subDatas = append(subDatas, n)
	subDatas = append(subDatas, d)
	return subDatas
}

func splitSubData(layer int, data string) []string {
	sept := ""
	for i := 1; i < layer; i++ {
		sept += "}"
	}
	sept += "|"

	subDatas := make([]string, 0)

	for {
		pos := strings.Index(data, sept)
		if pos == -1 {
			subDatas = append(subDatas, data)
			break
		} else {
			subData := data[0 : pos+layer-1]
			data = data[pos+layer:]
			subDatas = append(subDatas, subData)
		}
	}
	return subDatas
}

// 根据类型处理基础类型数据
func handleData(dataField *xField, data string) error {
	var retErr error
	if dataField.HasDefaultData && len(strings.TrimSpace(data)) == 0 {
		// 定义了默认值的空字段使用默认值
		retErr = nil
		return retErr
	}
	handleType := dataField.Type
	// 处理枚举和自定义类型的解析方式
	if dataField.IsCustomType() {
		if dataField.IsEnum() {
			// 枚举按整型处理
			handleType = "int"
			// 数据支持以#开始的注释
			first := strings.Index(data, "#")
			if first != -1 {
				data = data[:first]
			}
		} else {
			// 自定义结构体，保留原始数据
			dataField.Data = data
			retErr = nil
			return retErr
		}
	}
	switch handleType {
	case "int":
		ret, err := strconv.Atoi(data)
		dataField.Data = strconv.Itoa(ret)
		retErr = err
	case "float":
		ret, err := strconv.ParseFloat(data, 32)
		dataField.Data = strconv.FormatFloat(ret, 'f', 3, 32)
		retErr = err
	case "bool":
		ret, err := strconv.ParseBool(data)
		dataField.Data = strconv.FormatBool(ret)
		retErr = err
	case "string":
		// 字符串基础类型需要转义引号
		dataField.Data = strings.Replace(data, "\"", "\\\"", -1)
		retErr = nil
	default:
		retErr = errors.New("DataType " + dataField.Type + " is invalid for data " + data)
	}
	return retErr
}

func name2lower2Camel(name string) (string, string) {
	dotIndex := strings.LastIndex(name, ".")
	lower := name[:dotIndex]

	initial := strings.ToUpper(lower[0:1])
	other := lower[1:]
	for strings.Index(other, "_") != -1 {
		index := strings.Index(other, "_")
		replace := strings.ToUpper(other[index+1 : index+2])
		s := []string{other[:index], replace, other[index+2:]}
		other = strings.Join(s, "")
	}
	return lower, initial + other
}

func upperInitialChar(str string) string {
	initial := strings.ToUpper(str[0:1])
	other := str[1:]
	return initial + other
}
