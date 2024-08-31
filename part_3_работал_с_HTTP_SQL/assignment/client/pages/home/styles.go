package pages

var MainStyles =
// language=scss
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
.search-form-wrapper {
	display: flex;
	justify-content: space-between;
	margin-top: 1rem;
	width: 50vw;
	padding: 10px;
	background-color: #ffffff;
	box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	border-radius: 8px;
}

#search-input {
	padding: 0.75rem;
	font-size: 1rem;
	border-radius: 4px;
	border: 1px solid #ddd;
	transition: border-color 0.3s;
	width: 100%;
    margin: 0 10px;
}
#search-input:focus {
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
	width: 73px;
}
button:hover {
	background-color: #3fa6a9;
}
button:active {
	background-color: #349396;
}

.search-output{
	width: 50vw;
	background-color: #ffffff;
	min-height: 80vh;	
    padding: 10px;
	border-radius: 4px;
	margin-top: 1rem;
}

`
