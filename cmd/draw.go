package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

var rainbowColors = []string{colorRed, colorYellow, colorGreen, colorCyan, colorBlue, colorPurple}

var unicorns = map[string]string{
	"classic": `
        /\_____/\
       /  o   o  \
      ( ==  ^  == )
       )         (
      (           )
     ( (  )   (  ) )
    (__(__)___(__)__)
          ||
         /  \
        /    \
       / ,  , \
      /_/|  |\_\
`,
	"running": `
              ,,
             /  \__
    _______/   @  \___
   /            \_____\
  / /\  /\       \    |
 / /  \/  \  ,,  /    |
(_/ /\  /\ \/  \/    /
   /_/  \_\        _/
          \       /
           \_____/
`,
	"majestic": `
                    /
               ,--'~
          __,-'     \
    _,--''            \
   /      .  . ' .    |
  |    .  /|. |\ |\   |
  |   /  / || | \| \  |
   \ |  |  ||_|  |  | |
    \|  |  /  \  |  |/
     \  | /    \ |  /
      \_|/      \|_/
        \   /\  /
         '-'  ''
`,
	"tiny": `
  /\   /\
 (  o o  )
 =( Y )=
   )   (
  (_)-(_)
`,
}

var sayings = []string{
	"Believe in magic!",
	"Stay sparkly!",
	"You are one in a million!",
	"Keep galloping towards your dreams!",
	"Rainbow power!",
	"Spread glitter, not hate!",
	"Born to be magical!",
	"Life is short, eat the rainbow!",
}

var drawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draw an ASCII unicorn",
	Long:  "Draw a magical ASCII unicorn in your terminal. Choose from different styles and colors.",
	RunE: func(cmd *cobra.Command, args []string) error {
		style, _ := cmd.Flags().GetString("style")
		color, _ := cmd.Flags().GetString("color")
		rainbow, _ := cmd.Flags().GetBool("rainbow")
		say, _ := cmd.Flags().GetString("say")
		random, _ := cmd.Flags().GetBool("random")

		art, ok := unicorns[style]
		if !ok {
			return fmt.Errorf("unknown style %q, available: classic, running, majestic, tiny", style)
		}

		if random {
			styles := []string{"classic", "running", "majestic", "tiny"}
			art = unicorns[styles[rand.Intn(len(styles))]]
		}

		colorCode := resolveColor(color)

		if say == "" && random {
			say = sayings[rand.Intn(len(sayings))]
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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available unicorn styles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available unicorn styles:")
		for name := range unicorns {
			fmt.Printf("  - %s\n", name)
		}
	},
}

func resolveColor(name string) string {
	switch strings.ToLower(name) {
	case "red":
		return colorRed
	case "green":
		return colorGreen
	case "yellow":
		return colorYellow
	case "blue":
		return colorBlue
	case "purple", "magenta":
		return colorPurple
	case "cyan":
		return colorCyan
	case "white":
		return colorWhite
	default:
		return colorReset
	}
}

func printRainbow(art string) {
	lines := strings.Split(art, "\n")
	for i, line := range lines {
		color := rainbowColors[i%len(rainbowColors)]
		fmt.Println(color + line + colorReset)
	}
}

func speechBubble(text string) string {
	padding := 2
	width := len(text) + padding*2
	top := " " + strings.Repeat("_", width)
	middle := fmt.Sprintf("< %s >", text)
	bottom := " " + strings.Repeat("-", width)
	tail := "        \\"
	return fmt.Sprintf("%s\n%s\n%s\n%s", top, middle, bottom, tail)
}

func init() {
	drawCmd.Flags().StringP("style", "s", "classic", "Unicorn style (classic, running, majestic, tiny)")
	drawCmd.Flags().StringP("color", "c", "", "Color (red, green, yellow, blue, purple, cyan, white)")
	drawCmd.Flags().BoolP("rainbow", "r", false, "Rainbow mode!")
	drawCmd.Flags().StringP("say", "m", "", "Make the unicorn say something")
	drawCmd.Flags().BoolP("random", "R", false, "Pick a random style and saying")

	rootCmd.AddCommand(drawCmd)
	rootCmd.AddCommand(listCmd)
}
