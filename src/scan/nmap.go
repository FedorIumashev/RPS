package scan

import (
    "os"
    "os/exec"
)

func RunNmap(directory string) error {
    // Пример запуска Nmap
    cmd := exec.Command("nmap", "-sP", directory)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
