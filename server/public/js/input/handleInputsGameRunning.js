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

///**
// * Finds the shortest path from the player's position to the endpoint in a room grid.
// * @param {Player} player - The player object containing the current position.
// * @param {GameState} game_state - The room object with properties `worldHeight` and `worldWidth`.
// * @param {Point} endPoint - The endpoint to reach, with properties `x` and `y`.
// * @returns {Array} A two-element array where the first element is the step count to reach 
// *                  the endpoint, and the second element is the path as an array of points.
// *                  If no path exists, returns [-1, null].
// */
//function findShortestPath(player, game_state, endPoint) {
//	const nRows = game_state.WorldHeight;
//	const nCols = game_state.WorldWidth;
//
//	const queue = [player.Position];
//	const visited = Array.from({ length: nRows }, () => Array(nCols).fill(false));
//	visited[player.Position.X][player.Position.Y] = true;
//
//	/** @type {Point[][]} */
//	const steps = Array.from({ length: nRows }, () => Array(nCols).fill(0));
//	/** @type {Point[][]} */
//	const previous = Array.from({ length: nRows }, () => Array(nCols).fill(null));
//
//	const directions = [
//		{ x: -1, y: 0 }, { x: 1, y: 0 }, // Up, Down
//		{ x: 0, y: -1 }, { x: 0, y: 1 }, // Left, Right
//		{ x: -1, y: -1 }, { x: -1, y: 1 }, // Diagonal Up-Left, Up-Right
//		{ x: 1, y: -1 }, { x: 1, y: 1 }   // Diagonal Down-Left, Down-Right
//	];
//
//	while (queue.length > 0) {
//		const currentPoint = queue.shift();
//		if (currentPoint !== undefined) {
//			if (currentPoint.X === endPoint.X && currentPoint.Y === endPoint.Y) {
//				// Reconstruct path
//				let path = [];
//				for (let pth = currentPoint; pth !== player.Position; pth = previous[pth.X][pth.Y]) {
//					path.push(pth);
//				}
//				path.push(player.Position);
//				path.reverse();
//				return [steps[currentPoint.X][currentPoint.Y], path.slice(1)];
//			}
//		}
//
//		for (const direction of directions) {
//			if (currentPoint === undefined) {
//				break;
//			}
//			const newX = currentPoint.X + direction.x;
//			const newY = currentPoint.Y + direction.y;
//
//			if (isValid(newX, newY, nRows, nCols, game_state) && !visited[newX][newY]) {
//				visited[newX][newY] = true;
//				queue.push({ X: newX, Y: newY });
//				steps[newX][newY] = steps[currentPoint.X][currentPoint.Y] + 1;
//				previous[newX][newY] = currentPoint;
//			}
//		}
//	}
//
//	return [-1, null];
//}
//
///**
// * Checks if a point is valid within the room boundaries and does not hit obstacles.
// * @param {number} x - The x-coordinate.
// * @param {number} y - The y-coordinate.
// * @param {number} nRows - The total number of rows in the room.
// * @param {number} nCols - The total number of columns in the room.
// * @param {Object} room - The room object.
// * @returns {boolean} True if the point is valid and unvisited.
// */
//function isValid(x, y, nRows, nCols, room) {
//	return x >= 0 && x < nRows && y >= 0 && y < nCols && !room.isObstacle(x, y);
//}
