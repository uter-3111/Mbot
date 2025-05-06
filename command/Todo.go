package command

import (
	"fmt"
	"mbot/model"
)

// TodoCommand 待办事项命令
type TodoCommand struct {
	BaseCommand *model.BaseCommand
}

func (c *TodoCommand) GetName() string {
	return c.BaseCommand.Name
}

func NewTodoCommand() *TodoCommand {
	return &TodoCommand{
		BaseCommand: &model.BaseCommand{
			Name:        "todo",
			Description: "查看待办事项",
		},
	}
}

func (c *TodoCommand) Execute(_ []string) (body string, msgtype string) {
	return fmt.Sprintf(`{"text":"您有3条待办：1. 项目报告 2. 会议预约 3. 客户回访"}`), "text"
}
