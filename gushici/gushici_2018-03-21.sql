# ************************************************************
# Dump of table home_blog
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_blog`;

CREATE TABLE `home_blog` (
  `blog_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `cate_id` int(11) unsigned NOT NULL DEFAULT '0',
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `title_hash` varchar(30) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `author` char(60) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `update_user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `intro` varchar(100) COLLATE utf8_unicode_ci DEFAULT '',
  `cover` varchar(300) COLLATE utf8_unicode_ci DEFAULT NULL,
  `content` text COLLATE utf8_unicode_ci,
  `tags` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `from_url` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `url_hash` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `view_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '浏览量',
  `is_orig` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '1原创,2非原创',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态1正常2草稿3删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `love_num` int(11) unsigned NOT NULL DEFAULT '0',
  `comment_num` int(11) unsigned NOT NULL DEFAULT '0',
  `top` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1置顶',
  PRIMARY KEY (`blog_id`),
  KEY `status` (`status`),
  KEY `url_hash` (`url_hash`),
  KEY `user_id` (`user_id`),
  KEY `is_orig` (`is_orig`),
  KEY `update_user_id` (`update_user_id`),
  KEY `title_hash` (`title_hash`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table home_cate
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_cate`;

CREATE TABLE `home_cate` (
  `cate_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级id',
  `cate_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单名称',
  `is_nav` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `is_show` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示',
  `url` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'url',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `status` tinyint(4) unsigned NOT NULL COMMENT '状态',
  PRIMARY KEY (`cate_id`),
  KEY `status` (`status`),
  KEY `pid` (`pid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='分类表';



# Dump of table home_comment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_comment`;

CREATE TABLE `home_comment` (
  `com_id` int(11) NOT NULL AUTO_INCREMENT,
  `pcom_id` int(11) NOT NULL DEFAULT '0',
  `user_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `nick_name` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `blog_id` int(11) NOT NULL DEFAULT '0',
  `photo` varchar(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `content` text COLLATE utf8_unicode_ci,
  `path` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL COMMENT '状态1正常2草稿3删除',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`com_id`),
  KEY `status` (`status`),
  KEY `pcom_id` (`pcom_id`),
  KEY `huser_id` (`user_id`),
  KEY `create_time` (`create_time`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table home_feeling
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_feeling`;

CREATE TABLE `home_feeling` (
  `feel_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(200) DEFAULT '',
  `cover` varchar(150) DEFAULT NULL,
  `status` tinyint(3) unsigned DEFAULT '0' COMMENT '0屏蔽1置顶2正常',
  `create_time` int(10) unsigned DEFAULT NULL,
  `update_time` int(10) unsigned DEFAULT NULL,
  `type` tinyint(3) DEFAULT '1' COMMENT '1心情2开心一乐',
  PRIMARY KEY (`feel_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;



# Dump of table home_tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_tag`;

CREATE TABLE `home_tag` (
  `tag_id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '标签名称',
  `hash_num` bigint(20) NOT NULL DEFAULT '0' COMMENT 'hash num',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `use_num` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`tag_id`),
  KEY `status` (`status`),
  KEY `hash_num` (`hash_num`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='标签表';



# Dump of table home_tag_map
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_tag_map`;

CREATE TABLE `home_tag_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `blog_id` int(11) NOT NULL DEFAULT '0',
  `tag_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `blog_id` (`blog_id`),
  KEY `tag_id` (`tag_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='文章标签关联表';



# Dump of table home_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `home_user`;

CREATE TABLE `home_user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `photo` varchar(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `mark` varchar(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `hash_email` bigint(20) NOT NULL DEFAULT '0',
  `mobile` char(11) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `password` char(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `open_id` char(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `login_ip` char(20) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `login_time` int(11) unsigned NOT NULL DEFAULT '0',
  `login_count` int(11) unsigned NOT NULL DEFAULT '0',
  `get_pass_time` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `status` tinyint(4) unsigned NOT NULL COMMENT '状态 正常1 ',
  PRIMARY KEY (`user_id`),
  KEY `status` (`status`),
  KEY `hash_email` (`hash_email`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table sys_group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_group`;

CREATE TABLE `sys_group` (
  `group_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分组id',
  `group_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '分组名称',
  `intro` text COLLATE utf8_unicode_ci COMMENT '分组介绍',
  `roles` text COLLATE utf8_unicode_ci NOT NULL COMMENT '角色',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`group_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='分组表';



# Dump of table sys_log_act
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_log_act`;

CREATE TABLE `sys_log_act` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '日志',
  `class` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT '模块',
  `func` varchar(30) COLLATE utf8_unicode_ci NOT NULL COMMENT '功能操作',
  `content` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作描述',
  `user_id` int(11) NOT NULL COMMENT '操作人',
  `log_ip` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作ip',
  `log_time` int(11) NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='操作日志';



# Dump of table sys_log_login
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_log_login`;

CREATE TABLE `sys_log_login` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL COMMENT '登录id',
  `user_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `nick_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `login_ip` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT '登录ip',
  `login_time` int(11) unsigned NOT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='登录日志';



# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0' COMMENT '上级id',
  `name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单名称',
  `urlclass` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '所属类',
  `operate` text COLLATE utf8_unicode_ci NOT NULL COMMENT '功能方法',
  `icon` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '图标',
  `lock` tinyint(4) NOT NULL COMMENT '锁定',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`menu_id`),
  KEY `status` (`status`),
  KEY `pid` (`pid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='菜单表';



# Dump of table sys_notice
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_notice`;

CREATE TABLE `sys_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text COLLATE utf8_unicode_ci,
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '状态1已读2未读',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table sys_purview
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_purview`;

CREATE TABLE `sys_purview` (
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `menu_id` int(11) NOT NULL COMMENT '菜单id',
  `operate` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作权限 逗号分隔',
  KEY `role_id` (`role_id`),
  KEY `menu_id` (`menu_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='权限表';



# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `intro` text COLLATE utf8_unicode_ci COMMENT '角色介绍',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`role_id`),
  KEY `status` (`status`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='角色表';



# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '登录名',
  `nick_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '昵称',
  `group_id` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '分组id',
  `role_id` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色id',
  `photo` char(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(6) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码盐值',
  `email` char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `mobile` char(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned zerofill DEFAULT NULL,
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_ip` char(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `login_count` int(11) unsigned NOT NULL DEFAULT '0',
  `user_type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`user_id`),
  KEY `status` (`status`),
  KEY `email` (`email`),
  KEY `mobile` (`mobile`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户表';
