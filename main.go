package main

import (
	"github.com/zserge/webview"
	"log"
	"net/url"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("WebView Example")
	w.SetSize(640, 480, webview.HintNone)

	// Bind quit()
	w.Bind("quit", func() {
		log.Println("quit()")
		w.Terminate()
	})

	// Navigate (optional PathEscape to make sure this is not the problem).
	w.Navigate("data:text/html," + url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body>
			<h1>Hello, world!</h1>
			<p>This plays nice on Linux (Ubuntu 20.04) but not on Windows 10.</p>
			<button type="button" onclick="showInfo()">Info...</button>
			<button type="button" onclick="quit()">Quit</button>
 			
			<!-- Info-Box (hidden) -->
			<div id="infoBox" style="width: 600px; padding: 5px; 
				background-color: white; border: 2px solid #CCCCCC">
  				<p id="infoText">123<p>
  				<p style="text-align: center; margin-top: 20px">
    			<button type="button" onclick="hideInfo()">Close</button>
  			</p>
</div>

		</body>
		<script>
			var info = document.getElementById("infoBox");
			info.style.display = "none"; // Hide the box
			info.style.position = "absolute";
			info.style.zIndex = 999;
			info.style.marginTop = "10px";
 
			function showInfo() {
  				// Show box
				var p = document.getElementById("infoText");
				p.textContent = `+"`${navigator.userAgent}`"+`;
  				info.style.display = "";
			}

			function hideInfo() {
  				// Hide box
  				info.style.display = "none";
			}
		</script>
	</html>
	`))
	// Run it
	w.Run()
}
