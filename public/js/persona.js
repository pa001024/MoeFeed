// persona login
$.ajax({
	url: "https://login.persona.org/include.js",
	dataType: "script",
	cache: true,
	success: function() {
		window.currentUser = window.currentUser || null;//'genarate-dynamically@bycode.com';
		navigator.id.watch({
			loggedInUser: currentUser,
			onlogin: function(assertion){
				$.ajax({
					type: 'POST',
					url: '/loginbrowserid',
					data: { assertion: assertion },
					success: function(res) {
						window.location.reload();
					},
					error: function(res) {
						if (res.status == 302) {
							location = res.getResponseHeader("Location");
							return
						}
						alert("登入失败");
						console.log("登入失败", res);
					}
				});
			},
			onlogout: function(){
				$.ajax({
					type: 'GET',
					url: '/logout',
					success: function(res) {
						window.location.reload();
					},
					error: function(res) {
						if (res.status == 302) {
							location = res.getResponseHeader("Location");
							return
						}
						alert("登出失败");
						console.log("登出失败", res);
					}
				});
			},
		});
		var c = $(".persona-login");
		c.length && c.addClass("persona-loaded").click(function(d) {
			navigator.id.request();
			return false
		});
		c = $(".persona-logout");
		c.length && c.addClass("persona-loaded").click(function(d) {
			navigator.id.logout();
			return false
		});
	}
});