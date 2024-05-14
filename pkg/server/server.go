package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Title string
}

// NewRouter creates a new Gin router with predefined routes and configurations.
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Serve static files
	r.Static("/assets", filepath.Join("..", "..", "assets"))

	// Load templates
	templatePattern := filepath.Join("..", "..", "templates", "*.html")
	fmt.Println("Loading templates from:", templatePattern)
	r.LoadHTMLGlob(templatePattern)

	// Define routes
	r.GET("/", HomeHandler)
	r.GET("/about", AboutHandler)
	r.GET("/education", EducationHandler)
	r.GET("/skills", SkillsHandler)
	r.GET("/experience", ExperienceHandler)
	r.GET("/contact", ContactHandler)

	return r
}

// HomeHandler handles requests to the home page.
func HomeHandler(c *gin.Context) {
	renderTemplate(c, "home.html", "Home")
}

// AboutHandler handles requests to the about page.
func AboutHandler(c *gin.Context) {
	fmt.Println("ABOUT")
	renderTemplate(c, "about.html", "About Me")
}

// EducationHandler handles requests to the education page.
func EducationHandler(c *gin.Context) {
	renderTemplate(c, "education.html", "Education")
}

// SkillsHandler handles requests to the skills page.
func SkillsHandler(c *gin.Context) {
	renderTemplate(c, "skills.html", "Skills")
}

// ExperienceHandler handles requests to the experience page.
func ExperienceHandler(c *gin.Context) {
	renderTemplate(c, "experience.html", "Experience")
}

// ContactHandler handles requests to the contact page.
func ContactHandler(c *gin.Context) {
	renderTemplate(c, "contact.html", "Contact")
}

// renderTemplate is a helper function to render HTML templates with page data.
func renderTemplate(c *gin.Context, templateName string, title string) {
	c.HTML(http.StatusOK, templateName, gin.H{
		"title": title,
	})
}
