import { handleInput } from "./input.js";
import { Render } from "./render.js";

const fps = 30;
const frameInterval = 1000 / fps;
let lastFrameTime = 0;

// write type annotation in jsdocs for the variable bellow asserting that it cannot be null
/** @type {HTMLCanvasElement} */
const canvas = document.createElement("canvas");
const canvasHolder = document.querySelector('body > div');
if (canvasHolder !== null) {
	canvasHolder.appendChild(canvas);
}
canvas.classList.add("hidden");

let invite_code = "";
/** @type {?GameState} */
let game_state = null;
let this_player_id = "";
/** @type {WebSocket} */
let socket;

function connectSocket() {
	if (game_state === null) {
		return;
	}
	const socket_url =
		`ws://${window.location.host}/ws/${game_state.Id}/${this_player_id}`;
	socket = new WebSocket(socket_url);
	socket.onopen = function() {
		console.log("Websocket connected");
	};
	/**
	 * @param {{ data: string }} event
	 */
	function onMessage(event) {
		try {
			const data = JSON.parse(event.data);
			console.log(data);
			game_state = data;
		} catch (err) {
			console.log(err);
		}
	}
	socket.onmessage = onMessage;
}
async function JoinGame() {
	try {
		// @type {HTMLInputElement}
		const invite_code = /** @type {HTMLInputElement} */ (document.querySelector("#invite_code"));

		if (invite_code === null) {
			return;
		}

		const invite_code_value = invite_code.value;

		console.log(invite_code_value);
		const response = await fetch("/room/" + invite_code_value + "/join", {
			method: "POST",
		});

		const parsedResponse = await response.json();
		console.log(parsedResponse);
		game_state = parsedResponse;
		if (game_state !== null) {
			this_player_id = game_state.Players[1].Id;
		}
		connectSocket();
		BuildCanvas();
		setup();
		requestAnimationFrame(loop);
	} catch (error) {
		alert(error);
	}
}
async function StartGame() {
	try {
		const response = await fetch("/room", { method: "POST" });
		const parsedResponse = await response.json();
		invite_code = parsedResponse.invite_code;
		game_state = parsedResponse.game;
		if (game_state !== null) {
			this_player_id = game_state.Players[0].Id;
		}
		connectSocket();
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
/** @param {Mob} mob */
document.addEventListener("keydown", function(e) {
	if (game_state === null) {
		return
	}
	handleInput(e, game_state, socket);
});
/** @param {string} key */
const buttonStartGame = document.querySelector("#start_game");
if (buttonStartGame !== null) {
	buttonStartGame.addEventListener("click", StartGame);
}
const buttonJoinGame = document.querySelector("#join_game");
if (buttonJoinGame !== null) {
	buttonJoinGame.addEventListener("click", JoinGame);
}

