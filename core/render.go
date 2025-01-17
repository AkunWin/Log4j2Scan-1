package core

import (
	"bytes"
	"fmt"
	"github.com/KpLi0rn/Log4j2Scan/config"
	"github.com/KpLi0rn/Log4j2Scan/model"
)

func RenderHtml(resultList []*model.Result) []byte {
	if len(resultList) == 0 {
		return []byte("no result")
	}
	data := bytes.Buffer{}
	data.Write([]byte(config.TemplatePrefix))
	for i := 1; i < len(resultList)+1; i++ {
		data.Write([]byte(fmt.Sprintf("<tr>\n        "+
			"<th scope=\"row\">%d</th>\n        "+
			"<td>%s</td>\n        "+
			"<td><span class=\"badge badge-danger\">%s</span></td>\n        "+
			"<td><span class=\"badge badge-warning\">%s</span></td>\n    "+
			"</tr>", i, (resultList[i-1]).Host, (resultList[i-1]).Name, (resultList[i-1]).Finger)))
	}
	data.Write([]byte(config.TemplateSuffix))
	return data.Bytes()
}
