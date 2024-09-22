import './style.css'

document.addEventListener('DOMContentLoaded', () => {
  const passwordInput = document.querySelector('input[type="password"]') as HTMLInputElement;
  passwordInput.focus();

  const encryptedContent = document.querySelector('body')!.getAttribute('data-encrypted-content')!;
  console.log(encryptedContent);

  const form = document.querySelector('form') as HTMLFormElement;

  form.addEventListener('submit', (event) => {
    event.preventDefault();
    const password = passwordInput.value;
    console.log(password);
  });
});
