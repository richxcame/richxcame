/*
Copyright © 2025 Baygeldi Cholukov <baygeldicolukov@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	length       int
	useSymbol    bool
	useDigit     bool
	useUppercase bool
	useLowercase bool
)

const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var passgenCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Password generator",
	Long:  `Generate password as you want, and save it to your clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword(length, useSymbol, useDigit, useUppercase, useLowercase)
		if err := copyToClipboard(password); err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("Password copied to clipboard")
	},
}

func init() {
	rootCmd.AddCommand(passgenCmd)

	passgenCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	passgenCmd.Flags().BoolVarP(&useSymbol, "symbol", "s", true, "Use symbols")
	passgenCmd.Flags().BoolVarP(&useDigit, "digit", "d", true, "Use digits")
	passgenCmd.Flags().BoolVarP(&useUppercase, "uppercase", "u", true, "Use uppercase")
	passgenCmd.Flags().BoolVarP(&useLowercase, "lowercase", "c", true, "Use lowercase")
}

func generatePassword(length int, symbol, digit, uppercase, lowercase bool) string {
	var charset string

	if lowercase {
		charset += LowerLetters
	}
	if uppercase {
		charset += UpperLetters
	}
	if digit {
		charset += Digits
	}
	if symbol {
		charset += Symbols
	}

	if len(charset) == 0 {
		return "Error: No character sets selected"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

// copyToClipboard tries to copy text to the system clipboard
func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = in.Write([]byte(text))
	if err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
