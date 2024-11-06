import { PlayerState, State } from './enums.js';
import { handleInput } from "./input/index.js";
import { Render } from "./render/index.js";
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

/** @type {boolean} */
let isHost = false;
/** @type {WebSocket} */
let socket;
/** @type {number} */
const player_state = PlayerState.IsMoving;

/** @type {Point} */
let cursor = {
	X: 0,
	Y: 0
};
async function InitializeGame(isJoin = false) {
	try {
		let response, parsedResponse;
		const invite_code_input = /** @type {HTMLInputElement} */ (document.querySelector("#invite_code"));
		if (isJoin) {
			if (invite_code_input === null) {
				return;
			}
			response = await fetch(`/room/${invite_code_input.value}/join`, { method: "POST" });
		} else {
			response = await fetch("/room", { method: "POST" });
			isHost = true;
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
		if (isHost) {
			cursor = {
				X: game_state.Players[0].Position.X,
				Y: game_state.Players[0].Position.Y
			}
		} else {
			cursor = {
				X: game_state.Players[1].Position.X,
				Y: game_state.Players[1].Position.Y
			}
		}
		socket = connectSocket(game_state, this_player_id);
		if (socket === null) {
			throw new Error("Socket is null");
		}

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
	canvas.height = window.innerHeight;
	canvas.width = window.innerHeight;
}
function setup() {
	canvas.classList.remove("hidden");
}
function update() {
	canvas.height = window.innerHeight;
	canvas.width = window.innerHeight;
	if (game_state === null) {
		return;
	}
	game_state = mockGameState;
	Render(canvas, game_state, invite_code, cursor);

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
	handleInput(e, game_state, socket, cursor);
});
const mockGameState = {
	"Id": "bc81b13a-2c98-463b-b061-3ac9361f0154",
	"State": State.Running,
	"WorldWidth": 20,
	"WorldHeight": 20,
	"Players": [
		{
			"Id": "75b4c873-07cd-4298-9a0e-2bfd520858e7",
			"Ready": true,
			"IsHost": true,
			"Position": {
				"X": 0,
				"Y": 0
			},
			"Health": 10,
			"Coins": 0,
			"MoveCapacity": 10,
			"MovesRemaining": 10,
			"Strength": 10,
			"TotalShots": 10,
			"ShotsRemaining": 10,
			"IsDead": false
		},
		{
			"Id": "4acf7330-d562-4899-a242-9491bf27408c",
			"Ready": true,
			"IsHost": false,
			"Position": {
				"X": 19,
				"Y": 19
			},
			"Health": 10,
			"Coins": 0,
			"MoveCapacity": 10,
			"MovesRemaining": 10,
			"Strength": 10,
			"TotalShots": 10,
			"ShotsRemaining": 10,
			"IsDead": false
		}
	],
	"Mobs": [
		{
			"Health": 15,
			"Position": {
				"X": 6,
				"Y": 2
			},
			"Strength": 5,
			"CoinsToDrop": 5
		},
		{
			"Health": 15,
			"Position": {
				"X": 14,
				"Y": 18
			},
			"Strength": 5,
			"CoinsToDrop": 5
		},
		{
			"Health": 15,
			"Position": {
				"X": 10,
				"Y": 0
			},
			"Strength": 5,
			"CoinsToDrop": 5
		},
		{
			"Health": 15,
			"Position": {
				"X": 10,
				"Y": 19
			},
			"Strength": 5,
			"CoinsToDrop": 5
		}
	],
	"ShopItems": [],
	"Turn": 0
}
