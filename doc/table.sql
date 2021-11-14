-- 用户表
CREATE TABLE `tbl_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL COMMENT '用户名',
    `phone` varchar(32) DEFAULT '' COMMENT '手机号',
    `is_vip` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是会员(非会员0/会员1)',
    `user_pwd` varchar(256) NOT NULL COMMENT  '用户encoded密码',
    `create_time` datetime default NOW() COMMENT '创建时间',
    `update_time` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态(可用0/已删除1)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`user_name`),
    KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- 健身设备表
CREATE TABLE `tbl_equipment` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL COMMENT '健身设备名称',
    `photo_url` varchar(256) NOT NULL DEFAULT '' COMMENT '图片url地址',
    `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '健身设备详细说明',
    `extra1` varchar(64) NOT NULL DEFAULT '' COMMENT '扩展字段1',
    `extra2` varchar(64) NOT NULL DEFAULT '' COMMENT '扩展字段2',
    `extra3` varchar(64) NOT NULL DEFAULT '' COMMENT '扩展字段3',
    `create_time` datetime default NOW() COMMENT '创建时间',
    `update_time` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态(可用0/已删除1状态)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- 服务项目表
CREATE TABLE `tbl_service` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL COMMENT '服务项目名称',
    `cost` float(10) NOT NULL COMMENT '服务费用,元/小时',
    `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '服务项目详细说明',
    `tag` varchar(64) NOT NULL DEFAULT '' COMMENT '标签',
    `extra2` varchar(64) NOT NULL DEFAULT '' COMMENT '扩展字段2',
    `extra3` varchar(64) NOT NULL DEFAULT '' COMMENT '扩展字段3',
    `create_time` datetime default NOW() COMMENT '创建时间',
    `update_time` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态(可用0/已删除1状态)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- 用户与服务项目情况
CREATE TABLE `tbl_user_service` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL COMMENT '用户名',
    `service_name` varchar(64) NOT NULL COMMENT '服务项目名称',
    `service_cost` float(10) NOT NULL COMMENT '服务费用,元/小时',
    `service_time` int(11) NOT NULL COMMENT '使用时间,向上取整',
    `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:没有使用，1:使用过',
    `create_time` datetime default NOW() COMMENT '创建时间',
    `update_time` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
    `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态(可用0/已删除1状态)',
    PRIMARY KEY (`id`),
    key `idx_username` (`user_name`),
    KEY `idx_deleted` (`deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;