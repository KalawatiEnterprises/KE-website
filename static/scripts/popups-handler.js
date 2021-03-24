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
