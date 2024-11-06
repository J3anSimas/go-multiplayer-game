import { drawPlayer } from "./player.js";
import { drawMob } from "./mob.js";

/** 
 * @readonly
 * @enum {number}
*/
var Turn = {
	HostTurn: 0,
	GuestTurn: 1,
}
/**
 * @readonly
 * @enum {number}
 * */
var State = {
	WaitingForGuestConnection: 0,
	WaitingForPlayersToGetReady: 1,
	Running: 2,
	Paused: 3,
	GameOver: 4,
}

var ShopItemAttributeModifier = {
	StrengthAttribute: 0,
	MovementAttribute: 1,
	AttackVelocityAttribute: 2,
}
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
const ctx = canvas.getContext("2d");

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
		console.log(parsedResponse);
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
	/** @param {Status} game_state */
	switch (game_state.State) {
		case State.WaitingForGuestConnection:
			drawWaitingForGuest();
			break;
		case State.WaitingForPlayersToGetReady:
			drawWaitingForPlayerToGetReady();
			break;
		case State.Running:
			drawRunningGame();
			break;
		case State.GameOver:
			break;
	}
}
function drawWaitingForPlayerToGetReady() {
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
function drawWaitingForGuest() {
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
function drawRunningGame() {
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
/** @param {Mob} mob */
document.addEventListener("keydown", function(e) {
	if (game_state === null) {
		return
	}
	switch (game_state.State) {
		case 0:
			break;
		case 1:
			dealInputsWaitingForPlayerToGetReady(e.key);
			break;
	}
});
/** @param {string} key */
function dealInputsWaitingForPlayerToGetReady(key) {
	switch (key) {
		case " ":
			console.log("set ready");
			const payload = {
				cmd: "set_ready",
			};
			console.log(this_player_id);
			const stringData = JSON.stringify(payload);
			socket.send(stringData);
			break;
	}
}
const buttonStartGame = document.querySelector("#start_game");
if (buttonStartGame !== null) {
	buttonStartGame.addEventListener("click", StartGame);
}
const buttonJoinGame = document.querySelector("#join_game");
if (buttonJoinGame !== null) {
	buttonJoinGame.addEventListener("click", JoinGame);
}

