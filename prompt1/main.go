package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"prompt/utils/gpt"
)

const videoText = `
	哀牢山连国家地质科考队都有进无出的地方，他却独自一人两进两出。
	起因是一家名为山区的公司，为追求颜料的极致，进入哀牢山寻找一级矿，
	刚到山脚就已经迷雾重重，然后就是有熊出没，进山路只能按两年前探过的地方走，
	不然很容易迷失方向。刚进山就捡到巨大鸟毛，这可能是白腹锦鸡。深山野林里最害怕的是遇到熊，
	不是熊死，就是我活。徒步一个小时后来到曾经让他止步的地方。
	一条深不见底的裂缝河流在二零二二年时尝试用五十米登山绳捆住石头扔进河里依旧探不到底，
	于是制作简易木筏尝试过河，在木筏前进不到三十米后崩溃散架无法前进，最终狼狈无功而返。
	博主推测河堤可能是祖母绿矿脉，而哀牢山是传说中的哀牢。千年的演变，也埋藏着许多不为人知的东西。有了上次的经验，
	这次博主带了只皮艇在打气途中感觉身后河里有东西突然翻滚，推测是一条非常大的鱼，于是拿出鱼钩准备来一杆，事实究竟是何物？
	不出所料，是一条不知活了多少年的大鱼。随着划船前进，河里突然开始起雾了，没有任何的前兆，一种不好的预感油然而生，水里有东西随着咕噜咕噜的声音越来越近，
	不像是鱼类，而且水里的不明物动作很轻，体型非常非常大比刚来钓上来的鱼大十倍
	，一直跟随着小船要赶紧上岸了上岸后观察许久不明物才退走无奈只能徒步前行在确认方向后还要约二十分钟才能到在穿越斜坡密林后最终抵达目的地就是这种柔蓝色的矿石，
	竟不惜性命也要勇闯哀牢国。如果继续深入，就是真的出不来了。
`
const prompt1 = `Help me extract the video category and give it to me in Chinese`
const prompt2 = `Extract 3 to 4 keywords from this video text about to help with video classification.`
const prompt3 = `
	You will be provided with a block of text which is the content of a video, and your task is to give 5 keyword in Simplified Chinese to the video to attract audience.
	For example, 美食 | 旅行 | 阅读
`

type EnvCfg struct {
	ApiKey  string `env:"ApiKey"`
	ModelID string `env:"ModelID"`
}

var cfg EnvCfg

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("read .env file failed")
	}
	if err := env.Parse(&cfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}

}
func main() {
	fmt.Println(cfg)
	newGpt := gpt.NewGpt(cfg.ApiKey, cfg.ModelID)
	res, err := newGpt.ChatWithModel(context.TODO(), videoText, prompt3)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
