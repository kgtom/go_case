
### 一、流程
 app端——>登录授权appleSvc——>返回数据——>提交后端服务———>请求appleSvc验证授权码
 Ps :如果是web端，需要先在appleSve后头配置redirect_uri就是后台配置的接收code码的地址

### 二、验证

APP端登录授权返回：
User ID: 苹果用户唯一标识符，它在同一个开发者账号下的所有 App 下是一样的，我们可以用它来与后台的账号体系绑定起来（类似于微信的OpenID）。
Verification Data: 包括identityToken, authorizationCode。用于传给开发者后台服务器，然后开发者服务器再向苹果的身份验证服务端验证本次授权登录请求数据的有效性和真实性。
Account Information: 苹果用户信息，包括全名、邮箱等，登录时用户可以选择隐藏真实的邮件地址和随意修改姓名。
Real User Indicator: 用于判断当前登录的苹果账号是否是一个真实用户，取值有：unsupported、unknown、likelyReal。

#### 传入后台参数：identityToken, authorizationCode、userID

#### 验证一：获取 rest api apple公钥，解 identityToken，
理论：因为idnetityToken使用非对称加密 RSASSA【RSA签名算法】 和 ECDSA【椭圆曲线数据签名算法】，当验证签名的时候，利用公钥来解密，当解密内容与base64UrlEncode(header) + "." + base64UrlEncode(payload)的内容一致的时候，表示验证通过。


#### 验证二：给authorizationCode获取token,解出token中sub(即userID),比较入参userID与其是否一致。
理论：authorizationCode 只能使用一次且5分钟内

** 总结 **： 偏向使用第二种，web \app端都可以使用。

> reference:
*[juejin](https://juejin.im/post/5d551d11e51d4561cf15dfae)
*[easeapi](https://easeapi.com/blog/blog/88-sign-with-apple.html)
*[csdn](https://blog.csdn.net/wpf199402076118/article/details/99677412)
