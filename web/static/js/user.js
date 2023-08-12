document.addEventListener('DOMContentLoaded', function() {
	var logoutButton = document.querySelector('.btn-logout');

	if (logoutButton) {
		logoutButton.addEventListener('click', function() {
			Cookies.remove("auth-session")
		});
	}
});
