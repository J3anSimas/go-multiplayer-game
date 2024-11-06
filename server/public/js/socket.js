
/**
 * @param {GameState} game_state
 * @param {string} this_player_id
 * @returns {WebSocket}
 */
export function connectSocket(game_state, this_player_id) {
	const socket_url =
		`ws://${window.location.host}/ws/${game_state.Id}/${this_player_id}`;
	const socket = new WebSocket(socket_url);
	socket.onopen = function() {
		console.log("Websocket connected");
	};
	/**
	 * @param {{ data: string }} event
	 */
	function onMessage(event) {
		try {
			const data = JSON.parse(event.data);
			Object.assign(game_state, data);
			//game_state = data;
			console.log(game_state);
		} catch (err) {
			console.log(err);
		}
	}
	socket.onmessage = onMessage;
	return socket;
}
