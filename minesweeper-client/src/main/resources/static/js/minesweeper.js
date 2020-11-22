function clickSq(event, thisSquare) {
	if (!event) event = window.event;

    if (!event.which && event.button == 0) {
		alert("mouse-up ")
	} else if (event.shiftKey || event.button == 2) {
		alert("flag or unflag")
	} else {
	alert("nothing")
	}
	return false;
}