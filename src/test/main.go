package main

import (
    "github.com/sunnybug/go-tmlog/src/heiyeluren/tmlog"
    "runtime"
    "time"
)

// 为了充分利用多核CPU, 设置服务进程并发线程数
func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

/**
 * 测试服务进程的main函数
 */
func main() {

    /**
     * 以下为tmlog后端协程工作代码
     *
     * 说明: tmlog主工作协程必须在main协程里面启动工作, 否则无法完成后台日志打印的工作
     */

    // 传递给日志类的配置项 (配置项的含义参考 tmlog.go 代码里面的注释说明)
    // 说明: 这些配置实际可以放置到配置文件里，不过我们这里只是演示，直接就生成map了
    logConf := map[string]string{
        "log_notice_file_path":  "log/heiyeluren.log",
        "log_debug_file_path":   "log/heiyeluren.log",
        "log_trace_file_path":   "log/heiyeluren.log",
        "log_fatal_file_path":   "log/heiyeluren.log.wf",
        "log_warning_file_path": "log/heiyeluren.log.wf",
        "log_cron_time":         "day",
        "log_chan_buff_size":    "16",
        "log_flush_timer":       "10000",
        "log_debug_open":        "0",
        "log_level":             "31",
    }
    // 启动 tmlog 工作协程, 可以理解为tmlog的服务器端
    tmlog.Log_Run(logConf)

    //打印日志3: 在一个单独协程里打印日志
    go LogTest()

    // hold 住 main协程不退出
    // 说明: 如果想一直运行程序, 可以把下面这行开启
    //select {}

    //sleep
    time.Sleep(time.Second * 5)
}

/**
 * 单独协程函数测试打印日志
 *
 */
func LogTest() {
    logHandle3 := tmlog.NewLogger()

    logHandle3.WriteLogF(1, "[logger=logHandle3 msg='%s']", "xxxxx")
}
