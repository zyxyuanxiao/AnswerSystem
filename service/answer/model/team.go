package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Team struct {
	Team_id        int64 `orm:"pk"` //组id
	Team_name      string           //组名
	Refer_event_id int64            //参见的事件id,关联的事件的id
	Team_credit    float64          //本组在事件中的积分
}

func init() {
	orm.RegisterModel(new(Team))
}

func AddTeam(team_name string, refer_event_id int64) int64 {
	//login_name不能重复
	var t Team
	o := orm.NewOrm()

	t.Refer_event_id = refer_event_id
	t.Team_name = team_name
	t.Team_credit = 0
	id, err := o.Insert(&t)
	if err == nil {
		logs.Debug("AddTeam's id", id)
		return id
	} else {
		logs.Error("AddTeam's err", err)
		return -1
	}
}

func GetTeamById(team_id int64, event_id int64) Team {
	var t Team
	o := orm.NewOrm()
	o.QueryTable("team").Filter("team_id", team_id).Filter("Refer_event_id", event_id).One(&t)
	return t
}

func UpdateTeamCredit(team_id int64, team_credit float64) float64 {
	new_credit := 0.0
	//更新
	team := Team{Team_id: team_id}
	o := orm.NewOrm()
	if o.Read(&team) == nil {
		old_credit := team.Team_credit
		new_credit = old_credit + team_credit
		team.Team_credit = new_credit

		if num, err := o.Update(&team, "Team_credit"); err == nil {
			logs.Debug("UpdateTeamCredit's num", num)
		} else if err != nil {
			logs.Error("UpdateTeamCredit's err", err)
		}
	}
	return new_credit
}
