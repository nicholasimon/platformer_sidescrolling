package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//	backimg rl.Texture2D
	// player
	followplayer                               bool
	player, playerh, playerv, playerx, playery int
	// draw map
	drawblock, drawblocknext, drawblocknexth, drawblocknextv int
	// level
	levelw   = 1000
	levelh   = 120
	levela   = levelh * levelw
	levelmap = make([]string, levela)
	// core
	screenw, screenh, screena            int
	monh32, monw32                       int32
	monitorh, monitorw, monitornum       int
	grid16on, grid4on, debugon, lrg, sml bool
	framecount                           int
	mousepos                             rl.Vector2
	camera                               rl.Camera2D
)

// MARK: notes
/*
1920X1080 / 16px blocks
width 128 blocks (127+1)
height 68 blocks (67+1)

*/
func timers() { // MARK: timers

}
func updateall() { // MARK: updateall()
	if followplayer {
		updatedrawblockplayer()
	}
	if grid16on {
		grid16()
	}
	if grid4on {
		grid4()
	}
	timers()
	gravity()
	getpositions()
}
func gravity() { // MARK: gravity()
	if levelmap[player+levelw] == "." {

		player += levelw

	}
}
func updatedrawblockplayer() { // MARK: updatedrawblockplayer()
	if playerh > levelh-(screenh/2) {
		drawblocknext = player - (levelw * (screenh - 5))
		drawblocknext -= ((screenw / 2) - 1)
	} else {
		drawblocknext = player - (levelw * (screenh / 2))
		drawblocknext -= ((screenw / 2) - 1)
	}
}
func getpositions() { // MARK: getpositions()

	drawblocknexth = drawblocknext / levelw
	drawblocknextv = drawblocknext - (drawblocknexth * levelw)
	playerh = player / levelw
	playerv = player - (playerh * levelw)
}
func createlevel() { // MARK: createlevel()

	for a := 0; a < levela; a++ {
		levelmap[a] = "."
	}
	block := levela - levelw
	block -= levelw * 4
	for a := 0; a < (levelw * 4); a++ {
		levelmap[block] = "_"
		block++
	}
	block = levela - levelw
	block -= levelw * 8
	count := block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice() >= 11 {
			block = count
			platlen := rInt(8, 15)
			for b := 0; b < platlen; b++ {
				levelmap[block] = "#"
				if rolldice()+rolldice() == 12 {
					levelmap[block] = "*"
				}
				block++
			}
		}
		count++
	}
	block = levela - levelw
	block -= levelw * 12
	count = block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice() >= 11 {
			block = count
			platlen := rInt(8, 15)
			for b := 0; b < platlen; b++ {
				levelmap[block] = "#"
				if rolldice()+rolldice() == 12 {
					levelmap[block] = "*"
				}
				block++
			}
		}
		count++
	}
	block = levela - levelw
	block -= levelw * 16
	count = block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice() >= 11 {
			block = count
			platlen := rInt(8, 15)
			for b := 0; b < platlen; b++ {
				levelmap[block] = "#"
				if rolldice()+rolldice() == 12 {
					levelmap[block] = "*"
				}
				block++
			}
		}
		count++
	}
	block = levela - levelw
	block -= levelw * 20
	count = block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice()+rolldice() >= 17 {
			block = count
			platlen := rInt(6, 11)
			for b := 0; b < platlen; b++ {
				levelmap[block] = "#"
				if rolldice()+rolldice() == 12 {
					levelmap[block] = "*"
				}
				block++
			}
		}
		count++
	}
	block = levela - levelw
	block -= levelw * 24
	count = block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice()+rolldice() >= 17 {
			block = count
			platlen := rInt(6, 11)
			for b := 0; b < platlen; b++ {
				levelmap[block] = "#"
				if rolldice()+rolldice() == 12 {
					levelmap[block] = "*"
				}
				block++
			}
		}
		count++
	}

	block = levela - levelw
	block -= levelw * 30
	count = block
	for a := 0; a < (levelw); a++ {
		if rolldice()+rolldice()+rolldice() >= 17 {
			block = count
			roomlen := rInt(8, 15)
			rooma := roomlen * roomlen
			count2 := 0
			for b := 0; b < rooma; b++ {
				levelmap[block] = "#"
				block++
				count2++
				if count2 == roomlen {
					count2 = 0
					block -= levelw + roomlen
				}
			}
			block = count + 1
			block -= levelw
			roomlen -= 2
			rooma = roomlen * roomlen
			count2 = 0
			for b := 0; b < rooma; b++ {
				levelmap[block] = "."
				block++
				count2++
				if count2 == roomlen {
					count2 = 0
					block -= levelw + roomlen
				}
			}
			a += roomlen + 2
			count += roomlen + 2
		}
		count++
	}

}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	startsettings()
	createlevel()
	raylib()
}
func startplayer() { // MARK: startplayer()
	player = levela - levelw
	player -= levelw * 5
	player += screenw / 2

	drawblocknext = player
	drawblocknext -= levelw * (screenh - 5)
	drawblocknext -= ((screenw / 2) - 1)
}
func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "plat4mo turbo VII DX")
	setscreen()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "plat4mo turbo VII DX")

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	//	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		//mousepos = rl.GetMousePosition()
		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// rl.DrawTexture(backimg, 0, 0, rl.Red) // MARK: draw backimg
		rl.BeginMode2D(camera)

		// MARK: draw map layer 1
		count := 0
		drawx := int32(0)
		drawy := int32(0)
		drawblock = drawblocknext
		for a := 0; a < screena; a++ {

			checklevel := levelmap[drawblock]

			switch checklevel {
			case ".":
				//		rl.DrawRectangle(drawx, drawy, 15, 15, rl.Green)
			case "#":
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Red)
			case "*":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Red)
			case "_":
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Brown)
			}

			if player == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
				playerx = int(drawx)
				playery = int(drawy)
			}

			count++
			drawx += 16
			drawblock++
			if count == screenw {
				drawblock += levelw - screenw
				count = 0
				drawy += 16
				drawx = 0
			}

		}

		// MARK: draw map layer 2
		rl.EndMode2D() // MARK: draw no camera
		if debugon {
			debug()
		}
		rl.EndDrawing()
		input()
		updateall()

	}
	rl.CloseWindow()
}
func setscreen() { // MARK: setscreen()
	monitornum = rl.GetMonitorCount()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	setsizes()
}
func setsizes() { // MARK: setsizes()
	if monitorw >= 1600 {
		lrg = true
		sml = false
	} else if monitorw < 1600 && monitorw >= 1280 {
		lrg = false
		sml = true
	}

	screenw = (monitorw / 16) + 1
	screenh = (monitorh / 16) + 1
	screena = screenw * screenh
	startplayer()

}
func startsettings() { // MARK: start
	camera.Zoom = 1.0
	camera.Target.X = 0.0
	camera.Target.Y = 0.0
	//debugon = true
	grid16on = true
	//selectedmenuon = true
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Black, 0.7))
	rl.DrawFPS(monw32-290, monh32-100)

	screenhTEXT := strconv.Itoa(screenh)
	screenwTEXT := strconv.Itoa(screenw)
	playerxTEXT := strconv.Itoa(playerx)
	playeryTEXT := strconv.Itoa(playery)
	playerhTEXT := strconv.Itoa(playerh)
	playervTEXT := strconv.Itoa(playerv)

	rl.DrawText(screenwTEXT, monw32-290, 10, 10, rl.White)
	rl.DrawText("screenw", monw32-200, 10, 10, rl.White)
	rl.DrawText(screenhTEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("screenh", monw32-200, 20, 10, rl.White)
	rl.DrawText(playerxTEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("playerx", monw32-200, 30, 10, rl.White)
	rl.DrawText(playeryTEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("playery", monw32-200, 40, 10, rl.White)
	rl.DrawText(playerhTEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("playerh", monw32-200, 50, 10, rl.White)
	rl.DrawText(playervTEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("playerv", monw32-200, 60, 10, rl.White)
}
func input() { // MARK: keys input
	if rl.IsKeyPressed(rl.KeyUp) {
		if playerh > screenh {
			player -= levelw * 4
		}
	}
	if rl.IsKeyDown(rl.KeyRight) {
		if playerv < levelw-((screenw/2)+1) {
			player++
		}
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		if playerv > screenw/2+1 {
			player--
		}
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		if followplayer {
			followplayer = false
		} else {
			followplayer = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
			camera.Target.Y = float32(float32(playery) - (float32(monitorh) / 2.3))
			camera.Target.X = float32(playerx - (monitorw / 4))
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 4.0
			camera.Target.Y = float32(playery - (monitorh / 5))
			camera.Target.X = float32(playerx - (monitorw / 8))
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
			camera.Target.Y = 0
			camera.Target.X = 0
		} else if camera.Zoom == 4.0 {
			camera.Zoom = 2.0
			camera.Target.Y = float32(float32(playery) - (float32(monitorh) / 2.3))
			camera.Target.X = float32(playerx - (monitorw / 4))
		}
	}
	if rl.IsKeyDown(rl.KeyKp8) {
		if drawblocknexth > 0 {
			drawblocknext -= levelw
		}
	}
	if rl.IsKeyDown(rl.KeyKp2) {
		if drawblocknexth < levelh-(screenh+1) {
			drawblocknext += levelw
		}
	}
	if rl.IsKeyDown(rl.KeyKp6) {
		if drawblocknextv < levelw-(screenw+1) {
			drawblocknext++
		}
	}
	if rl.IsKeyDown(rl.KeyKp4) {
		if drawblocknextv > 0 {
			drawblocknext--
		}
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if grid16on {
			grid16on = false
		} else {
			grid16on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if grid4on {
			grid4on = false
		} else {
			grid4on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

}
func grid16() { // MARK: grid16()
	for a := 0; a < monitorw; a += 16 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.Green, 0.1))
	}
	for a := 0; a < monitorh; a += 16 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.Green, 0.1))
	}
}
func grid4() { // MARK: grid4()
	for a := 0; a < monitorw; a += 4 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.DarkGreen, 0.1))
	}
	for a := 0; a < monitorh; a += 4 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.DarkGreen, 0.1))
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
