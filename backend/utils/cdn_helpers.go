package utils

func GetWrappedHtmlContent(contentSrcUrl string) string {
	// Create a wrapped HTML with an iframe
	wrappedHTML := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Video Embed</title>
			<style>
				iframe {
					width: 100%;
					height: 100vh; /* Use viewport height for the iframe */
					border: none;
				}
				html, body {
					margin: 0;
					padding: 0;
					width: 100%;
					height: 100%;
					overflow: hidden; /* Hide overflow to prevent scrollbars */
				}
			</style>
		</head>
		<body>
			<iframe src="` + contentSrcUrl + `" allowFullScreen></iframe>
		</body>
		</html>`

	return wrappedHTML
}
