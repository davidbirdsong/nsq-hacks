package main

type MsgBody struct {
	Url    string
	Host   string `json:"header_host"`
	Id     string `json:"id_request"`
	Source string `json:"id_source"`
}
