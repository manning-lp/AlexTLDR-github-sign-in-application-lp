package main

type userData struct {
	Login       string
	accessToken string
}

var sessionsStore = make(map[string]userData)
