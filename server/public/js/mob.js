
/**
	* @param {Mob} mob
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
*/
export function drawMob(mob, canvas, game_state) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	const working_height = canvas.height * 0.9;
	const working_width = canvas.width * 0.9;
	//ctx.strokeRect(1, 1, working_width, working_height)
	const worldWidth = game_state.WorldWidth;
	const worldHeight = game_state.WorldHeight;
	const squareWidth = working_width / worldWidth;
	ctx.fillStyle = "green";
	ctx.beginPath();
	const posX = working_width * mob.Position.X / worldWidth + squareWidth / 2;
	const posY = working_height * mob.Position.Y / worldHeight +
		squareWidth / 2;

	ctx.arc(posX, posY, (squareWidth / 2) * 0.7, 0, 2 * Math.PI);
	ctx.fill();
}
