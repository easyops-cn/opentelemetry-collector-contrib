package infoscraper

type System struct {
	Distribution string
	Arch         string
	Version      string
}

func getSystemInfo() (System, error) {
	return _getSystemInfo()
}
