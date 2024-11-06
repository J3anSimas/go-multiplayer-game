/** 
 * @param {string} key 
 * @param {WebSocket} socket 
 * @param {GameState} game_state
 * @param {Point} cursor
 * */
export function handleInputsGameRunning(key, socket, game_state, cursor) {
	console.log(key);
	switch (key) {
		case "w":
		case "ArrowUp":
			cursor.Y != 0 ? cursor.Y-- : cursor.Y;
			break;
		case "s":
		case "ArrowDown":
			cursor.Y != game_state.WorldHeight ? cursor.Y++ : cursor.Y;
			break;
		case "a":
		case "ArrowLeft":
			cursor.X != 0 ? cursor.X-- : cursor.X;
			break;
		case "d":
		case "ArrowRight":
			cursor.X != game_state.WorldHeight ? cursor.X++ : cursor.X;
			break;
		case "f":
			const payload = {
				cmd: "end_turn",
			};
			const stringData = JSON.stringify(payload);
			socket.send(stringData);
			break;
	}
}
