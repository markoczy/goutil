package files

import
(
	"os"
)

func Exists(aPath string) (bool, error) {
	f, err := os.Open(aPath)
	defer f.Close()
	if os.IsNotExist(err) { return false, nil }
	if err!= nil { return false, err }
	return true, nil
}


func Stats(aPath string) (os.FileInfo, error) {
	f, err := os.Open(aPath)
	defer f.Close()
	if err!= nil { return nil, err }
	stats, err := f.Stat()
	if err!= nil { return nil, err }
	return stats, nil
}

func Size(aPath string) (int64, error) {
	stats, err := Stats(aPath)
	if err!= nil { return -1, err }
	return stats.Size(), nil
}
