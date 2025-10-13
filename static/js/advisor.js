  document.getElementById('advisorForm').addEventListener('submit', function(e) {
            e.preventDefault();
            const formData = new FormData(this);
            console.log('Advisor Data:', Object.fromEntries(formData));
            document.getElementById('response').style.display = 'block';
            this.reset();
        });