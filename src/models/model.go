package models

import (
	"samh_common_lib/base"
)

type ActivityRequest struct {
	base.SamhBaseRequest
	FetchActivityType FetchActivityTypeCode `form:"fetch_activity_type" json:"fetch_activity_type" binding:"-"`
	ActivityId        int64                 `form:"activity_id" json:"activity_id" binding:"-"`
	ActivityType      ActivityTypeCode      `form:"activity_type" json:"activity_type" binding:"-"`
}

type ActivityResponse struct {
	ActivityArr []*Activity `json:"activity_arr"`
}

type FetchActivityTypeCode int

const (
	FetchActivityTypeCode_All FetchActivityTypeCode = iota
	FetchActivityTypeCode_Id
	FetchActivityTypeCode_Type
)

type ActivityTypeCode int

const (
	ActivityTypeCode_Vip             ActivityTypeCode = 1
	ActivityTypeCode_Daily           ActivityTypeCode = 2
	ActivityTypeCode_ThirdVipRechage ActivityTypeCode = 1000
)

type Activity struct {
	ActivityId   int64  `json:"activity_id" xorm:"not null pk autoincr BIGINT(20)"`
	Type         int    `json:"type" xorm:"default 0 comment('1-会员,2-日常,1000+为三方的，如1000-会员充值') INT(11)"`
	Title        string `json:"title" xorm:"default '' VARCHAR(200)"`
	Content      string `json:"content" xorm:"default '' VARCHAR(1000)"`
	LinkUrl      string `json:"link_url" xorm:"default '活动链接，一般是h5网页' VARCHAR(200)"`
	UseStatus    int    `json:"use_status" xorm:"default 0 comment('0-进行中，1-已下线') INT(11)"`
	CreateTime   int64  `json:"create_time" xorm:"default 0 BIGINT(20)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 0 BIGINT(20)"`
	StartTime    int64  `json:"start_time" xorm:"default 0 BIGINT(20)"`
	EndTime      int64  `json:"end_time" xorm:"default 0 BIGINT(20)"`
	Status       int    `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
	ThirdId      int64  `json:"third_id" xorm:"default 0 comment('来源三方的id，如虚拟商品、good_id') INT(11)"`
	FilterRuleId int    `json:"filter_rule_id" xorm:"default 0 comment('过滤规则ID') INT(11)"`
}

type ShowRequest struct {
	Uid           int64             `form:"uid" json:"uid" binding:"required"`
	DeviceId      string            `form:"udid" json:"udid" binding:"required"`
	FetchShowType FetchShowTypeCode `form:"fetch_show_type" json:"fetch_show_type" binding:"-"`
	ShowType      int               `form:"show_type" json:"show_type" binding:"-"`
}

type FetchShowTypeCode int

const (
	FetchShowTypeCode_All FetchShowTypeCode = iota
	FetchShowTypeCode_ShowType
)

type ShowResponse struct {
	Status base.SamhResponseCode `json:"status"`
	Msg    string                `json:"msg"`
	Data   *struct {
		ShowArr []*Show `json:"show_arr"`
	} `json:"data"`
}

type Show struct {
	ShowId       int64  `json:"show_id" xorm:"not null pk autoincr BIGINT(20)"`
	JumpType     int    `json:"jump_type" xorm:"default 0 comment('1-app自己请求活动信息再请求参加活动,2-跳转h5取link_url，3-跳转app某页') INT(11)"`
	Title        string `json:"title" xorm:"default 'toast用的标题' VARCHAR(200)"`
	Content      string `json:"content" xorm:"default 'toast用的内容' VARCHAR(1000)"`
	ImageUrl     string `json:"image_url" xorm:"default '图标、图片网址' VARCHAR(200)"`
	ShowType     int    `json:"show_type" xorm:"default 0 comment('1-首页公告，2-VIP页公告，3-首页弹窗，4-浮窗') INT(11)"`
	ShowPolicyId int64  `json:"show_policy_id" xorm:"default 0 comment('对应的显示策略ID') BIGINT(20)"`
	OrderSn      int    `json:"order_sn" xorm:"default 0 comment('顺序，小的在前,仅用于同一show_type时排层级，极少用') INT(11)"`
	UseStatus    int    `json:"use_status" xorm:"default 0 comment('0-进行中，1-已下线') INT(11)"`
	CreateTime   int64  `json:"create_time" xorm:"default 0 BIGINT(20)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 0 BIGINT(20)"`
	StartTime    int64  `json:"start_time" xorm:"default 0 BIGINT(20)"`
	EndTime      int64  `json:"end_time" xorm:"default 0 BIGINT(20)"`
	Status       int    `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
	ActivityId   int64  `json:"activity_id" xorm:"default 0 comment('对应的活动ID，活动里有过滤规则，如不能再参与此活动则不此show不返回') INT(11)"`
	//
	ShowPolicyItem *ShowPolicy `json:"show_policy"`
	RewardRuleId   int64       `json:"reward_rule_id"`
	LinkUrl        string      `json:"link_url"`
}

type ShowPolicy struct {
	ShowPolicyId int64  `json:"show_policy_id" xorm:"not null pk autoincr BIGINT(20)"`
	Title        string `json:"title" xorm:"default '标题' VARCHAR(200)"`
	Content      string `json:"content" xorm:"default '内容' VARCHAR(1000)"`
	Type         int    `json:"type" xorm:"default 0 comment('显示类型，1：一直显示，2：秒，3：分，4：时，5：间隔天，根据开始时间定义 如果是10点开始，则是今天10点到明天十点只能显示一次，6：自然天，0点到24点 只显示一次，7：周，8：月，9：启动时显示，次数是下面的number') INT(11)"`
	TimeInterval int64  `json:"time_interval" xorm:"default 0 comment('显示时间间隔,单位根据上面的type可为秒、分、时、天、周、月') BIGINT(20)"`
	Number       int    `json:"number" xorm:"default 0 comment('显示次数，注：无限为2147483647(int32 max)') INT(11)"`
	CreateTime   int64  `json:"create_time" xorm:"default 0 BIGINT(20)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 0 BIGINT(20)"`
	Status       int    `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
}

type JoinActivityRequest struct {
	base.SamhBaseRequest
	ActivityId    int64  `form:"activity_id" json:"activity_id" binding:"required"`
	RewardRuleId  int64  `form:"reward_rule_id" json:"reward_rule_id" binding:"required"`
	PayType       int    `form:"pay_type" json:"pay_type" binding:"-"`
	IsPaypal      int    `form:"ispaypal" json:"ispaypal" binding:"-"` //是必需的，但为0的时候这框架认为没有，故不能用required
	ClientType    string `form:"client-type" json:"client-type" binding:"-"`
	ClientVersion string `form:"client-version" json:"client-version" binding:"-"`
}

type JoinActivityResponse struct {
	Status base.SamhResponseCode `json:"status"`
	Msg    string                `json:"msg"`
	Data   *struct {
		Rewarded       int         `json:"rewarded"`
		RewardSucc     int         `json:"reward_succ"`
		Title          string      `json:"title"`
		Content        string      `json:"content"`
		RewardRuleType int         `json:"reward_rule_type" xorm:"default 0 comment('1-不需要充值什么的，直接请求奖励,2-需要充值之类的，直接请求VIP服务那边并把返回的data放到extra_data里') INT(11)"`
		ExtraData      interface{} `json:"extra_data"`
	} `json:"data"`
}

type VipRechargeRequest struct {
	base.SamhBaseRequest
	ActivityId    int64  `form:"activity_id" json:"activity_id"`
	RewardRuleId  int64  `form:"reward_rule_id" json:"reward_rule_id"`
	ThirdId       int64  `form:"third_id" json:"third_id"`
	PayType       int    `form:"pay_type" json:"pay_type"`
	IsPaypal      int    `form:"ispaypal" json:"ispaypal"`
	ClientType    string `form:"client_type" json:"client_type"`
	ClientVersion string `form:"client_version" json:"client_version"`
}

type VipRechargeResponse struct {
	base.SamhResponse
}
