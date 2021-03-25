const popupForm = document.getElementById("popup-form-container");
const form = document.getElementById("popup-form");

function showPopupForm() {
	popupForm.style.display = "flex";

	// hide popup when background is clicked
	popupForm.addEventListener('click', (e) => {
		if (e.target.id == "popup-form-container") {
			popupForm.style.display = "none";
		}
	});
}

// raise while inputting text
if (/Android|webOS|iPhone|iPod/i.test(navigator.userAgent)) {
	form.addEventListener('mouseover', (e) => {
		if (e.target.className == "textbox") {
			form.style.marginTop = "-15rem";
		}
	});
}
// the opposite of above
if (/Android|webOS|iPhone|iPod/i.test(navigator.userAgent)) {
	form.addEventListener('mouseout', (e) => {
		if (e.target.className == "textbox") {
			form.style.marginTop = "0rem";
		}
	});
}
