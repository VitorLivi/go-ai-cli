package data

type Model struct {
	Id string
}

var GeminiModels = []Model{
	{Id: "gemini-2.5-pro"},
	{Id: "gemini-2.5-flash"},
	{Id: "gemini-1.5-pro"},
	{Id: "gemini-1.5-flash"},
	{Id: "gemini-pro"},
}

var GPTModels = []Model{
	{Id: "gpt-4"},
	{Id: "gpt-3.5-turbo"},
	{Id: "gpt-3.5-turbo-16k"},
}
