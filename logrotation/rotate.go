package logrotation

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Rotate
func (R *Rotor) Rotate() (err error) {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	return R.rotateInsecure()
}

// Rotate
func (R *Rotor) rotateInsecure() (err error) {
	if !fileExists(R.filepath) {
		fh, err := os.OpenFile(R.filepath, os.O_SYNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, R.Permissions)
		if err != nil {
			return fmt.Errorf("cannot create new logfile: %v", err)
		}
		R.file = fh
		return nil
	}

	if R.Retention > 0 {
		dir := filepath.Dir(R.filepath) + string(os.PathSeparator)
		basename := filepath.Base(R.filepath)

		length := strconv.Itoa(int(math.Ceil(float64(R.Retention) / 10)))
		formatstring := "%s.%0" + length + "d"

		for i := R.Retention; i >= 0; i-- {
			movetarget := dir + fmt.Sprintf(formatstring, basename, i+1) + R.compressExt
			filename := dir + fmt.Sprintf(formatstring, basename, i) + R.compressExt
			if !fileExists(filename) && i != 0 {
				continue
			}

			if i == R.Retention {
				err = os.Remove(filename)
				if err != nil {
					return fmt.Errorf("cannot remove oldest file: %v", err)
				}
				continue
			}

			if i == 0 {
				err = R.file.Close()
				if err != nil {
					return fmt.Errorf("cannot close logfile for rotation: %v", err)
				}

				sourceFile, err := os.OpenFile(R.filepath, os.O_EXCL|os.O_RDONLY, R.Permissions)
				if err != nil {
					return fmt.Errorf("cannot open logfile for rotation: %v", err)
				}

				destinationFile, err := os.OpenFile(movetarget, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, R.Permissions)
				if err != nil {
					return fmt.Errorf("cannot open archive logfile: %v", err)
				}

				err = R.compressor(destinationFile, sourceFile)
				if err != nil {
					return fmt.Errorf("could not compress file: %v", err)
				}

				err = sourceFile.Close()
				if err != nil {
					return fmt.Errorf("cannot close sourcefile after rotation: %v", err)
				}
				err = destinationFile.Close()
				if err != nil {
					return fmt.Errorf("cannot close destinationfile after rotation: %v", err)
				}

				fh, err := os.OpenFile(R.filepath, R.fileFlags|os.O_TRUNC, R.Permissions)
				if err != nil {
					return fmt.Errorf("cannot create new logfile: %v", err)
				}
				R.file = fh
				return nil
			}

			err = os.Rename(filename, movetarget)
			if err != nil {
				return fmt.Errorf("cannot move file '%s' to '%s': %v", filename, movetarget, err)
			}
		}
	} else {
	}
	return
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
