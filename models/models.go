package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	// some comment
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
)

// DBNAME dbnme
// SQLITE3DRIVER driver
const (
	DBNAME        = "data/beeblog.db"
	SQLITE3DRIVER = "sqlite3"
)

// Category define cetetory model
type Category struct {
	ID              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

// Topic define topic model
type Topic struct {
	ID              int64
	UID             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Category        string
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastuserID int64
}

// RegisterDB xxx
func RegisterDB() {
	if !com.IsExist(DBNAME) {
		os.MkdirAll(path.Dir(DBNAME), os.ModePerm)
		os.Create(DBNAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(SQLITE3DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", SQLITE3DRIVER, DBNAME, 10)
}

// AddCategory create category
func AddCategory(name string) error {
	o := orm.NewOrm()

	category := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")

	err := qs.Filter("title", name).One(category)

	if err == nil {
		return err
	}
	_, err = o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCategories get all categories
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	categories := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&categories)
	return categories, err
}

// DelCategory delete category
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	category := &Category{ID: cid}
	_, err = o.Delete(category)
	return err
}

// AddTopic create topic
func AddTopic(title, category, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:     title,
		Content:   content,
		Category:  category,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

// GetAllTopics xxx
func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

// GetTopic xxx
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

// ModifyTopic modify topic
func ModifyTopic(tid, title, category, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Category = category
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}

// DeleteTopic delet topic
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tidNum}
	_, err = o.Delete(topic)
	return err
}
