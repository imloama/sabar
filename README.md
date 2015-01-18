# sabar
go语言权限框架，打算试试模仿shiro写一个相同功能的权限框架，减少功能。



使用示例与shiro基本相同

UsernamePasswordToken token = new UsernamePasswordToken(username, password);
token.setRememberMe(true);
Subject currentUser = SecurityUtils.getSubject();
currentUser.login(token);


import(
    "github.com/itwarcraft/sabar"
)


currentUser := sabar.GetSubject()
currentUser.Login(sabar.UsernamePasswordToken{Name:"admin",Password:[]byte("admin")})

isLogin := currentUser.isAuthorized()
