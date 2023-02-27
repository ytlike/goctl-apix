CREATE TABLE `qbq_open_platform_application_center`.`application_base_info`
(
    `id`                 bigint(20) NOT NULL COMMENT '主键',
    `app_name`           varchar(64)  NOT NULL DEFAULT '' COMMENT '应用名称',
    `app_code`           varchar(255) NOT NULL DEFAULT '' COMMENT '应用编码',
    `app_developer_id`   bigint(20) NOT NULL COMMENT '开发者用户id',
    `app_developer_name` varchar(20)           DEFAULT NULL COMMENT '开发者名称',
    `app_icon`           varchar(255)          DEFAULT NULL COMMENT '应用图标',
    `app_key`            varchar(32)           DEFAULT NULL COMMENT 'appKey',
    `app_secret`         varchar(32)           DEFAULT NULL COMMENT 'appSecret',
    `app_ability_url`    varchar(255)          DEFAULT NULL COMMENT '能力环境地址',
    `paas_app_id`        varchar(64)           DEFAULT NULL COMMENT '能力平台的AppId',
    `app_desc`           varchar(255)          DEFAULT NULL COMMENT '应用描述',
    `app_status`         tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态：0-停用；1-启用；',
    `watermark`          varchar(100)          DEFAULT NULL COMMENT '水印内容',
    `watermark_position` tinyint(2) DEFAULT NULL COMMENT '水印位置 1-首页；2-每一页；',
    `deleted`            tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否删除：0-未删除；1-已删除',
    `create_time`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_unique_app_code` (`app_code`),
    KEY                  `idx_developer_id` (`app_developer_id`),
    KEY                  `idx_app_key` (`app_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;