-- ----------------------------
-- Table structure for sch_visit
-- ----------------------------
DROP TABLE IF EXISTS `sch_visit`;
CREATE TABLE `sch_visit` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `visit_no` varchar(64) NULL DEFAULT NULL COMMENT '拜访编号',
  `visit_theme` varchar(200) NULL DEFAULT NULL COMMENT '拜访主题',
  `visit_type` varchar(64) NULL DEFAULT NULL COMMENT '拜访类型',
  `visit_address` text NULL COMMENT '拜访地址',
  `visit_address_echo` text NULL COMMENT '拜访地址（回显）',
  `start_time` datetime NULL DEFAULT NULL COMMENT '预计开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '预计结束时间',
  `visitor` bigint NULL DEFAULT NULL COMMENT '拜访人员',
  `follow_up_personnel` text NULL COMMENT '随访人员',
  `visit_reason` text NULL COMMENT '拜访事由',
  `visit_result` text NULL COMMENT '拜访结果',
  `visit_status` varchar(64) NULL DEFAULT NULL COMMENT '拜访日程状态',
  `effective_status` varchar(64) NULL DEFAULT NULL COMMENT '生效状态',
  `schedule_type` varchar(64) NULL DEFAULT NULL COMMENT '拜访日程类型',
  `visit_plan_template` varchar(64) NULL DEFAULT NULL COMMENT '拜访模板',
  `customer_id` bigint NULL DEFAULT NULL COMMENT '拜访客户对象',
  `contact_id` bigint NULL DEFAULT NULL COMMENT '拜访联系人对象',
  `leads_id` bigint NULL DEFAULT NULL COMMENT '拜访线索对象',
  `business_id` bigint NULL DEFAULT NULL COMMENT '拜访商机对象',
  `related_partner` bigint NULL DEFAULT NULL COMMENT '关联合作伙伴',
  `visit_plan_id` bigint NULL DEFAULT NULL COMMENT '拜访计划（关联）',
  `follow_ids` text NULL COMMENT '关联随访日程',
  `visit_obj` text NULL COMMENT '拜访对象(JSON)',
  `need_visits` tinyint(1) NULL DEFAULT 0 COMMENT '是否需要售前陪同拜访',
  `pre_sales_manager` bigint NULL DEFAULT NULL COMMENT '售前负责人',
  `transfer_reason` text NULL COMMENT '移交原因说明',
  `urgency` varchar(64) NULL DEFAULT NULL COMMENT '紧急程度',
  `importance` varchar(64) NULL DEFAULT NULL COMMENT '重要性',
  `project_stage` varchar(64) NULL DEFAULT NULL COMMENT '项目阶段',
  `lead_stage` varchar(64) NULL DEFAULT NULL COMMENT '线索阶段',
  `visiting_arrangements` text NULL COMMENT '拜访安排',
  `participants_from_the_opposite` text NULL COMMENT '对方单位参加人员',
  `other_desc` varchar(255) NULL DEFAULT NULL COMMENT '其他说明',
  `tag` text NULL COMMENT '数据标签',
  `audit_code` varchar(255) NULL DEFAULT NULL COMMENT '审批编号',
  `audit_status` varchar(64) NULL DEFAULT NULL COMMENT '审批状态',
  `audit_start_time` datetime NULL DEFAULT NULL COMMENT '审批发起时间',
  `audit_end_time` datetime NULL DEFAULT NULL COMMENT '审批结束时间',
  `audit_start_user` bigint NULL DEFAULT NULL COMMENT '审批发起人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_visit_no` (`visit_no`) USING BTREE,
  KEY `idx_visitor` (`visitor`) USING BTREE,
  KEY `idx_customer_id` (`customer_id`) USING BTREE,
  KEY `idx_contact_id` (`contact_id`) USING BTREE,
  KEY `idx_leads_id` (`leads_id`) USING BTREE,
  KEY `idx_business_id` (`business_id`) USING BTREE,
  KEY `idx_related_partner` (`related_partner`) USING BTREE,
  KEY `idx_visit_plan_id` (`visit_plan_id`) USING BTREE,
  KEY `idx_visit_status` (`visit_status`) USING BTREE,
  KEY `idx_effective_status` (`effective_status`) USING BTREE,
  KEY `idx_audit_status` (`audit_status`) USING BTREE,
  KEY `idx_start_time` (`start_time`) USING BTREE,
  KEY `idx_end_time` (`end_time`) USING BTREE,
  KEY `idx_visit_type` (`visit_type`) USING BTREE,
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='拜访';

-- ----------------------------
-- Table structure for sch_visit_record
-- ----------------------------
DROP TABLE IF EXISTS `sch_visit_record`;
CREATE TABLE `sch_visit_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `record_no` varchar(64) NULL DEFAULT NULL COMMENT '拜访记录编号',
  `visit_user_id` bigint NULL DEFAULT NULL COMMENT '拜访人员',
  `customer_id` bigint NULL DEFAULT NULL COMMENT '拜访客户',
  `contact_id` bigint NULL DEFAULT NULL COMMENT '拜访联系人',
  `leads_id` bigint NULL DEFAULT NULL COMMENT '拜访线索',
  `schedule_id` bigint NULL DEFAULT NULL COMMENT '拜访日程',
  `visit_plan_id` bigint NULL DEFAULT NULL COMMENT '关联拜访计划',
  `partner_id` bigint NULL DEFAULT NULL COMMENT '拜访合作伙伴',
  `finish_time` datetime NULL DEFAULT NULL COMMENT '完成时间',
  `visit_duration` varchar(300) NULL DEFAULT NULL COMMENT '拜访时长',
  `visit_content` text NULL COMMENT '拜访内容',
  `problem_and_solution` text NULL COMMENT '问题及解决方案',
  `customer_intention` text NULL COMMENT '客户意向',
  `next_plan` text NULL COMMENT '下一步计划',
  `status` varchar(64) NULL DEFAULT NULL COMMENT '拜访记录状态',
  -- 签到信息
  `check_in_status` varchar(64) NULL DEFAULT NULL COMMENT '签到状态',
  `check_in_time` datetime NULL DEFAULT NULL COMMENT '签到时间',
  `check_in_longitude` varchar(300) NULL DEFAULT NULL COMMENT '签到经度',
  `check_in_latitude` varchar(300) NULL DEFAULT NULL COMMENT '签到纬度',
  `check_in_address` text NULL COMMENT '签到地址',
  `check_in_address_echo` varchar(300) NULL DEFAULT NULL COMMENT '签到地址（回显）',
  `check_in_picture` text NULL COMMENT '签到照片',
  `check_in_description` text NULL COMMENT '签到补充说明',
  `check_in_exception_reason` varchar(300) NULL DEFAULT NULL COMMENT '签到异常原因',
  -- 签退信息
  `check_out_status` varchar(64) NULL DEFAULT NULL COMMENT '签退状态',
  `check_out_time` datetime NULL DEFAULT NULL COMMENT '签退时间',
  `check_out_longitude` varchar(300) NULL DEFAULT NULL COMMENT '签退经度',
  `check_out_latitude` varchar(300) NULL DEFAULT NULL COMMENT '签退纬度',
  `check_out_address` text NULL COMMENT '签退地址',
  `check_out_address_echo` varchar(300) NULL DEFAULT NULL COMMENT '签退地址（回显）',
  `check_out_picture` text NULL COMMENT '签退照片',
  `check_out_description` text NULL COMMENT '签退补充说明',
  `check_out_exception_reason` varchar(300) NULL DEFAULT NULL COMMENT '签退异常原因',
  -- 附加信息
  `scene_photo` text NULL COMMENT '客户现场照片',
  `travel_summary_report` text NULL COMMENT '出差总结',
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_record_no` (`record_no`) USING BTREE,
  KEY `idx_visit_user_id` (`visit_user_id`) USING BTREE,
  KEY `idx_customer_id` (`customer_id`) USING BTREE,
  KEY `idx_leads_id` (`leads_id`) USING BTREE,
  KEY `idx_schedule_id` (`schedule_id`) USING BTREE,
  KEY `idx_visit_plan_id` (`visit_plan_id`) USING BTREE,
  KEY `idx_status` (`status`) USING BTREE,
  KEY `idx_finish_time` (`finish_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='拜访记录';

-- ----------------------------
-- Table structure for sch_task
-- ----------------------------
DROP TABLE IF EXISTS `sch_task`;
CREATE TABLE `sch_task` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `task_no` varchar(64) NULL DEFAULT NULL COMMENT '任务编号',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `title` varchar(200) NULL DEFAULT NULL COMMENT '任务标题',
  `task_describe` text NULL COMMENT '任务描述',
  `task_type` varchar(64) NULL DEFAULT NULL COMMENT '任务类型',
  `element` varchar(64) NULL DEFAULT NULL COMMENT '要素',
  -- 执行人员
  `executor` bigint NULL DEFAULT NULL COMMENT '执行人',
  `cooperate_users` text NULL COMMENT '协同人',
  `supervisor` text NULL COMMENT '督办人',
  `handover_cooperateuser` tinyint(1) NULL DEFAULT 0 COMMENT '将原执行人设置为协同人',
  -- 优先级与重要性
  `priority` varchar(64) NULL DEFAULT '1' COMMENT '优先级',
  `importance` varchar(64) NULL DEFAULT NULL COMMENT '重要性',
  `urgency` varchar(64) NULL DEFAULT NULL COMMENT '紧急程度',
  -- 时间相关
  `start_time` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `exact_date` datetime NULL DEFAULT NULL COMMENT '指定日期',
  `end_exact_date` datetime NULL DEFAULT NULL COMMENT '结束指定日期',
  `finish_date` datetime NULL DEFAULT NULL COMMENT '完成日期',
  `task_send_time` datetime NULL DEFAULT NULL COMMENT '任务下发时间',
  -- 提醒设置
  `start_reminder_setting` varchar(64) NULL DEFAULT NULL COMMENT '开始提醒设置',
  `end_reminder_setting` varchar(64) NULL DEFAULT NULL COMMENT '结束提醒设置',
  `send_create_remind` tinyint(1) NULL DEFAULT 0 COMMENT '是否发送过创建提醒',
  -- 关联对象
  `relate_obj` varchar(64) NULL DEFAULT NULL COMMENT '关联对象',
  `customer_id` bigint NULL DEFAULT NULL COMMENT '关联客户',
  `lead_id` bigint NULL DEFAULT NULL COMMENT '关联线索',
  `phase_id` varchar(64) NULL DEFAULT NULL COMMENT '项目阶段',
  `lead_phase_id` varchar(64) NULL DEFAULT NULL COMMENT '线索阶段',
  `task_item` varchar(64) NULL DEFAULT NULL COMMENT '任务事项',
  -- 完成情况
  `finish_status` varchar(64) NULL DEFAULT NULL COMMENT '完成状态',
  `task_finish_detail` text NULL COMMENT '任务完成情况',
  `finish_user` bigint NULL DEFAULT NULL COMMENT '任务完成人',
  `quality_score` varchar(64) NULL DEFAULT NULL COMMENT '质量评分',
  `speed` varchar(64) NULL DEFAULT NULL COMMENT '进度',
  -- 附件
  `attachment` text NULL COMMENT '附件',
  -- 研判相关
  `assess_content` text NULL COMMENT '研判内容',
  `assess_view` varchar(64) NULL DEFAULT NULL COMMENT '研判意见',
  `assess_depment` bigint NULL DEFAULT NULL COMMENT '研判部门',
  `assess_user` bigint NULL DEFAULT NULL COMMENT '研判人员',
  `assess_time` datetime NULL DEFAULT NULL COMMENT '研判时间',
  `industry_group` bigint NULL DEFAULT NULL COMMENT '研判行业分组',
  -- 问题相关（区县对接）
  `issue_type` varchar(64) NULL DEFAULT NULL COMMENT '问题类型',
  `issue_content` text NULL COMMENT '问题描述',
  `area_dock_content` text NULL COMMENT '区级对接情况',
  `area_advice` text NULL COMMENT '区级建议',
  `money_value` varchar(64) NULL DEFAULT NULL COMMENT '问题金额',
  `issue_id` varchar(64) NULL DEFAULT NULL COMMENT '市局项目问题轮次ID',
  -- 标签
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_task_no` (`task_no`) USING BTREE,
  KEY `idx_executor` (`executor`) USING BTREE,
  KEY `idx_customer_id` (`customer_id`) USING BTREE,
  KEY `idx_lead_id` (`lead_id`) USING BTREE,
  KEY `idx_start_time` (`start_time`) USING BTREE,
  KEY `idx_end_time` (`end_time`) USING BTREE,
  KEY `idx_finish_status` (`finish_status`) USING BTREE,
  KEY `idx_task_type` (`task_type`) USING BTREE,
  KEY `idx_priority` (`priority`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务';

-- ----------------------------
-- Table structure for sch_conference
-- ----------------------------
DROP TABLE IF EXISTS `sch_conference`;
CREATE TABLE `sch_conference` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `conference_no` varchar(64) NULL DEFAULT NULL COMMENT '会议编号',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `title` varchar(200) NULL DEFAULT NULL COMMENT '会议标题',
  `conference_type` varchar(64) NULL DEFAULT NULL COMMENT '会议类型',
  `type` varchar(64) NULL DEFAULT NULL COMMENT '日程类型',
  -- 时间相关
  `all_day` tinyint(1) NULL DEFAULT 0 COMMENT '是否全天',
  `start_time` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `start_time_all` datetime NULL DEFAULT NULL COMMENT '开始时间-全天',
  `end_time_all` datetime NULL DEFAULT NULL COMMENT '结束时间-全天',
  -- 提醒设置
  `reminder_setting` varchar(64) NULL DEFAULT NULL COMMENT '提醒设置',
  `reminder_setting_time` datetime NULL DEFAULT NULL COMMENT '提醒设置指定日期',
  `reminder_setting_time_all` datetime NULL DEFAULT NULL COMMENT '提醒设置指定日期-全天',
  -- 参与人员
  `participant_users` text NULL COMMENT '参与人',
  `schedule_participation_status` varchar(64) NULL DEFAULT NULL COMMENT '日程参与状态',
  -- 会议详情
  `host_leader` varchar(64) NULL DEFAULT NULL COMMENT '主持领导',
  `agenda` text NULL COMMENT '议程',
  `schedule_describe` text NULL COMMENT '日程描述',
  `activity_address` text NULL COMMENT '会议地点',
  -- 关联对象
  `relate_obj` varchar(64) NULL DEFAULT NULL COMMENT '关联对象',
  `lead_id` bigint NULL DEFAULT NULL COMMENT '关联线索',
  `customer_id` bigint NULL DEFAULT NULL COMMENT '关联客户',
  `project_phase` varchar(64) NULL DEFAULT NULL COMMENT '项目阶段',
  `leads_phase` varchar(64) NULL DEFAULT NULL COMMENT '线索阶段',
  -- 优先级
  `urgency` varchar(64) NULL DEFAULT NULL COMMENT '紧急程度',
  `importance` varchar(64) NULL DEFAULT NULL COMMENT '重要性',
  -- 附件与标签
  `attachment` text NULL COMMENT '附件',
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_conference_no` (`conference_no`) USING BTREE,
  KEY `idx_conference_type` (`conference_type`) USING BTREE,
  KEY `idx_customer_id` (`customer_id`) USING BTREE,
  KEY `idx_lead_id` (`lead_id`) USING BTREE,
  KEY `idx_start_time` (`start_time`) USING BTREE,
  KEY `idx_end_time` (`end_time`) USING BTREE,
  KEY `idx_title` (`title`(100)) USING BTREE,
  KEY `idx_urgency` (`urgency`) USING BTREE,
  KEY `idx_importance` (`importance`) USING BTREE,
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会议';

-- ----------------------------
-- Table structure for sch_activity
-- ----------------------------
DROP TABLE IF EXISTS `sch_activity`;
CREATE TABLE `sch_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `activity_no` varchar(64) NULL DEFAULT NULL COMMENT '活动编号',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `title` varchar(200) NULL DEFAULT NULL COMMENT '活动标题',
  `activity_type` varchar(64) NULL DEFAULT NULL COMMENT '活动类型',
  `type` varchar(64) NULL DEFAULT NULL COMMENT '日程类型',
  -- 时间相关
  `all_day` tinyint(1) NULL DEFAULT 0 COMMENT '是否全天',
  `start_time` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `start_time_all` datetime NULL DEFAULT NULL COMMENT '开始时间-全天',
  `end_time_all` datetime NULL DEFAULT NULL COMMENT '结束时间-全天',
  -- 提醒设置
  `reminder_setting` varchar(64) NULL DEFAULT NULL COMMENT '提醒设置',
  `reminder_setting_time` datetime NULL DEFAULT NULL COMMENT '提醒设置指定日期',
  `reminder_setting_time_all` datetime NULL DEFAULT NULL COMMENT '提醒设置指定日期-全天',
  -- 参与人员
  `participant_users` text NULL COMMENT '参与人',
  `number_of_participants` int(11) NULL DEFAULT NULL COMMENT '参与人数',
  `schedule_participation_status` varchar(64) NULL DEFAULT NULL COMMENT '日程参与状态',
  -- 活动详情
  `activity_address` text NULL COMMENT '活动地点',
  `activity_related_parties` text NULL COMMENT '活动相关方',
  `schedule_describe` text NULL COMMENT '日程描述',
  `qr_code` text NULL COMMENT '二维码',
  `qr_code_url` varchar(300) NULL DEFAULT NULL COMMENT '二维码跳转地址',
  -- 活动规模与级别
  `activity_level` varchar(64) NULL DEFAULT NULL COMMENT '活动级别',
  `scale_of_people` int(11) NULL DEFAULT NULL COMMENT '规模人次',
  `industry_classification` varchar(64) NULL DEFAULT NULL COMMENT '行业分类',
  -- 关联对象
  `relate_obj` varchar(64) NULL DEFAULT NULL COMMENT '关联对象',
  `lead_id` bigint NULL DEFAULT NULL COMMENT '关联线索',
  `customer_id` bigint NULL DEFAULT NULL COMMENT '关联客户',
  `project_phase` varchar(64) NULL DEFAULT NULL COMMENT '项目阶段',
  `leads_phase` varchar(64) NULL DEFAULT NULL COMMENT '线索阶段',
  -- 优先级
  `urgency` varchar(64) NULL DEFAULT NULL COMMENT '紧急程度',
  `importance` varchar(64) NULL DEFAULT NULL COMMENT '重要性',
  -- 附件与标签
  `attachment` text NULL COMMENT '附件',
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_activity_no` (`activity_no`) USING BTREE,
  KEY `idx_activity_type` (`activity_type`) USING BTREE,
  KEY `idx_customer_id` (`customer_id`) USING BTREE,
  KEY `idx_lead_id` (`lead_id`) USING BTREE,
  KEY `idx_start_time` (`start_time`) USING BTREE,
  KEY `idx_end_time` (`end_time`) USING BTREE,
  KEY `idx_activity_level` (`activity_level`) USING BTREE,
  KEY `idx_title` (`title`(100)) USING BTREE,
  KEY `idx_industry_classification` (`industry_classification`) USING BTREE,
  KEY `idx_urgency` (`urgency`) USING BTREE,
  KEY `idx_importance` (`importance`) USING BTREE,
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动';

-- ----------------------------
-- Table structure for sch_activity_parties
-- ----------------------------
DROP TABLE IF EXISTS `sch_activity_parties`;
CREATE TABLE `sch_activity_parties` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `parties_no` varchar(64) NULL DEFAULT NULL COMMENT '相关方编号',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  -- 类型信息
  `type_str` varchar(100) NULL DEFAULT NULL COMMENT '类型(英文)',
  `type_name_str` varchar(100) NULL DEFAULT NULL COMMENT '类型(中文)',
  `related_type` varchar(64) NULL DEFAULT NULL COMMENT '类型(枚举单选)',
  `activity_role` varchar(64) NULL DEFAULT NULL COMMENT '角色定位',
  -- 机构信息
  `investment_name` varchar(200) NULL DEFAULT NULL COMMENT '投资机构',
  `finance_name` varchar(300) NULL DEFAULT NULL COMMENT '金融机构',
  `research_name` varchar(300) NULL DEFAULT NULL COMMENT '新研机构',
  `scene_name` varchar(300) NULL DEFAULT NULL COMMENT '场景',
  -- 关联对象名称（冗余存储，便于展示）
  `leads_name` varchar(300) NULL DEFAULT NULL COMMENT '线索',
  `project_name` varchar(300) NULL DEFAULT NULL COMMENT '项目',
  -- 关联ID
  `relation_id` varchar(64) NULL DEFAULT NULL COMMENT '机构库/项目/线索ID',
  `activity_id` bigint NULL DEFAULT NULL COMMENT '活动ID',
  `relation_dept_id` bigint NULL DEFAULT NULL COMMENT '相关联的部门ID',
  -- 标签
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_activity_id` (`activity_id`) USING BTREE,
  KEY `idx_related_type` (`related_type`) USING BTREE,
  KEY `idx_activity_role` (`activity_role`) USING BTREE,
  KEY `idx_relation_id` (`relation_id`) USING BTREE,
  KEY `idx_type_str` (`type_str`) USING BTREE,
  KEY `idx_create_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动相关方';

-- ----------------------------
-- Table structure for sch_activity_sign_in
-- ----------------------------
DROP TABLE IF EXISTS `sch_activity_sign_in`;
CREATE TABLE `sch_activity_sign_in` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` int(11) unsigned DEFAULT '0' COMMENT '租户ID',
  `sign_in_no` varchar(64) NULL DEFAULT NULL COMMENT '签到编号',
  -- 参与状态
  `schedule_participation_status` varchar(64) NULL DEFAULT NULL COMMENT '日程参与状态',
  -- 签到信息
  `sign_in_status` varchar(64) NULL DEFAULT NULL COMMENT '签到状态',
  `sign_in_time` datetime NULL DEFAULT NULL COMMENT '签到时间',
  `sign_in_type` varchar(64) NULL DEFAULT NULL COMMENT '签到方式',
  -- 参与人信息
  `participant_user` bigint NULL DEFAULT NULL COMMENT '参与人',
  `participant_dept_id` bigint NULL DEFAULT NULL COMMENT '参与人所属部门',
  `participant_name` varchar(64) NULL DEFAULT NULL COMMENT '参与人姓名',
  `participant_position` varchar(64) NULL DEFAULT NULL COMMENT '参与人职位',
  -- 关联
  `schedule_info_id` bigint NULL DEFAULT NULL COMMENT '关联日程',
  `activity_role` varchar(64) NULL DEFAULT NULL COMMENT '活动相关方角色定位',
  -- 标签
  `tag` text NULL COMMENT '数据标签',
  -- 审计字段
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tenant_id` (`tenant_id`) USING BTREE,
  KEY `idx_schedule_info_id` (`schedule_info_id`) USING BTREE,
  KEY `idx_participant_user` (`participant_user`) USING BTREE,
  KEY `idx_sign_in_status` (`sign_in_status`) USING BTREE,
  KEY `idx_sign_in_time` (`sign_in_time`) USING BTREE,
  KEY `idx_activity_role` (`activity_role`) USING BTREE,
  KEY `idx_schedule_participation_status` (`schedule_participation_status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动签到';