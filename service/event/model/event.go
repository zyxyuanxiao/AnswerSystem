package model

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Event struct {
	Event_id          int64 `orm:"pk"`
	Manage_id         int
	Event_title       string
	Event_description string
	Event_time        string
	Event_num         string
	Event_type        string
	Problem_random    bool
	Option_random     bool
	Answer_time       string
	Credit_rule       string
	Participant_num   int
}

type CreditRule struct {
	Single_score    float64 `json:"single_score,string"`    //单选题分数
	Multi_score     float64 `json:"multi_score,string"`     //多选题分数
	Fill_score      float64 `json:"fill_score,string"`      //填空题分数
	Judge_score     float64 `json:"judge_score,string"`     //判断题分数
	Person_score    float64 `json:"person_score,string"`    //当日本人全对额外加分
	Team_score      float64 `json:"team_score,string"`      //当日团队全对额外加分
	Team_score_up   float64 `json:"team_score_up,string"`   //团队总积分上限
	Person_score_up float64 `json:"person_score_up,string"` //个人总积分上限
}

type ProblemNum struct {
	Single   int `json:"single,string"`   //单选题数量
	Multiple int `json:"multiple,string"` //多选题数量
	Fill     int `json:"fill,string"`     //填空题数量
	Judge    int `json:"judge,string"`    //判断题数量
}

type EventTime struct {
	Start_time string `json:"start_time"` //包括开始时间
	End_time   string `json:"end_time"`   //结束时间
	Answer_day string `json:"answer_day"` //答题时间,0,1,2,3,4,5,6
}

func init() {
	orm.RegisterModel(new(Event)) // 注册模型，建立User类型对象，注册模型时，需要引入包
}

func GetEventByEventId(event_id int64) (event *Event) {
	e := Event{Event_id: event_id}
	o := orm.NewOrm()
	err := o.Read(&e, "Event_id")
	if err != nil {
		return nil
	} else {
		return &e
	}
}

func GetCreditRuleByEventId(event_id int) CreditRule {
	var params Event
	var creditRule CreditRule
	o := orm.NewOrm()
	o.QueryTable("event").Filter("Event_id", event_id).One(&params, "Credit_rule")
	if params.Credit_rule != "" {
		err := json.Unmarshal([]byte(params.Credit_rule), &creditRule)
		if err != nil {
			logs.Error("GetCreditRuleByEventId's err:", err)
		}
	}
	return creditRule
}

func GetProblemNumByEventId(event_id int) ProblemNum {
	var params Event
	var problemNum ProblemNum
	o := orm.NewOrm()
	o.QueryTable("event").Filter("Event_id", event_id).One(&params, "event_num")
	if params.Event_num != "" {
		err := json.Unmarshal([]byte(params.Event_num), &problemNum)
		if err != nil {
			logs.Error("GetProblemNumByEventId's err:", err)
		}
	}
	return problemNum
}

func GetEventListByManageIdAndOffst(manage_id int64, offset int, limit int) []Event {
	var e []Event
	o := orm.NewOrm()
	offset = offset - 1
	o.QueryTable("event").Filter("manage_id", manage_id).Offset(offset * limit).Limit(limit).All(&e, "Event_id", "Event_title", "Event_description", "Event_type")
	return e
}

func AddNewEvent(e Event) (string, int64) {
	o := orm.NewOrm()
	id, err := o.Insert(&e)
	if err != nil {
		logs.Error("AddNewEvent's err", err)
		return "faild", 0
	} else {
		return "success", id
	}
}
