### Week02 作业题目：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
* 需要 wrap 这个 error 抛给上层
* 因为是应用调用了其它库（sql)，所以需要 wrap 并往上抛
* 但 biz 不需要 wrap, 直接往上抛 error；但如果 biz 自己的逻辑有 error, 需要 wrap 再往上抛
* service 类似