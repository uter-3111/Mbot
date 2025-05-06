package command

var Registry *CommandRegistry

func init() {
	Registry = NewCommandRegistry()

	// 注册命令
	Registry.Register(NewWeatherCommand())
	Registry.Register(NewTodoCommand())
	Registry.Register(NewStockCommand())
}
