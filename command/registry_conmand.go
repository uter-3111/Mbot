package command

import (
	"fmt"
	"mbot/model"
	"strings"
)

// CommandRegistry 命令注册中心
type CommandRegistry struct {
	commands map[string]model.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]model.Command),
	}
}

// Register 动态注册命令
func (cr *CommandRegistry) Register(cmd model.Command) {
	cr.commands[cmd.GetName()] = cmd
}

// ParseAndExecute 解析并执行命令
func (cr *CommandRegistry) ParseAndExecute(input string) (body string, msgtype string) {
	parts := strings.Fields(input)
	if len(parts) <= 1 {
		return fmt.Sprintf(`{"text":"无效指令"}`), "text"
	}
	cmdName := strings.TrimPrefix(parts[1], "/")
	cmd, exists := cr.commands[cmdName]
	if !exists {
		return fmt.Sprintf(`{"text":"未知指令，支持：%v"}`, cr.ListCommands()), "text"
	}

	args := parts[1:]
	return cmd.Execute(args)
}

// ListCommands 列出所有注册命令
func (cr *CommandRegistry) ListCommands() []string {
	var cmds []string
	for name := range cr.commands {
		if name == "stock" {
			continue
		}
		cmds = append(cmds, "/"+name)
	}
	return cmds
}
