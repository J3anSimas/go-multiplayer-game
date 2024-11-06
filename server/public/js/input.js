import { State } from "./enums.js";
/**
 * @param {KeyboardEvent} e 
 * @param {GameState} game_state 
 * @param {WebSocket} socket 
 * */
export function handleInput(e, game_state, socket) {
	switch (game_state.State) {
		case State.WaitingForGuestConnection:
			break;
		case State.WaitingForPlayersToGetReady:
			dealInputsWaitingForPlayerToGetReady(e.key, socket);
			break;
		case State.Running:
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
function dealInputsWaitingForPlayerToGetReady(key, socket) {
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
