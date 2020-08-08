package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zombispormedio/fizzbuzz-server/pkg/fizzbuzz"
)

var rootCmd = &cobra.Command{
	Use:   "fizzbuzz",
	Short: "A generator for fizzbuzz",
	Long:  `Cobra is a CLI that generate fizzbuzz ranges and convert numbers to fizzbuzz.`,
}

var convertCmd = &cobra.Command{
	Use:  "convert [number]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Invalid number")
			return
		}

		fmt.Println(fizzbuzz.ItoFizzBuzz(value))
	},
}

var rangeCmd = &cobra.Command{
	Use: "range [from] [to]",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			result := fizzbuzz.CreateFizzbuzzRange()
			fmt.Println(strings.Join(result, "\n"))
			return
		}

		from, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid number")
			return
		}

		if len(args) == 1 {
			result := fizzbuzz.CreateFizzbuzzRange(from)
			fmt.Println(strings.Join(result, "\n"))
			return
		}

		to, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("Invalid number")
			return
		}

		result := fizzbuzz.CreateFizzbuzzRange(from, to)
		fmt.Println(strings.Join(result, "\n"))
	},
}

func main() {
	rootCmd.AddCommand(rangeCmd)
	rootCmd.AddCommand(convertCmd)
	rootCmd.Execute()
}
