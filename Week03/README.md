学习笔记

errgroup控制开始以及signal processing

结束时也需要用waitgroup 等待两个server 都完成gracefully shutdown的过程， 不加shutdownWaitGroup的处理最后可能等不到两个server 都gracefully shutdown, 主goroutine 可能已经退出。在http handler特别加了一些sleep, 来模拟需要force shutdown 的情况。