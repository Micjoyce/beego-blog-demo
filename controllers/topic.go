package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

// TopicController xxx
type TopicController struct {
	beego.Controller
}

// Get xxx
func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}

// Post create topic
func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 301)
}

// Add xxx
func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}
