// Copyright 2015 itwarcraft@gmail.com  All rights reserved.

/*
   base example:

   import(
       "fmt"
       "github.com/itwarcraft/sabar"
   )

   func main(){
       //配置相关权限信息：用户名、密码、角色、权限
       sabar.InitRealm(&sabar.PropsRealm{
            Name:"admin",
            Password:"admin",
            Roles:[]string{"admin","normal"},
            Permissions: []string{"/admin","/index"},
            })
//使用配置文件，从而获取上边的信息   ！未实现
//         sabar.InitRealm(&sabar.IniRealm{
//                FilePath:'../realm.ini',
//            })

       currentUser := sabar.GetSubject()
       token := sabar.UsernamePasswordToken{Name:"admin",Password:[]byte("admin"),}
       err := currentUser.Login(token)
       if err!=nil {
            fmt.Println(err.Error())
            return
       }

        hasPerm := currentUser.CheckPermission("/admin")
        if hasPerm{
            fmt.Println("user has permission /admin")
        }else{
            fmt.Println("user donot has permission /admin")
        }
        fmt.Println("end")
   }



    example1:

    import(
       "fmt"
       "github.com/itwarcraft/sabar"
   )



   func main(){
        //从数据库中查询对应的信息   并配置密码加密处理策略
       sabar.InitRealm(&DbRealm{})


       currentUser := sabar.GetSubject()
       token := sabar.UsernamePasswordToken{Name:"admin",Password:[]byte("admin"),}
       err := currentUser.Login(token)
       if err!=nil {
            fmt.Println(err.Error())
            return
       }

        hasPerm := currentUser.CheckPermission("/admin")
        if hasPerm{
            fmt.Println("user has permission /admin")
        }else{
            fmt.Println("user donot has permission /admin")
        }
        fmt.Println("end")
   }

    //用于实现sabar.Realm相应的接口
   type DbRealm struct{

   }

   func (realm *DbRealm)GetName(){
    return "dbRealm"
   }
   //根据当前的登陆信息（用户名与密码），查询数据库，得到当前用户的数据库信息
   func (self *DbRealm)GetAuthcInfo(token AuthcToken) (sabar.AuthcInfo, error){

   }
    //如果校验登陆成功，则如何获取对应的授权信息，其中的principal来源于AuthcInfo的GetPrincipal()方法
   func (self *DbRealm)GetAuthzInfo(principal interface{})sabar.AuthzInfo{

   }
   //密码加密处理方法
   func Encrypt(info *sabar.AuthcInfo) (*sabar.AuthcInfo, error){

   }





*/

package sabar
