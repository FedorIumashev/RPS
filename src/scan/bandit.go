package scan

import (
    "os"
    "os/exec"
)

func RunBandit(directory string) error {
    // Пример запуска Bandit
    cmd := exec.Command("bandit", "-r", directory)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
