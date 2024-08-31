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

#form {
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

#clear-button{
	margin-left: 10px;
	background-color: #f05050;
}
#clear-button:hover{
	background-color: #e03d3d;
}

.search-output{
	position: relative;
	width: 50vw;
	background-color: #ffffff;
	min-height: 80vh;	
    padding: 10px;
    border-radius: 4px;
    margin-top: 1rem;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.dog-card {
    background-color: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    padding: 15px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.dog-image {
    width: 100%;
    height: auto;
    border-radius: 8px;
}

.dog-name {
    font-size: 1.5rem;
    margin: 10px 0;
    color: #333;
}

.dog-breed-group, .dog-bred-for, .dog-life-span, .dog-temperament, .dog-weight, .dog-height {
    font-size: 1rem;
    color: #666;
    margin: 5px 0;
}

.spinner {
    display: none;
	position: absolute;
	left: 45%;
	top: 45%;
    margin: 20px auto;
    border: 8px solid #f3f3f3; 
    border-top: 8px solid #56c6ca; 
    border-radius: 50%;
    width: 60px;
    height: 60px;
    animation: spin 1s linear infinite;
}

@keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}
`
