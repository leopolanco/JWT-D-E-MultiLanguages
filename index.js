window.addEventListener('DOMContentLoaded', () => {
	document
		.getElementById('action-select')
		.addEventListener('change', selectChange)

	document.getElementById('submit-btn').addEventListener('click', submit)
})

const selectChange = (e) => {
	const action = e.target.value
	const contentArea = document.getElementById('contentarea')
	contentArea.setAttribute('data-action', action)
	contentArea.innerText = ''
    document.getElementById('result').innerText = ''

	if (action === 'decode') {
		const area = document.createElement('textarea')
		area.id = 'content'
		area.placeholder = 'Enter the encoded JWT'
		area.style = 'margin: 0px; width: 197px; height: 220px;'
		contentArea.appendChild(area)
	} else {
		createKeyValueInput(contentArea)
		const createBtn = document.createElement('button')
		createBtn.innerText = 'Add'
		createBtn.addEventListener('click', (e) => {
			createKeyValueInput(contentArea)
		})
		contentArea.appendChild(createBtn)
	}
}

const submit = async (e) => {
	e.preventDefault()
	const language = document.getElementById('language-select').value
	const resultArea = document.getElementById('result')
	const contentArea = document.getElementById('contentarea')
	const action = document
		.getElementById('contentarea')
		.getAttribute('data-action')

	let result = 'Sorry there was an error'
	let content = ''

	// Get the content
	// In decode we only take the string
	// But for encode we need to build the object
	if (action === 'decode') {
		content = contentArea.getElementsByTagName('textarea')[0].value
	}
	if (action === 'encode') {
		content = {}
		const inputs = contentArea.getElementsByTagName('input')
		for (let i = 0; i < inputs.length; i++) {
			const key = inputs[i].value
			const value = inputs[++i].value
			if (key && value) {
				content[key] = value
			}
		}
	}
	result = await postToEndpoint(action, content)
	resultArea.innerText = result
}

const createKeyValueInput = (parent) => {
	const wrapper = document.createElement('div')
	wrapper.style = 'margin-bottom:0.5rem'
	const key = document.createElement('input')
	key.type = 'text'
	key.placeholder = 'Key'
	wrapper.appendChild(key)

	const value = document.createElement('input')
	value.type = 'text'
	value.placeholder = 'Value'
	wrapper.appendChild(value)

	const del = document.createElement('button')
	del.innerText = '-'
	del.addEventListener('click', (e) => {
		e.target.parentElement.remove()
	})
	wrapper.appendChild(del)

	parent.prepend(wrapper)
}

const postToEndpoint = async (action, data) => {
	const url = `http://127.0.0.1:8000/${action}`
	if (!data) return 'No data provided'
	try {
		const res = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})

		return await res.json()
	} catch (error) {
		console.error(error)
		return 'Sorry there was an error'
	}
}
