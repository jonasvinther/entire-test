package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var beavers = map[string]string{
	"classic": `
    __  _
   / _)/ )
  / /   /
 ( (   /
  \ \_/|
   \__ |
   |  ||
   |  ||___
   |  |    )
   |__|___/
  (______)
`,
	"chewing": `
        _     _
       ( \---/ )
       (  o o  )
       /| = |\
      / |___| \
     (  (   )  )
      '-'   '-'
     /  \ | /  \
    /    \|/    \
   (  ~~ /_\ ~~  )
    \___/   \___/
       |     |
      _|     |_
     [_]     [_]
`,
	"swimming": `
        _.---._
      .'       '.
     /   0   0   \
    |      __     |
    |    /    \   |
     \  |  __ |  /
   ~~~'-.____.-'~~~
  ~  ~  ~  ~  ~  ~
`,
	"builder": `
      __. __.
     (  \/ \/  )
     |\  /\  /|
     | \/  \/ |
     |  \  /  |
     |___|___|
    /    |||    \
   /   ==|||==   \
  /_____|_|_|_____\
        /   \
       / ~~~ \
      |_______|
`,
	"tiny": `
  (\ /)
  ( . .)
  c(")(")
`,
	"sunglasses": `
  (\ /)
  (-0-0-)
  c(")(")
`,
}

var beaverSayings = []string{
	"Dam, that's good!",
	"I chew what I want!",
	"Building dreams, one log at a time.",
	"Semi-aquatic and proud!",
	"Gnaw your problems away!",
	"Keep calm and build dams.",
	"Lodge goals only.",
	"Wood you believe it?",
}

var beaverCmd = &cobra.Command{
	Use:   "beaver",
	Short: "Draw ASCII beavers",
	Long:  "Draw a magnificent ASCII beaver in your terminal.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var beaverDrawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draw an ASCII beaver",
	Long:  "Draw a magnificent ASCII beaver in your terminal. Choose from different styles and colors.",
	RunE: func(cmd *cobra.Command, args []string) error {
		style, _ := cmd.Flags().GetString("style")
		color, _ := cmd.Flags().GetString("color")
		rainbow, _ := cmd.Flags().GetBool("rainbow")
		say, _ := cmd.Flags().GetString("say")
		random, _ := cmd.Flags().GetBool("random")

		art, ok := beavers[style]
		if !ok {
			return fmt.Errorf("unknown style %q, available: classic, chewing, swimming, builder, tiny", style)
		}

		if random {
			styles := []string{"classic", "chewing", "swimming", "builder", "tiny"}
			art = beavers[styles[rand.Intn(len(styles))]]
		}

		colorCode := resolveColor(color)

		if say == "" && random {
			say = beaverSayings[rand.Intn(len(beaverSayings))]
		}

		if say != "" {
			bubble := speechBubble(say)
			fmt.Println(colorCode + bubble + colorReset)
		}

		if rainbow {
			printRainbow(art)
		} else {
			fmt.Print(colorCode + art + colorReset)
		}

		return nil
	},
}

var beaverListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available beaver styles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available beaver styles:")
		for name := range beavers {
			fmt.Printf("  - %s\n", name)
		}
	},
}

func init() {
	beaverDrawCmd.Flags().StringP("style", "s", "classic", "Beaver style (classic, chewing, swimming, builder, tiny)")
	beaverDrawCmd.Flags().StringP("color", "c", "", "Color (red, green, yellow, blue, purple, cyan, white)")
	beaverDrawCmd.Flags().BoolP("rainbow", "r", false, "Rainbow mode!")
	beaverDrawCmd.Flags().StringP("say", "m", "", "Make the beaver say something")
	beaverDrawCmd.Flags().BoolP("random", "R", false, "Pick a random style and saying")

	beaverCmd.AddCommand(beaverDrawCmd)
	beaverCmd.AddCommand(beaverListCmd)
	rootCmd.AddCommand(beaverCmd)
}
