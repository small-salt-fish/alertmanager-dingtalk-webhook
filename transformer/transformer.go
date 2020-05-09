package transformer

import (
	"bytes"
	"fmt"
	"reflect"
	"github.com/small-salt-fish/alertmanager-dingtalk-webhook/model"
)

func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {

	var status = notification.Status
	labels := notification.CommonLabels
	url := notification.ExternalURL
	
	if status == "firing"{
		status = "<font color=#FF0000>告警</font>"
	} else {
		status = "<font color=#008B00>恢复</font>"
	}


	annotations := notification.CommonAnnotations
	robotURL = annotations["dingtalkRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 【%s】 %s 告警通知: \n",status,labels["project"]))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("> 告警名称: %s \n", annotations["summary"]))

		buffer.WriteString(fmt.Sprintf("\n > 告警状态: %s \n", status))
		buffer.WriteString(fmt.Sprintf("\n > 告警实例: %s \n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n > 告警级别: %s \n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n > 告警时间: %s \n", alert.StartsAt.Format("15:04:05")))
		buffer.WriteString(fmt.Sprintf("\n > 告警详情: %s \n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n > 告警地址: [链接详情](%s)\n", url))
	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: fmt.Sprintf(labels["project"]+" "+annotations["summary"]),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}
	println(reflect.TypeOf(markdown).Name())

	return
}

