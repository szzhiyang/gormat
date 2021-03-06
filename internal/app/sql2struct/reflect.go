/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : reflect
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"encoding/json"
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"gormat/configs"
	"gormat/pkg/Sql2struct"
	"strings"
)

func Reflect(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	dataType := widget.NewMultiLineEntry()
	dataType.SetText(strings.ReplaceAll(options.Reflect, ",", ",\n"))
	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.Reflect = strings.ReplaceAll(dataType.Text, ",\n", ",")
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.Json, jsons, "sql2struct"); err == nil {
				configs.Json = data
				dialog.ShowInformation("成功", "保存成功", win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: "数据类型转换", Widget: dataType},
		},
	}
}
