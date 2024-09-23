import './style.css'
import { decrypt } from './decrypt'

document.addEventListener('DOMContentLoaded', () => {
  const passwordInput = document.querySelector('input[type="password"]') as HTMLInputElement;
  passwordInput.focus();

  const encryptedContent = document.querySelector('body')!.getAttribute('data-encrypted-content')!;

  const form = document.querySelector('form') as HTMLFormElement;

  form.addEventListener('submit', async (event) => {
    event.preventDefault();
    const password = passwordInput.value;

    try {
      const decryptedContent = await decrypt(encryptedContent, password)
      document.write(decryptedContent)
      document.close()
    } catch (error) {
      passwordInput.onanimationend = () => passwordInput.classList.remove('error')
      passwordInput.classList.add('error')
    }
  });
});
