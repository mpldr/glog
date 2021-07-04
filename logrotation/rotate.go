package logrotation

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// Rotate triggers a rotation, independent of whether the limit size is reached
// or not.
func (R *Rotor) Rotate() (err error) {
	R.fileMtx.Lock()
	defer R.fileMtx.Unlock()

	return R.rotateInsecure()
}

// rotateInsecure is doing the actual rotating. It exists to not unlock the
// mutex between last write and rotation.
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
		return R.rotateWithRetention()
	}
	return R.rotateWithRemain()
}

// rotateWithRetention keeps entire logfiles and just does a common file
// rotation where the original logfile is compressed, and older archived
// logs have their number increased. If a log exceeds the retention limit after
// the rotation, it is deleted.
func (R *Rotor) rotateWithRetention() (err error) {
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
			break
		}

		err = os.Rename(filename, movetarget)
		if err != nil {
			return fmt.Errorf("cannot move file '%s' to '%s': %v", filename, movetarget, err)
		}
	}
	return nil
}

// rotateWithRemain takes the last x% (x being the set KeptPercent in the
// struct), truncates the logfile and writes the kept percentage of the logfile
// back. Thereby creating more of a floating-window log.
func (R *Rotor) rotateWithRemain() (err error) {
	err = R.file.Close()
	if err != nil {
		return fmt.Errorf("cannot close logfile for rotation: %v", err)
	}

	sourceFile, err := os.OpenFile(R.filepath, os.O_EXCL|os.O_RDONLY, R.Permissions)
	if err != nil {
		return fmt.Errorf("cannot open logfile for rotation: %v", err)
	}

	fi, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat sourcefile: %v", err)
	}

	offset := int64(float64(fi.Size()) * (1 - float64(R.KeptPercent)/100))
	_, err = sourceFile.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("cannot seek to percentage: %v", err)
	}

	b := []byte{0}

	for b[0] != '\n' && err == nil {
		_, err = sourceFile.Read(b)
	}

	buf := bytes.NewBuffer([]byte{})

	_, err = buf.ReadFrom(sourceFile)
	if err != nil {
		return fmt.Errorf("cannot read old data: %v", err)
	}

	fh, err := os.OpenFile(R.filepath, R.fileFlags|os.O_TRUNC, R.Permissions)
	if err != nil {
		return fmt.Errorf("cannot create new logfile: %v", err)
	}

	_, err = fh.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("cannot put old content into file: %v", err)
	}

	R.file = fh

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
