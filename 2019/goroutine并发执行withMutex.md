
### 经常使用并发场景

* 根据userIDs 获取用户信息，因为对方没有提供批量获取用户信息的方法，所以我们开启goroutie并发获取
~~~
result := make([]*proto.User, 0)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for uid := range UserIDs {
		wg.Add(1)
		go func(userID uint32) {
			defer wg.Done()
			userInfo, err := b.getOneUserInfo(ctx, userID)
			if err == nil && userInfo != nil {
				mutex.Lock()
				result = append(result, userInfo)
				mutex.Unlock()
			}
		}(uid)
	}
	wg.Wait()

	if err!=nil{
		//handle err
	}

~~~

* 每次最多50userIDs 获取用户在线信息，可以采用并发获取

~~~
//  maxCheckUserOnlineStatusCnt 是常量 50.
  rsp.UserOnlineList = make([]*proto.UserOnlineStatus, 0)
	rsp.UserOfflineList = make([]*proto.UserOfflineStatus, 0)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for len(req.UserIDs) > 0 {
		wg.Add(1)
		reqUserIDs := make([]uint32, 0)

		if len(req.UserIDs) > maxCheckUserOnlineStatusCnt {
			reqUserIDs = req.UserIDs[:maxCheckUserOnlineStatusCnt]
			req.UserIDs = req.UserIDs[maxCheckUserOnlineStatusCnt:]
		} else {
			reqUserIDs = req.UserIDs[:]
			req.UserIDs = nil
		}
		go func(userIDs []uint32) {
			defer wg.Done()
			checkCtx := context.Background()
			userOnlineMap, err := cmessenger.Cli.QueryUserOnlineStatus(checkCtx, userIDs)
			if err != nil {
				log.WithTrace(checkCtx).Error("call QueryUserOnlineStatusMap err", log.Fields{"req": req.UserIDs, "err": err})
				return
			}

			for _, v := range userIDs {
				if userOnlineMap[v] != nil {

					if userOnlineMap[v].Status {
						on := &proto.UserOnlineStatus{
							UserID:     v,
							OnlineTime: userOnlineMap[v].Time,
						}

						mutex.Lock()
						rsp.UserOnlineList = append(rsp.UserOnlineList, on)
						mutex.Unlock()

					} else {
						off := &proto.UserOfflineStatus{
							UserID:      v,
							OfflineTime: userOnlineMap[v].Time,
						}
						mutex.Lock()
						rsp.UserOfflineList = append(rsp.UserOfflineList, off)
						mutex.Unlock()
					}

				}
			}
		}(reqUserIDs)
	}
	wg.Wait()

~~~
