

## http 的G泄露(或者G数量暴涨)

如果 没有 resp.Body.Close()，则会泄露 一个read 和write G,在没有收到结束信号之前，在for+select 机制里，都会堵塞住，导致G及时退出.当请求量过大时，G数量会暴涨
