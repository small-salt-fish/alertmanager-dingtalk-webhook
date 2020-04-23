package transformer

import (
	"bytes"
	"fmt"

	"github.com/small-salt-fish/alertmanager-dingtalk-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {

	//groupKey := notification.GroupKey
	status := notification.Status
	labels := notification.CommonLabels


	annotations := notification.CommonAnnotations
	robotURL = annotations["dingtalkRobot"]

	var buffer bytes.Buffer

//	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))
	buffer.WriteString(fmt.Sprintf("### 告警通知: \n"))

//	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
//		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("+ 告警名称: %s\n", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("+ 告警状态: %s\n", status))
		buffer.WriteString(fmt.Sprintf("+ 告警实例: %s\n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("+ 告警级别: %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("+ 告警详情：%s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("+ 告警时间：%s", alert.StartsAt.Format("15:04:05")))
	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
//			Title: fmt.Sprintf("通知组：%s(当前状态:%s)", groupKey, status),
			Title: fmt.Sprintf("告警通知："),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}
