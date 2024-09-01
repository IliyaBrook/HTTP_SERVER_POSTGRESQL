package pages

import "fmt"

var MainScripts = fmt.Sprintf(
	//language=html
	`
<script>
	document.addEventListener('DOMContentLoaded', function() {
		const datalist = document.getElementById('breeds-datalist');
		const searchInput = document.getElementById('search-input');
		const submitButton = document.getElementById('submit-button');
		const clearButton = document.getElementById('clear-button');
		const searchOutput = document.querySelector('.search-output');
		const spinner = document.getElementById('loading-spinner');
		
		function showSpinner() {
			spinner.style.display = 'block';
		}

		function hideSpinner() {
			spinner.style.display = 'none';
		}
		
		fetch('/getBreeds')
			.then(response => response.json())
			.then(dogBreeds => {				
				dogBreeds.forEach(breed => {
						let option = document.createElement('option');
						option.value = breed;
						datalist.appendChild(option);
					});
			})
			.catch(error => console.error('Error fetching select options:', error));
	
		function setDisabled(value) {
			let pointerEvents;
			if (value) {
				pointerEvents = 'none'
			}else {
				pointerEvents = 'auto'
			}
			submitButton.disabled = value;
			submitButton.style.pointerEvents = pointerEvents;
		}
		setDisabled(true);
		
		searchInput.addEventListener('change', onChange);
		submitButton.addEventListener('click', submitForm);
		clearButton.addEventListener('click', clearOutput);
		
		function clearOutput() {
			searchOutput.innerHTML = '';
            searchInput.value = ''; 
		}
		
		function onChange(event) {
			const value = event.target.value.trim();
			if (value.length > 0) {
				setDisabled(false)
			}else {
				setDisabled(true)
			}
		}
		
		function submitForm() {
			showSpinner();
			fetch("/getDogs?breed=" + searchInput.value)
				.then(response => {
					hideSpinner();
					if (response.ok) {
						return response.json()
					}else {
						console.error("getDogs error:", response)
					}
				})
				.then(dogs => {
					if (dogs.length > 0) {
						dogs.forEach(dog => {
							const dogCard = document.createElement('div');
							dogCard.classList.add('dog-card');
							dogCard.innerHTML = 
							'<img src="' + dog.image.url + '" alt="' + dog.name + '" class="dog-image">' +
							'<h3 class="dog-name">' + dog.name + '</h3>' +
							'<p class="dog-breed-group"><strong>Breed Group:</strong> ' + dog?.breed_group + '</p>' +
							'<p class="dog-bred-for"><strong>Bred For:</strong> ' + dog?.bred_for + '</p>' +
							'<p class="dog-life-span"><strong>Life Span:</strong> ' + dog?.life_span + '</p>' +
							'<p class="dog-temperament"><strong>Temperament:</strong> ' + dog?.temperament + '</p>' +
							'<p class="dog-weight"><strong>Weight:</strong> ' + dog.weight?.imperial + ' lbs (' + dog.weight?.metric + ' kg)</p>' +
							'<p class="dog-height"><strong>Height:</strong> ' + dog.height?.imperial + ' in (' + dog.height?.metric + ' cm)</p>';
							searchOutput.appendChild(dogCard);
						});
					}
				})
		}
	});
	
</script>
`)
