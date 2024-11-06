import { drawMob } from "../mob.js";
import { drawPlayer } from "../player.js";
/** 
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
	* @param {number} working_width
	* @param {Point} cursor
*/
export function drawRunningGame(canvas, game_state, working_width = canvas.width * 0.9, cursor) {
	const ctx = canvas.getContext("2d");
	if (ctx === null || game_state === null) {
		return;
	}

	drawBoard(canvas, game_state, working_width, cursor);
	drawButtons(canvas, game_state, working_width);
	game_state.Players.forEach((player) => {
		if (game_state === null) {
			return;
		}
		drawPlayer(player, canvas, game_state);
	});
	game_state.Mobs.forEach((mob) => {
		if (game_state === null) {
			return;
		}
		drawMob(mob, canvas, game_state);
	});

	ctx.fillStyle = "white";
	ctx.font = "bold 18px Arial";
	ctx.textAlign = "center";
	ctx.textAlign = "left";
	let text = `Vida: ${game_state.Players[0].Health} | Moedas: ${game_state.Players[0].Coins
		} | Força: ${game_state.Players[0].Strength} | Movimentos Restantes: ${game_state.Players[0].MovesRemaining
		}`;
	ctx.fillText(text, working_width * 0.01, canvas.height * 0.93);
	text = `Vida: ${game_state.Players[1].Health} | Moedas: ${game_state.Players[1].Coins
		} | Força: ${game_state.Players[1].Strength} | Movimentos Restantes: ${game_state.Players[1].MovesRemaining
		}`;
	ctx.fillText(text, working_width * 0.01, canvas.height * 0.98);
}
/** 
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
	* @param {number} working_width
	* @param {Point} cursor
*/
function drawBoard(canvas, game_state, working_width, cursor) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	ctx.fillStyle = "black"; ctx.fillRect(0, 0, canvas.width, canvas.height);
	const worldWidth = game_state.WorldWidth;
	const worldHeight = game_state.WorldHeight;
	const squareWidth = working_width / worldWidth;
	for (let i = 0; i < worldWidth; i++) {
		for (let j = 0; j < worldHeight; j++) {
			//if (cursor.X === undefined || cursor.Y === undefined) {
			//	return
			//}
			if (cursor.X === i && cursor.Y === j) {
				ctx.fillStyle = "#1D6CF7";
				ctx.fillRect(
					i * squareWidth,
					j * squareWidth,
					squareWidth,
					squareWidth,
				);
			} else {
				ctx.strokeStyle = "white";
				ctx.strokeRect(
					i * squareWidth,
					j * squareWidth,
					squareWidth,
					squareWidth,
				);
			}
		}
	}
}
/**
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
	* @param {number} working_width
* */
function drawButtons(canvas, game_state, working_width) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	const xPos = (working_width + (canvas.width - working_width) / 2) * 0.97;
	let yPos = canvas.height * 0.05;
	const buttonWidth = (canvas.width - working_width) * 0.7;
	drawButton(ctx, xPos, yPos, buttonWidth, "Attack", "blue");
	yPos += 60;
	drawButton(ctx, xPos, yPos, buttonWidth, "Mover", "blue");
	yPos += 60;
	drawButton(ctx, xPos, yPos, buttonWidth, "Comprar", "blue");
	yPos += 60;
	drawButton(ctx, xPos, yPos, buttonWidth, "Finalizar", "blue");
	//ctx.drawImage(img, xPos, yPos, 50, 50);
}
/**
	* @param {CanvasRenderingContext2D} ctx
	* @param {number} posX
	* @param {number} posY
	* @param {number} width
	* @param {string} text
	* @param {string} color
*/
function drawButton(ctx, posX, posY, width, text, color) {
	ctx.fillStyle = color || "green";
	//ctx.beginPath();
	//ctx.arc(posX + width / 2, posY + width / 2, width * 0.5, 0, 2 * Math.PI);
	//ctx.fill();
	ctx.fillStyle = "color";
	ctx.font = `bold ${width * 0.3}px Arial`;
	ctx.textAlign = "center";
	ctx.textBaseline = "middle";
	ctx.fillText(text[0], posX + width / 2, posY);
	ctx.fillText(text, posX + width / 2, posY + width / 3);
}
