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
func (r *Rotor) Rotate() (err error) {
	r.fileMtx.Lock()
	defer r.fileMtx.Unlock()

	return r.rotateInsecure()
}

// rotateInsecure is doing the actual rotating. It exists to not unlock the
// mutex between last write and rotation.
func (r *Rotor) rotateInsecure() (err error) {
	if !fileExists(r.filepath) {
		fh, err := os.OpenFile(r.filepath, os.O_SYNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, r.Permissions)
		if err != nil {
			return fmt.Errorf("cannot create new logfile: %w", err)
		}
		r.file = fh

		return nil
	}

	if r.Retention > 0 {
		return r.rotateWithRetention()
	}

	return r.rotateWithRemain()
}

// rotateWithRetention keeps entire logfiles and just does a common file
// rotation where the original logfile is compressed, and older archived
// logs have their number increased. If a log exceeds the retention limit after
// the rotation, it is deleted.
func (r *Rotor) rotateWithRetention() (err error) {
	for i := r.Retention; i >= 0; i-- {
		err = r.rotateFile(i)
		if err != nil {
			return
		}
	}

	return nil
}

func (r *Rotor) rotateFile(i int) (err error) {
	dir := filepath.Dir(r.filepath) + string(os.PathSeparator)
	basename := filepath.Base(r.filepath)
	length := strconv.Itoa(int(math.Ceil(float64(r.Retention) / 10)))
	formatstring := "%s.%0" + length + "d"

	movetarget := dir + fmt.Sprintf(formatstring, basename, i+1) + r.compressExt
	filename := dir + fmt.Sprintf(formatstring, basename, i) + r.compressExt
	if !fileExists(filename) && i != 0 {
		return nil
	}

	if i == r.Retention {
		err = os.Remove(filename)
		if err != nil {
			return fmt.Errorf("cannot remove oldest file: %w", err)
		}

		return nil
	}

	if i == 0 {
		err = r.file.Close()
		if err != nil {
			return fmt.Errorf("cannot close logfile for rotation: %w", err)
		}

		sourceFile, err := os.OpenFile(r.filepath, os.O_EXCL|os.O_RDONLY, r.Permissions)
		if err != nil {
			return fmt.Errorf("cannot open logfile for rotation: %w", err)
		}

		destinationFile, err := os.OpenFile(movetarget, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, r.Permissions)
		if err != nil {
			return fmt.Errorf("cannot open archiwe logfile: %w", err)
		}

		err = r.compressor(destinationFile, sourceFile)
		if err != nil {
			return fmt.Errorf("could not compress file: %w", err)
		}

		err = sourceFile.Close()
		if err != nil {
			return fmt.Errorf("cannot close sourcefile after rotation: %w", err)
		}
		err = destinationFile.Close()
		if err != nil {
			return fmt.Errorf("cannot close destinationfile after rotation: %w", err)
		}

		fh, err := os.OpenFile(r.filepath, r.fileFlags|os.O_TRUNC, r.Permissions)
		if err != nil {
			return fmt.Errorf("cannot create new logfile: %w", err)
		}
		r.file = fh

		return nil
	}

	err = os.Rename(filename, movetarget)
	if err != nil {
		return fmt.Errorf("cannot move file '%s' to '%s': %v", filename, movetarget, err)
	}

	return nil
}

// rotateWithRemain takes the last x% (x being the set KeptPercent in the
// struct), truncates the logfile and writes the kept percentage of the logfile
// back. Thereby creating more of a floating-window log.
func (r *Rotor) rotateWithRemain() (err error) {
	err = r.file.Close()
	if err != nil {
		return fmt.Errorf("cannot close logfile for rotation: %w", err)
	}

	sourceFile, err := os.OpenFile(r.filepath, os.O_EXCL|os.O_RDONLY, r.Permissions)
	if err != nil {
		return fmt.Errorf("cannot open logfile for rotation: %w", err)
	}

	fi, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("cannot stat sourcefile: %w", err)
	}

	offset := int64(float64(fi.Size()) * (1 - float64(r.KeptPercent)/100))
	_, err = sourceFile.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("cannot seek to percentage: %w", err)
	}

	b := []byte{0}

	for b[0] != '\n' && err == nil {
		_, err = sourceFile.Read(b)
	}

	buf := bytes.NewBuffer([]byte{})

	_, err = buf.ReadFrom(sourceFile)
	if err != nil {
		return fmt.Errorf("cannot read old data: %w", err)
	}

	fh, err := os.OpenFile(r.filepath, r.fileFlags|os.O_TRUNC, r.Permissions)
	if err != nil {
		return fmt.Errorf("cannot create new logfile: %w", err)
	}

	_, err = fh.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("cannot put old content into file: %w", err)
	}

	r.file = fh

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
