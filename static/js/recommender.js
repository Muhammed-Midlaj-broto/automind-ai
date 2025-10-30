document.getElementById('recommenderForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(this);
    const budget = formData.get('budget').trim();
    const useCase = formData.get('useCase').trim();
    const terrain = formData.get('terrain').trim();

    if (!budget || !useCase || !terrain) {
        alert('Please fill in all fields!');
        return;
    }

    const responseDiv = document.getElementById('response');
    responseDiv.style.display = 'block';
    responseDiv.innerText = "⏳ Getting AI recommendations...";

    try {
        const res = await fetch('/recommender', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                budget: budget,
                useCase: useCase,
                terrain: terrain
            })
        });

        const data = await res.json();

        if (res.ok) {
            responseDiv.innerText = data.answer; // Display AI-generated response
        } else {
            responseDiv.innerText = "❌ Error: " + (data.error || "Something went wrong");
        }

    } catch (err) {
        responseDiv.innerText = "❌ Error: " + err.message;
    }

    //this.reset(); // clear input fields
});
