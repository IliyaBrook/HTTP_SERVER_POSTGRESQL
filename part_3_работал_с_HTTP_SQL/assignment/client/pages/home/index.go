package pages

import "fmt"

//goland:noinspection JSUnresolvedReference
var MainHtml = fmt.Sprintf(
	//language=html
	`
<!DOCTYPE html>
<html lang='eng'>
	<head>
		<meta charset='UTF-8'>
		<meta name='viewport' content='width=device-width, initial-scale=1.0'>
		<title>Home Page</title>
		<style>
			%s
		</style>
	</head>
<body>
<h1>Submit a request</h1>
<form id='testForm'>
	<label for='data'/>
	<input type='text' name='data' id='data' placeholder='Enter something'>
	<button type='button' onclick='submitForm()'>Submit</button>
</form>
	%s
</body>
</html>
`, MainStyles, MainScripts)
