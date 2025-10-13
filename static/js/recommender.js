document.getElementById('recommenderForm').addEventListener('submit', function(e) {
      e.preventDefault();
      const formData = new FormData(this);
      console.log('Recommender Data:', Object.fromEntries(formData));
      document.getElementById('response').style.display = 'block';
      this.reset();
    });