package pages

import "fmt"

var MainScripts = fmt.Sprintf(
	//language=html
	`
<script>
	function submitForm() {
		const data = document.getElementById('data').value
		console.log("data from input:", data)
		// fetch('http://localhost:8081/submit', {
		// 	method: 'POST',
		// 	headers: {
		// 		'Content-Type': 'application/json'
		// 	},
		// 	body: JSON.stringify({ data: data })
		// })
		// 	.then(response => response.text())
		// 	.then(data => {
		// 		alert('Response: ' + data)
		// 	})
	}
</script>
`)
