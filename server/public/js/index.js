import { handleInput } from "./input.js";
import { Render } from "./render.js";
import { connectSocket } from "./socket.js";

const fps = 30;
const frameInterval = 1000 / fps;
let lastFrameTime = 0;

// write type annotation in jsdocs for the variable bellow asserting that it cannot be null
/** @type {HTMLCanvasElement} */
const canvas = document.createElement("canvas");
const canvasHolder = document.querySelector('body > div');
if (canvasHolder === null) {
	throw new Error("Canvas holder not found");
}
canvasHolder.appendChild(canvas);
canvas.classList.add("hidden");

let invite_code = "";
/** @type {?GameState} */
let game_state = null;
let this_player_id = "";
/** @type {WebSocket} */
let socket;

async function InitializeGame(isJoin = false) {
	try {
		let response, parsedResponse;
		// @type {HTMLInputElement}
		const invite_code_input = /** @type {HTMLInputElement} */ (document.querySelector("#invite_code"));
		if (isJoin) {
			if (invite_code_input === null) {
				return;
			}
			response = await fetch(`/room/${invite_code_input.value}/join`, { method: "POST" });
		} else {
			response = await fetch("/room", { method: "POST" });
		}

		parsedResponse = await response.json();
		game_state = parsedResponse.game || parsedResponse;

		if (game_state === null) {
			throw new Error("Game state is null");
		}
		this_player_id = isJoin ? game_state.Players[1].Id : game_state.Players[0].Id;
		invite_code = isJoin ? invite_code_input.value : parsedResponse.invite_code;
		if (game_state === null) {
			throw new Error("Game state is null");
		}
		socket = connectSocket(game_state, this_player_id);
		BuildCanvas();
		setup();
		requestAnimationFrame(loop);

	} catch (error) {
		alert(error);
	}
}
/** @param {number} timestamp */
function loop(timestamp) {
	const timeSinceLastFrame = timestamp - lastFrameTime;
	if (timeSinceLastFrame >= frameInterval) {
		update();
		lastFrameTime = timestamp;
	}
	requestAnimationFrame(loop);
}
function BuildCanvas() {
	canvas.classList.remove("hidden");
	const initial_buttons = document.querySelector("#initial_buttons");
	if (initial_buttons !== null) {
		initial_buttons.classList.add("hidden");
	}
	canvas.width = 600;
	canvas.height = 600;
}
function setup() {
	canvas.classList.remove("hidden");
}
function update() {
	if (game_state === null) {
		return;
	}
	Render(canvas, game_state, invite_code);

}
const buttonStartGame = document.querySelector("#start_game");
if (buttonStartGame !== null) {
	buttonStartGame.addEventListener("click", InitializeGame.bind(null, false));
}
const buttonJoinGame = document.querySelector("#join_game");
if (buttonJoinGame !== null) {
	buttonJoinGame.addEventListener("click", InitializeGame.bind(null, true));
}

document.addEventListener("keydown", function(e) {
	if (game_state === null) {
		return
	}
	handleInput(e, game_state, socket);
});
