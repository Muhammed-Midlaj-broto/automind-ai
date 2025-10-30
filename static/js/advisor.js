document.getElementById('advisorForm').addEventListener('submit', async function (e) {
  e.preventDefault();

  const formData = new FormData(this);
  const issue = formData.get('issue').trim();
  const responseDiv = document.getElementById('response');

  if (!issue) {
    alert('Please describe your issue!');
    return;
  }

  responseDiv.style.display = 'block';
  responseDiv.innerText = "⏳ Getting AI advice...";

  try {
    const res = await fetch('/advisor/ask', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ prompt: issue })
    });

    const data = await res.json();

    if (res.ok) {
      responseDiv.innerText = "💡 " + data.answer;
    } else {
      responseDiv.innerText = "❌ Error: " + (data.error || "Something went wrong");
    }

  } catch (err) {
    responseDiv.innerText = "❌ Error: " + err.message;
  }

  this.reset();
});
