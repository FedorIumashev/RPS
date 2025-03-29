// File: src/testing_file.go
// testing_file.go - Файл для тестирования безопасности кода
package src

import (
    "bytes"
    "os/exec"
)

// Функция запуска Bandit для одного файла
func RunBanditFile(file string) (string, error) {
	cmd := exec.Command("bandit", "-f", "json", "-o", "bandit-report.json", file)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
