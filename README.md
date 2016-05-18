# go-deeplink
Provides deeplink for android and ios app

### INSTALLATION

    go get github.com/michaelfemi81/go-deeplink
### API DOCUMENTATION
``` 
package main
import (
"github.com/michaelfemi81/go-deeplink"
)
func main() {

  http.HandleFunc("/", doDeep)
  http.ListenAndServe(":1234", nil)
}


func doDeep(w http.ResponseWriter ,r *http.Request) {
  dl := deeplink.DeepLink{}
  
dl.Init("http://yourfallback.com","com.myapp.myappname");

  err:=dl.DoDeepLink(w,r)
  if(err!=nil){
    fmt.Println(err)
  }
}
```