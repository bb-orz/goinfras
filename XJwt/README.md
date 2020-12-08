# JWT TokenUtils

> 基于 github.com/dgrijalva/jwt-go 包

### Gin Documentation
> Documentation https://godoc.org/github.com/dgrijalva/jwt-go

### JWT Base Example
```
tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

type MyCustomClaims struct {
    Foo string `json:"foo"`
    jwt.StandardClaims
}

// sample token is expired.  override time so it parses as valid
at(time.Unix(0, 0), func() {
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("AllYourBase"), nil
    })

    if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
        fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
    } else {
        fmt.Println(err)
    }
})
```

### Starter Usage
```
goinfras.RegisterStarter(XJwt.NewStarter(middlewares...))

```

### Usage
```
userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: "", Gender: 1}
token, err := XJwt.XTokenUtils().Encode(userClaim)
So(err, ShouldBeNil)
Println("Token String", token)
Println("Token Service Decode And Validate:")
claim, err := XJwt.XTokenUtils().Decode(token)
Println("Validate Error:", err)
So(err, ShouldBeNil)
Println("Token Claim:", claim)
```