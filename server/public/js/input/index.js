import { State } from "../enums.js";
import { handleInputsGameRunning } from "./handleInputsGameRunning.js";
/**
 * @param {KeyboardEvent} e 
 * @param {GameState} game_state 
 * @param {WebSocket} socket 
 * @param {Point} cursor
 * */
export function handleInput(e, game_state, socket, cursor) {
	console.log(game_state.State, State.Running);
	switch (game_state.State) {
		case State.WaitingForGuestConnection:
			break;
		case State.WaitingForPlayersToGetReady:
			handleInputsWaitingForPlayerToGetReady(e.key, socket);
			break;
		case State.Running:
			handleInputsGameRunning(e.key, socket, game_state, cursor);
			break;
		case State.Paused:
			break;
		case State.GameOver:
			break;
	}
}
/** 
 * @param {string} key 
 * @param {WebSocket} socket 
 * */
function handleInputsWaitingForPlayerToGetReady(key, socket) {
	switch (key) {
		case " ":
			const payload = {
				cmd: "set_ready",
			};
			const stringData = JSON.stringify(payload);
			socket.send(stringData);
			break;
	}
}
