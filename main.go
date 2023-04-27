package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"goframe2-skeleton/internal/cmd"
	_ "goframe2-skeleton/internal/logic"
	_ "goframe2-skeleton/internal/packed"
)

func main() {
	cmd.ServerBoot()
}
