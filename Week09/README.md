学习笔记

以前用Java和Nodejs 写过的练手TCP Server移植到Golang, 用了之前学到的Gracefully shutdown的技巧。有些resource cleanup 做的还不完善。


```
/l                               # list all active sessions
/q                               # exit session
/b {message}                     # broadcast excluding the sender
/c {session_id} {message}        # direct messsage
```