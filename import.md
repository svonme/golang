# import

### 导入包文件
  import "fmt"
  导入 fmt 包， fmt.Println("hello world")

### 相对路径
  import "./xxx"
  加载当前文件同一目录下的 xxx 模块文件

### 绝对路径
  import "model/xxx"
  加载 ${gopath}/src/model/xxx 模块文件

### 省略调用 (点, 不建议使用，容易混淆资源具体属于某个模块)
  import (
    . "fmt"
  )
  导入 fmt 之后在调用这个包的函数时，你可以省略前缀的包名
  fmt.Println("hello world") 可以写成 Println("hello world")

### 别名
  import (
    f "fmt"
  )
  f.Println("hello world")

### _
  import (
    _ "xxx/xxx"
  )
  _ 操作其实是引入该包, 而不直接使用包里面的函数, 而是调用了该包里面的init函数
