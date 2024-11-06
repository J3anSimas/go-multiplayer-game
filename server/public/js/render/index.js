import { State } from "../enums.js";
import { drawRunningGame } from "./draw_running_game.js";

/** 
 * @param {HTMLCanvasElement} canvas 
 * @param {GameState} game_state 
 * @param {string} invite_code
 * @param {Point} cursor
 * */
export function Render(canvas, game_state, invite_code, cursor) {
	switch (game_state.State) {
		case State.WaitingForGuestConnection:
			drawWaitingForGuest(canvas, invite_code);
			break;
		case State.WaitingForPlayersToGetReady:
			drawWaitingForPlayerToGetReady(canvas, game_state);
			break;
		case State.Running:
			drawRunningGame(canvas, game_state, canvas.width * 0.9, cursor);
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

