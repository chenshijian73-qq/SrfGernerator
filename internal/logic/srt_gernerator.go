package logic

import (
	"bufio"
	"fmt"
	"github.com/chenshijian73-qq/cobra_scaffold/pkg/log"
	"os"
	"time"
)

func SrtGernerate(interval int, file, outSrt string) {
	document, err := os.Open(file)
	pkg.CheckAndExit(err)
	defer document.Close()

	// 创建字幕文件
	srtFile, err := os.Create(outSrt)
	pkg.CheckAndExit(err)
	defer srtFile.Close()

	t := time.Duration(interval) * time.Second

	var startTime time.Duration
	var endTime time.Duration = t

	scanner := bufio.NewScanner(document)
	lineIndex := 1

	for scanner.Scan() {
		// 写入字幕序号
		fmt.Fprintf(srtFile, "%d\n", lineIndex)

		// 写入字幕时间
		fmt.Fprintf(srtFile, "%s --> %s\n", formatDuration(startTime), formatDuration(endTime))

		// 写入字幕文本
		fmt.Fprintf(srtFile, "%s\n\n", scanner.Text())

		// 更新时间
		startTime = endTime
		endTime = startTime + t

		lineIndex++
	}

	err = scanner.Err()
	pkg.CheckAndExit(err)
}

// formatDuration 将时间间隔格式化为 SRT字幕格式
func formatDuration(d time.Duration) string {
	hours := int64(d.Hours())
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60
	milliseconds := int64(d.Milliseconds()) % 1000

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}
