/*************************************************************************
> File Name: web.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 23:49:03 星期三
> Content: This is a desc
*************************************************************************/

package main

import "github.com/sgs921107/go_framework/app"

// @version 0.0.1
// @description Listen and Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	app.ListenAndServer("0.0.0.0:8080")
}
