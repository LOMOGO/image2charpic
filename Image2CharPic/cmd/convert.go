package cmd

import (
	"Image2CharPic/convert"
	"github.com/spf13/cobra"
)

var convCmd = &cobra.Command{
	Use: "conv",
	Short: "图片转文字",
	Run: func(cmd *cobra.Command, args []string) {
		convert.Conv(source, target, text, size)
	},
}

var source, target, text string
var size int8

func init() {
	convCmd.Flags().StringVarP(&source,"source", "", "", "原图片的地址")
	convCmd.Flags().StringVarP(&target,"target", "", "", "字符图片的输出地址")
	convCmd.Flags().StringVarP(&text,"text", "", "hello", "图片文字化的文字(目前只支持英文、数字)")
	convCmd.Flags().Int8VarP(&size,"size", "", 7, "文字的大小")
}
