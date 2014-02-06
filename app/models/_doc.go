/*
模型层

whatever



数据结构(层次)
--------------

Common
  Account // 通行证
    EmailVerify // 验证
    Auth // 社会化登录
    Activity // 用户登录/修改密码等操作记录
  System // 系统信息
    Cron // 计划任务
    Queue // 任务队列
BBS
  User // 论坛用户
  Forum // 论坛
    Topic // 主题
      TopicData // 额外信息 如加分/顶/赞/分享等
    Post // 帖子
      PostData // 额外信息 如加分/顶/赞/分享等
      PostTag
Platform
  User
  Auth // 第三方网站OAuth信息
Wiki
  User // Wiki用户
Blog
  User // 博客用户
Yo
  User // 微博用户
Translate
  User // 翻译用户











*/
package models
