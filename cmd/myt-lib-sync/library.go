package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/mytmlt/myt-lib-sync/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	libPath string
	libType string
)

const librariesKey = "libraries"

var libraryCommand = &cobra.Command{
	Use:   "library",
	Short: "Manage download libraries",
}

var libraryCreateCmd = &cobra.Command{
	Use: "create", Short: "Create a library",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lib := config.LibraryConfig{Type: libType, Path: libPath, CreatedAt: time.Now()}
		err := createLib(args[0], lib)
		exitOnError(err)
	},
}

var libraryAddCmd = &cobra.Command{
	Use: "add", Short: "Add media to a library",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var libraryListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List libraries",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		libs, err := readLibs()
		exitOnError(err)
		printLibs(libs)
	},
}

var libraryDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm", "d"},
	Short:   "Delete a library (doesn't delete the directory and its contents)",
	Args:    cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		libs, err := readLibs()
		exitOnError(err)

		_, ok := libs[args[0]]

		if !ok {
			err := fmt.Errorf("library with name %v not found", args[0])
			exitOnError(err)
		}

		delete(libs, args[0])

		viper.Set(librariesKey, libs)
		err = viper.WriteConfig()
		exitOnError(err)
		printLibs(libs)
	},
}

func createLib(name string, lib config.LibraryConfig) error {
	libs, err := readLibs()
	if err != nil {
		return err
	}

	lib.Path, err = homedir.Expand(lib.Path)
	if err != nil {
		return err
	}

	for n, l := range libs {
		if lib.Path == l.Path {
			return fmt.Errorf("a library named %q already exists at path %q", n, lib.Path)
		}
	}
	_, ok := libs[name]

	if ok {
		return fmt.Errorf("library with name %q already exists", name)
	}

	libs[name] = lib
	viper.Set(librariesKey, libs)
	err = viper.WriteConfig()

	return err
}

func readLibs() (map[string]config.LibraryConfig, error) {
	libs := make(map[string]config.LibraryConfig)

	if err := viper.UnmarshalKey(librariesKey, &libs); err != nil {
		return nil, fmt.Errorf("reading libraries config: %w", err)
	}

	return libs, nil
}

func printLibs(libs map[string]config.LibraryConfig) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Name", "Path", "Type", "Create at"})

	names := make([]string, 0, len(libs))

	for name := range libs {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		table.Append([]string{
			name, libs[name].Path, libs[name].Type, libs[name].CreatedAt.Local().String(),
		})
	}

	table.Render()
}

func init() {
	libraryCommand.AddCommand(libraryCreateCmd, libraryListCmd, libraryDeleteCmd, libraryAddCmd)
	libraryCreateCmd.Flags().StringVarP(&libPath, "path", "p", ".", "The output dir of the library")
	libraryCreateCmd.Flags().StringVarP(&libType, "type", "t", "video", "Type of content video or music")
	rootCmd.AddCommand(libraryCommand)
}
