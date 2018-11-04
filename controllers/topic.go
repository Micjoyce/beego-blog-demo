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
	tid := c.Input().Get("tid")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 301)
}

// Add xxx
func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}

// View xxx
func (c *TopicController) View() {
	topic, err := models.GetTopic(c.Ctx.Input.Params()["0"])
	c.TplName = "topic_view.html"
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Params()["0"]
}

// Modify modify article
func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)

	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

// Delete delete topic
func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(c.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}
