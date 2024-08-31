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
<form id='form'>
	<div class='search-form-wrapper'>
		<input type='text' name='search-input' placeholder='Search by breed' id='search-input'>
		<button type='button' onclick='submitForm()'>
			Submit
		</button>
	</div>
	<div class='search-select-wrapper'>
		<select id='search-select'>
			<option value='' disabled selected>Select dog breed</option>
		</select>
	</div>
	<div class='search-output'>
		
	</div>
</form>
	%s
</body>
</html>
`, MainStyles, MainScripts)
