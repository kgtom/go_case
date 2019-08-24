## 一 、OOP 面向接口编程，不依赖与具体的实现，依赖于抽象的接口

* 对外使用接口，内部私有结构体具体的实现
~~~
type XXXServicer interface {
	List(ctx context.Context, userID, offset, cnt uint32) ([]uint32, uint32, error)
	
}

func NewXXXServicer(xxx proto.XXXType) XXXServicer {
	return &xxxService{
		xxxType: xxx,
	}
}

type xxxService struct {
	xxxType proto.XXXType
}

//列表
func (f *xxxService) List(ctx context.Context, userID, offset, cnt uint32) ([]uint32, uint32, error) {
}
~~~
