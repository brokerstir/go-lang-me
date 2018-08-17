package files

import "os"

func Size(fileName string) (int64, error) {
  fileInfo, err := os.Stat(fileName)
  if err != nil {
    // RETURN THE VALUES 0 AND err
    return 0, err
  } else {
  	// RETURN THE VALUES fileInfo.Size() AND nil
  	return fileInfo.Size(), nil
  }
}
