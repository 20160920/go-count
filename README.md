# go-count
channel 来实现协程安全的计数器，使用 channel 充当队列，对计数器的操作(读、写)都缓存在队列中，按顺序操作。
