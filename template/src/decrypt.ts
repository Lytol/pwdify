const algorithm = 'AES-GCM'
const nonceLength = 12

export async function decrypt(data: string, password: string): Promise<string> {
  const key = await passwordToKey(password)

  const [iv, ciphertext] = splitData(data)

  const decryptedData = new Uint8Array(
    await window.crypto.subtle.decrypt({ name: algorithm, iv }, key, ciphertext),
  )

  const decoder = new TextDecoder()
  return decoder.decode(decryptedData)
}

const passwordToKey = async (password: string): Promise<CryptoKey> => {
  const hashed = await window.crypto.subtle.digest('SHA-256', new TextEncoder().encode(password))

  return window.crypto.subtle.importKey('raw', hashed, algorithm, false, ['decrypt'])
}

function hexToByteArray(hex: string): Uint8Array {
  const matchArray = hex.match(/.{1,2}/g)
  if (!matchArray) {
    throw new Error('Invalid hex string')
  }

  return Uint8Array.from(matchArray.map(byte => parseInt(byte, 16)));
}

function splitData(data: string): [Uint8Array, Uint8Array] {
  const nonce = hexToByteArray(data.slice(0, nonceLength * 2))
  const ciphertext = hexToByteArray(data.slice(nonceLength * 2))

  return [nonce, ciphertext]
}