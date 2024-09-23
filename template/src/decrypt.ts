const algorithm = 'AES-GCM'
const nonceLength = 12

export async function decrypt(data: string, passwordHash: Uint8Array): Promise<string> {
  const key = await hashToKey(passwordHash)
  const [iv, ciphertext] = splitData(data)

  const decryptedData = new Uint8Array(
    await window.crypto.subtle.decrypt({ name: algorithm, iv }, key, ciphertext),
  )

  const decoder = new TextDecoder()
  return decoder.decode(decryptedData)
}

export async function passwordHash(password: string): Promise<Uint8Array> {
  return new Uint8Array(
    await window.crypto.subtle.digest('SHA-256', new TextEncoder().encode(password))
  )
}

const hashToKey = (hash: Uint8Array): Promise<CryptoKey> => {
  return window.crypto.subtle.importKey('raw', hash, algorithm, false, ['decrypt'])
}

export function hexToBytes(hex: string): Uint8Array {
  const matchArray = hex.match(/.{1,2}/g)
  if (!matchArray) {
    throw new Error('Invalid hex string')
  }

  return Uint8Array.from(matchArray.map(byte => parseInt(byte, 16)))
}

export function bytesToHex(bytes: Uint8Array): string {
  return Array.from(bytes, byte => ('0' + (byte & 0xFF).toString(16)).slice(-2)).join('')
}

function splitData(data: string): [Uint8Array, Uint8Array] {
  const nonce = hexToBytes(data.slice(0, nonceLength * 2))
  const ciphertext = hexToBytes(data.slice(nonceLength * 2))

  return [nonce, ciphertext]
}
