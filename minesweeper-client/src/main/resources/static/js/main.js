function clickSq(event, thisTile) {
	if (!event) event = window.event;
	debugger;
	if (!event.shiftKey && event.which === 1 && event.button === 0) {
		play(thisTile.parentNode.rowIndex - 1, thisTile.cellIndex)
		console.log("Left Click")
	} else if (!event.shiftKey && event.which === 3 && event.button === 2) {
		mark(thisTile.parentNode.rowIndex - 1, thisTile.cellIndex)
		console.log("Right Click")
	} else if (event.shiftKey && event.which === 3 && event.button === 2) {
		console.log("Shift Right Click")
	} else {
		console.log("Invalid movement")
	}

	return false;
}

function play(row, column) {
	executeAnchor(`${basePath}/play?row=${row}&column=${column}`)
}

function mark(row, column) {
	executeAnchor(`${basePath}/mark?row=${row}&column=${column}`)
}



function executeAnchor(url) {
	document.getElementById("myAnchor").href = url
	window.location = document.getElementById("myAnchor").href;
}


function buildBasePath() {
	let url = window.location.href;
	let arr = url.split("/");
	let protocol = arr[0];
	let hostPort = arr[2];
	let contextRoot = arr[3];

	var basePath = `${protocol}//${hostPort}/${contextRoot}`
	return basePath;
}