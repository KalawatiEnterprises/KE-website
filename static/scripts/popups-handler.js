const popupForm = document.getElementById("popup-form-container");

function showPopupForm() {
	popupForm.style.display = "flex";

	// hide popup when background is clicked
	popupForm.addEventListener('click', (e) => {
		if (e.target.id == "popup-form-container") {
			popupForm.style.display = "none";
		}
	});
}

if (/Android|webOS|iPhone|iPod/i.test(navigator.userAgent)) {
	popupForm.addEventListener('mouseover', (e) => {
		if (e.target.className == "textbox") {
			// make it move the form upwards
			popupForm.style.display = "none";
		}
	});
}
