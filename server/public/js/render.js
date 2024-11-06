import { drawMob } from "./mob.js";
import { drawPlayer } from "./player.js";
import { State } from "./enums.js";

/** 
 * @param {HTMLCanvasElement} canvas 
 * @param {GameState} game_state 
 * @param {string} invite_code
 * */
export function Render(canvas, game_state, invite_code) {
	switch (game_state.State) {
		case State.WaitingForGuestConnection:
			drawWaitingForGuest(canvas, invite_code);
			break;
		case State.WaitingForPlayersToGetReady:
			drawWaitingForPlayerToGetReady(canvas, game_state);
			break;
		case State.Running:
			drawRunningGame(canvas, game_state);
			break;
		case State.GameOver:
			break;
	}
}
/** 
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
*/
function drawWaitingForPlayerToGetReady(canvas, game_state) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	if (game_state === null) {
		return;
	}
	ctx.fillStyle = "#7d4c2d";
	ctx.fillRect(0, 0, canvas.width, canvas.height);
	ctx.fillStyle = "white";
	ctx.font = "bold 30px Arial";
	ctx.textAlign = "center";
	ctx.textBaseline = "middle";
	ctx.fillText("Aguardando Pronto", canvas.width / 2, canvas.height / 2 - 50);
	ctx.font = "bold 18px Arial";
	ctx.fillText(
		"Pressione ESPAÇO quando estiver pronto",
		canvas.width / 2,
		canvas.height / 2,
	);

	ctx.textAlign = "left";
	let text = game_state.Players[0].Ready ? "Pronto" : "Não Pronto";
	ctx.fillText(text, canvas.width * 0.1, canvas.height * 0.9);
	ctx.textAlign = "right";
	text = game_state.Players[1].Ready ? "Pronto" : "Não Pronto";
	ctx.fillText(text, canvas.width * 0.9, canvas.height * 0.9);
}

/** 
	* @param {HTMLCanvasElement} canvas
	* @param {string} invite_code
*/
function drawWaitingForGuest(canvas, invite_code) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	ctx.fillStyle = "#7d4c2d";
	ctx.fillRect(0, 0, canvas.width, canvas.height);
	ctx.fillStyle = "white";
	ctx.font = "bold 30px Arial";
	ctx.textAlign = "center";
	ctx.textBaseline = "middle";
	ctx.fillText(
		"AGUARDANDO JOGADOR",
		canvas.width / 2,
		canvas.height / 2 - 50,
	);
	ctx.fillText(
		"INVITE CODE: " + invite_code.toUpperCase(),
		canvas.width / 2,
		canvas.height / 2,
	);
}

/** 
	* @param {HTMLCanvasElement} canvas
	* @param {GameState} game_state
*/
function drawRunningGame(canvas, game_state) {
	const ctx = canvas.getContext("2d");
	if (ctx === null) {
		return;
	}
	if (game_state === null) {
		return;
	}
	ctx.fillStyle = "black";
	ctx.fillRect(0, 0, canvas.width, canvas.height);
	ctx.strokeStyle = "white";
	const working_height = canvas.height * 0.9;
	const working_width = canvas.width * 0.9;
	//ctx.strokeRect(1, 1, working_width, working_height)
	const worldWidth = game_state.WorldWidth;
	const worldHeight = game_state.WorldHeight;
	const squareWidth = working_width / worldWidth;
	for (let i = 0; i < worldWidth; i++) {
		for (let j = 0; j < worldHeight; j++) {
			ctx.strokeRect(
				i * squareWidth,
				j * squareWidth,
				squareWidth,
				squareWidth,
			);
		}
	}
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
