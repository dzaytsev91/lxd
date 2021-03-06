package config

// DeviceNamed contains the name of a device and its config.
type DeviceNamed struct {
	Name   string
	Config Device
}

// DevicesSortable is a sortable slice of device names and config.
type DevicesSortable []DeviceNamed

func (devices DevicesSortable) Len() int {
	return len(devices)
}

func (devices DevicesSortable) Less(i, j int) bool {
	a := devices[i]
	b := devices[j]

	// First sort by types.
	if a.Config["type"] != b.Config["type"] {
		// In VMs, network interface names are derived from PCI
		// location. As a result of that, we must ensure that nic devices will
		// always show up at the same spot regardless of what other devices may be
		// added. Easiest way to do this is to always have them show up first.
		if a.Config["type"] == "nic" {
			return true
		}

		if b.Config["type"] == "nic" {
			return false
		}

		return a.Config["type"] < b.Config["type"]
	}

	// Special case disk paths.
	if a.Config["type"] == "disk" && b.Config["type"] == "disk" {
		if a.Config["path"] != b.Config["path"] {
			return a.Config["path"] < b.Config["path"]
		}
	}

	// Fallback to sorting by names.
	return a.Name < b.Name
}

func (devices DevicesSortable) Swap(i, j int) {
	devices[i], devices[j] = devices[j], devices[i]
}
