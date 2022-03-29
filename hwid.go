package hwid

import (
	"fmt"
	"math/big"
	"net"
	"runtime"
)

var (
	// This map contains the preferred primary ethernet interface name from which we
	// want to get the ethernet hardware address from in order to identify the gateway.
	// Note that this is dependent on OS version and configuration, so the same OS may
	// have different names.  But these are guesses.  If they do not work then the user
	// will have to specify the interface used to determine hardware identity on the
	// command line.
	interfacePreferenceMap = map[string]string{
		"arm/linux":     "eth0",
		"amd64/darwin":  "en0",
		"arm64/darwin":  "en0",
		"windows/386":   "Ethernet",
		"windows/amd64": "Ethernet",
		"windows/arm":   "Ethernet",
		"windows/arm64": "Ethernet",
	}
)

// ID returns an ID based on the ethernet address of the primary
// network interface adapter. You can optionally specify which
// adapter as a string parameter (it will ignore all but the first).
func ID(interfaceName ...string) (string, error) {
	ifaceName := interfacePreferenceMap[runtime.GOARCH+"/"+runtime.GOOS]

	if len(interfaceName) > 0 && interfaceName[0] != "" {
		ifaceName = interfaceName[0]
	}

	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return "", fmt.Errorf("interface [%s]: %v", ifaceName, err)
	}

	if len(iface.HardwareAddr) < 6 {
		return "", fmt.Errorf("weird hardware address: %s", iface.HardwareAddr.String())
	}

	return big.NewInt(0).SetBytes(iface.HardwareAddr).Text(36), nil
}
