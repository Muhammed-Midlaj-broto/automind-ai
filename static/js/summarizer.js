document.getElementById('summarizerForm').addEventListener('submit', async function(e) {
  e.preventDefault();

  const formData = new FormData(this);
  const data = Object.fromEntries(formData);
  const responseDiv = document.getElementById('response');

  responseDiv.style.display = 'block';
  responseDiv.textContent = '⏳ Getting AI summary...';

  try {
    const res = await fetch('/summarizer/ask', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });

    const result = await res.text();
    responseDiv.textContent = `📋 AI Summary:\n${result}`;
  } catch (error) {
    responseDiv.textContent = '❌ Error: Could not fetch AI summary.';
    console.error(error);
  }

  this.reset();
});
