document.getElementById('predictorForm').addEventListener('submit', async function(e) {
  e.preventDefault();

  const formData = new FormData(this);
  const data = Object.fromEntries(formData);
  const responseDiv = document.getElementById('response');

  // Show loading message
  responseDiv.style.display = 'block';
  responseDiv.innerText = "⏳ Getting AI prediction...";

  try {
    const res = await fetch('/predictor/ask', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

    if (res.ok) {
      // Parse JSON response from backend
      const result = await res.json();
      responseDiv.innerText = result.answer || "No prediction received.";
    } else {
      const errorResult = await res.json();
      responseDiv.innerText = "❌ Error: " + (errorResult.error || "Something went wrong.");
    }
  } catch (err) {
    responseDiv.innerText = "❌ Error: " + err.message;
  }

  // Optional: reset the form after submission
  // this.reset();
});
