replace into samh_operation.activity(type,title,http://thyrsi.com/t6/659/1547780651x2890174213.png,link_url,end_time) values(1,'会员-一分钱充值分享','http://thyrsi.com/t6/659/1547780651x2890174213.png','link_url',9999999999);
replace into samh_operation.activity(type,title,http://thyrsi.com/t6/659/1547780651x2890174213.png,link_url,start_time,end_time) values(1,'春节一分钱充值','http://image.samh.xndm.tech/operation/activity/group1/image/android/drawable-xhdpi/homepage-floating.png','link_url',1547481600,1548777600);
replace into samh_operation.reward_rule(title,rechange,vip_type,vip_time) values('春节一分钱分享',1,1,259200),('春节一分钱充值奖励规则',1,1,1296000);
replace into samh_operation.reward_tip(reward_tip_id,type,describes,title,content) values(1,0,'春节一分钱充值分享','奖励成功标题','奖励成功内容'),(1,1,'春节一分钱充值分享','已奖励过标题','已奖励过内容');
replace into samh_operation.activity_reward_rule(activity_id,reward_rule_id,reward_tip_id) values(1,1,1),(2,2,1);
replace into samh_operation.float_window(type,title,http://thyrsi.com/t6/659/1547780651x2890174213.png,link_url,end_time,activity_id) values(1,'标题','http://image.samh.xndm.tech/operation/activity/group1/image/android/drawable-xhdpi/homepage-floating.png','link_url',9999999999,1);
replace into samh_operation.activity_group(title,content) values('活动组标题','活动组内容');
replace into samh_operation.activity_group_activity(activity_group_id,activity_id,order_sn) values(1,1,1),(1,2,0);
