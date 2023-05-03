package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ledongthuc/pdf"
	"github.com/spf13/cobra"
)

func isPDFFile(path string) bool {
	return strings.ToLower(filepath.Ext(path)) == ".pdf"
}

func searchInFile(path string, search string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	f_info, _ := f.Stat()
	f_size := f_info.Size()
	r, err := pdf.NewReader(f, int64(f_size))
	if err != nil {
		fmt.Printf("Error reading PDF file %s: %s\n", path, err)
		return
	}
	if r == nil {
		fmt.Printf("Nil reader for PDF file %s\n", path)
		return
	}

	for i := 1; i <= r.NumPage(); i++ {
		page := r.Page(i)
		text, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		if strings.Contains(text, search) {
			fmt.Printf("Found search string '%s' in file '%s'\n", search, path)
			return
		}
	}

}

func searchInDirectory(dir string, search string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			if isPDFFile(entry.Name()) {
				wg.Add(1)
				go searchInFile(filepath.Join(dir, entry.Name()), search, wg)
			}
		} else if strings.Contains(entry.Name(), "C:\\$Recycle.Bin\\") {
			continue
		} else {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go searchInDirectory(subdir, search, wg)
		}
	}
}

var (
	search string
	dir    string
)

var rootCmd = &cobra.Command{
	Use:   "pdfsearch",
	Short: "Search for a string in PDF files in a directory",
	Long: `pdfsearch is a CLI tool that searches for a given string in all PDF files in a directory
and its subdirectories.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if search == "" {
			fmt.Println("error: no search string specified")
			os.Exit(1)
		}

		if dir == "" {
			fmt.Println("error: no directory path provided")
			os.Exit(1)
		}

		err := os.Chdir(dir)
		if err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go searchInDirectory(".", search, &wg)
		wg.Wait()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "the string to search for")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "the directory to search in")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
