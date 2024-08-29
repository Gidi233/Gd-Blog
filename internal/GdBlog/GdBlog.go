// Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Gidi233/Gd-Blog.

package GdBlog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGdBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// specify the name of the command
		Use: "Gd-Blog",
		// specify the short description of the command
		Short: "Gd-Blog is a simple blog system",
		// specify the long description of the command
		Long: `Gd-Blog is a simple but not easy blog system,
Find more Gd-Blog information at:
	https://github.com/Gidi233/Gd-Blog#readme`,

		// when error occurs, the command will not print usage information
		SilenceUsage: true,
		// specify the run function to execute when cmd.Execute() is called
		// if the function fails, an error message will be returned
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		//
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	return cmd
}

func run() error {
	fmt.Println("Hello, Gd-Blog!")
	return nil
}
