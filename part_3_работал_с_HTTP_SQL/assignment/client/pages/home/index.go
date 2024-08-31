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
	<div class='searchFormWrapper'>
	    <div class='searchForm'>
    	    	<div class='search-input-wrapper'>
    	    	    <label for='data'/>
            		<input type='text' name='data' id='data' placeholder='Enter something'>
    	    	<div/>
        		<div class='search-button-wrapper'>
        		    <button type='button' onclick='submitForm()'>Submit</button>
        		</div>
    	</div>
		<div class='content-output'>
				
		</div>
	</div>
</form>
	%s
</body>
</html>
`, MainStyles, MainScripts)
