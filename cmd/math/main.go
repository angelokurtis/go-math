package main

import (
	"fmt"
	"github.com/angelokurtis/go-math"
	"github.com/spf13/cobra"
	"strconv"
)

func main() {
	cmd := &cobra.Command{Use: "math"}
	cmd.AddCommand(&cobra.Command{
		Use:  "add",
		Args: numberArgs,
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			check(err)

			b, err := strconv.Atoi(args[1])
			check(err)

			fmt.Printf("%d + %d = %d\n", a, b, go_math.Add(a, b))
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:  "sub",
		Args: numberArgs,
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			check(err)

			b, err := strconv.Atoi(args[1])
			check(err)

			fmt.Printf("%d - %d = %d\n", a, b, go_math.Sub(a, b))
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:  "mul",
		Args: numberArgs,
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			check(err)

			b, err := strconv.Atoi(args[1])
			check(err)

			fmt.Printf("%d * %d = %d\n", a, b, go_math.Mul(a, b))
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:  "div",
		Args: numberArgs,
		Run: func(cmd *cobra.Command, args []string) {
			a, err := strconv.Atoi(args[0])
			check(err)

			b, err := strconv.Atoi(args[1])
			check(err)

			fmt.Printf("%d / %d = %d\n", a, b, go_math.Div(a, b))
		},
	})
	_ = cmd.Execute()
}

func numberArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("accepts %d arg(s), received %d", 2, len(args))
	}
	for _, arg := range args {
		_, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("accepts numbers, received %q", arg)
		}
	}
	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
