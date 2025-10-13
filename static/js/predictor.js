document.getElementById('predictorForm').addEventListener('submit', function(e) {
      e.preventDefault();
      const formData = new FormData(this);
      console.log('Predictor Data:', Object.fromEntries(formData));
      document.getElementById('response').style.display = 'block';
      this.reset();
    });