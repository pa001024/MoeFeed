/*
组合层

1. 用于替代DAO层 强化数据领域之间的关系和形态联合
2. 在传统CRUD的基础上进行组合层的Filter封装 实现Cache/Auth/Transfer等功能
3. 加入部分组合逻辑 简化业务操作
4. 加入Factory工厂层 抽象业务逻辑中对模型的依赖 (MVVM模式)


TODO:
共享连接对象

*/
package repository
