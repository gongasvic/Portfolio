package main

import (
    "io/ioutil"
   // "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "html/template"
    "strings"
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
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
    		"title" : "Welcome to my Portfolio",
    		})
}

func viewHandler(r *gin.Context){
	//compare to a certain list of possibilities if it is in display else discard to 404
    title := r.Param("action")[1:]
    if title != "" {
      
        s := []string{title, "tmpl"}
        z := strings.Join(s, ".")
        test, err := r.HTML(http.StatusOK, z, gin.H{
            "title" : title,
            })
        if err != nil {
            noRoute(r)
        }
        defer test.Body.Close()
      
    }else{
        r.HTML(http.StatusOK, "index.tmpl", gin.H{
                "title" : "index",
                })
    }
}

func noRoute(r *gin.Context){
        r.HTML(http.StatusOK, "404.tmpl", gin.H{
                "title" : "Something is not right",
        })
}

func setTemplates(){
    template.Must(template.ParseFiles(
        "./templates/index.tmpl",
        "./templates/404.tmpl",
        "./templates/view/index.tmpl",
        "./templates/view/about.tmpl"))
}


func main() {
    // Creates a router without any middleware by default
    r := gin.New()
    r.Use(gin.Logger())

    html := setTemplates()
    r.SetHTMLTemplate(html)

    r.GET("/", indexHandler)
    r.GET("/view/*action", viewHandler)
    r.NoRoute(noRoute)
    r.Run(":8080")
}