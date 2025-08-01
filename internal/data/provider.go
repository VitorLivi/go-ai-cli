package data

type Provider struct {
	Id     string
	Name   string
	Models []Model
}

var Providers = []Provider{
	{Id: "gemini", Name: "Gemini", Models: GeminiModels},
	{Id: "gpt", Name: "Chat GPT", Models: GPTModels},
}
