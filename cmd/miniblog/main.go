// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/miniblog.

package main

import (
	"github.com/ra1n6ow/miniblog/internal/miniblog"
	// 通过导入匿名包 go.uber.org/automaxprocs 来使程序自动设置 GOMAXPROCS 以匹配 Linux 容器 CPU 配额。
	//  通过正确设置容器的 CPU 配额，可以解决 GOMAXPROCS 可能设置过大，导致生成线程过多，从而导致严重的上下文切换，浪费 CPU，降低程序性能的潜在问题。
	// 参考1：https://pandaychen.github.io/2020/02/28/GOMAXPROCS-POT/
	// 参考2：https://pandaychen.github.io/2020/02/29/AUTOMAXPROCS-ANALYSIS/
	//_ "go.uber.org/automaxprocs"
	"os"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
