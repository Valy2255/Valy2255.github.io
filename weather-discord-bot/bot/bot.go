package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Valy2255/weather-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string


func messageHandler(s *discordgo.Session,m *discordgo.MessageCreate){
	if m.Author.ID== BotId{
		return
	}
	if strings.HasPrefix(m.Content,"!weather"){
		location:=strings.TrimSpace(strings.TrimPrefix(m.Content,"!weather"))
		if location==""{
			s.ChannelMessageSend(m.ChannelID,"Please specify a location.")
			return
		}
		weather,err:=getWeather(location,config.ApiKey)
		if err!=nil{
			s.ChannelMessageSend(m.ChannelID,"Error getting weather data")
			return
		}
		s.ChannelMessageSend(m.ChannelID,weather)
	}
}

func getWeather(location string,apiKey string) (string,error){
	url:=fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s",location,apiKey)
    resp,err:=http.Get(url)
	if err!=nil{
		return "",err
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		return "", fmt.Errorf("API request failed with status code %d",resp.StatusCode)
	}

	body,err:=io.ReadAll(resp.Body)
	if err!=nil{
		return "",err
	}

	var data map[string]interface{}
	err=json.Unmarshal(body,&data)
	if err!=nil{
		return "",err
	}

	weather:=data["weather"].([]interface{})[0].(map[string]interface{})
	description:=weather["description"].(string)
	temp:=data["main"].(map[string]interface{})["temp"].(float64)
	name:=data["name"].(string)

	return fmt.Sprintf("The weather in %s is %s with a temperature of %.1f°C",name,description,temp),nil

}

func Start(){
	goBot,err:=discordgo.New("Bot "+config.Token)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}

	u,err:=goBot.User("@me")

	if err!=nil{
		fmt.Println(err.Error())
		return
	}

	BotId=u.ID

	goBot.AddHandler(messageHandler)

	err=goBot.Open()

	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running. Press CTRL-C to exit.")

}

