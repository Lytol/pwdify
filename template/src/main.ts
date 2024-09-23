import './style.css'
import { bytesToHex, decrypt, hexToBytes, passwordHash } from './decrypt'

const sessionStorageKey = 'pwdify'

const form = () => document.querySelector('form') as HTMLFormElement
const passwordInput = () => document.querySelector('input[type="password"]') as HTMLInputElement

document.addEventListener('DOMContentLoaded', () => {
  // If a session storage key exists, try to use it
  const hash = getPasswordHash()
  if (hash) {
    decryptPage(hash)
    return
  }

  // Focus the password input
  passwordInput().focus()

  form().addEventListener('submit', async (event) => {
    event.preventDefault()
    const password = passwordInput().value
    const hash = await passwordHash(password)
    decryptPage(hash)
  })
})

async function decryptPage(hash: Uint8Array) {
  const encryptedContent = document.querySelector('body')!.getAttribute('data-encrypted-content')!

  try {
    const decryptedContent = await decrypt(encryptedContent, hash)
    document.write(decryptedContent)
    document.close()
    savePasswordHash(hash)
  } catch (error) {
    const input = passwordInput()
    input.onanimationend = () => input.classList.remove('error')
    input.classList.add('error')
    resetPasswordHash()
  }
}

function savePasswordHash(hash: Uint8Array) {
  const hashString = bytesToHex(hash)
  window.sessionStorage.setItem(sessionStorageKey, hashString)
}

function getPasswordHash(): Uint8Array | null {
  const hashString = window.sessionStorage.getItem(sessionStorageKey)
  if (!hashString) {
    return null
  }
  return hexToBytes(hashString)
}

function resetPasswordHash() {
  window.sessionStorage.removeItem(sessionStorageKey)
}