package pages

import "fmt"

//goland:noinspection all
var MainStyles = fmt.Sprintf(
	//language=scss
	`
body {
	background: linear-gradient(to right, #CEE5E7, #A8DBE5);
	height: 100vh;
	display: flex;
	justify-content: center;
	font-family: Arial, sans-serif;
	color: #333;
	margin: 0;
}
#form{
	height: 80vh;
}

label {
	font-size: 1rem;
	margin-bottom: 0.5rem;
	color: #555;
}

input[type='text'] {
	padding: 0.75rem;
	font-size: 1rem;
	border-radius: 4px;
	border: 1px solid #ddd;
	transition: border-color 0.3s;
}

input[type='text']:focus {
	border-color: #A8DBE5;
	outline: none;
}

button {
	padding: 0.75rem;
	font-size: 1rem;
	background-color: #56c6ca;
	color: #fff;
	border: none;
	border-radius: 4px;
	cursor: pointer;
	transition: background-color 0.3s;
}

button:hover {
	background-color: #3fa6a9;
}

button:active {
	background-color: #349396;
}

.content-output{
	
}

.searchFormWrapper {
	display: flex;
	justify-content: space-around;
	margin-top: 5rem;
	width: 50vw;
	padding: 10px;
	background-color: #ffffff;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	border-radius: 8px;
}
.searchForm{
	display: flex;
	flex-direction: column;
	width: 100%%;
	background-color: chocolate;
}
.search-input-wrapper{
	background-color: #ffffff;
}
.search-button-wrapper{
	
}

`)
