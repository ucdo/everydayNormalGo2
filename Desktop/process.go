package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

var user32 = syscall.NewLazyDLL("user32.dll")
var findWindow = user32.NewProc("FindWindowW")
var getProcessId = user32.NewProc("GetWindowThreadProcessId")

var (
	procOpenProcess        = syscall.NewLazyDLL("kernel32.dll").NewProc("OpenProcess")
	procReadProcessMemory  = syscall.NewLazyDLL("kernel32.dll").NewProc("ReadProcessMemory")
	procWriteProcessMemory = syscall.NewLazyDLL("kernel32.dll").NewProc("WriteProcessMemory")
)

func main() {
	name := "Plants vs. Zombies"
	//name = "Clash for Windows"
	findWindowByName(name)
}

func findWindowByName(name string) {
	// 将字符串转换为Windows API需要的UTF-16格式
	titlePtr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 调用FindWindow函数获取窗口句柄
	hwnd, _, err := findWindow.Call(0, uintptr(unsafe.Pointer(titlePtr)))
	if hwnd == 0 {
		fmt.Println("Window not found.")
		return
	}

	fmt.Println("Window handle: ", hwnd)

	// 变量用于接收进程ID
	var processID uint32
	// 调用GetWindowThreadProcessId函数获取进程ID
	threadID, _, err := getProcessId.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&processID)))
	if threadID == 0 || processID == 0 {
		log.Panic("Failed to get process ID:", err)
	}
	log.Println("Thread ID:", threadID)
	log.Println("Process ID:", processID)
	access := uint32(0x1000) // PROCESS_ALL_ACCESS

	// 打开进程
	handle, _, err := procOpenProcess.Call(uintptr(access), uintptr(0), uintptr(processID))
	if handle == 0 {
		fmt.Println("Failed to open process:", err)
		return
	}
	//log.Panic(handle)
	// 要读取或写入的内存地址
	baseAddress := uintptr(0x1D0445D0) // 这需要你知道确切的地址
	buffer := make([]byte, 4)          // 读取/写入的数据缓冲区
	//bytesRead := uint32(0)
	bytesWritten := uint32(0)

	//// 读取内存
	//readSuccess, _, err := procReadProcessMemory.Call(
	//	handle,
	//	uintptr(baseAddress),
	//	uintptr(unsafe.Pointer(&buffer[0])),
	//	uintptr(len(buffer)),
	//	uintptr(unsafe.Pointer(&bytesRead)),
	//)
	//if readSuccess == 0 {
	//	fmt.Println("Failed to read memory:", err)
	//	return
	//}

	// 写入内存
	writeSuccess, _, err := procWriteProcessMemory.Call(
		handle,
		baseAddress,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	if writeSuccess == 0 {
		fmt.Println("Failed to write memory:", err)
		return
	}

	// 关闭进程句柄
	syscall.CloseHandle(syscall.Handle(handle))
}
