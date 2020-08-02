

~~~
svc := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StatsHandler(zipkingrpc.NewServerHandler(zkTracer)),
	)
  
  
  func StartMetric(ctx context.Context, server *grpc.Server, metricPort int) {
	grpc_prometheus.Register(server)              //注册prometheus监控到服务上
	grpc_prometheus.EnableHandlingTimeHistogram() //开启直方图
	if metricPort != 0 {
		http.Handle("/metrics", promhttp.Handler()) //几行代码实现了一个exporter，内部使用默认收集器
		go http.ListenAndServe(":"+strconv.Itoa(metricPort), nil)
	}
}
~~~
