package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	HexOld  = "6e533e406a45f0b6372f3ea1071700000c7120127cd915cef8ed1a3f2c5b"
	HexNew  = "785782391ad0b9169f17415dd35f00002790175204e3aa65ea10cff20818"
	License = `
name: swaince
company: relax`
)

var cfg = &CrtConf{}
var crtCmd = &cobra.Command{
	Use:   "crt",
	Short: "active SecureCRT",
	Long:  "active SecureCRT",
	Run:   active,
}

func init() {
	rootCmd.AddCommand(crtCmd)
	crtCmd.Flags().StringVar(&cfg.Path, "path", DefaultPath(), "--path /usr/local/bin/SecureCRT")
	crtCmd.Flags().StringVar(&cfg.HexOld, "old", HexOld, fmt.Sprintf("--old %s", HexOld))
	crtCmd.Flags().StringVar(&cfg.HexNew, "new", HexNew, fmt.Sprintf("--old %s", HexNew))
}

func active(cmd *cobra.Command, args []string) {

	f, err := ioutil.ReadFile(cfg.Path)
	if err != nil {
		panic(err)
	}

	err = backFile()

	e := hex.EncodeToString(f)
	r := strings.ReplaceAll(e, HexOld, HexNew)

	c, err := hex.DecodeString(r)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(cfg.Path, c, fs.ModeExclusive)
	if err != nil {
		panic(err)
	}
	fmt.Println(License)
}

func backFile() error {
	// back file
	return os.Rename(cfg.Path, cfg.Path+".back")
}

type CrtConf struct {
	Path   string
	HexOld string
	HexNew string
}

func DefaultPath() (path string) {
	switch runtime.GOOS {
	case "linux":
		path = "/usr/bin/SecureCRT"
	case "windows":
		//path = `C:/Program Files/VanDyke Software/Clients/SecureCRT.exe`
		path = `D:/soft/VanDykeSoftware/Clients/SecureCRT.exe`
	case "darwin":
		path = `/Applications/SecureCRT.app/Contents/MacOS/SecureCRT`
	}
	return path
}
