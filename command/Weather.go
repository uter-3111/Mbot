package command

import (
	"fmt"
	"mbot/model"
)

type WeatherCommand struct {
	BaseCommand *model.BaseCommand
}

func (c *WeatherCommand) GetName() string {
	return c.BaseCommand.Name
}

func NewWeatherCommand() *WeatherCommand {
	return &WeatherCommand{
		BaseCommand: &model.BaseCommand{
			Name:        "weather",
			Description: "查询天气（格式：/天气 <城市>）",
		},
	}
}

func (c *WeatherCommand) Execute(args []string) (body string, msgtype string) {
	city := "北京" // 默认值
	if len(args) > 1 {
		city = args[1]
	}
	return fmt.Sprintf(`{"text":"%s今日晴，气温15-25℃""}`, city), "text"
}
