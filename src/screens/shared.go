package shared

import (
	"example/raylib-game/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Font rl.Font
var SecondaryFont rl.Font
var FxClick rl.Sound

const (
	FontSmallTextSize  float32 = 16
	FontMediumTextSize float32 = 24
	FontBigTextSize    float32 = 32
	FontHugeTextSize   float32 = 42
)

var AppSettings settings.Settings
var Scores settings.Scores

// Logo variables
var LogoIcon rl.Texture2D
var IconRect rl.Rectangle
var TextRect rl.Rectangle

// Gamepad variables
var gamepadButtonCooldown float32

const (
	Logo int = iota
	Title
	Options
	Gameplay
	Ending
	Leaderboard
)

const (
	ButtonUnchanged int = iota
	ButtonConfirm
	ButtonGoBack
)

const Unchanged int = -1

// Load the shared assets
func LoadSharedAssets() error {
	// Set up the font
	Font = rl.LoadFont("resources/fonts/montserrat_semibold.ttf")
	rl.GenTextureMipmaps(&Font.Texture)
	rl.SetTextureFilter(Font.Texture, rl.FilterBilinear)

	SecondaryFont = rl.LoadFont("resources/fonts/cartograph_cf_italic.ttf")
	rl.GenTextureMipmaps(&SecondaryFont.Texture)
	rl.SetTextureFilter(SecondaryFont.Texture, rl.FilterBilinear)

	// Load the necessary settings and scores
	if err := AppSettings.LoadFromFile(); err != nil {
		return err
	}

	if err := Scores.LoadFromFile(); err != nil {
		return err
	}

	// Logo textures
	LogoIcon = rl.LoadTexture("resources/icons/logo_old.png")
	IconRect = rl.NewRectangle(30, 25, 45, 45)
	TextRect = rl.NewRectangle(82, 27, 250, 50)

	return nil
}

// A function used to navigate the UI using keyboard buttons
func UpdateMovement(current int, availableButtons int) (int, int) {
	if gamepadButtonCooldown <= 0.0 {
		switch rl.GetGamepadButtonPressed() {
		case 3: // PS3 gamepad down
			current++
			if current == availableButtons {
				current = 0
			}
		case 1: // PS3 gamepad up
			current--
			if current == -1 {
				current = availableButtons - 1
			}
		case 7: // PS3 gamepad confirm
			return current, ButtonConfirm
		case 6: // PS3 gamepad go back
			return current, ButtonGoBack
		}

		gamepadButtonCooldown = 0.2
	} else {
		if gamepadButtonCooldown > 0 {
			gamepadButtonCooldown -= 0.01
		}
	}

	switch rl.GetKeyPressed() {
	case rl.KeyDown, rl.KeyTab:
		current++
		if current == availableButtons {
			current = 0
		}
	case rl.KeyUp:
		current--
		if current == -1 {
			current = availableButtons - 1
		}
	case rl.KeyEnter:
		return current, ButtonConfirm
	case rl.KeyEscape:
		return current, ButtonGoBack
	}

	return current, ButtonUnchanged
}
