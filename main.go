package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/wertick01/dclib/internals/app"
	"github.com/wertick01/dclib/internals/cfg"
)

func main() {
	config, err := cfg.GetConfig() //грузим конфигурацию
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст для работы контекстнозависимых частей системы

	c := make(chan os.Signal, 1)   //создаем канал для сигналов системы
	signal.Notify(c, os.Interrupt) //в случае поступления сигнала завершения - уведомляем наш канал

	server := app.NewServer(*config, ctx) // создаем сервер

	go func() { //горутина для ловли сообщений системы
		oscall := <-c //если таки что то пришло
		log.Printf("system call:%+v\n", oscall)
		server.Shutdown() //выключаем сервер
		cancel()          //отменяем контекст
	}()

	server.Serve()
}
