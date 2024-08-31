package pages

import "fmt"

var MainScripts = fmt.Sprintf(
	//language=html
	`
<script>
	document.addEventListener('DOMContentLoaded', function() {
	fetch('/getBreeds')
		.then(response => response.json())
		.then(data => {
			console.log('Select options breeds:', data);
			// const select = document.getElementById('search-select');
			// data.forEach(item => {
			// 	let option = document.createElement('option');
			// 	option.value = item;
			// 	option.textContent = item;
			// 	select.appendChild(option);
			// });
		})
		.catch(error => console.error('Error fetching select options:', error));
	});
	function submitForm() {
		const data = document.getElementById('search-input').value
		const selectedOption = document.getElementById('search-select').value;
		console.log("data from input:", data)
		console.log("Selected option:", selectedOption);

	}
</script>
`)
