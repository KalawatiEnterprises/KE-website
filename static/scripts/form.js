// copied code idk jquery
// will re-write this properly later
$(document).ready(function () {
	$("form").submit(function (event) {
		var formData = {
			name: $("#fullname").val(),
			email: $("#email").val(),
			phone: $("#phone").val(),
			message: $("#message").val(),
		};
	$.ajax({
		type: "POST",
		url: "contact",
		data: formData,
		dataType: "json",
		encode: true,
	}).done(function (data) {
		if (data == "ye") return showFormConfirmation()
	});
		event.preventDefault();
	});
});

const showFormConfirmation = () => {
	console.log("recorded the response")

	document.getElementById("popup-form-container").style.display = "none";
	const box = document.getElementById("form-messagebox")
	box.style.display = "inline";
	box.innerHTML = "<p>Your response has been recorded!</p>";

	setTimeout(() => {
		console.log("setTimeout block entered")
		box.style.display = "none";
		box.innerHTML = "";
		document.getElementById("contact-button").style.color = "gray";
	}, 3000);
}
