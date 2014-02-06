/*
控制器

第一次重构:

大幅度简化repo层代码
实现对业务进行边界细分(domain层)


域(Domain)说明

通用域(Common) 提供SSO登录验证(Account)
  ├ 平台域(Platform) 提供平台基础设施
  └ 管理域(Admin) 提供系统管理功能


*/
package controllers
