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
		const searchOutput = document.querySelector('.search-output');
				const fakeDogsArray = [
    {
        "weight": {
            "imperial": "38 - 50",
            "metric": "17 - 23"
        },
        "height": {
            "imperial": "23 - 26",
            "metric": "58 - 66"
        },
        "id": 8,
        "name": "Alaskan Husky",
        "bred_for": "Sled pulling",
        "breed_group": "Mixed",
        "life_span": "10 - 13 years",
        "temperament": "Friendly, Energetic, Loyal, Gentle, Confident",
        "reference_image_id": "-HgpNnGXl",
        "image": {
            "id": "-HgpNnGXl",
            "width": 500,
            "height": 500,
            "url": "https://cdn2.thedogapi.com/images/-HgpNnGXl.jpg"
        }
    },
    {
        "weight": {
            "imperial": "65 - 100",
            "metric": "29 - 45"
        },
        "height": {
            "imperial": "23 - 25",
            "metric": "58 - 64"
        },
        "id": 9,
        "name": "Alaskan Malamute",
        "bred_for": "Hauling heavy freight, Sled pulling",
        "breed_group": "Working",
        "life_span": "12 - 15 years",
        "temperament": "Friendly, Affectionate, Devoted, Loyal, Dignified, Playful",
        "reference_image_id": "dW5UucTIW",
        "image": {
            "id": "dW5UucTIW",
            "width": 1023,
            "height": 769,
            "url": "https://cdn2.thedogapi.com/images/dW5UucTIW.jpg"
        }
    }
]

		
	// fetch('/getBreeds')
	// 		.then(response => response.json())
	// 		.then(dogBreeds => {
	// 			console.log("breeds log:", dogBreeds) 
	//			
	// 			dogBreeds.forEach(breed => {
	// 					let option = document.createElement('option');
	// 					option.value = breed;
	// 					datalist.appendChild(option);
	// 				});
	// 		})
	// 		.catch(error => console.error('Error fetching select options:', error));
	
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
			
			function onChange(event) {
				const value = event.target.value.trim();
				if (value.length > 0) {
					setDisabled(false)
				}else {
					setDisabled(true)
				}
			}
			
			function submitForm() {
				// fetch("/getDogs?breed=" + searchInput.value)
				// 	.then(response => {
				// 		if (response.ok) {
				// 			return response.json()
				// 		}else {
				// 			console.error("getDogs error:", response)
				// 		}
				// 	})
				// 	.then(dogs => {
				// 		console.log("dogs array:", dogs)
				// 	})
			}
			
			fakeDogsArray.forEach(dog => {
			const dogCard = document.createElement('div');
			dogCard.classList.add('dog-card');
			dogCard.innerHTML = 
			'<img src="' + dog.image.url + '" alt="' + dog.name + '" class="dog-image">' +
			'<h3 class="dog-name">' + dog.name + '</h3>' +
			'<p class="dog-breed-group"><strong>Breed Group:</strong> ' + dog.breed_group + '</p>' +
			'<p class="dog-bred-for"><strong>Bred For:</strong> ' + dog.bred_for + '</p>' +
			'<p class="dog-life-span"><strong>Life Span:</strong> ' + dog.life_span + '</p>' +
			'<p class="dog-temperament"><strong>Temperament:</strong> ' + dog.temperament + '</p>' +
			'<p class="dog-weight"><strong>Weight:</strong> ' + dog.weight.imperial + ' lbs (' + dog.weight.metric + ' kg)</p>' +
			'<p class="dog-height"><strong>Height:</strong> ' + dog.height.imperial + ' in (' + dog.height.metric + ' cm)</p>';
			searchOutput.appendChild(dogCard);
		});
	});
	
</script>
`)
