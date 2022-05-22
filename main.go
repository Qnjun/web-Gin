package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func sayHello(w http.ResponseWriter, r *http.Request) {
	bool, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(bool))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
*/

func wangYe(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Golang",
	})
}

func main() {

	r := gin.Default() //返回默认的路由引擎
	// 用户使用Get请求访问 /hello时，执行wanYe这个函数
	r.GET("/hello", wangYe)
	//启动服务
	//r.Run()

	//使用匿名函数的形式来,实现
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	r.Run()

}
