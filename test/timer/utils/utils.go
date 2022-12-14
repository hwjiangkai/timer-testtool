// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

const (
	FormatJSON = "json"
)

func IsFormatJSON(cmd *cobra.Command) bool {
	v, err := cmd.Flags().GetString("format")
	if err != nil {
		return false
	}
	return strings.ToLower(v) == FormatJSON
}

func CmdFailedf(cmd *cobra.Command, format string, a ...interface{}) {
	errStr := format
	if a != nil {
		errStr = fmt.Sprintf(format, a...)
	}
	if IsFormatJSON(cmd) {
		m := map[string]string{"ERROR": errStr}
		data, _ := json.Marshal(m)
		color.Red(string(data))
	} else {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"ERROR"})
		t.AppendRow(table.Row{errStr})
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 2, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		t.SetOutputMirror(os.Stdout)
		t.Render()
	}

	os.Exit(-1)
}

func CmdFailedWithHelpNotice(cmd *cobra.Command, format string) {
	color.White(format)
	color.Cyan("\n============ see below for right usage ============\n\n")
	_ = cmd.Help()
	os.Exit(-1)
}
