package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ListTimezones() {
	zoneInfoPath := "/usr/share/zoneinfo"
	err := filepath.Walk(zoneInfoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, _ := filepath.Rel(zoneInfoPath, path)
			fmt.Println(relPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error listing timezones: %v\n", err)
		os.Exit(1)
	}
}

func SetTimezone(timezone string) {
	tzPath := fmt.Sprintf("/usr/share/zoneinfo/%s", timezone)
	if _, err := os.Stat(tzPath); os.IsNotExist(err) {
		fmt.Printf("Timezone '%s' does not exist\n", timezone)
		return
	}

	// direct symlink creation (works on some systems)
	if err := os.Symlink(tzPath, "/etc/localtime"); err == nil {
		if updateErr := updateTimezoneFile(timezone); updateErr == nil {
			fmt.Printf("Set timezone in /etc/timezone to: %s\n", timezone)
			fmt.Println("Note: /etc/localtime not updated - system may need reboot")
			return
		}
		fmt.Printf("Successfully set timezone to: %s\n", timezone)
		return
	}

	// timedatectl
	if _, err := exec.LookPath("timedatectl"); err == nil {
		cmd := exec.Command("timedatectl", "set-timezone", timezone)
		if output, err := cmd.CombinedOutput(); err == nil {
			fmt.Printf("Successfully set timezone to: %s\n", timezone)
			return
		} else {
			fmt.Printf("timedatectl failed: %s\n", output)
		}
	}

	// modify through /etc/timezone only
	if err := updateTimezoneFile(timezone); err == nil {
		fmt.Printf("Set timezone in /etc/timezone to: %s\n", timezone)
		fmt.Println("Note: /etc/localtime not updated - system may need reboot")
		return
	}

	fmt.Println("All timezone update methods failed")
}

func updateTimezoneFile(timezone string) error {
	return os.WriteFile("/etc/timezone", []byte(timezone+"\n"), 0644)
}

func GetCurrentTimezone() {
	if link, err := os.Readlink("/etc/localtime"); err == nil {
		if strings.Contains(link, "/usr/share/zoneinfo/") {
			tz := strings.SplitN(link, "/usr/share/zoneinfo/", 2)[1]
			fmt.Printf("Current timezone: %s\n", tz)
			return
		}
		fmt.Printf("Current timezone symlink: %s\n", link)
		return
	}

	if data, err := os.ReadFile("/etc/timezone"); err == nil {
		fmt.Printf("Current timezone: %s", string(data))
		return
	}

	fmt.Println("Could not determine current timezone")
	os.Exit(1)
}
