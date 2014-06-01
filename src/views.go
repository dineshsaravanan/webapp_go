package main

import (
	"html/template"
	"bytes"
)

const template_base_path = "templates/";

func parseTemplate(filePath string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetPage(filePath string, data interface{}) (out []byte) {
	//TODO: Logic to build different parts of the page depending on the app
	filePath = template_base_path + filePath
	out, err := parseTemplate(filePath, data)
	if err != nil {
		panic("View missing template file in the path " + filePath)
	}
	return out
}

func GetHomeView(data interface{}) string {
	path := "home.html"
	return string(GetPage(path, data));
}

func GetLoginView(data interface{}) string {
	path := "login.html"
	return string(GetPage(path, data));
}








