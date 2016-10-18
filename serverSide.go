package main

import (
    "io/ioutil"
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}


func indexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "GodPage.tmpl" , gin.H{
    		"title" : "Welcome to my Portfolio",
    		})
}

func noRoute(r *gin.Context){
        r.HTML(http.StatusOK, "404.tmpl", gin.H{
                "title" : "Something is not right",
        })
}


func main() {
    // Creates a router without any middleware by default
    r := gin.New()
    r.Use(gin.Logger())

    html := template.Must(template.ParseFiles(
        "./templates/header.tmpl",
        "./templates/menu.tmpl",
        "./templates/404.tmpl",
        "./templates/welcome.tmpl",
        "./templates/GodPage.tmpl",
        "./templates/skills.tmpl",
        "./templates/index.tmpl",
        "./templates/projects.tmpl"))

    r.SetHTMLTemplate(html)

    r.GET("/*action", indexHandler)
    r.NoRoute(noRoute)
    r.Run(":8080")
}