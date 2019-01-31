create database samh_operation;
use samh_operation;

DROP TABLE IF EXISTS samh_operation.float_window;
CREATE TABLE samh_operation.float_window (
  float_window_id BIGINT NOT NULL AUTO_INCREMENT,
  type INT DEFAULT 0 comment '1-首页',
  title VARCHAR(200) DEFAULT '',
  icon_url VARCHAR(200) DEFAULT '',
  link_url VARCHAR(200) DEFAULT '',
  use_status INT DEFAULT 0 comment '0-进行中，1-已下线',
  create_time BIGINT default 0,
  update_time BIGINT default 0,
  start_time BIGINT default 0,
  end_time BIGINT default 0,
  status INT DEFAULT 0 comment '0-正常，1-已删除',
  PRIMARY KEY (float_window_id),
  activity_id BIGINT DEFAULT 0 comment '活动ID'
) default charset=utf8mb4 comment '浮窗表';

DROP TABLE IF EXISTS samh_operation.activity;
CREATE TABLE samh_operation.activity (
  activity_id BIGINT NOT NULL AUTO_INCREMENT,
  type INT DEFAULT 0 comment '1-会员,1000-三方的，如虚拟商品',
  title VARCHAR(200) DEFAULT '',
  content VARCHAR(1000) DEFAULT '',
  icon_url VARCHAR(200) DEFAULT '',
  link_url VARCHAR(200) DEFAULT '',
  use_status INT DEFAULT 0 comment '0-进行中，1-已下线',
  create_time BIGINT default 0,
  update_time BIGINT default 0,
  start_time BIGINT default 0,
  end_time BIGINT default 0,
  status INT DEFAULT 0 comment '0-正常，1-已删除',
  PRIMARY KEY (activity_id),
  third_id INT default 0 comment '来源三方的id，如虚拟商品'
) default charset=utf8mb4 comment '活动表';

DROP TABLE IF EXISTS samh_operation.reward_rule;
CREATE TABLE samh_operation.reward_rule (
  reward_rule_id BIGINT NOT NULL AUTO_INCREMENT,
  title VARCHAR(200) DEFAULT '',
  content VARCHAR(1000) DEFAULT '',
  rechange INT default 0 comment '单位是分，充值多少',
  vip_type INT default 0 comment '1-普通，2-白金，3-钻石',
  vip_time BIGINT default 0 comment '单位是s',
  exp INT default 0,
  coin INT default 0,
  create_time BIGINT default 0,
  update_time BIGINT default 0,
  status INT DEFAULT 0 comment '0-正常，1-已删除',
  PRIMARY KEY (reward_rule_id)
) default charset=utf8mb4 comment '活动奖励规则表';

DROP TABLE IF EXISTS samh_operation.reward_tip;
CREATE TABLE samh_operation.reward_tip (
  reward_tip_id BIGINT NOT NULL DEFAULT 0,
  type INT DEFAULT 0 comment '0-奖励成功,1-已奖励过',
  describes VARCHAR(200) DEFAULT '' comment '描述本提示属于的活动等',
  title VARCHAR(200) DEFAULT '' comment '提示标题',
  content VARCHAR(1000) DEFAULT '' comment '提示内容',
  status INT DEFAULT 0 comment '0-正常，1-已删除',
  PRIMARY KEY (reward_tip_id,type)
) default charset=utf8mb4 comment '活动奖励规则表';

--  关联表的就可以直接delete掉，而不是更新一个状态表示已删除，不然下面的唯一键就用不了，不能一直删并且会把status加入索引浪费性能
DROP TABLE IF EXISTS samh_operation.activity_reward_rule;
CREATE TABLE samh_operation.activity_reward_rule (
  activity_reward_rule_id BIGINT NOT NULL AUTO_INCREMENT,
  activity_id BIGINT NOT NULL DEFAULT 0 comment '活动ID',
  reward_rule_id BIGINT NOT NULL DEFAULT 0,
  reward_tip_id BIGINT default 0 comment '奖励提示id',
  PRIMARY KEY (activity_reward_rule_id),
  UNIQUE KEY (activity_id,reward_rule_id)
) default charset=utf8mb4 comment '活动跟奖励规则对应表';

DROP TABLE IF EXISTS samh_operation.reward_record;
CREATE TABLE samh_operation.reward_record (
  reward_record_id BIGINT NOT NULL AUTO_INCREMENT,
  uid BIGINT NOT NULL DEFAULT 0,
  udid VARCHAR(50) NOT NULL DEFAULT '',
  activity_id BIGINT DEFAULT 0 comment '活动ID',
  reward_rule_id BIGINT DEFAULT 0 comment '活动奖励规则ID',
  create_time BIGINT default 0,
  update_time BIGINT default 0,
  status INT DEFAULT 0 comment '0-奖励失败,1-奖励成功',
  PRIMARY KEY (reward_record_id)
) default charset=utf8mb4 comment '奖励纪录表';

DROP TABLE IF EXISTS samh_operation.activity_group;
CREATE TABLE samh_operation.activity_group (
  activity_group_id BIGINT NOT NULL AUTO_INCREMENT comment '活动组ID',
  title VARCHAR(200) DEFAULT '',
  content VARCHAR(1000) DEFAULT '',
  create_time BIGINT default 0,
  update_time BIGINT default 0,
  status INT DEFAULT 0 comment '0-正常，1-已删除',
  PRIMARY KEY (activity_group_id)
) default charset=utf8mb4 comment '活动组表';

DROP TABLE IF EXISTS samh_operation.activity_group_activity;
CREATE TABLE samh_operation.activity_group_activity (
  activity_group_activity_id BIGINT NOT NULL AUTO_INCREMENT,
  activity_group_id BIGINT NOT NULL DEFAULT 0 comment '活动组ID',
  activity_id BIGINT NOT NULL DEFAULT 0 comment '活动ID',
  order_sn INT DEFAULT 0 comment '顺序，小的在前',
  PRIMARY KEY (activity_group_activity_id),
  UNIQUE KEY (activity_group_id,activity_id)
) default charset=utf8mb4 comment '活动组跟活动关联表';
